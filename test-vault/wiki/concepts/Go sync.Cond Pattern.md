---
title: "Go sync.Cond Pattern"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-advanced-concepts.md"
related:
  - "[[Go Channel Concurrency Patterns]]"
  - "[[Go Escape Analysis]]"
tags:
  - go
  - concurrency
  - sync
---

# Go sync.Cond Pattern

`sync.Cond` is the right tool when you need to **broadcast a state change to an arbitrary number of waiting goroutines**. Channels handle N consumers awkwardly when N is dynamic; `sync.Cond` is designed for exactly this.

## Core Pattern

```go
type Queue struct {
    mu    sync.Mutex
    cond  *sync.Cond
    items []Job
}

func NewQueue() *Queue {
    q := &Queue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *Queue) Push(j Job) {
    q.mu.Lock()
    q.items = append(q.items, j)
    q.cond.Broadcast() // wake all waiting consumers
    q.mu.Unlock()
}

func (q *Queue) Pop() Job {
    q.mu.Lock()
    defer q.mu.Unlock()
    for len(q.items) == 0 {
        q.cond.Wait() // atomically releases lock and suspends
    }
    j := q.items[0]
    q.items = q.items[1:]
    return j
}
```

## Critical Details

**`Wait()` is atomic.** It releases the mutex and suspends the goroutine in a single operation. When it returns, the lock is re-acquired. This atomicity prevents the race between "check condition" and "go to sleep."

**Spurious wakeups exist.** `Wait()` may return even when the condition is not true. Always re-check the condition in a `for` loop, not an `if`:

```go
// Correct
for len(q.items) == 0 {
    q.cond.Wait()
}

// Wrong — spurious wakeup will cause a panic
if len(q.items) == 0 {
    q.cond.Wait()
}
```

## Broadcast vs Signal

| Method | Wakes |
|--------|-------|
| `Broadcast()` | All waiters |
| `Signal()` | Exactly one waiter |

Use `Broadcast` when the state change is relevant to multiple waiters (e.g., new items in a shared queue). Use `Signal` when only one waiter can make progress (e.g., a single consumer pipeline).

## When to Prefer sync.Cond Over Channels

| Scenario | Better tool |
|----------|-------------|
| Fixed number of consumers | `chan` (close to broadcast, fan-out goroutine) |
| Dynamic N consumers waiting on one condition | `sync.Cond` |
| One-shot signal (done/cancel) | `chan struct{}` or `context.Context` |
| Collecting errors from concurrent tasks | `errgroup` |

The channel-based equivalent for N dynamic consumers requires either N separate channels, a fan-out goroutine, or a closed channel (one-shot only) — none compose as cleanly.

## See Also

- [[Go Channel Concurrency Patterns]] — channels, errgroup, done-channel vs context
- [[Go Escape Analysis]] — sync.Pool for reducing allocation pressure on reusable objects
