---
tags: [concept, go]
cluster: go
aliases: ["go package design", "go API design", "package naming", "go surface area"]
related: ["Go Functional Options Pattern", "Go Modules and Packages", "Go Interfaces"]
sources:
  - "[[archive/videos/2026-06-04-dave-cheney-practical-go]]"
---

# Go Package Design

Every exported identifier is a cost paid forever. Package design is minimizing surface area while maximizing expressiveness.

## Package Naming

- Package name is part of every exported identifier: `http.Client`, not `http.HTTPClient`
- Names should be short, lowercase, not stutter (`bytes.ByteBuffer` → `bytes.Buffer`)
- Avoid utility packages (`util`, `helpers`, `common`) — they accumulate unrelated things

## Single Responsibility

A package should provide a single, focused capability. If you can't name a package without "util", the design needs revisiting — the contents belong in the packages that own them.

## Naming Identifiers

- Variables: what they hold, not their type. `var users []User`, not `var userSlice []User`
- Receiver names: 1-2 letters derived from type name, consistent across all methods: `func (c *Client) Do()` not `func (client *Client) Do()`
- Short names for short-lived variables (loop indices: `i`, `j`); longer names for longer-lived or exported names

## Package-Level Variables

Prefer zero-value initialization over `var x = someFunc()` for package-level variables. Side effects in init expressions create implicit dependencies that are hard to track.

## Connections

- [[Go Functional Options Pattern]] — functional options are the canonical solution to backward-compatible API design
- [[Go Modules and Packages]] — the module path provides global namespace; package naming operates within that
- [[Go Interfaces]] — package design determines what interfaces are exported; the accept-interfaces-return-structs rule shapes the public API
