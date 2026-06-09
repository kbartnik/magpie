---
tags: [concept, go]
cluster: go
aliases: ["go modules", "go packages", "MVS", "go.mod", "module versioning"]
related: ["Go Tooling", "Go Interfaces", "Go Generics Type Constraints"]
sources:
  - "[[archive/books/2026-06-04-learning-go-2e]]"
---

# Go Modules and Packages

## Modules vs Packages

A **module** is a collection of packages with a shared version (defined by `go.mod`). A **package** is a directory of `.go` files that compile together. Every Go file belongs to exactly one package; packages belong to one module.

## MVS — Minimum Version Selection

Go's dependency resolution algorithm: when multiple modules require a dependency at different versions, Go selects the **minimum version** that satisfies all requirements. This is deliberately conservative — no automatic upgrades to latest, no semver ranges. Builds are reproducible.

## go.mod Structure

```
module github.com/example/mymod

go 1.22

require (
    github.com/some/dep v1.3.2
)
```

`go.sum` records the expected cryptographic hashes of all dependency versions — ensures builds use the exact bytes they were tested with.

## Workspaces

`go.work` enables multi-module development without `replace` directives. Useful when developing a library and its consumer simultaneously.

## Internal Packages

A package path containing `/internal/` can only be imported by code rooted at the parent of `internal/`. Enforces module-internal encapsulation.

## Connections

- [[Go Tooling]] — `go mod tidy`, `go mod vendor`, `govulncheck` operate on the module graph
- [[Go Generics Type Constraints]] — the `golang.org/x/exp/constraints` package illustrates how experimental APIs graduate through the module system
