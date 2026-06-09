---
tags: [concept, go]
cluster: go
aliases: ["go tools", "golangci-lint", "staticcheck", "govulncheck", "go generate"]
related: ["Go Modules and Packages", "Go Testing Patterns", "Go Memory Model"]
sources:
  - "[[archive/books/2026-06-04-learning-go-2e]]"
---

# Go Tooling

## Static Analysis

**staticcheck** — the most comprehensive Go linter. Catches bugs, performance issues, and style problems the compiler misses. Run before CI.

**golangci-lint** — meta-linter that runs multiple linters (including staticcheck) in parallel with a single config file.

**govulncheck** — checks your module's dependencies against the Go vulnerability database. Essential for security-conscious workflows.

## Build Tags

```go
//go:build linux && amd64
```

Conditionally include files based on OS, architecture, or custom tags. Used for platform-specific implementations and test utilities.

## go:generate

`//go:generate <command>` in source files, run with `go generate ./...`. Drives code generation: mockery, stringer, protobuf, etc. Not run automatically by build — must be triggered explicitly.

## Escape Analysis

`go build -gcflags="-m"` prints escape analysis decisions. Shows which variables escape to heap. Used to identify unexpected allocations in hot paths.

## Cross-Compilation

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux ./cmd/myapp
```

Go cross-compiles to any supported target without external toolchains.

## Connections

- [[Go Modules and Packages]] — tooling operates on the module graph; `go mod tidy` keeps go.mod consistent with imports
- [[Go Memory Model]] — `-gcflags="-m"` reveals escape decisions grounded in the memory model
- [[Go Testing Patterns]] — golangci-lint includes test-specific linters; govulncheck checks test dependencies too
