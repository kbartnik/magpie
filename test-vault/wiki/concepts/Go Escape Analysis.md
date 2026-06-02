---
title: "Go Escape Analysis"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-structs-deep-dive.md"
related:
  - "[[Go Struct Memory Layout]]"
  - "[[Mechanical Sympathy]]"
tags:
  - go
  - performance
  - memory
---

# Go Escape Analysis

The Go compiler decides at compile time whether a variable lives on the **stack** (fast, no GC) or the **heap** (slower, GC-managed). This decision is called escape analysis.

## The Rule

A variable **escapes to the heap** when its lifetime must outlast the function call that created it. The most common trigger: returning a pointer.

```go
// Escapes — pointer returned, outlives the function
func createUser() *User {
    u := User{Name: "isu"}
    return &u
}

// Stays on stack — value used and discarded within the function
func process() {
    u := User{Name: "isu"}
    fmt.Println(u.Name)
}
```

## Why It Matters

- Heap allocations create work for the garbage collector
- GC pauses introduce latency spikes in high-throughput systems
- Fewer heap allocations = lower GC pressure = more predictable latency

## Inspecting Escape Decisions

```bash
go build -gcflags="-m" ./...
```

The compiler prints each escape decision. Look for `"escapes to heap"` on variables you want to keep on the stack.

More verbose output:

```bash
go build -gcflags="-m -m" ./...
```

## Common Escape Triggers

| Cause | Example |
|-------|---------|
| Returning a pointer | `return &localVar` |
| Storing in an interface | `var i interface{} = myStruct` |
| Passing to a function that stores the pointer | `go func(p *T)` in a goroutine |
| Closures capturing a variable | `func() { use(x) }` |

## Reducing Heap Allocations

- Pass values instead of pointers for small structs
- Use `sync.Pool` to reuse heap-allocated objects
- Use `[]byte` buffers on the stack with `var buf [N]byte` rather than `make([]byte, N)` when N is known and small

## sync.Pool: Reusing Allocations

`sync.Pool` reduces allocation pressure for **short-lived, reusable objects** (buffers, scratch structs). The GC can evict the entire pool at any time — never use it for objects that must persist.

```go
var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

func processRequest(data []byte) []byte {
    buf := bufPool.Get().(*bytes.Buffer)
    buf.Reset() // always reset before use — previous content is still there
    defer bufPool.Put(buf)
    buf.Write(data)
    return buf.Bytes()
}
```

Two failure modes:
1. **Forgetting `Reset()`** — previous user's data is still in the object
2. **Using the object after `Put()`** — data race; the pool owns it now

## defer Cost

`defer` adds a small but real overhead per call — a deferred function record on the goroutine stack. In hot paths (>10K calls/sec), this accumulates:

```go
// Benchmark to measure the difference
func BenchmarkWithDefer(b *testing.B) {
    var mu sync.Mutex
    for i := 0; i < b.N; i++ { mu.Lock(); defer mu.Unlock() }
}
```

Run with `go test -bench=. -benchmem`. **Rule:** use `defer` freely outside hot paths; avoid in tight loops or high-frequency handlers.

Also: `defer` with named return values can modify the return value — useful for middleware error wrapping, confusing otherwise:

```go
func double(n int) (result int) {
    defer func() { result *= 2 }() // modifies the named return
    result = n
    return // returns n*2
}
```
