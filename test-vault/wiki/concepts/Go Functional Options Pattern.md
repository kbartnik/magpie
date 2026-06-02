---
title: "Go Functional Options Pattern"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-advanced-concepts.md"
related:
  - "[[Go Context Patterns]]"
  - "[[Go Error Handling]]"
tags:
  - go
  - patterns
  - api-design
---

# Go Functional Options Pattern

Popularized by [[Dave Cheney]]. Solves the problem of growing constructor parameters while keeping the API backward-compatible and self-documenting.

## The Pattern

```go
type Client struct {
    addr    string
    timeout time.Duration
    retries int
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
    return func(c *Client) {
        c.timeout = d
    }
}

func WithRetries(n int) Option {
    return func(c *Client) {
        c.retries = n
    }
}

func NewClient(addr string, opts ...Option) *Client {
    c := &Client{
        addr:    addr,
        timeout: 30 * time.Second, // sensible defaults
        retries: 3,
    }
    for _, opt := range opts {
        opt(c)
    }
    return c
}
```

Call sites are readable and explicit:

```go
client := NewClient(
    "grpc.service.internal:50051",
    WithTimeout(5*time.Second),
    WithRetries(5),
)
```

## Advantages Over a Config Struct

| | Config struct | Functional options |
|-|---------------|--------------------|
| Defaults | Scattered at call sites or in a `DefaultConfig()` | Live in `NewClient` |
| Adding options | May break callers using struct literals | Always backward-compatible |
| Self-documentation | Field names only | Function names are expressive |
| Validation | Must check after construction | Can validate inside each `Option` |

## Error-Returning Variant

When an option involves I/O or validation, return an error:

```go
type Option func(*Client) error

func WithTLSCert(certPath string) Option {
    return func(c *Client) error {
        cert, err := tls.LoadX509KeyPair(certPath, certPath)
        if err != nil {
            return fmt.Errorf("loading TLS cert: %w", err)
        }
        c.tls = &tls.Config{Certificates: []tls.Certificate{cert}}
        return nil
    }
}

// NewClient collects errors:
func NewClient(addr string, opts ...Option) (*Client, error) {
    c := &Client{addr: addr, timeout: 30 * time.Second}
    for _, opt := range opts {
        if err := opt(c); err != nil {
            return nil, err
        }
    }
    return c, nil
}
```

## See Also

- [[Go Error Handling]] — wrapping errors with `%w` inside option functions
- [[Go Context Patterns]] — context is the other main mechanism for threading request state; options are for construction-time config
