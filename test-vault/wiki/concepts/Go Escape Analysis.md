---
tags: [concept, go]
cluster: go
aliases: ["escape analysis", "stack vs heap", "heap allocation", "go allocations", "gcflags -m"]
related: ["Go Memory Model", "Go Tooling", "Go Interfaces"]
sources:
  - "[[archive/papers/2026-06-04-go-memory-model]]"
---

# Go Escape Analysis

The compiler determines whether each variable can be **stack-allocated** (cheaper, no GC pressure) or must be **heap-allocated** (GC-managed).

## When Variables Escape to Heap

- Returned by pointer from a function (outlives the function's stack frame)
- Stored in a slice, map, or interface (runtime access via pointer)
- Captured by a closure (goroutine captures a variable)
- Passed to an interface method (concrete type is boxed)
- Too large for the stack (~8MB limit)

## Inspecting Escape Decisions

```bash
go build -gcflags="-m" ./...
```

Output per variable: `./main.go:12:9: x escapes to heap` or `does not escape`.

## Common Escape Triggers

```go
// Escapes: returned pointer
func newT() *T { return &T{} }

// Escapes: interface boxing
var i interface{} = myStruct{}

// Escapes: goroutine capture
go func() { fmt.Println(x) }()
```

## Reducing Allocations

- Pass values instead of pointers for small structs
- Avoid interface parameters in hot paths (boxing allocates)
- Preallocate slices with known capacity: `make([]T, 0, n)`

## Connections

- [[Go Memory Model]] — escape analysis determines which variables live on heap; memory model guarantees apply to heap variables
- [[Go Tooling]] — `-gcflags="-m"` is the primary tool for inspecting escape decisions
- [[Go Interfaces]] — interface values cause escape because the concrete type is stored as a pointer; hot-path interfaces allocate
