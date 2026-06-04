---
title: "Contexts and Structs: How to Use context.Context in Go"
author:
  - "Jean de Klerk"
  - "Matt T. Proud"
url: "https://go.dev/blog/context-and-structs"
description: "The Go team's authoritative guidance on when and how to use context.Context — including the common misuse patterns and the rules for passing context across API boundaries."
tags:
  - "clippings"
  - "go"
  - "concurrency"
published: 2021-02-24T00:00:00Z
created: 2026-06-05T09:00:00-04:00
---

# Contexts and Structs: How to Use context.Context in Go

By Jean de Klerk, Matt T. Proud
Published: 2021-02-24

>[!summary]
>The official Go blog post on context.Context usage. context.Context carries deadlines, cancellation signals, and request-scoped values across API boundaries. It should be passed as the first parameter to functions, never stored in structs, and used for cancellation and metadata — not for dependency injection.

## The Rules

**Rule 1: Pass Context as the first parameter, named `ctx`.**
```go
func DoSomething(ctx context.Context, arg Arg) error { ... }
```

**Rule 2: Do not store Contexts in structs.**
Context belongs to a call path, not an object. Storing context in a struct makes it impossible to use different contexts for different calls on the same object, and makes the context's lifetime ambiguous.

**Rule 3: Pass `context.Background()` only at the top of a call tree.**
Background is the root context — zero deadline, never cancelled, no values. Use it at main(), test entry points, and server handlers before deriving a request-scoped context.

**Rule 4: Use `context.TODO()` when you're not sure which context to use.**
`TODO` signals "this code needs a real context, but I'm not sure what yet." It's a placeholder visible to static analysis tools.

## Cancellation

`context.WithCancel` returns a derived context and a cancel function. When cancel is called, the context is cancelled and any operation respecting `ctx.Done()` can terminate:

```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel()  // always call cancel to release resources
```

**`context.WithTimeout` and `context.WithDeadline`** are the most common patterns for HTTP clients, database queries, and any operation that should fail fast rather than hang.

## Context Values: The Controversial Part

`context.WithValue` stores a key-value pair in the context. The rule: use only for request-scoped data that crosses process boundaries (trace IDs, authentication tokens, request IDs). Do not use it for optional function parameters, dependency injection, or anything you could pass as a regular argument.

The key must be an unexported type to prevent collisions:
```go
type contextKey string
const traceIDKey contextKey = "trace-id"
```

## Structured Logging with Context

A common pattern: attach a `*slog.Logger` (or structured logger) to the context with a trace ID already embedded, so all log calls downstream automatically carry the request trace. This is the narrow exception to "don't inject dependencies via context" — the logger is metadata about the request, not a functional dependency.

## Deep Read

**Key Insight:** `context.Context` is a cancellation tree, not a dependency container. Its purpose is to let callers communicate "you have until time T" or "you can stop now" to the functions they call. Using it for dependency injection violates this contract — dependencies don't change per-request, don't need cancellation, and should be tested independently.

**What Surprised Me:** The rule "don't store context in structs" has a real exception: when the struct *is* a request or a unit of work with a natural lifetime. An HTTP handler can legitimately hold a context for its lifetime. The distinction is whether the struct represents a call or an object — calls have context; objects don't.

**Open Questions:**
- Structured logging via context (attaching a logger with a trace ID) is widely used but technically violates the "request metadata only" rule. Is there a cleaner pattern, or is this the accepted pragmatic exception?
- `context.WithValue` uses `interface{}` for both key and value — there's no type safety at retrieval sites. The unexported key type pattern prevents collisions but not type errors. Is there a proposal to type-parameterize context values?
- Cancellation propagates downward through the context tree but not upward. If a child operation fails and you want to cancel the parent, you need a separate mechanism (errgroup, a manual cancel call). Is there a standard pattern for "cancel parent if child fails"?

**Wikilink Candidates:**
- [[Go Context Patterns]] — context.Context for cancellation and deadlines; WithCancel/WithTimeout/WithDeadline; context values for request metadata only; unexported key type pattern; not yet a wiki page
- [[Go Structured Logging]] — attaching a slog.Logger to context with trace ID; Logger.With for request-scoped fields; not yet a wiki page

**Connections:**
- [[Go Channel Concurrency Patterns]] — the done-channel pattern from pre-context Go is the ancestor of context.WithCancel; context replaced done channels as the standard cancellation mechanism
- [[Go Worker Pool Pattern]] — context cancellation is the standard way to shut down a worker pool; workers range on their work channel but also select on ctx.Done()
- [[Go Error Handling]] — context cancellation errors (context.Canceled, context.DeadlineExceeded) are sentinel errors that callers check with errors.Is
