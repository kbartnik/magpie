---
tags: [concept, go, concurrency]
cluster: go
aliases: ["context.Context", "go context", "context cancellation", "context values", "WithCancel", "WithTimeout"]
related: ["Go Channel Concurrency Patterns", "Go Worker Pool Pattern", "Go Error Handling", "Go Structured Logging"]
sources:
  - "[[archive/clippings/2026-06-04-go-context-patterns]]"
---

# Go Context Patterns

`context.Context` carries cancellation signals, deadlines, and request-scoped metadata across API boundaries.

## The Rules

1. **First parameter named `ctx`:** `func DoSomething(ctx context.Context, arg Arg) error`
2. **Never store in structs** — context belongs to a call path, not an object
3. **`context.Background()`** only at the top of a call tree (main, test entry, server handler)
4. **`context.TODO()`** as a placeholder when the right context is unclear

## Cancellation

```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel() // always call cancel to release resources

// With deadline:
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
```

Workers check `ctx.Done()`:
```go
select {
case work := <-jobs:
    process(work)
case <-ctx.Done():
    return ctx.Err()
}
```

## Context Values

`context.WithValue` — for request-scoped metadata only (trace IDs, auth tokens). **Not** for dependency injection. Keys must be unexported types to prevent collisions:

```go
type contextKey string
const traceKey contextKey = "trace-id"
```

## Connections

- [[Go Channel Concurrency Patterns]] — `context.WithCancel` replaced the done-channel cancellation pattern
- [[Go Worker Pool Pattern]] — context cancellation is the standard worker pool shutdown mechanism
- [[Go Error Handling]] — `context.Canceled` and `context.DeadlineExceeded` are sentinel errors
- [[Go Structured Logging]] — attaching a slog.Logger to context with trace ID is the narrow acceptable use of context values
