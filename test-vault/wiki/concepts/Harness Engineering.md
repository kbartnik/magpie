---
title: "Harness Engineering"
type: concept
status: active
created: 2026-05-20
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-20-what-is-harness-engineering.md"
  - "archive/clippings/2026-05-20-harness-engineering-memory-gap.md"
  - "archive/clippings/2026-06-01-no-vibes-allowed-context-engineering.md"
related:
  - "Agentic Workflow Patterns"
  - "JSON Schema Discipline"
  - "Context Rot"
tags:
  - agentic-systems
  - harness-engineering
---

# Harness Engineering

The discipline of designing, building, and operating the engineered runtime that wraps a large language model and converts its raw text output into reliable system behavior.

**Formula:** `Agent = Model + Harness`

## What the Harness Does (Six Dimensions)

The harness does six things the model cannot do on its own:

| Dimension | What it handles |
|-----------|----------------|
| **Context assembly** | Shapes what the model sees on each call |
| **Tool contracts and validators** | Decides what the model is allowed to do |
| **Memory and durable state** | Remembers what happened across calls |
| **Observability** | Watches what the model produced (drift detection, evaluation gates) |
| **Recovery** | Rollback, retry, replay when something goes wrong |
| **Orchestration** | Coordinates when multiple models or agents are involved |

![[resources/media/119a8ebf7edd49c645da5fe58714664c_MD5.webp]]

Of the six agent dimensions (Perception, Brain, Memory, Planning, Action, Collaboration), four are harness concerns. Perception and Brain are largely shaped by the model.

## Named Failure Modes

These have names now — design your harness to compensate:

- **[[Context Rot]]** — performance decay as context fills with stale or low-signal tokens
- **Context panic** — agent skips steps and short-circuits plans under context pressure
- **Lost-in-the-middle** — information buried in the middle of long prompts is under-attended vs. beginning/end
- **U-shaped attention** — the generalization: attention follows a U-curve over long contexts

## Harness Engineering vs. Context Engineering

Dex Horthy (AI Engineer 2025) draws a distinction worth preserving: **context engineering** is the general discipline of managing what goes into a model's context window. **Harness engineering** is the codebase-specific layer — how you customize the integration points of a particular tool (Claude Code hooks, CLAUDE.md files, skills, Codex config, Cursor rules) for your specific repository and workflow.

The distinction matters because context engineering principles (compaction, dumb zone avoidance, phase boundaries) are universal, but harness engineering work is necessarily local: your team's branching strategy, your codebase's documentation hierarchy, your sub-agent configurations. A team that learns context engineering principles still needs to do harness engineering work to apply them to their environment.

## Origin

- **May 2024**: SWE-agent paper (Princeton) introduced Agent-Computer Interface (ACI). Held GPT-4 Turbo constant, changed only the interface layer (file search cap, stateful viewer, linter, context manager). SWE-bench performance: 3.8% → 12.47%. Proved the runtime around the model can outperform model upgrades.
- **Late 2025**: Anthropic published engineering posts using "harness" as a term of art (*Effective Context Engineering*, *Effective Harnesses for Long-Running Agents*)
- **February 2026**: Mitchell Hashimoto coined the discipline name. OpenAI used it the same month. Terminology converged.

The cockpit metaphor (Fitts & Jones, 1947 USAF study): when operators repeat the same mistake, the environment is the variable — redesign the cockpit, not the pilot.

## Mechanical Sympathy for LLMs

Harness engineering is [[Mechanical Sympathy]] applied to a new substrate. Just as LMAX Disruptor (2011) optimized for CPU cache lines and branch prediction, harness engineering optimizes for LLM failure modes: context rot, context panic, lost-in-the-middle, U-shaped attention.

## What the Harness Is Not

Not a try/catch wrapper. Not a guardrail preventing bad outputs. It is the engineered environment that enables capable models to do larger, longer, more autonomous work than they could do alone.

## See Also

- [[Agentic Workflow Patterns]] — the practice layer; specific patterns for agent coordination
- [[JSON Schema Discipline]] — schema design is part of harness design
- [[Context Rot]] — the primary failure mode harness memory management prevents
- [[Mechanical Sympathy]] — the underlying engineering principle
- [[Graphify]] — implements the PreToolUse hook injection pattern; a concrete harness component
- [[Claude Code Memory Architecture]] — the memory dimension implemented: four-layer system for cross-session persistence
- [[Claude Code Hooks]] — the architectural constraints dimension implemented: deterministic lifecycle control
- [[Claude Code Skills]] — the capability acquisition dimension implemented: PDA-based skill packaging
- [[Agent Memory Architectures]] — decision framework for the harness memory dimension: RAG vs LLM Wiki vs Fat Skills
- [[LLM Wiki Pattern]] — the "compiled, auditable artifact" approach to harness memory; Karpathy's pattern and implementations
- [[Digital Phenotyping]] — structural parallel: observability/memory applied to human behavioral data rather than AI system logs; same instrumentation → inference → intervention loop
- [[Agentic Identity and Zero Trust]] — vault-as-middleware is the "tool contracts and validators" harness dimension extended to backend identity bridging; short-lived credential issuance is a harness responsibility
