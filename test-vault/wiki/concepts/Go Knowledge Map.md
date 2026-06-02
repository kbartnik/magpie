---
title: "Go Knowledge Map"
type: hub
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - archive/books/2026-05-31-learning-go-2e.md
related:
  - preflight-sync-go
tags:
  - go
  - moc
---

# Go Knowledge Map

Navigational hub for the vault's Go knowledge cluster. Pages are ordered by conceptual dependency — later entries build on earlier ones. Each annotation notes *why it matters*, not what it contains.

**Active project:** preflight-sync-go — a Go TUI + file-sync tool. The concurrency and context tiers below are the most live.

---

## Tier 1 — Types and Values

The substrate everything else builds on. Understand these before reading any other Go page.

- Go Defined Types — Go's enum substitute: a defined type + iota constant block gives the compiler discriminant enforcement. Not just syntax — it determines what operations are legal.
- Go Struct Memory Layout — field ordering controls padding and struct size. Matters when allocating large numbers of structs or passing by value vs pointer.
- Go Empty Struct Pattern — `struct{}` as a zero-byte presence marker. Foundation for sets, channel signals, and visited-node tracking; shows up everywhere in concurrency code.

---

## Tier 2 — Memory and Pointers

Why Go code looks the way it does at function boundaries.

- Go Escape Analysis — the compiler decides stack vs heap. Returning a pointer to a small value is *slower* than returning the value (~30ns vs ~10ns), because it forces a heap allocation. "Use pointers for performance" is wrong for small types.
- Go Memory Model — the sequenced-before / synchronized-before / happened-before trinity. Required reading before writing any concurrent code: defines what each primitive (channel, mutex, atomic, WaitGroup, Once) actually guarantees.

---

## Tier 3 — Abstraction

How Go models behaviour without inheritance.

- Go Interfaces — implicit implementation (no `implements`). The key rule: *accept interfaces, return structs*. The caller defines the interface; the implementer doesn't need to know it exists. Nil interface gotcha: a non-nil interface can hold a nil concrete value.
- Go Defined Types — (see Tier 1) defined types can carry methods; type declarations are not inheritance.
- Go Functional Options Pattern — variadic `Option func` type for backward-compatible config. The idiomatic Go substitute for named/optional parameters and builder patterns.
- Go Generics Type Constraints — `~T` underlying-type constraints; type inference via unification. Use generics to reduce repetition, not to achieve runtime polymorphism (that's interfaces).

---

## Tier 4 — Errors

Go's explicit error model.

- [[Go Error Handling]] — `(value, error)` return convention; `%w` wrapping; `errors.Is`/`errors.As` for unwrapping chains; sentinel errors for "no further processing" states. The compiler enforces that errors are read.

---

## Tier 5 — Concurrency

The most distinctive tier. Read in order — each page assumes the previous.

- Go Channel Internals — the runtime data structure (`hchan`), sudog pooling, send/receive fast paths, scheduler integration. Explains *why* channels behave as they do.
- Go Channel Concurrency Patterns — `select`, channel close, done-channel vs `context.Context`, worker pool, backpressure. The practical patterns built on the internals. *(stub — awaiting preflight context.Context refactor)*
- Go sync.Cond Pattern — for N-dynamic waiters where channels are awkward. `Wait()` atomically releases the lock; spurious wakeup loop is required. Use channels unless you need broadcast to an unknown number of goroutines.
- Go Context Patterns — `context.Value` for request-scoped metadata only; cancellation and deadlines as the primary use. Unexported key type prevents collisions. Pass as the first parameter — always explicit, never ambient.

---

## Tier 6 — I/O and Standard Library

Practical patterns for the things Go programs actually do.

- [[Go File IO]] — `WalkDir`, `os.Stat` stability, `defer close`, NFS `Sync()`. The mechanics of safe filesystem interaction.
- Go Structured Logging — `log/slog` (Go 1.21); Logger/Handler split; `LogAttrs` for zero-allocation hot paths; `Logger.With` for request-scoped fields; `LogValuer` for redaction.

---

## Tier 7 — Testing

- [[Go Testing Patterns]] — table-driven tests, `testify` assert/require, dependency injection via function parameters, `httptest`, integration tests via build tags, data race detector (`-race`), fuzzing, benchmarks.

---

## Tier 8 — Module System and Tooling

Infrastructure. Usually read once, then referenced.

- [[Go Modules and Packages]] — module/package hierarchy; minimal version selection (MVS picks the *minimum* satisfying version, not latest); major version import path rules (`/v2`); `internal/` packages; avoiding `init()`.
- Go Workspaces — `go work` for simultaneous multi-module development. Local only; don't commit `go.work`.
- Go Tooling — `staticcheck`, `golangci-lint`, `govulncheck`; build tags for conditional compilation; `go:generate`; file embedding; cross-compilation with `GOOS`/`GOARCH`.

---

## Canonical Source

Learning Go (Jon Bodner, O'Reilly 2024) → `archive/books/2026-05-31-learning-go-2e.md`

Tier ordering above mirrors the book's chapter sequence. The Deep Read in that archive file has the key insight, what was surprising, and the three open questions worth returning to.

---

## Coverage Gaps

Pages not yet in the vault:
- Go reflection (`reflect` package) — Ch.16 of Learning Go; rarely needed but occasionally essential
- Go HTTP patterns — `net/http` server/client, `ResponseController`, middleware chaining
- Go benchmarking patterns — `testing.B`, `b.ReportAllocs`, `b.ResetTimer`; `pprof` for profiling

See Go Program Instrumentation for the synthesis on the full instrumentation sequence (-race → testing.B → -benchmem → pprof → slog) and what the pprof gap means in practice.

Open design questions: Go Generics for context.WithValue, errgroup All-Errors Collection, CGo Goroutine Scheduler Interaction, GODEBUG Governance Model.
