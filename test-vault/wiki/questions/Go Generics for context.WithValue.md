---
title: "Go Generics for context.WithValue"
question: "Can Go generics provide a type-safe replacement for context.WithValue's untyped key interface — and why hasn't the Go team shipped one?"
type: question
status: open
domain: go
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/books/2026-05-31-learning-go-2e.md"
related:
  - "[[Go Context Patterns]]"
  - "[[Go Knowledge Map]]"
tags:
  - go
  - generics
---

# Go Generics for context.WithValue

*Can Go generics provide a type-safe replacement for context.WithValue's untyped key interface — and why hasn't the Go team shipped one?*

The canonical Go pattern for `context.WithValue` uses unexported key types to prevent collisions between packages. This works but is verbose and relies on convention rather than type safety — callers can still pass the wrong unexported key type and get a silent nil at extraction. A generic `Context[T]` type would provide compile-time type safety for context values.

Community proposals exist (`context/v2` discussions, third-party typed context libraries). The Go team hasn't shipped a solution as of Go 1.22. The obstacle may be backward compatibility (existing code passes `context.Context` everywhere; a `Context[T]` would require new function signatures), interface proliferation (every function taking context would need type parameters), or the judgment that the current pattern is "good enough" and the added complexity isn't justified.

Whether generics actually help or introduce more complexity in practice is worth examining — typed context may push the type management complexity into function signatures rather than eliminating it.

## See Also

- [[Go Context Patterns]] — canonical usage; unexported key type pattern; context.Value for metadata only
