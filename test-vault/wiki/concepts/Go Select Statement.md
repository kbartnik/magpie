---
tags: [concept, go, concurrency]
cluster: go
aliases: ["go select", "channel select", "non-blocking channel operations"]
related: ["Go Channel Concurrency Patterns", "Go Worker Pool Pattern", "Go Context Patterns"]
sources:
  - "[[archive/videos/2026-06-04-go-concurrency-patterns]]"
---

# Go Select Statement

`select` waits on multiple channel operations simultaneously. When multiple cases are ready, one is chosen at random.

## Basic Form

```go
select {
case v := <-ch1:
    // handle v
case ch2 <- x:
    // sent x
case <-ctx.Done():
    return ctx.Err()
default:
    // non-blocking: runs if no channel is ready
}
```

## Timeout Pattern

```go
select {
case result := <-workCh:
    return result
case <-time.After(5 * time.Second):
    return errors.New("timeout")
}
```

## Cancellation Pattern

```go
for {
    select {
    case work := <-workCh:
        process(work)
    case <-ctx.Done():
        return
    }
}
```

## Random Selection

When multiple cases are ready simultaneously, Go selects one at random. This prevents starvation but means you cannot rely on priority ordering.

## Connections

- [[Go Channel Concurrency Patterns]] — select is the mechanism enabling fan-in and done-channel patterns
- [[Go Context Patterns]] — `ctx.Done()` in a select case is the standard cancellation idiom
- [[Go Worker Pool Pattern]] — worker goroutines use select to drain work channels while respecting cancellation
