---
type: concept
tags: [go, tooling, dependency-management]
---

# Go Modules and Packages

Go modules are the unit of versioning and distribution. A module is a tree of packages rooted at a `go.mod` file.

**Source:** [[archive/books/2026-06-04-learning-go-2e|Learning Go, 2nd Edition]] Ch. 10

---

## Module System

**`go.mod`** declares the module path, Go version, and dependencies. The module path is both the import root and the identifier used by the package registry.

```
module github.com/user/mymodule

go 1.22

require (
    github.com/some/dep v1.2.3
)
```

**Minimal Version Selection (MVS).** When two dependencies require different versions of a third, Go picks the *minimum* version that satisfies all requirements — not the latest. This makes builds reproducible without a lockfile being strictly necessary (though `go.sum` records cryptographic hashes for verification).

**Semantic versioning.** Major version changes (`v2+`) require a different import path suffix (`github.com/user/mymodule/v2`). This allows v1 and v2 to coexist in the same build — different packages, different imports.

## Packages

A package is a directory of `.go` files sharing the same `package` declaration. The package name is the last segment of the import path by convention (but not required).

**Internal packages.** A package path containing `/internal/` can only be imported by code within the parent of `internal/`. Enforces API boundaries without making types unexported.

**`init()` functions.** Each file can have an `init()` function that runs once at program startup, after all variable initializations. Multiple `init()` functions per package are allowed; order within a package is by source file name. Avoid side-effect-heavy `init()` — prefer explicit initialization.

## Workspaces (`go.work`)

Go workspaces (`go work`) let multiple modules be developed together locally without publishing. The `go.work` file redirects imports for specified modules to local paths.

```
go 1.22

use (
    ./mymodule
    ./mylib
)
```

Useful for: developing a library and its consumer simultaneously; testing a fork before it's published.

## Key Commands

| Command | Purpose |
|---------|---------|
| `go mod tidy` | Remove unused deps, add missing ones |
| `go mod vendor` | Copy deps into `vendor/` for offline builds |
| `go get pkg@version` | Add or upgrade a dependency |
| `go work init` | Create a workspace |
| `go list -m all` | List all modules in the build |

---

## Related

- [[Go Tooling]] — linting, vulnerability scanning, build tags
