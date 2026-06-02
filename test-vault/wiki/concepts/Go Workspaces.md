---
title: "Go Workspaces"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-2026-04-08-go-workspaces.md"
related:
  - "preflight-sync-go"
tags:
  - go
  - modules
  - development
---

# Go Workspaces

Go workspaces (Go 1.18+) allow working with multiple modules simultaneously without editing `go.mod` files. A workspace-local `go.work` file lists the modules in scope; the Go toolchain treats each as a root module when resolving dependencies.

## The Problem They Solve

Before workspaces, developing against a local version of a dependency required a `replace` directive in `go.mod`:

```go
replace golang.org/x/example => ../example
```

This directive had to be removed before committing — a manual step that was easy to forget. Workspaces eliminate it: `go.work` is workspace-local state (typically gitignored), not project state.

## Setup

```bash
# Create workspace at the parent directory level
go work init

# Add modules to the workspace
go work use ./hello
go work use ./example
```

The resulting `go.work` file:

```
go 1.18

use (
  ./example
  ./hello
)
```

With both modules in scope, `hello`'s import of `golang.org/x/example` resolves to `./example` locally instead of the downloaded version. No `go.mod` changes needed.

## Key Properties

- `go.work` is workspace-local — add to `.gitignore`
- Each `use` entry is treated as a root module for dependency resolution
- Works with all standard toolchain commands: `go run`, `go build`, `go test`
- `go work sync` updates `go.sum` files for all modules in the workspace

## When to Use

- Developing a library and a consumer simultaneously
- Testing a proposed upstream change before filing a PR
- Monorepo-style development where modules share a parent directory but stay separate modules

## See Also

- [[preflight-sync-go]] — Go project where this pattern applies when working with shared packages
- [[Go File IO]], [[Go Error Handling]] — other Go concept pages
