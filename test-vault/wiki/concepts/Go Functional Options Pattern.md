---
tags: [concept, go]
cluster: go
aliases: ["functional options", "option functions", "Dave Cheney options", "go configuration pattern"]
related: ["Go Interfaces", "Go Package Design", "Go Error Handling"]
sources:
  - "[[archive/videos/2026-06-04-dave-cheney-practical-go]]"
---

# Go Functional Options Pattern

Backward-compatible, self-documenting configuration for APIs with optional parameters. Originated by Dave Cheney.

## Pattern

```go
type Server struct {
    addr    string
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func WithMaxConnections(n int) Option {
    return func(s *Server) { s.maxConn = n }
}

func NewServer(addr string, opts ...Option) (*Server, error) {
    s := &Server{
        addr:    addr,
        timeout: 30 * time.Second, // sensible default
        maxConn: 100,
    }
    for _, o := range opts { o(s) }
    return s, nil
}
```

## Why It Works

- **Zero-value default:** `NewServer("addr")` works without any options
- **Backward compatible:** New options don't break existing callers
- **Self-documenting:** Option names appear at the call site: `NewServer("addr", WithTimeout(5*time.Second))`
- **Testable:** Options are functions — easily overridden or inspected in tests

## Error-Returning Variant

When option validation is needed:
```go
type Option func(*Server) error
```

## Connections

- [[Go Interfaces]] — option functions operate on concrete types, not interfaces; complement to the interface pattern
- [[Go Package Design]] — functional options are the canonical solution to API backward-compatibility; every exported option is a stable contract
- [[Go Error Handling]] — the error-returning option variant propagates validation failures through the constructor
