---
title: "Learning Go, 2nd Edition"
type: books
captured-date: 2026-05-31
source-url: ""
author: "Jon Bodner"
publisher: "O'Reilly"
year: 2024
isbn: "978-1-098-13929-2"
---

# Learning Go, 2nd Edition — Jon Bodner

**Source:** `2026-05-31-learning-go-2e.pdf` (same directory)
**Coverage:** Full book — 16 chapters: environment setup, types, composite types, control structures, functions, pointers, types/methods/interfaces, generics, errors, modules/packages, tooling, concurrency, standard library, context, testing, reflect/unsafe/cgo

## Deep Read

**Key Insight:** Go's most distinctive design choices — implicit interfaces, explicit error values, context as an explicit parameter, goroutines communicating via channels — all emerge from the same philosophy: make dependencies, failures, and data flow *visible in the type system* rather than hiding them in exceptions, thread-locals, or inheritance hierarchies. "Boring Go" is a deliberate design goal, not a limitation.

**What Surprised Me:** Returning a pointer to a small struct is *slower* than returning the value — ~30ns for a pointer vs ~10ns for a 100-byte struct on an i7, because pointer returns force heap allocation and increase GC pressure. The performance crossover to pointer being faster doesn't happen until around 10MB of data. The conventional wisdom "use pointers for performance" is wrong for the vast majority of Go code.

**Open Questions:**
- The Go scheduler integrates with the network poller to detect I/O-blocked goroutines — how does this interact with CGo calls, which block OS threads rather than goroutines? Does heavy CGo use undermine the scheduler's efficiency guarantees?
- Go 1.22 introduced loop-variable scoping changes via a GODEBUG-based mechanism that allows backward-incompatible behavior fixes without breaking semver. How far does this governance model extend — what's the bar for adding a new GODEBUG flag vs. making a clean breaking change?
- `context.WithValue` uses `any` for both key and value, relying on unexported key types for collision avoidance. Is there a generics-based replacement being considered, or is the current pattern considered idiomatic enough to stay?

**Wikilink Candidates:**
- [[Go Interfaces]] — implicit interface implementation, type-safe duck typing, nil gotchas, accept interfaces/return structs rule — no dedicated page yet (created)
- [[Go Modules and Packages]] — module system, workspaces, minimal version selection, versioning, go.mod — no dedicated page yet (created)
- [[Go Tooling]] — staticcheck, golangci-lint, govulncheck, build tags, go generate, embedding — no dedicated page yet (created)

**Connections:**
- [[Go Channel Concurrency Patterns]] — Ch.12 is the canonical source; covers goroutine scheduler, CSP model, select, backpressure, WaitGroups, mutexes
- [[Go Context Patterns]] — Ch.14 maps directly; covers WithValue, WithCancel, WithDeadline, middleware HTTP pattern
- [[Go Error Handling]] — Ch.9 is the authoritative source; wrapping with %w, Is/As, sentinel errors, panic/recover
- [[Go Escape Analysis]] — Ch.6 on stack vs heap, pointer escape rules, GC pressure
- [[Go Generics Type Constraints]] — Ch.8; type terms, comparable, type inference, standard library additions
- [[Go Testing Patterns]] — Ch.15; table tests, httptest, integration tests, fuzzing, benchmarks, data race detector
- [[Go Memory Model]] — Ch.6 on slice internals, map-as-pointer, stack growth
- [[preflight-sync-go]] — active Go project this book directly informs; concurrency (Ch.12) and context (Ch.14) are the most relevant chapters
- [[Language Implementation Patterns]] — both books now in vault span theory→practice: LIP for parsing/compiler design, Learning Go for idiomatic language use
