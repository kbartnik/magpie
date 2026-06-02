---
title: "Go Modules and Packages"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - archive/books/2026-05-31-learning-go-2e.md
related:
  - Go Workspaces
  - Go Tooling
tags:
  - go
  - modules
  - tooling
---

# Go Modules and Packages

## Module Path as Namespace

The module path (e.g. `github.com/kbartnik/magpie`) is an **import namespace**, not a live URL. It controls how packages within the module import each other and how external consumers reference the module. It can be set before the repository exists on the hosting service — the path just needs to match once you push. Choosing it is a one-time, permanent decision: changing it later breaks all importers.

## Hierarchy

- **Repository** — version control unit (one repo, typically one module)
- **Module** — unit of versioning and distribution, defined by `go.mod`
- **Package** — unit of compilation, a directory of `.go` files with a shared `package` declaration

## go.mod

```
module github.com/user/project

go 1.22

require (
    github.com/some/dep v1.2.3
)
```

- The `go` directive specifies the *minimum* Go version required to build the module. Since Go 1.21 this is a hard minimum, not advisory.
- The `require` directive pins direct and indirect dependencies.

## Minimal Version Selection (MVS)

Go's dependency resolution deliberately picks the *minimum* version that satisfies all requirements, not the latest. This means builds are reproducible and predictable — upgrading a dependency requires an explicit change to `go.mod`, not an automatic pull.

## Versioning Rules

- `v0.x.x` and `v1.x.x`: no import path change
- `v2+`: import path must include the major version: `github.com/user/project/v2`
- This makes major version upgrades an explicit code change, not just a manifest change

## Workspaces

Go workspaces (`go work`) allow simultaneous development across multiple modules without replacing `require` directives. Useful when modifying a library and its consumer in tandem:

```
go work init ./mylib ./myapp
```

The `go.work` file is local and should not be committed. See Go Workspaces.

## Key Commands

| Command | Purpose |
|---------|---------|
| `go mod tidy` | Add missing, remove unused dependencies |
| `go get pkg@version` | Add/update a dependency |
| `go mod vendor` | Copy dependencies into `vendor/` |
| `go list -m all` | Show all module dependencies |

## Organizing a Module

- One package per directory
- Package name ≠ directory name is allowed but confusing — avoid
- `internal/` packages are importable only within the parent module
- Circular dependencies are a compile error — use interfaces to break cycles
- Avoid `init()` functions; they make initialization order implicit and hard to test

## Module Proxy

By default, Go fetches modules through `proxy.golang.org`, which caches versions and provides a checksum database. Private repositories require configuring `GONOSUMCHECK` and `GONOSUMDB`.

## See Also

- Go Workspaces — multi-module development
- Go Tooling — build tags, go generate, vulnerability scanning
