---
title: "CGo Goroutine Scheduler Interaction"
question: "What is the practical upper bound on concurrent CGo calls before goroutine scheduler thread-pinning degrades overall program performance?"
type: question
status: open
domain: go
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/books/2026-05-31-learning-go-2e.md"
related:
  - "[[Go Channel Concurrency Patterns]]"
  - "[[Go Channel Internals]]"
  - "[[Go Knowledge Map]]"
tags:
  - go
  - performance
  - cgo
---

# CGo Goroutine Scheduler Interaction

*What is the practical upper bound on concurrent CGo calls before goroutine scheduler thread-pinning degrades overall program performance?*

Goroutines executing CGo calls are pinned to an OS thread for the duration of the call, bypassing Go's M:N scheduling. If many goroutines make concurrent CGo calls, the runtime creates a thread per call, which can exhaust the thread pool and degrade scheduler performance for the entire program.

The Go runtime's `GOMAXPROCS` setting limits the number of OS threads actively running Go code, but CGo calls bypass this limit — they run on additional threads outside the GOMAXPROCS budget. The interaction between `GOMAXPROCS`, CGo thread counts, and scheduler performance is documented at the mechanism level but not at the engineering guidance level: there's no well-published "at N concurrent CGo calls, performance degrades by X% under these conditions" characterization.

The practical guidance for programs mixing Go and CGo heavily is: limit concurrent CGo calls, use a semaphore to bound parallelism, and profile under load. But the specific threshold is workload-dependent.

## See Also

- [[Go Channel Concurrency Patterns]] — goroutine scheduling; context cancellation patterns
- [[Go Channel Internals]] — the scheduler internals that CGo bypasses
