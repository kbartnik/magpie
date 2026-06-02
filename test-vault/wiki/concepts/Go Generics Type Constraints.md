---
title: "Go Generics Type Constraints"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-go-type-parameters.md"
  - "archive/clippings/2026-05-31-go-type-inference.md"
related:
  - "Go Defined Types"
tags: [go, generics, types, type-inference]
---

# Go Generics Type Constraints

How Go generics constraints work — the `~` tilde syntax, type parameter deconstruction, and how the compiler resolves type arguments via inference.

## The ~ Tilde: Underlying Type Constraints

```go
// Only matches the exact literal type []string — MySlice would fail
func Clone[S []E, E any](s S) S   // BAD for named slice types

// Matches any type whose underlying type is []E — MySlice passes
func Clone[S ~[]E, E any](s S) S  // CORRECT (this is slices.Clone)
```

`~T` means "any type whose underlying type is T." Without `~`, named types like `type MySlice []string` do not satisfy the constraint `[]string` even though they are backed by it.

**Why `~` is explicit:** For predeclared types like `int`, the type is its own underlying type. `[T ~int]` means "any type backed by int" (including `type MyAge int`). Without the `~`, `[T int]` would mean only the exact type `int`. Making `~` explicit keeps the two meanings unambiguous.

**Rule:** `~T` is only valid when `T` is a type literal or predeclared type — never a named defined type. `[T ~MySlice]` is a compile error because `MySlice` is not the underlying type of anything (its own underlying type is `[]string`).

## Type Parameter Deconstruction

Break a composite type into multiple parameters to name and constrain each component:

```go
// Slices
func Clone[S ~[]E, E any](s S) S

// Maps
func Clone[M ~map[K]V, K comparable, V any](m M) M

// Constrain the element type
func WithStrings[S ~[]E, E interface{ String() string }](s S) (S, []string)
```

This pattern lets you:
- Preserve the named slice/map type in the return value (callers get back `MySlice`, not `[]string`)
- Place independent constraints on each component (`K comparable`, `E cmp.Ordered`)

## Type Inference

Type inference lets callers omit type arguments the compiler can deduce:

```go
type List []int
var list List

// Without inference:
slices.Sort[List, int](list)

// With inference — looks like non-generic code:
slices.Sort(list)
```

The compiler infers `S = List` from the argument, then infers `E = int` from the constraint `S ~[]E` (since `under(List) = []int`).

### How It Works: Type Equations

Type inference is constraint solving. The compiler collects type equations from three sources:

| Source | Equation form |
|--------|--------------|
| Explicit type arguments | `P ≡ A` (identity) |
| Function arguments | `T(param) :≡ T(arg)` (assignability) |
| Constraints | `P ∈ constraint` (satisfaction) |

It then solves the equations by **unification** — recursive structural matching:

```
// Given: func Clone[S ~[]E, E any](s S) S
// Call:  Clone(myList)  where myList is List

Equations:
  S :≡ List         // from argument
  S ∈ ~[]E          // from constraint
  E ∈ any           // from constraint

Solving:
  S ➞ List          // from assignability equation
  under(S) = []int ≡ []E  // from constraint, so:
  E ➞ int
```

### Unification

Unification walks two type trees recursively. When it finds a bound type parameter (`S`, `E`) against a concrete type, it records the inference. If the same parameter appears twice and different types are inferred, unification fails.

```
// Solves A ➞ string, B ➞ byte, C ➞ int
X: map[A]struct{i int; s []B}
Y: map[string]struct{i C; s []byte}
```

### Go 1.21: Cross-Function Inference

Since Go 1.21, inference works when passing uninstantiated generic functions as arguments:

```go
func myEq[P comparable](x, y P) bool { return x == y }

// Compiler infers both CompactFunc's S/E and myEq's P:
result := slices.CompactFunc(list, myEq)
// same as: slices.CompactFunc[List, int](list, myEq[int])
```

### Self-Recursive Generics

For a self-recursive generic function like `fact[P](n P)`, the compiler renames type parameters in the recursive call's signature before inference, turning the circular equation `P :≡ P` into `Q :≡ P` (solvable).

### Untyped Constants

Typed arguments take precedence over untyped constants for inference. When only untyped numeric constants are present, Go 1.21 picks the "later" default type in the ordering `int → rune → float64 → complex128`:

```go
foo(1, 2.0)  // Go 1.21: P ➞ float64 (not an error as in Go 1.20)
```

## See Also

- [[Go Defined Types]] — underlying types, named type semantics, why `~` is needed
- [[Go Escape Analysis]] — type parameters don't change escape analysis; generic instantiation inlines the type
