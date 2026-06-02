---
title: "Go Channel Internals"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-go-channels-internals.md"
related:
  - "Go Channel Concurrency Patterns"
  - "Go Escape Analysis"
  - "Mechanical Sympathy"
  - "Go sync.Cond Pattern"
tags: [go, concurrency, channels, runtime, internals]
---

# Go Channel Internals

How Go channels work under the hood — the runtime structures, memory model implications, and scheduler integration. For usage patterns see [[Go Channel Concurrency Patterns]].

## hchan: The Channel Struct

Every `make(chan T, N)` allocates an `hchan` on the heap:

```go
type hchan struct {
    qcount   uint           // elements currently in buffer
    dataqsiz uint           // buffer capacity
    buf      unsafe.Pointer // circular ring buffer
    elemsize uint16

    closed uint32

    sendx uint32  // next write index (wraps mod dataqsiz)
    recvx uint32  // next read index

    recvq waitq   // goroutines blocked on receive
    sendq waitq   // goroutines blocked on send

    lock mutex    // spin-mutex hybrid
}
```

Key properties:
- Buffer is a contiguous ring array — cache-friendly for sequential access
- **Every send/receive acquires the lock** — lock contention, not data transfer, is the bottleneck under high concurrency
- The spin-mutex tries a brief spin before yielding to OS, reducing context switch overhead for low-contention paths

## sudog: Parked Goroutine Record

When a goroutine blocks on a channel, the runtime creates (or reuses from a pool) a `sudog`:

```go
type sudog struct {
    g    *g              // the blocked goroutine
    elem unsafe.Pointer  // value being sent or received
    c    *hchan          // which channel
    next *sudog          // linked list pointer
    // ...
}
```

`sudog`s are pooled — in programs with many goroutines blocking/unblocking, this avoids GC pressure from constant short-lived allocations.

A single goroutine in a `select` statement can have multiple simultaneous `sudog`s registered across channels. When one case fires, the runtime cancels and recycles the others.

## Send/Receive Fast Paths

### Direct stack copy (fastest)
If a receiver is already waiting in `recvq`, the sender copies the value **directly from sender stack to receiver stack** — no heap buffer involved. This is why unbuffered channels can be faster than buffered ones under low contention: buffered channels require two copies (sender→buffer, buffer→receiver), unbuffered requires one.

### Buffered path
If no receiver is waiting and buffer has space: copy into `buf[sendx]`, advance `sendx` mod capacity, increment `qcount`. Symmetric on receive.

### Block path
Buffer full (send) or buffer empty with no sender (receive): create `sudog`, enqueue on `sendq`/`recvq`, park the goroutine. The scheduler removes it from the run queue and runs another goroutine on that P.

## Memory Model: Happens-Before Rules

The Go memory model gives channel operations specific ordering guarantees:

- A **send** happens-before the corresponding **receive completes**
- For **unbuffered** channels only: a receive happens-before the send completes
- A **close** happens-before a receive that returns the zero value

**Critical subtlety:** the guarantee covers memory writes that happen *before* the send or close — not after:

```go
// Safe: x=42 happens before send, which happens-before receive
go func() { x = 42; ch <- 1 }()
<-ch
fmt.Println(x) // guaranteed to see 42

// Unsafe: x=99 happens after the send — NOT ordered relative to receiver
go func() { ch <- 1; x = 99 }()
<-ch
fmt.Println(x) // NOT guaranteed to see 99
```

## Scheduler Integration

Channels are woven into the G/M/P scheduler:

- **G** = goroutine, **M** = OS thread, **P** = logical processor (run queue manager)
- When a goroutine blocks on a channel, it's removed from P's run queue and stored in the channel's wait queue — the P immediately picks up another G to run
- When unblocked (matching operation arrives), the goroutine is placed back onto a P's local run queue
- Unblocked goroutines may resume on a different P than where they blocked, potentially causing cache misses

**Thundering herd:** closing a channel unblocks all `recvq` waiters simultaneously. In systems with many goroutines waiting on a single done channel, this can flood the scheduler.

## select Internals

`select` with N cases:
1. Inspects all channels for immediate readiness
2. If multiple ready: picks one pseudo-randomly (guaranteed by spec — prevents starvation)
3. If none ready and no `default`: registers a `sudog` on *every* channel, parks goroutine
4. When any channel fires: dequeues the goroutine, cancels `sudog`s on all other channels

The random selection when multiple cases are ready is why you cannot rely on `select` to drain channels in a predictable order.

## See Also

- [[Go Channel Concurrency Patterns]] — select, done channel, context.Context, errgroup usage patterns
- [[Go Escape Analysis]] — `hchan` is always heap-allocated; why channels can't be stack-allocated
- [[Mechanical Sympathy]] — direct stack copy as hardware-cooperative design
- [[Go sync.Cond Pattern]] — similar mutex+wait-queue model; `sync.Cond` for N-dynamic waiters without channel overhead
- [[Go Memory Model]] — the formal happens-before rules that channel operations implement; the model this page explains at the runtime level
