---
title: "Go Interfaces"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - archive/books/2026-05-31-learning-go-2e.md
related:
  - Go Defined Types
  - Go Error Handling
  - Go Generics Type Constraints
tags:
  - go
  - types
  - design
---

# Go Interfaces

## Implicit Implementation

Go interfaces are implemented *implicitly*: a concrete type satisfies an interface if its method set is a superset of the interface's method set. There is no `implements` declaration. This is type-safe duck typing — the compiler verifies compatibility at assignment, not at declaration.

```go
type Logic interface {
    Process(data string) string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
    // business logic
}

// No declaration needed — LogicProvider implicitly satisfies Logic
c := Client{L: LogicProvider{}}
```

## Accept Interfaces, Return Structs

The core Go API design rule: business logic should accept interfaces (flexible, testable) but return concrete types (extensible without breaking callers).

- **Accepting interfaces:** declares exactly what the function needs, enables swapping implementations, enables mocking in tests
- **Returning structs:** adding a new field or method to a returned struct is a minor (non-breaking) change; adding a method to a returned interface breaks all existing implementations

## Interfaces Are Defined by the Caller

Go's key insight over Java-style explicit interfaces: the *caller* defines the interface to express what it needs, not the implementer. This means:
- New implementations can satisfy existing interfaces without modifying the interface
- Small, focused interfaces (`io.Reader`, `io.Writer`) are composable via embedding
- Interfaces can be defined in the same package as the code that uses them

## Interfaces and nil

An interface value has two internal fields: a type pointer and a value pointer. An interface is `nil` only if *both* are nil. A common gotcha:

```go
var pointerCounter *Counter      // nil pointer
var incrementer Incrementer      // nil interface — both fields nil
incrementer = pointerCounter     // non-nil! type field is now set

fmt.Println(incrementer == nil)  // false — type is set even though value is nil
```

Invoking a method on a nil interface panics. Invoking a method on a non-nil interface with a nil concrete value may or may not panic depending on the method's nil handling.

## Interface Embedding

Interfaces can embed other interfaces, composing capabilities:

```go
type ReadCloser interface {
    Reader  // embeds io.Reader
    Closer  // embeds io.Closer
}
```

## Type Assertions and Type Switches

Use type assertions to access the concrete type behind an interface:

```go
val, ok := i.(ConcreteType)  // ok = false if assertion fails; no panic
```

Use type switches when handling multiple possible types. Use sparingly — if you find yourself asserting frequently, the abstraction may be wrong.

## See Also

- [[Go Defined Types]] — underlying type system
- [[Go Error Handling]] — `error` is the canonical small interface
- [[Go Generics Type Constraints]] — interfaces as type constraints
