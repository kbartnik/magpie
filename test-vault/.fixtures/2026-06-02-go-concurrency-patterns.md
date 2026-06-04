---
title: "Go Concurrency Patterns"
type: video
captured-date: 2026-06-02
source-url: "https://www.youtube.com/watch?v=f6kdp27TYZs"
---

# Go Concurrency Patterns

**Speaker:** Rob Pike
**Event:** Google I/O 2012
**Duration:** ~51 min

## Overview

Foundational talk on Go's concurrency model. Pike walks through the core primitives — goroutines and channels — and builds up a vocabulary of reusable patterns: generators, fan-in multiplexing, worker pools, timeouts, and cancellation. The unifying theme is that channels are values: you can pass them around, return them from functions, and compose them into larger patterns just like any other type.

## Core Primitives

**Goroutines:** Independently executing functions launched with `go`. Cheap — you can have thousands. Multiplexed onto OS threads by the Go scheduler. Not threads, not coroutines, something in between.

**Channels:** Typed conduits for communication between goroutines. Sending blocks until a receiver is ready (unbuffered), or until the buffer has space (buffered). This synchronization property is what makes channels safe by design — no shared memory, no locks needed for the patterns in this talk.

**The CSP model:** Go's concurrency is based on Hoare's Communicating Sequential Processes. The key idea: don't communicate by sharing memory; share memory by communicating.

## Patterns

**Generator:** A function that returns a channel. The goroutine producing values owns the channel; the caller just ranges over it. Decouples producer from consumer.

```go
func boring(msg string) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
        }
    }()
    return c
}
```

**Fan-in (multiplexing):** Merge multiple channels into one. A `fanIn` function takes N input channels and returns a single output channel. Each input gets its own goroutine forwarding values to the output.

**Select statement:** The `select` block waits on multiple channel operations simultaneously, proceeding with whichever is ready first. If multiple are ready, one is chosen at random. This is the primitive underlying timeout and cancellation patterns.

```go
select {
case v := <-c1:
    fmt.Println(v)
case v := <-c2:
    fmt.Println(v)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

**Timeout:** `time.After` returns a channel that receives after a duration. Composing it with `select` adds a deadline to any channel operation — no special timeout API needed.

**Done channel / cancellation:** Pass a `done` channel into goroutines. When the caller closes it, all goroutines select on `done` and exit. The pattern that became `context.Context`.

**Worker pool:** N goroutines reading from a shared work channel. Bounded concurrency without a semaphore — the channel's buffer is the bound.

## Deep Read

**Key Insight:** The generator pattern makes concurrency look like iteration. `for v := range boring("hello")` reads like a synchronous loop but runs the producer in a separate goroutine. This is the same compositional thinking that makes Unix pipes powerful — small pieces, standard interfaces, composable.

**What Surprised Me:** Pike shows that the fan-in pattern is itself just a generator: `fanIn` returns a channel. You can fan-in the outputs of fan-ins. The patterns are fractal — the same shape at every level, which is why complex pipelines remain readable.

**Open Questions:**
- The done-channel pattern Pike shows here became `context.WithCancel` in Go 1.7. How closely does the stdlib implementation mirror his approach — does `context.Context` still use a channel internally, or did it evolve to something more efficient?
- Worker pools here use a shared input channel. `errgroup` and `semaphore` from `golang.org/x/sync` solve similar problems. When does the simple channel-based pool break down and require those packages?
- The talk predates generics. A typed `fanIn[T any](channels ...<-chan T) <-chan T` would be cleaner than `interface{}`. Has a generics-native concurrency utility library emerged in the community?

**Wikilink Candidates:**
- [[Go Channel Concurrency Patterns]] — canonical reference for the generator, fan-in, and done-channel patterns; this talk is the primary source
- [[Go Select Statement]] — `select` mechanics and random selection behavior; not yet a dedicated page
- [[Go Worker Pool Pattern]] — bounded concurrency via buffered channel; not yet a dedicated page

**Connections:**
- [[Go Interfaces]] — channel types are interfaces under the hood (`<-chan T` vs `chan<- T`); the generator pattern returns an interface-like read-only channel
- [[Go Modules and Packages]] — `golang.org/x/sync` (errgroup, semaphore) extends the patterns shown here; module system context matters
- [[Go Context Patterns]] — the done-channel pattern in this talk is the direct predecessor of `context.WithCancel`; the evolution from manual done channels to context is a key Go history story
