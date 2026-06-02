---
title: "Go Struct Memory Layout"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-go-structs-deep-dive.md"
related:
  - "[[Go Escape Analysis]]"
  - "[[Mechanical Sympathy]]"
tags:
  - go
  - performance
  - memory
---

# Go Struct Memory Layout

Go aligns struct fields in memory according to each field's alignment requirement (typically equal to its size). This can insert invisible **padding bytes** between fields, inflating the struct's total size.

## The Padding Problem

```go
type A struct {
    X int8   // 1 byte
    Y int64  // 8 bytes — must be 8-byte aligned
}
// Size = 16 bytes (1 byte + 7 padding + 8 bytes)
```

The 7 padding bytes between `X` and `Y` are wasted space. The compiler inserts them so `Y` starts at an 8-byte boundary.

## The Fix: Largest Fields First

```go
type A struct {
    Y int64  // 8 bytes
    X int8   // 1 byte
}
// Size = 16 bytes (8 + 1 + 7 trailing padding)
```

Same size in this small example, but in larger structs with many mixed-size fields, putting the largest fields first can significantly reduce total size.

## General Rule

Order fields **largest to smallest alignment** to minimize internal padding:

1. `int64`, `float64`, `uint64`, pointers (8 bytes)
2. `int32`, `float32`, `uint32` (4 bytes)
3. `int16`, `uint16` (2 bytes)
4. `int8`, `uint8`, `bool`, `byte` (1 byte)

## Inspecting Struct Size

```go
import "unsafe"
fmt.Println(unsafe.Sizeof(MyStruct{}))
```

## Why It Matters

- Smaller structs mean more fit in a CPU cache line (64 bytes typical)
- Fewer cache misses = faster iteration over slices of structs
- Reduced GC pressure when many instances are allocated
- Connects directly to [[Mechanical Sympathy]] — hardware caches operate on fixed-size lines

## Prefer Typed Structs Over Maps

For performance-sensitive code, a typed struct beats `map[string]interface{}`:

```go
// Slow — reflection + allocations for every marshal/unmarshal
type Event struct {
    Data map[string]interface{}
}

// Fast — fixed layout, no reflection overhead
type Event struct {
    UserID int64  `json:"user_id"`
    Action string `json:"action"`
}
```
