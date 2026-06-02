---
title: "Go Memory Model"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-go-memory-model.md"
related:
  - "Go Channel Internals"
  - "Go Channel Concurrency Patterns"
  - "Go sync.Cond Pattern"
tags: [go, concurrency, memory-model, happens-before]
---

# Go Memory Model

The Go memory model defines when a memory write by one goroutine is guaranteed to be visible to another. Without these guarantees, modern CPUs and compilers are free to reorder operations in ways that break concurrent programs.

The canonical caution from the spec:

> *If you must read the rest of this document to understand the behavior of your program, you are being too clever. Don't be clever.*

## The Three-Tier Ordering Model

### 1. Sequenced-before (within a goroutine)

Within a single goroutine, operations follow logical code order. The compiler may reorder instructions internally but must preserve the observable outcome for that goroutine. This is what makes single-threaded code safe.

### 2. Synchronized-before (across goroutines)

When a **synchronizing write** in one goroutine is observed by a **synchronizing read** in another, a synchronized-before relationship is established. Synchronizing operations are:

| Write | Read |
|-------|------|
| Channel send | Channel receive |
| Mutex `Unlock` | Mutex `Lock` |
| Atomic store / CAS | Atomic load / CAS |

### 3. Happened-before (the guarantee)

Happened-before = sequenced-before + synchronized-before combined. If operation A happens-before operation B, then all memory writes visible at A are visible at B.

If you cannot establish a happens-before chain between a write and a read, they are **concurrent** — the program has a data race and the result is undefined.

## Guarantees Per Primitive

### Goroutine launch

The `go` statement happens-before the goroutine's body starts. Anything written before `go func()` is visible inside the goroutine.

```go
msg := "ready"
go func() {
    fmt.Println(msg) // guaranteed to print "ready"
}()
```

No guarantee about goroutine *completion* without explicit sync — the goroutine may still be running after the spawner moves on.

### Channels

- **Unbuffered:** send happens-before the corresponding receive completes; receive happens-before the send completes.
- **Buffered (capacity N):** the *i*th receive happens-before the *(i+N)*th send. The ith and (i+N-1)th send may proceed concurrently.
- **Close:** close happens-before any receive that returns the zero value due to the close.

```go
var shared int
done := make(chan struct{})

go func() {
    shared = 42
    close(done) // happens-before receivers unblocking
}()

<-done
fmt.Println(shared) // guaranteed to see 42
```

### Mutexes

`Unlock` in goroutine G1 happens-before any subsequent `Lock` in goroutine G2. All writes G1 made inside the critical section are visible to G2 after it acquires the lock.

### Atomics (`sync/atomic`)

An atomic store happens-before an atomic load that observes the stored value. This chains with other memory writes ordered before the store:

```go
counter.Store(10)
signal.Store(true)

// in another goroutine:
for !signal.Load() {}
fmt.Println(counter.Load()) // guaranteed to see 10
```

### Package initialization

`init()` functions run in import order (C → B → A if A imports B imports C). All `init()` functions complete before `main()` starts. Goroutines spawned inside `init()` are **not** guaranteed to complete before `main()` — they run independently.

### sync.Once

`once.Do(f)` completion happens-before any subsequent return from `once.Do` in any goroutine. Safe lazy initialization: every goroutine that calls `Do` after the first will block until `f` completes and then see `f`'s writes.

### sync.WaitGroup

`wg.Done()` happens-before the return of `wg.Wait()`. All writes inside the goroutine that calls `Done` are visible after `Wait` returns.

```go
wg.Add(1)
go func() {
    defer wg.Done()
    value = 123 // visible after wg.Wait()
}()
wg.Wait()
fmt.Println(value) // guaranteed to see 123
```

## The Common Mistake: Writing After the Sync Point

The happens-before guarantee applies only to writes **before** the synchronizing operation — not after:

```go
// Safe: x=42 is before the send
go func() { x = 42; ch <- 1 }()
<-ch
fmt.Println(x) // guaranteed 42

// Unsafe: x=99 is after the send
go func() { ch <- 1; x = 99 }()
<-ch
fmt.Println(x) // NOT guaranteed — data race
```

## Practical Rules

1. If two goroutines access the same variable and at least one writes, access must be synchronized.
2. Don't rely on timing, buffering, or "it works in practice" — establish explicit happens-before chains.
3. Use `-race` to detect violations: `go test -race ./...`

## See Also

- [[Go Channel Internals]] — how channels implement happens-before at the runtime level (hchan, sudog, scheduler)
- [[Go Channel Concurrency Patterns]] — select, done channel, context.Context patterns built on these guarantees
- [[Go sync.Cond Pattern]] — broadcast/wait for N-dynamic waiters; how Cond's lock establishes happens-before
