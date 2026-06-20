---
tags: [concept, go]
cluster: go
aliases: ["slog", "go logging", "structured logs", "log/slog", "Logger.With"]
related: ["Go Context Patterns", "Go Tooling"]
sources:
  - "[[archive/clippings/2026-06-04-go-context-patterns]]"
---

# Go Structured Logging

Go 1.21 added `log/slog` — the standard structured logging package. Structured logs emit key-value pairs, enabling log aggregation and querying.

## Core API

```go
logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

logger.Info("request completed",
    "method", r.Method,
    "path", r.URL.Path,
    "status", 200,
    "duration_ms", 42,
)
```

## Logger.With for Request Context

`Logger.With` returns a derived logger with pre-attached fields. Attach request-scoped fields once and pass the derived logger:

```go
reqLogger := logger.With(
    "trace_id", traceID,
    "user_id", userID,
)
// reqLogger carries trace_id and user_id on all subsequent calls
```

## Context Pattern

Attach the derived logger to context for propagation through call chains:

```go
ctx = context.WithValue(ctx, loggerKey, reqLogger)

// Retrieve:
func LoggerFromCtx(ctx context.Context) *slog.Logger {
    if l, ok := ctx.Value(loggerKey).(*slog.Logger); ok { return l }
    return slog.Default()
}
```

## Performance

`logger.LogAttrs(ctx, slog.LevelInfo, "msg", slog.String("k","v"))` avoids reflection-based attribute boxing. Use for hot paths.

## Connections

- [[Go Context Patterns]] — attaching a logger with trace ID to context is the acceptable exception to "no DI via context"
- [[Go Tooling]] — structured logs integrate with log aggregation tools; JSON handler output is machine-parseable
