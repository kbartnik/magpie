---
title: "Go Error Handling"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources:
  - "archive/books/2026-05-31-learning-go-2e.md (Ch.9)"
related: ["Go Defined Types"]
tags: [go, errors, preflight-sync-go]
---

# Go Error Handling

Go has no exceptions. Functions that can fail return `(value, error)`. The caller checks `if err != nil` every time — explicit, no hidden control flow.

## The (value, error) pattern

```go
func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("load config: %w", err)
    }
    return cfg, nil
}
```

`if err != nil` is the most common two-word phrase in Go.

## Error wrapping with %w

`fmt.Errorf("load config: %w", err)` wraps the original error, preserving it for `errors.Is()` / `errors.As()` checks up the call stack.

Convention: prefix with the operation name so error chains read like a call stack:

```
load config: open file: no such file or directory
```

This is analogous to Rust's `?` operator but explicit every time.

## errors.Join

Combines multiple errors into a single value (Go 1.20+):

```go
var errs []error
errs = append(errs, fmt.Errorf("file1: %w", err1))
errs = append(errs, fmt.Errorf("file2: %w", err2))
return errors.Join(errs...)
```

Returns `nil` if the slice is empty — clean runs return `nil` naturally. Idiomatic for batch operations where you want to process all items and report all failures rather than stopping at the first error.

## Where to exit

Only `main` calls `os.Exit`. Library functions return errors. The Python version of preflight-sync calls `sys.exit(1)` in seven places — the Go version surfaces all of those as returned errors and lets `main` decide what to do.

## Sentinel errors vs custom error types

**Sentinel errors** — simple named error values, checked with `errors.Is`:

```go
var ErrDiskFull = errors.New("disk full")

if errors.Is(err, ErrDiskFull) {
    // err is ErrDiskFull, or wraps it somewhere in the chain
}
```

Use when the caller only needs to know *which kind* of error occurred.

**Custom error types** — structs with an `Error() string` method, checked with `errors.As`:

```go
type DiskFullError struct {
    Filename string
    Err      error
}

func (e *DiskFullError) Error() string {
    return fmt.Sprintf("destination disk full while copying %s", e.Filename)
}

func (e *DiskFullError) Unwrap() error { return e.Err }
```

Use when the caller needs structured data from the error:

```go
var diskFull *models.DiskFullError
if errors.As(err, &diskFull) {
    fmt.Println(diskFull.Filename)
}
```

`errors.As` unwraps the error chain automatically via `Unwrap()` — no manual unwrapping needed.

## The error interface

The built-in `error` interface has exactly one method:

```go
type error interface {
    Error() string
}
```

Any type with `Error() string` is automatically an `error` — no declaration needed. This is Go's implicit interface system in action. The `Unwrap() error` method is not part of a formal interface — the `errors` package discovers it by convention using reflection.

C# analogy: like implementing `IError` without ever writing `implements`. Any type can satisfy an interface retroactively — more flexible than C#'s explicit declarations.

## errors.Is vs errors.As

- `errors.Is(err, target)` — checks if `err` *is* a specific value anywhere in the chain. Used with sentinel errors.
- `errors.As(err, &target)` — checks if `err` *is of a specific type* anywhere in the chain, and populates `target` with it. Used with custom error types.

Both walk the error chain via `Unwrap` automatically.

*Related: [[Go Defined Types]] (sentinel errors follow the same named-constant pattern), [[Go Structured Logging]] (errors are surfaced via `slog.Any("err", err)` in structured log output)*
