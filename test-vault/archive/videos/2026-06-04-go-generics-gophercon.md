---
title: "Typing [Generic] Code in Go — GopherCon 2022"
type: video
captured-date: 2026-06-05
source-url: "https://www.youtube.com/watch?v=Pa_e9EeCdy8"
---

# Typing [Generic] Code in Go — GopherCon 2022

**Speaker:** Robert Griesemer
**Event:** GopherCon 2022
**Duration:** ~45 min

## Overview

Griesemer — one of the three original Go designers — explains the generics design in Go 1.18. The talk covers the problem generics solve, the design decisions made, and the constraints system that makes it work. This is the authoritative explanation of why Go generics look the way they do.

## The Problem Generics Solve

Before Go 1.18, writing a function that works on multiple types required either:
1. Duplicating the code for each type (error-prone, maintenance burden)
2. Using `interface{}` / `any` (loses type safety, requires runtime type assertions)
3. Generating code with `go generate` (works but is complex)

Generics allow writing a function once, with a type parameter, and having the compiler check type safety at instantiation time.

## Type Parameters

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b { return a }
    return b
}
```

`T` is a type parameter. The constraint `constraints.Ordered` specifies which types are valid: any type that supports `<`, `>`, `<=`, `>=`. The compiler checks at each call site that the provided type satisfies the constraint.

## Constraints and the `~` Operator

The key innovation is the `~T` syntax for **underlying type constraints**:

```go
type Celsius float64
type Fahrenheit float64

// ~float64 matches any type whose underlying type is float64
type Temperature interface {
    ~float64
}
```

Without `~`, a constraint `float64` only matches the `float64` type itself. With `~float64`, it also matches `Celsius`, `Fahrenheit`, and any other defined type built on `float64`. This is critical for Go's defined type idiom.

## When to Use Generics

Griesemer is explicit: generics are not always the answer. Use generics when:
- You're writing a data structure that must work for multiple element types (linked list, stack, set)
- You're writing an algorithm that works on any type satisfying a constraint (Min, Max, Map, Filter)
- You need to write the same code multiple times for different types and the duplication is the only reason

Do not use generics as a substitute for interfaces. Interfaces are the right tool when behavior varies by implementation; generics are the right tool when behavior is uniform but types vary.

## Type Inference

The compiler can infer type arguments in most cases:

```go
xs := []int{3, 1, 4}
min := Min(xs[0], xs[1])  // T inferred as int
```

Type inference makes generic code at the call site look like non-generic code, preserving Go's legibility.

## Deep Read

**Key Insight:** The `~T` underlying-type constraint is the conceptual center of Go generics. It preserves Go's defined-type idiom (where `type Celsius float64` and `float64` are different types with different semantics) while allowing generic algorithms to work across them. Without `~`, generics would force users to abandon defined types to use generic libraries.

**What Surprised Me:** Griesemer explicitly says that generics do not replace interfaces and that mixing them is often wrong. The rule of thumb: if you need a *heterogeneous* collection (different types in the same container), use interfaces. If you need a *homogeneous* collection (all elements of the same type, unknown at library-write time), use generics. Many early generic proposals got this wrong and produced awkward APIs.

**Open Questions:**
- Type inference works well for function calls but has limits for methods — you can't infer type parameters on method receivers. This limits how generic types can express behavior. Is this a fundamental constraint of the unification-based inference algorithm, or a limitation of the current implementation?
- The `constraints` package was initially added to `golang.org/x/exp` and the standard library absorbed only `comparable`. The rest (Ordered, Integer, Float) remain experimental. What's the bar for promoting a constraint to stdlib?
- Generic functions produce specialized machine code per instantiation (monomorphization) in theory, but the current Go compiler uses a different strategy (GC shape stenciling). What's the performance difference, and is full monomorphization on the roadmap?

**Wikilink Candidates:**
- [[Go Generics Type Constraints]] — type parameters; constraints interface; ~ underlying type operator; type inference via unification; when to use generics vs interfaces; not yet a wiki page

**Connections:**
- [[Go Interfaces]] — generics and interfaces solve adjacent problems; the decision point (heterogeneous vs homogeneous) is the key; interface embedding inside constraints is the bridge
- [[Go Modules and Packages]] — the `golang.org/x/exp/constraints` package situation illustrates how experimental APIs graduate to stdlib; module versioning is the mechanism
- [[Go Tooling]] — `-gcflags` can show instantiation decisions; the GC shape stenciling strategy affects profiling output for generic code
