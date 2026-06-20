---
tags: [concept, go]
cluster: go
aliases: ["golang interfaces", "implicit interfaces", "duck typing Go", "interface satisfaction"]
related: ["Go Modules and Packages", "Go Error Handling", "Go Generics Type Constraints", "Go Functional Options Pattern"]
sources:
  - "[[archive/books/2026-06-04-learning-go-2e]]"
  - "[[archive/books/2026-06-04-go-programming-language]]"
---

# Go Interfaces

Go interfaces are satisfied implicitly — no `implements` keyword. Any type that has the required methods satisfies the interface automatically.

## Key Rules

**Accept interfaces, return structs.** Function parameters should be interface types (accept the widest useful type); return values should be concrete types (give callers the most information).

**Nil interface gotcha.** A nil pointer wrapped in an interface is not nil. `var p *MyType = nil; var i MyInterface = p` — `i != nil` is true even though the underlying value is nil. Check interface nilness by checking the concrete value.

**Interface embedding.** Interfaces can embed other interfaces: `io.ReadWriter` embeds `io.Reader` and `io.Writer`. This composes behaviors without inheritance.

**The empty interface (`any`).** `interface{}` / `any` accepts all types but loses type safety. Prefer concrete types or constrained generics.

## Type Assertions and Switches

```go
// Type assertion — panics if wrong type
s := i.(string)

// Safe type assertion
s, ok := i.(string)

// Type switch
switch v := i.(type) {
case string: ...
case int: ...
}
```

## Connections

- [[Go Generics Type Constraints]] — generics and interfaces solve adjacent problems; the decision point: heterogeneous collections → interfaces; homogeneous collections → generics
- [[Go Error Handling]] — `error` is an interface; sentinel errors, custom error types, and errors.As all work through interface satisfaction
- [[Go Functional Options Pattern]] — option functions operate on concrete struct types, not interfaces; the interplay illustrates when to use each
