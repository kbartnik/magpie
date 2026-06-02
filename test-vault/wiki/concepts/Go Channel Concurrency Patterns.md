---
title: "Go Channel Concurrency Patterns"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-31
sources:
  - "[[preflight-sync-go]] stopCh → context.Context refactor (2026-05-30)"
  - "archive/books/2026-05-31-learning-go-2e.md (Ch.12)"
related:
  - "Fsnotify Filesystem Watching"
  - "Bubbletea Async Patterns"
tags: [go, concurrency, channels, preflight-sync-go]
---

# Go Channel Concurrency Patterns

## select: Multiplexing Channel Operations

`select` blocks until one of its cases is ready, then runs that case. If multiple cases are ready simultaneously, Go picks one at random — this is deliberate, not a bug.

```go
select {
case msg := <-workCh:
    // process msg
case <-ctx.Done():
    return ctx.Err()
default:
    // non-blocking: runs immediately if no other case is ready
}
```

`default` makes a select non-blocking. Without it, `select` waits indefinitely for any case to be ready.

**Channel close semantics:** a closed channel always returns immediately with the zero value and `ok=false`. A write to a closed channel panics. This is what makes close-based cancellation work — once closed, every goroutine blocking on that channel unblocks and stays unblocked.

## Done Channel (`chan struct{}`) — Cancellation Before context

The pre-context idiom uses a plain channel as a broadcast stop signal:

```go
stopCh := make(chan struct{})

// in goroutine:
select {
case <-stopCh:
    return
case work := <-workCh:
    process(work)
}

// to cancel all listeners:
close(stopCh)
```

`close` broadcasts to all listeners simultaneously — every goroutine blocking on `<-stopCh` unblocks at once. This is the key behaviour: a channel receive on a closed channel never blocks.

`chan struct{}` is preferred over `chan bool` because `struct{}` is a zero-byte type with no heap allocation, and the type itself communicates "this is a signal, not a value" — see [[Go Struct Memory Layout]].

**Limitation:** `chan struct{}` is convention, not contract. The type carries no information about intent, timeout, or cause.

## context.Context — The Standard Cancellation Primitive

`context.Context` solves the same problem as `chan struct{}` but makes intent explicit and adds composability:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // always cancel to release resources

// pass ctx downstream:
go worker(ctx, workCh)

// to cancel:
cancel()
```

Inside a goroutine, `ctx.Done()` returns the underlying channel — same select pattern as before:

```go
select {
case <-ctx.Done():
    return ctx.Err() // context.Canceled or context.DeadlineExceeded
case work := <-workCh:
    process(work)
}
```

`ctx.Err()` tells you *why* the context ended — cancelled by `cancel()` returns `context.Canceled`; an expired deadline returns `context.DeadlineExceeded`.

**Why context over `chan struct{}`:**
- Intent is in the type — any function taking `context.Context` signals it respects cancellation
- Composable — `context.WithTimeout`, `context.WithDeadline`, and `context.WithCancel` nest; a child context is cancelled when any ancestor is cancelled
- Carries metadata — `context.WithValue` (use sparingly; only for request-scoped data, not function parameters)

`context.Background()` is the root context — never cancelled, no deadline. Use it at the top of a call tree (e.g. in `main`) and pass derived contexts downstream.

## Bubbletea: Delivering a Context to the Model

Bubbletea models are **value types** — `Update` receives and returns a copy of the struct. `context.CancelFunc` (a `func()`) and `context.Context` (an interface) are both reference types, so copying the struct copies the handle, not the state — all copies share the same underlying context. This makes them safe to store on a value model.

The wrinkle: `Init()` returns `tea.Cmd`, not `(tea.Model, tea.Cmd)`, so it cannot directly mutate the model. The idiomatic solution is a **setup message**:

```go
type ctxInitMsg struct {
    ctx    context.Context
    cancel context.CancelFunc
}

func (m AppModel) Init() tea.Cmd {
    ctx, cancel := context.WithCancel(context.Background())
    return tea.Batch(
        tea.EnterAltScreen,
        func() tea.Msg { return ctxInitMsg{ctx: ctx, cancel: cancel} },
    )
}

// in Update:
case ctxInitMsg:
    m.ctx = msg.ctx
    m.cancel = msg.cancel
    return m, nil
```

`tea.Cmd` is just `func() tea.Msg` — returning an inline function literal that returns a message is the standard way to fire an immediate command. The model holds a no-op placeholder cancel (`func() {}`) until `ctxInitMsg` arrives, so a quit keypress in the narrow startup window is safe.

This pattern generalises: any initialisation that needs to land on the model goes through a setup message from `Init()`.

## errgroup: Structured Fan-Out

`golang.org/x/sync/errgroup` is the clean solution for running N concurrent tasks, waiting for all, and propagating the first error. Raw `sync.WaitGroup` + error channel is easy to deadlock or miss errors.

```go
import "golang.org/x/sync/errgroup"

func fetchAll(ctx context.Context, ids []string) ([]Result, error) {
    g, ctx := errgroup.WithContext(ctx)
    results := make([]Result, len(ids))

    for i, id := range ids {
        i, id := i, id // capture loop variables
        g.Go(func() error {
            r, err := fetch(ctx, id)
            if err != nil {
                return err
            }
            results[i] = r
            return nil
        })
    }

    return results, g.Wait()
}
```

When any goroutine returns an error, the derived `ctx` is cancelled — other goroutines that respect context will exit early.

### Why Not WaitGroup + Error Channel?

The naive hand-rolled version has a subtle bug: goroutines keep running after the first error, wasting resources on slow dependencies (network calls, external APIs). Adding `context.WithCancel` fixes it but grows to ~40 lines of orchestration code for a conceptually simple goal. `errgroup` collapses that to 10. The worked example is K8s readiness probes — sequential `db.Ping(); redis.Ping()` runs in `O(N × latency)` on every probe; `errgroup` makes it `O(max latency)`.

For **bounded concurrency** (fan-out N tasks, K at a time), pair with a semaphore:

```go
sem := semaphore.NewWeighted(10) // max 10 concurrent

g.Go(func() error {
    if err := sem.Acquire(ctx, 1); err != nil {
        return err
    }
    defer sem.Release(1)
    return doWork(ctx)
})
```

## See Also

- [[Fsnotify Filesystem Watching]] — uses `select` over `watcher.Events` and `ctx.Done()` for graceful shutdown
- [[Bubbletea Async Patterns]] — two-message pattern for non-blocking channel sends into the bubbletea runtime
- [[Go sync.Cond Pattern]] — when N-dynamic waiters need broadcast notification; `sync.Cond` beats channels here
- [[Go Context Patterns]] — `ctx.Done()` deep dive; context.Value misuse
- [[Go Empty Struct Pattern]] — `struct{}` as zero-cost channel signal type; the standard `done chan struct{}` idiom
- [[Go Channel Internals]] — `hchan` struct, `sudog` pooling, send/receive fast paths, happens-before rules, scheduler integration
