---
title: "Mechanical Sympathy"
type: concept
status: active
created: 2026-05-20
updated: 2026-05-20
sources:
  - "archive/clippings/2026-05-20-what-is-harness-engineering.md"
related:
  - "Harness Engineering"
tags:
  - agentic-systems
  - engineering-principles
---

# Mechanical Sympathy

The discipline of writing software that adapts to the substrate on which it runs, rather than fighting it.

Coined by racing driver Jackie Stewart ("you can't drive a car fast unless you understand how it works"), brought into software engineering by Martin Thompson around 2011 with the LMAX Disruptor — demonstrating millions of operations per second on commodity hardware by respecting CPU cache lines, branch prediction, memory hierarchy, and false sharing.

## Applied to LLMs

[[Harness Engineering]] is mechanical sympathy applied to a new substrate. The substrate is the LLM, the context memory, and the attention budget:

| Hardware mechanical sympathy | Harness engineering (AI agents) |
|------------------------------|----------------------------------|
| Respect CPU cache lines | Respect context window limits |
| Avoid branch mispredictions | Avoid context rot and context panic |
| Minimize false sharing | Isolate subagent contexts |
| Reduce page faults | Compress or drop stale observations |

![[resources/media/e7944e91eab91a6cc407acdc3c1eeb62_MD5.webp]]

Just as the LMAX Disruptor adapted code to how the CPU *actually* behaves (not how you wish it behaved), harness engineering adapts agent architecture to how LLMs *actually* behave — with attention that degrades mid-context, sessions that start blank, and reasoning that collapses under context pressure.

The SWE-agent paper (2024) applied this principle explicitly: held the model constant, changed only the interface, and tripled performance.

## See Also

- [[Harness Engineering]] — the discipline this principle anchors
- [[Context Rot]] — the primary LLM failure mode mechanical sympathy addresses
