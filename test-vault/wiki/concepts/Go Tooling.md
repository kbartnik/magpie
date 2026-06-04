---
type: concept
tags: [go, tooling, quality]
---

# Go Tooling

Go ships with a strong standard toolchain, supplemented by a small set of widely-adopted third-party tools.

**Source:** [[archive/books/2026-06-04-learning-go-2e|Learning Go, 2nd Edition]] Ch. 11

---

## Static Analysis

**`go vet`** — built-in; catches provably incorrect code (mismatched `Printf` verbs, unreachable code, suspicious composite literals). Always runs before `go test`.

**`staticcheck`** — the essential third-party linter. Catches a much broader class of bugs than `go vet`: deprecated API usage, impossible conditions, redundant type conversions, performance issues. Treat its findings as errors.

**`golangci-lint`** — meta-linter that runs staticcheck and many others in one pass. Configurable via `.golangci.yml`. The standard choice for CI pipelines — runs fast via caching and parallelism.

**`govulncheck`** — scans the module graph for known CVEs, and crucially, only reports vulnerabilities in code paths that are actually called. Avoids the false-positive flood of dep-tree scanners.

## Build Tags

Build tags constrain which files are included in a build.

```go
//go:build linux && amd64
```

Placed at the top of a file (before `package`). Can combine OS, architecture, and custom tags. Used for platform-specific implementations, test-only code, and feature flags in build pipelines.

```bash
go build -tags integration ./...   # include files tagged "integration"
```

## Code Generation (`go generate`)

`go generate` runs arbitrary commands embedded in source files:

```go
//go:generate stringer -type=Direction
```

Run with `go generate ./...`. Conventionally used for: generating interface mocks, string methods for enums (`stringer`), protobuf/gRPC bindings, and embedding assets. Generated files should be committed; they're part of the build.

## Embedding (`//go:embed`)

Embed files and directories into the binary at compile time:

```go
//go:embed static/*
var staticFiles embed.FS
```

Supports single files, directory trees, and glob patterns. Embedded content is read-only at runtime.

## Testing Tools

| Tool | Purpose |
|------|---------|
| `go test -race` | Data race detector — runs instrumented binary; ~5-20x slower, use in CI |
| `go test -fuzz` | Fuzzing — mutates inputs to find crashes; Go 1.18+ |
| `go test -bench` | Microbenchmarks — `BenchmarkFoo(b *testing.B)` functions |
| `go test -cover` | Coverage report |
| `httptest.NewServer` | Spins up a real HTTP server for handler tests |

## Formatting

`gofmt` and `goimports` (adds/removes imports) are non-negotiable. Most editors run them on save. CI should fail on unformatted files (`gofmt -l . | grep .`).

---

## Related

- [[Go Modules and Packages]] — module system, `go mod tidy`, workspaces
- [[Go Testing Patterns]] — table tests, integration tests, fuzzing
