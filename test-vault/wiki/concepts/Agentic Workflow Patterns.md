---
tags: [concept, agents, llm]
cluster: agents
aliases: ["agentic patterns", "AI workflow patterns", "agent orchestration", "multi-agent patterns"]
related: ["JSON Schema Discipline", "LLM Tool Calling", "Context Engineering"]
sources:
  - "[[archive/clippings/2026-06-04-eleven-agentic-patterns]]"
---

# Agentic Workflow Patterns

Eleven named patterns for multi-agent system design, each identified by a **trigger phrase** in a scenario. Matching trigger to pattern — not memorizing names — is the core skill.

## The Patterns

| Pattern | Trigger |
|---------|---------|
| **Prompt Chaining** | Steps are known in advance; each step has a simpler job than the whole |
| **Routing** | Inputs fall into distinct categories; one-size-fits-all degrades across them |
| **Parallelization (Sectioning)** | Independent subtasks can run concurrently |
| **Parallelization (Voting)** | Consensus needed through diverse perspectives |
| **Orchestrator-Workers** | Open-ended; subtask list cannot be predefined, plan emerges from exploration |
| **Evaluator-Optimizer** | Clear evaluation criteria exist AND iterative feedback measurably improves output |
| **Hub-and-Spoke** | Multi-agent work where subagents must not share context directly |
| **Dynamic Adaptive Decomposition** | Unknown problem space; plan reveals itself as findings come in |
| **Sequential Prompt Chaining (Coordinator-Scoped)** | Multi-aspect work that decomposes into a known, predictable order |
| **Programmatic Prerequisites** | Ordering must be guaranteed; prompt-based ordering has unacceptable failure rate |
| **Graceful Degradation** | Independent subtask failed; successful work must not be discarded |

## Pattern Diagrams

![[resources/media/129d60b226da713f7d5fa0aab3153a64_MD5.webp]]
![[resources/media/12c8f5f88a035387e33067ad88ee360a_MD5.webp]]
![[resources/media/1736686ac91552ef6b415ba156af189a_MD5.webp]]
![[resources/media/56d9d7461b1130e79cb44b27c3c70197_MD5.webp]]
![[resources/media/62db0763d7cc3b4628bc8e5f5b782773_MD5.webp]]
![[resources/media/646fe23f05e709641161f466ed1d72fd_MD5.webp]]
![[resources/media/6b186be725a46cfb1dc97e7f3c3d2b61_MD5.webp]]
![[resources/media/809c5b04973e9ef33bb651551dc9a41e_MD5.webp]]
![[resources/media/97eeb06f18564de219f7516e9af76e33_MD5.webp]]
![[resources/media/9db9e0c3bbb82dc68ac3f688868f175e_MD5.webp]]
![[resources/media/9eb9e3a4c7fecbf9d599ad3a66c514b4_MD5.webp]]
![[resources/media/dc02a6edf6b6f8f79b853d7e809a476f_MD5.webp]]
![[resources/media/f482a9e111b046d9dea21e268a1f8dab_MD5.webp]]
![[resources/media/f82111ff83c26cd7fb30584c22814d39_MD5.webp]]

## Connections

- [[JSON Schema Discipline]] — when to enforce schema at handoff boundaries
- [[LLM Tool Calling]] — tool use is the mechanism behind orchestrator-workers and most multi-agent patterns
- [[Context Engineering]] — hub-and-spoke isolates context deliberately; all patterns depend on context management
- [[Executive Function]] — AI agents externalize the planning and task-decomposition that EF would otherwise require
- [[Scaffolding]] — orchestrator-worker pattern is a scaffold for users who can't decompose tasks independently
