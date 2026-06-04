---
title: "The Go Memory Model"
type: paper
captured-date: 2026-06-05
source-url: "https://go.dev/ref/mem"
author: "Russ Cox, Go Team"
publisher: "golang.org"
year: 2022
---

# The Go Memory Model

**Source:** go.dev/ref/mem (revised 2022 with explanatory additions from Russ Cox's companion post)
**Coverage:** Happens-before ordering, synchronization primitives (goroutines, channels, sync package, atomic operations), compiler and hardware reordering

## Overview

The Go memory model defines the conditions under which a goroutine reading a variable is guaranteed to observe the write produced by a different goroutine. Without synchronization, the Go compiler and the underlying CPU are free to reorder memory operations — a program that "looks correct" under naive sequential reading may have data races.

## The Core Rule

> Within a single goroutine, reads and writes must behave as if they executed in the order specified by the program.
> Across goroutines, there is no ordering guarantee unless synchronization is used.

The key concept is **happens-before**: if event A happens-before event B, then B is guaranteed to observe A's effects. If there is no happens-before relationship, there is no guarantee either way — even if the program happens to work correctly most of the time.

## Synchronization Guarantees

**Goroutine creation:** `go f()` happens-before f begins executing.

**Channel operations:**
- A send on a channel happens-before the corresponding receive completes
- The closing of a channel happens-before a receive that returns a zero value (channel closed)
- For buffered channels: the kth receive from a channel of capacity C happens-before the (k+C)th send completes

**Sync package:**
- `sync.Mutex`: Unlock n happens-before Lock n+1
- `sync.WaitGroup`: Add happens-before Done happens-before Wait returns
- `sync.Once`: the first call to f in `once.Do(f)` returns before any other call to `once.Do` returns

## Common Incorrect Patterns

Double-checked locking without atomics is a data race. Reading a variable written by another goroutine without synchronization is a data race — even if the types involved seem "atomic" (bool, int) on the underlying hardware.

The `go run -race` detector catches these at runtime but requires the race condition to actually execute during the test run.

## Escape Analysis

The compiler performs **escape analysis** to determine whether a variable can be stack-allocated or must be heap-allocated. Variables that are referenced after the function returns, passed to interfaces, or captured in closures escape to the heap. Use `go build -gcflags="-m"` to see escape decisions. Reducing heap allocations reduces GC pressure.

## Deep Read

**Key Insight:** The memory model is not about what code does on your machine today — it's about what the language *guarantees* across all conforming implementations. Code that "works" without synchronization is relying on current compiler and hardware behavior, not on the language spec. The race detector is the only reliable way to find these bugs.

**What Surprised Me:** The 2022 revision explicitly weakened some guarantees relative to the informal understanding. Specifically: the spec now says that programs with data races have "undefined behavior" in the same sense as C/C++ — the compiler is allowed to do anything, including deleting the racy access entirely if it can prove it's dead code. This is stricter than "the race might produce a stale value"; it means the compiler can legally make the program behave in ways that have nothing to do with what any thread wrote.

**Open Questions:**
- The channel synchronization rules are elegant but asymmetric: the send happens-before the receive *completes*, not before it *starts*. For buffered channels the rule involves capacity. Are these rules intuitive enough that Go programmers use them correctly by default, or does the race detector catch violations regularly?
- Escape analysis is documented as "use `-m` to inspect" but not as a stable contract. Can a future compiler version change escape decisions for existing code in ways that change observable behavior (e.g., allocating on stack where heap was used, affecting GC pause timing)?
- The memory model says synchronized-before is a partial order, not a total order. For distributed systems built in Go (where you want linearizability), the language-level memory model is insufficient — you need additional coordination. Is there a standard pattern for building linearizable operations on top of Go primitives?

**Wikilink Candidates:**
- [[Go Memory Model]] — happens-before ordering; synchronization guarantees per primitive; data race undefined behavior; not yet a wiki page
- [[Go Escape Analysis]] — stack vs heap allocation decisions; gcflags=-m inspection; closure capture triggers escape; not yet a wiki page

**Connections:**
- [[Go Channel Concurrency Patterns]] — channel synchronization guarantees are grounded in the memory model; the done-channel and worker pool patterns rely on these guarantees
- [[Go Worker Pool Pattern]] — the buffered channel capacity rule (kth receive happens-before k+Cth send) is the formal basis for why buffer size = concurrency bound
- [[Go Interfaces]] — interface values cause escape to heap; this is a common source of unexpected allocations in hot paths
