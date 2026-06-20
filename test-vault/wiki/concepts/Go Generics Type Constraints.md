---
tags: [concept, go]
cluster: go
aliases: ["go generics", "type parameters", "type constraints", "tilde operator", "underlying type constraints", "go 1.18"]
related: ["Go Interfaces", "Go Modules and Packages", "Go Tooling"]
sources:
  - "[[archive/videos/2026-06-04-go-generics-gophercon]]"
---

# Go Generics Type Constraints

Type parameters allow writing functions and types that work across multiple types, with compiler-checked type safety at each instantiation.

## Basic Syntax

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b { return a }
    return b
}

min := Min(3, 5)      // T inferred as int
min := Min(3.0, 5.0)  // T inferred as float64
```

## The `~` Underlying Type Operator

```go
type Celsius float64

// Without ~: only matches float64, not Celsius
// With ~: matches float64 AND any type built on float64
type Temperature interface { ~float64 }
```

`~T` means "any type whose underlying type is T." Critical for Go's defined-type idiom — without `~`, generic libraries would force users to abandon defined types.

## When to Use Generics vs Interfaces

| Use generics when... | Use interfaces when... |
|---|---|
| Homogeneous collection (all same type) | Heterogeneous collection (different types) |
| Uniform algorithm, types vary | Behavior varies by implementation |
| Avoiding code duplication | Abstracting over behavior |

## Type Inference

The compiler infers type arguments at call sites in most cases — generic code looks like non-generic code to callers.

## Connections

- [[Go Interfaces]] — generics and interfaces are complementary; mixing them is often wrong; the decision point is heterogeneous vs homogeneous
- [[Go Modules and Packages]] — `golang.org/x/exp/constraints` contains experimental constraints; module versioning controls promotion to stdlib
- [[Go Tooling]] — GC shape stenciling (not full monomorphization) is the current compilation strategy; affects profiling of generic code
