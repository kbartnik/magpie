---
tags: [concept, go, concurrency]
cluster: go
aliases: ["go memory model", "happens-before", "data race", "go synchronization guarantees", "goroutine synchronization"]
related: ["Go Channel Concurrency Patterns", "Go Worker Pool Pattern", "Go Escape Analysis", "Go Tooling"]
sources:
  - "[[archive/papers/2026-06-04-go-memory-model]]"
---

# Go Memory Model

Defines when a goroutine reading a variable is guaranteed to observe writes by another goroutine. Without synchronization, compiler and CPU may reorder operations.

## Happens-Before

If event A *happens-before* event B, B is guaranteed to observe A's effects. Without a happens-before relationship, there is no ordering guarantee.

## Synchronization Guarantees

| Primitive | Guarantee |
|---|---|
| `go f()` | goroutine creation happens-before f starts |
| Channel send | send happens-before corresponding receive completes |
| Channel close | close happens-before receive of zero value |
| Buffered channel | kth receive happens-before (k+C)th send (C = capacity) |
| `sync.Mutex` | Unlock n happens-before Lock n+1 |
| `sync.WaitGroup` | Done happens-before Wait returns |
| `sync.Once` | first f() returns before any other Do() returns |

## Data Race = Undefined Behavior

Programs with data races have undefined behavior — not "may produce stale values" but "the compiler can legally do anything." Use `go run -race` to detect races at runtime.

## Connections

- [[Go Channel Concurrency Patterns]] — channel operations have explicit happens-before guarantees; the done-channel and worker-pool patterns rely on them
- [[Go Worker Pool Pattern]] — buffered channel capacity rule is grounded in the memory model
- [[Go Escape Analysis]] — memory model guarantees apply to heap-allocated variables; escape analysis determines which variables are heap-allocated
- [[Go Tooling]] — `-race` flag enables the race detector; `-gcflags="-m"` shows escape decisions
