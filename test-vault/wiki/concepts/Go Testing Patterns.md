---
tags: [concept, go]
cluster: go
aliases: ["go tests", "table-driven tests", "testify", "go testing", "test helpers"]
related: ["Go Error Handling", "Go Interfaces", "Go Tooling"]
sources:
  - "[[archive/books/2026-06-04-go-programming-language]]"
  - "[[archive/books/2026-06-04-learning-go-2e]]"
---

# Go Testing Patterns

## Table-Driven Tests

The idiomatic Go testing pattern — define test cases as a slice of structs, iterate and run each:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            assert.Equal(t, tt.expected, got)
        })
    }
}
```

## testify

`github.com/stretchr/testify` adds `assert` and `require`:
- `assert.*` — logs failure and continues
- `require.*` — logs failure and stops (for fatal preconditions)

## Dependency Injection via Functions

Inject dependencies as function values rather than mocking frameworks:

```go
type Fetcher func(ctx context.Context, url string) ([]byte, error)

func Process(ctx context.Context, fetch Fetcher) error { ... }

// In tests:
Process(ctx, func(ctx context.Context, url string) ([]byte, error) {
    return []byte(`{"ok":true}`), nil
})
```

## Test Coverage

`go test -cover ./...` — reports line coverage. `go test -coverprofile=coverage.out && go tool cover -html=coverage.out` — visual report.

## Connections

- [[Go Error Handling]] — table-driven tests typically include error cases; `assert.ErrorIs` wraps `errors.Is`
- [[Go Interfaces]] — interfaces enable DI patterns that make code testable without mocking frameworks
- [[Go Tooling]] — golangci-lint includes test linters; `go test -race` detects data races
