---
title: "Go Empty Struct Pattern"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-structs-deep-dive.md"
related:
  - "[[Go Channel Concurrency Patterns]]"
  - "[[Go Struct Memory Layout]]"
tags:
  - go
  - idioms
  - concurrency
---

# Go Empty Struct Pattern

`struct{}` is a zero-byte type. It allocates no memory and carries no information — only presence.

```go
fmt.Println(unsafe.Sizeof(struct{}{})) // 0
```

## Mental Model

Use `struct{}` when **existence matters but value does not**. Think of it as a tick mark, a marker, or a signal.

## Use Cases

### Sets

Go has no built-in set. `map[T]struct{}` is the idiomatic implementation:

```go
// bool wastes 1 byte per entry
visited := map[string]bool{}

// struct{} costs nothing
visited := map[string]struct{}{}
visited["apple"] = struct{}{}

if _, ok := visited["apple"]; ok {
    // present
}
```

### Channel Signaling

When a goroutine needs to signal completion but has no data to send, `chan struct{}` is cleaner and more expressive than `chan bool`:

```go
done := make(chan struct{})

go func() {
    // do work
    done <- struct{}{}
}()

<-done
```

The type explicitly says: "this channel carries a signal, not a value."

### Goroutine Coordination (Counting Completions)

```go
done := make(chan struct{})
for i := 0; i < 3; i++ {
    go func() {
        // work
        done <- struct{}{}
    }()
}
for i := 0; i < 3; i++ {
    <-done
}
```

### Dedup / Visited Tracking

```go
visited := make(map[int]struct{})
visited[nodeID] = struct{}{}

if _, ok := visited[nodeID]; ok {
    continue
}
```

## Subtle Behavior: Address Identity

Because `struct{}` has zero size, the compiler may assign the same address to all zero-size values:

```go
a := struct{}{}
b := struct{}{}
fmt.Println(&a == &b) // may print true
```

Do not take the address of `struct{}` and compare it for identity — the result is undefined in practice. Use `struct{}` only as a value, never as an identity token.

## See Also

- [[Go Channel Concurrency Patterns]] — `chan struct{}` for done/cancel signaling
- [[Go Struct Memory Layout]] — why zero-size matters in the broader memory layout picture
