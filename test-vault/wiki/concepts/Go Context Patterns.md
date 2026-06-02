---
title: "Go Context Patterns"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-advanced-concepts.md"
  - "archive/books/2026-05-31-learning-go-2e.md (Ch.14)"
related:
  - "[[Go Channel Concurrency Patterns]]"
  - "[[Go Functional Options Pattern]]"
  - "[[Go Error Handling]]"
tags:
  - go
  - concurrency
  - api-design
---

# Go Context Patterns

`context.Context` carries deadlines, cancellation signals, and request-scoped metadata across API boundaries. It is frequently misused as a dependency injection mechanism — this causes type-unsafety and runtime failures.

## The Rule

> `context.Value` is for metadata that *describes* a request, not for dependencies that *handle* it.

**Right:** trace IDs, request IDs, auth principals, correlation IDs — things that cross API boundaries and don't affect program correctness if absent.

**Wrong:** database connections, loggers, configuration, service clients — these are dependencies and belong in constructors or function parameters.

## Correct Use: Request-Scoped Metadata

```go
// Use an unexported custom type as the key — prevents collisions with other packages
type contextKey string
const requestIDKey contextKey = "requestID"

func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFrom(ctx context.Context) (string, bool) {
    id, ok := ctx.Value(requestIDKey).(string)
    return id, ok
}
```

**Why unexported custom key type?** Using a plain `string` or exported constant as a key means any package can accidentally collide with it. An unexported type is package-private — only your package can set or retrieve it.

## Wrong Use: Dependency Injection

```go
// Wrong — no type safety, nil dereference risk in production
db := ctx.Value("db").(*sql.DB)

// Right — explicit, compiler-checked, testable
type UserService struct {
    db     *sql.DB
    logger *slog.Logger
}

func NewUserService(db *sql.DB, logger *slog.Logger) *UserService {
    return &UserService{db: db, logger: logger}
}
```

Explicit dependency injection:
- The compiler verifies the types
- Tests can inject fakes without constructing a context
- `grep` finds every usage; context threading hides them

## Cancellation and Deadlines

The primary use of context — propagating cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel() // always cancel to free resources

result, err := doWork(ctx)
```

Goroutines that respect context should check `ctx.Err()` or select on `ctx.Done()`:

```go
select {
case <-ctx.Done():
    return ctx.Err()
case result := <-workCh:
    return result
}
```

## See Also

- [[Go Channel Concurrency Patterns]] — `ctx.Done()` as a cancellation channel; errgroup propagates context cancellation to child goroutines
- [[Go Functional Options Pattern]] — the right mechanism for construction-time configuration (not context)
- [[Go Structured Logging]] — `slog.InfoContext(ctx, ...)` extracts trace IDs and request metadata from context
