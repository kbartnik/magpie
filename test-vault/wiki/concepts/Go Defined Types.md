---
title: "Go Defined Types"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Bubbletea Elm Architecture", "Go Error Handling"]
tags: [go, types, preflight-sync-go]
---

# Go Defined Types

A **defined type** gives a new name to an underlying type, creating a distinct type that the compiler treats as separate — even though the representation is identical.

```go
type PathStyle string

const (
    PathStyleSeason PathStyle = "season"
    PathStyleTitle  PathStyle = "title"
)
```

`PathStyle` and `string` have the same memory layout, but the compiler rejects assignments between them without an explicit cast. This is stronger than a type alias (`type PathStyle = string`), which creates no distinction at all.

## Why this matters

- **Invalid values are caught at compile time**, not at runtime. A function accepting `PathStyle` rejects a plain `string` without a cast — typos and missing constants become build errors.
- **Map keys typed as the defined type refuse plain strings**, making the dispatch table self-documenting and the compiler your validation layer:

```go
var pathBuilders = map[PathStyle]pathBuilder{
    PathStyleSeason: buildSeasonPath,
    PathStyleTitle:  buildTitlePath,
}
```

Adding a new path style means adding one constant and one map entry. `BuildDestPath` never needs to change.

- **Defined types communicate intent at call sites.** A function parameter typed as `AppMode` tells the reader exactly what kind of value is expected — not just any string.

## Go's approach vs enums

Go has no `enum` keyword. The defined-type-plus-constants pattern is the idiomatic substitute. The underlying type stays `string` (or `int`, or whatever fits), which means:
- TOML/JSON decoders can populate the field directly without custom unmarshaling
- The value is human-readable in logs and config files
- You validate the value in code after decoding (the decoder doesn't know your constants)

## In this codebase

`PathStyle`, `AppMode`, `CopyState`, and `FileStatus` all follow this pattern. They live in `internal/models` or alongside the domain they describe so both the producing and consuming packages can reference them without an import cycle.

*Related: [[Bubbletea Elm Architecture]] (AppMode, AppModel), [[Go Error Handling]] (sentinel errors follow the same constant pattern), [[Go Generics Type Constraints]] (the `~T` underlying-type constraint uses the same underlying-type concept)*
