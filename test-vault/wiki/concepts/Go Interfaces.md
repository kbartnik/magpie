---
type: concept
tags: [go, types, design]
---

# Go Interfaces

Go interfaces are **implicitly satisfied** — a type implements an interface simply by having the required methods. No `implements` declaration. This is type-safe duck typing: the compiler checks the contract at compile time without requiring explicit coupling between the implementing type and the interface definition.

**Source:** [[archive/books/2026-06-04-learning-go-2e|Learning Go, 2nd Edition]] Ch. 7

---

## Core Rules

**Accept interfaces, return structs.** Function parameters should be interfaces (maximum flexibility for callers); return types should be concrete (callers get a fully specified value with no surprises). This rule emerges directly from the implicit satisfaction model — requiring an interface as a parameter means any type with those methods works.

**Implicit satisfaction means zero coupling.** The type and the interface can live in completely separate packages and never reference each other. The interface can be defined by the consumer (where it's used), not the producer (where the type is defined). This is the inverse of most OO languages.

## The nil Interface Gotcha

An interface value has two parts: the **type** and the **value**. An interface is only `nil` when both are nil.

```go
var p *MyType = nil       // p is nil
var i MyInterface = p     // i is NOT nil — it has a type (*MyType) with a nil value
fmt.Println(i == nil)     // false — this surprises almost everyone
```

Returning a concrete `nil` pointer as an interface causes this. Return the interface type directly when the function signature returns an interface.

## Interface Composition

Interfaces compose by embedding:

```go
type ReadWriter interface {
    Reader  // embeds io.Reader
    Writer  // embeds io.Writer
}
```

The Go standard library is built on small, composable interfaces (`io.Reader`, `io.Writer`, `io.Closer`) rather than large ones.

## Type Assertions and Type Switches

```go
// Assert a specific type (panics if wrong; use two-value form)
v, ok := i.(MyType)

// Switch over possible types
switch v := i.(type) {
case *MyType:   // v is *MyType here
case int:       // v is int here
default:        // v is the original interface type
}
```

---

## Related

- [[Go Error Handling]] — `error` is an interface; the nil gotcha applies
- [[Go Generics Type Constraints]] — type constraints are interface syntax
