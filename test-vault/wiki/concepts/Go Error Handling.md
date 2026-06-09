---
tags: [concept, go]
cluster: go
aliases: ["go errors", "error wrapping", "sentinel errors", "errors.Is", "errors.As", "go error handling"]
related: ["Go Interfaces", "Go Testing Patterns", "Go Context Patterns", "Go HTTP Client Patterns"]
sources:
  - "[[archive/books/2026-06-04-go-programming-language]]"
  - "[[archive/books/2026-06-04-learning-go-2e]]"
---

# Go Error Handling

Go errors are values — the `error` interface with a single `Error() string` method. No exceptions; errors propagate explicitly as return values.

## The Idiom

```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doSomething: %w", err)
}
```

`%w` wraps the original error, preserving it for inspection.

## Sentinel Errors

Named error values for comparison:
```go
var ErrNotFound = errors.New("not found")

if errors.Is(err, ErrNotFound) { ... }
```

`errors.Is` unwraps the error chain to find a matching target.

## Custom Error Types

```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string { return e.Message }

var ve *ValidationError
if errors.As(err, &ve) { ... }
```

`errors.As` finds the first error in the chain that matches the target type.

## Defer for Cleanup

```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil { return err }
    defer f.Close()
    // ...
}
```

`defer` guarantees cleanup even if the function returns early due to error.

## Connections

- [[Go Interfaces]] — `error` is an interface; the entire error system is built on interface satisfaction
- [[Go Context Patterns]] — `context.Canceled` and `context.DeadlineExceeded` are sentinel errors checked with `errors.Is`
- [[Go Testing Patterns]] — table-driven tests typically include error cases; `testify/assert.ErrorIs` wraps `errors.Is`
