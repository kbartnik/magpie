---
title: "Go Structured Logging"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-go-slog-tour.md"
related:
  - "Go Context Patterns"
  - "Go Error Handling"
tags: [go, logging, observability, stdlib]
---

# Go Structured Logging

`log/slog` is the structured logging package added to the Go standard library in Go 1.21. It provides key-value structured output while staying interoperable with existing third-party logging packages.

## Core API

```go
// Top-level functions use the default logger
slog.Info("request received", "method", "GET", "path", "/api/v1/users")
slog.Error("db query failed", "err", err, "query", q)

// Create a custom logger with a specific handler
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello", "user", username)
```

Four named levels: `Debug`, `Info`, `Warn`, `Error`. Levels are integers — `Info` is 0, `Warn` is 4. Custom intermediate levels are valid.

## Handlers

The `Handler` interface is the backend. `Logger` is the frontend that calls it. This split lets existing packages (Zap, logr, hclog) plug in a `slog.Handler` without rewriting callers.

Built-in handlers:
- `slog.NewTextHandler(w, opts)` — `key=value` text format
- `slog.NewJSONHandler(w, opts)` — JSON objects, one per log call

The `Handler` interface has four methods: `Enabled`, `Handle`, `WithAttrs`, `WithGroup`. `Enabled` is called first on every log event — returning false drops the event before any formatting work.

## Performance: LogAttrs and Logger.With

The alternating key-value syntax (`slog.Info("msg", "k", v)`) is convenient but allocates. For hot paths:

```go
// Typed Attr avoids interface{} boxing
slog.LogAttrs(ctx, slog.LevelInfo, "request",
    slog.String("method", r.Method),
    slog.Int("status", code),
    slog.Duration("latency", elapsed),
)
```

`Logger.With` pre-formats shared attributes once and returns a new logger with them baked in. Use this for request-scoped loggers:

```go
reqLogger := logger.With(
    slog.String("request_id", rid),
    slog.String("user_id", uid),
)
// reqLogger carries request_id and user_id on every call
reqLogger.Info("handler started")
reqLogger.Info("db query", "rows", n)
```

The `WithAttrs` method on `Handler` is called by `With` — a well-implemented handler pre-serializes these attributes once rather than repeating them on every `Handle` call.

## Context Integration

Pass `context.Context` to include trace IDs and other request metadata:

```go
slog.InfoContext(ctx, "operation complete", "duration", elapsed)
```

Cancelling the context does **not** suppress the log write — the context is for extracting values, not for gating output.

## LogValuer: Custom Formatting

Types can control their log representation by implementing `LogValuer`:

```go
type Secret string

func (s Secret) LogValue() slog.Value {
    return slog.StringValue("[REDACTED]")
}
```

Useful for: redacting sensitive fields, logging a struct as a group of key-value pairs, computing expensive representations lazily.

## Groups

Attributes can be nested under a group name for disambiguation:

```go
logger.Info("request",
    slog.Group("http",
        slog.String("method", "GET"),
        slog.Int("status", 200),
    ),
)
// JSON: {"msg":"request","http":{"method":"GET","status":200}}
```

## See Also

- [[Go Context Patterns]] — passing context to log functions for trace ID extraction
- [[Go Error Handling]] — logging errors with `slog.Any("err", err)`
