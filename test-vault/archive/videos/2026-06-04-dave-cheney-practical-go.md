---
title: "Practical Go: Real World Advice for Writing Maintainable Go Programs"
type: video
captured-date: 2026-06-05
source-url: "https://www.youtube.com/watch?v=EXrEd2-b9BY"
---

# Practical Go: Real World Advice for Writing Maintainable Go Programs

**Speaker:** Dave Cheney
**Event:** GopherCon Singapore 2019
**Duration:** ~60 min

## Overview

A synthesis of Go idioms and design principles accumulated across years of production Go. Cheney frames the entire talk around a single axis: every design decision in Go code is a tradeoff between the cost of writing it and the cost of maintaining it. The talk covers identifier naming, package design, API design, error handling philosophy, and the functional options pattern.

## Identifier Naming

Names are the primary documentation mechanism. A name should describe *what* it does, not *how* it does it. Avoid stutter (`bytes.ByteBuffer`), avoid type names in variable names (`var userSlice []User` → `var users []User`), and keep short names for short-lived variables. The receiver name should be one or two letters derived from the type name — consistent across all methods on that type.

## Package Design

Packages should provide a single, focused capability. The package name is part of every identifier exported from it — `http.Client` not `http.HTTPClient`. Avoid utility packages (`util`, `helpers`, `common`) — they accumulate unrelated things and obscure dependency structure. If you can't name a package without a word like "util", you've identified a design smell.

## API Design and Functional Options

Public APIs are contracts. Every exported name is a cost paid forever: you cannot remove it without breaking callers. Design APIs that are easy to use correctly and difficult to use incorrectly.

The **functional options pattern** solves the configuration problem for APIs with optional parameters:

```go
type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func NewServer(addr string, opts ...Option) (*Server, error) {
    s := &Server{addr: addr, timeout: 30 * time.Second}
    for _, o := range opts { o(s) }
    return s, nil
}
```

This is backward-compatible (new options don't break existing callers), self-documenting (option names appear at call sites), and testable (options are functions, easily mocked).

## Error Philosophy

Error handling should be at the outermost layer where you have enough context to make a decision. Don't log and return — either handle the error or return it to your caller. The `github.com/pkg/errors` pattern (now partially absorbed into the standard library via `%w`) enables wrapping errors with context without losing the original.

## Deep Read

**Key Insight:** Every exported identifier is a cost paid forever. API design is fundamentally about minimizing the surface area that must remain stable, while maximizing the expressiveness of the things you do expose. The functional options pattern is the canonical Go solution to this problem: it makes the zero-value configuration work, makes extension backward-compatible, and names each option at the call site.

**What Surprised Me:** Cheney argues against using `var x = value` at package scope and instead recommends `var x Type` (zero value) plus initialization. The reasoning is subtle: package-level var initialization order is determined by the compiler, and side effects in init expressions can create surprising dependency chains. Explicit initialization makes the dependency graph visible.

**Open Questions:**
- The functional options pattern requires allocating a closure per option. For hot paths (e.g., creating many short-lived objects), is the allocation overhead measurable? Does the Go compiler inline or escape-analyze these closures?
- Cheney recommends keeping package names short and non-generic. But as ecosystems grow, name collisions become more likely. How does Go module versioning (where the module path is the identifier) affect this advice?
- The "errors should be handled once" rule conflicts with structured logging: you often want to log at each layer with different context. Is the recommended pattern to annotate errors with context fields rather than log statements?

**Wikilink Candidates:**
- [[Go Functional Options Pattern]] — the option func pattern; backward-compatible configuration; zero-value default; not yet a wiki page
- [[Go Package Design]] — single-responsibility packages; no utility packages; package name is part of exported names; not yet a wiki page

**Connections:**
- [[Go Interfaces]] — API design principles apply directly to interface design; accept interfaces/return structs rule is part of this same philosophy
- [[Go Error Handling]] — Cheney's error philosophy (handle once, annotate with context) is the pragmatic companion to the syntactic patterns in Go Error Handling
- [[Go Modules and Packages]] — package naming discipline and dependency visibility are the human-facing side of the module system's mechanical guarantees
