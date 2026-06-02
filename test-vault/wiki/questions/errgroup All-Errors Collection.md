---
title: "errgroup All-Errors Collection"
question: "Is there a clean, idiomatic Go pattern for collecting all errors from parallel goroutines — or is this a genuine errgroup API gap requiring extra infrastructure?"
type: question
status: open
domain: go
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-30-go-errgroup-readiness-checks.md"
related:
  - "[[Go Channel Concurrency Patterns]]"
  - "[[Go Knowledge Map]]"
tags:
  - go
  - concurrency
---

# errgroup All-Errors Collection

*Is there a clean, idiomatic Go pattern for collecting all errors from parallel goroutines — or is this a genuine errgroup API gap requiring extra infrastructure?*

`golang.org/x/sync/errgroup` collects the first non-nil error and cancels the context. For parallel operations where you need all errors (not just the first), you need extra infrastructure: a mutex-protected slice and a custom error aggregator, with errgroup providing only cancellation.

The question is whether there's a clean idiomatic Go solution that doesn't require reaching outside the errgroup abstraction. Third-party libraries exist (e.g., `errgroup` variants that collect all errors), but the standard library hasn't provided one. The Go team's position may be: if you need all errors, you're doing error-first fan-out wrong; you should instead log errors as you go and return the final count/summary.

Whether this is a genuine API gap or a design decision that encourages a different pattern is worth understanding before writing the custom infrastructure.

## See Also

- [[Go Channel Concurrency Patterns]] — errgroup patterns; WaitGroup-vs-errgroup comparison
