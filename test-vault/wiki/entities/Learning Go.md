---
title: "Learning Go"
type: book
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-learning-go-2e.md"
related:
  - Go Knowledge Map
  - preflight-sync-go
tags: [book, go, jon-bodner, oreilly]
---

# Learning Go

**Author:** [[Jon Bodner]]
**Edition:** Second (January 2024)
**Publisher:** O'Reilly — ISBN 978-1-098-13929-2

An idiomatic introduction to Go for developers who already know at least one other language. The 2nd edition adds generics coverage, structured logging (`log/slog`), fuzzing, workspaces, and the GODEBUG compatibility mechanism introduced in Go 1.22.

## Coverage (16 chapters — fully ingested)

1. Setting Up Your Go Environment
2. Predeclared Types and Declarations
3. Composite Types
4. Blocks, Shadows, and Control Structures
5. Functions
6. Pointers — stack vs heap, GC pressure, pointer-return performance counter-intuition
7. Types, Methods, and Interfaces — implicit implementation, accept interfaces/return structs
8. Generics — type constraints, type inference, standard library additions
9. Errors — wrapping, `errors.Is`/`As`, sentinel errors, panic/recover
10. Modules, Packages, and Imports — MVS, workspaces, versioning rules
11. Go Tooling — `staticcheck`, `golangci-lint`, `govulncheck`, build tags, embedding
12. Concurrency in Go — goroutines, channels, select, context, backpressure patterns
13. The Standard Library — `io`, `time`, `encoding/json`, `net/http`, `log/slog`
14. The Context — propagation, cancellation, deadlines, HTTP middleware pattern
15. Writing Tests — table-driven, benchmarks, fuzzing, httptest, race detector
16. Here Be Dragons: Reflect, Unsafe, and Cgo

## Key Concepts

- [[Go Knowledge Map]] — compiled navigation map for all 22 Go concept pages, tier-ordered by this book's chapter sequence
- [[Go Interfaces]] — implicit implementation, duck typing, nil gotchas
- [[Go Modules and Packages]] — MVS, module versioning, workspaces
- [[Go Tooling]] — static analysis, vulnerability scanning, build tags
- [[Go Error Handling]] — canonical source for Go's explicit error model
- [[Go Channel Concurrency Patterns]] — Ch.12 canonical source
- [[Go Context Patterns]] — Ch.14 canonical source

## See Also

- [[preflight-sync-go]] — active Go project this book directly informs
- [[Go Program Instrumentation]] — synthesis on the instrumentation sequence for Go programs
- [[Language Implementation Patterns]] — companion book in vault; LIP covers language tool design, Learning Go covers idiomatic language use
