---
title: "Go Tooling"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - archive/books/2026-05-31-learning-go-2e.md
related:
  - Go Modules and Packages
  - Go Testing Patterns
tags:
  - go
  - tooling
  - quality
---

# Go Tooling

## Core Commands

| Command | Purpose |
|---------|---------|
| `go build` | Compile packages and dependencies |
| `go run` | Compile and run (prototyping / small programs only) |
| `go fmt` | Format source code (non-negotiable; always run) |
| `go vet` | Report likely mistakes (composite literals, printf format errors, etc.) |
| `go test` | Run tests, benchmarks, fuzz targets |
| `go install` | Build and install a binary to `$GOPATH/bin` |
| `go generate` | Run code generators declared in `//go:generate` comments |

## Static Analysis

**`staticcheck`** (`honnef.co/go/tools`) — the most comprehensive linter. Catches unused code, incorrect API usage, performance anti-patterns.

**`revive`** — configurable, drop-in replacement for `golint` (which is deprecated).

**`golangci-lint`** — meta-linter that runs multiple linters in one pass. The standard choice for CI pipelines. Configure via `.golangci.yml`.

**`govulncheck`** — scans module graph for known vulnerabilities in the Go vulnerability database. Only reports vulnerabilities in code that is actually called (not just imported).

## go vet vs Linters

`go vet` ships with Go and catches definite bugs. Linters like `staticcheck` go further, catching style issues, deprecated API use, and suspicious patterns that aren't necessarily bugs. Both should run in CI.

## Build Tags

Control what gets compiled using build constraints:

```go
//go:build linux && amd64
```

Common uses: platform-specific code, integration tests, debug builds.

```go
//go:build integration
```

Run with: `go test -tags integration ./...`

## go:generate

Declare code generators inline:

```go
//go:generate stringer -type=Direction
```

Run with `go generate ./...`. Common uses: generating string methods for enums (`stringer`), mock generation, protobuf compilation.

## Embedding Files

Embed static assets directly into the binary:

```go
//go:embed templates/*
var templateFS embed.FS
```

Files are included at compile time; no runtime filesystem dependency. The `embed.FS` type implements `fs.FS`.

## Cross-Compilation

Go cross-compiles cleanly:

```bash
GOOS=linux GOARCH=arm64 go build ./...
```

No additional toolchain needed (unlike CGo, which requires the target's C compiler).

## Reading Binary Build Info

```go
info, ok := debug.ReadBuildInfo()
```

Available at runtime; includes module versions, VCS commit, build settings. Useful for `--version` flags.

## See Also

- [[Go Modules and Packages]] — module system, go.mod, MVS
- [[Go Testing Patterns]] — go test, benchmarks, fuzzing, race detector
