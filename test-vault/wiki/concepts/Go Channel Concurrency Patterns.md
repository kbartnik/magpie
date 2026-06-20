---
tags: [concept, go, concurrency]
cluster: go
aliases: ["go channels", "goroutine patterns", "generator pattern", "fan-in pattern", "done channel"]
related: ["Go Select Statement", "Go Worker Pool Pattern", "Go Context Patterns", "Go Memory Model", "Go Channel Internals"]
sources:
  - "[[archive/videos/2026-06-04-go-concurrency-patterns]]"
---

# Go Channel Concurrency Patterns

Channels are first-class values in Go — they can be passed as function parameters and returned from functions. This means patterns compose.

## Generator Pattern

A function that returns a channel and launches a goroutine to write to it. The caller receives values as if iterating a sequence, but the generation is concurrent.

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums { out <- n }
        close(out)
    }()
    return out
}
```

## Fan-In

Merge multiple input channels into one. A fan-in function is itself a generator. Enables N producers → 1 consumer.

## Done Channel (Pre-context)

Before `context.Context`, a `done` channel signaled cancellation. Workers `select` on both their work channel and `done`; closing `done` unblocks all workers simultaneously.

`context.WithCancel` is the modern replacement — it generalizes done channels with deadlines and values.

## Key Insight

Channels are values — patterns compose fractally. A generator that produces generators is valid. Fan-in is composable with generators. Buffer capacity sets the concurrency bound for worker pools.

## Connections

- [[Go Select Statement]] — select is the mechanism for simultaneously waiting on multiple channels
- [[Go Worker Pool Pattern]] — worker pools use buffered channels as the work queue
- [[Go Context Patterns]] — context.WithCancel replaced the done-channel cancellation pattern
- [[Go Memory Model]] — channel operations have specific happens-before guarantees
