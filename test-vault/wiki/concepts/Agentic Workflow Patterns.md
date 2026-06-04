---
type: concept
tags: [ai, agents, architecture, patterns]
---

# Agentic Workflow Patterns

The eleven patterns underlying production multi-agent systems, organized by trigger and structural shape. The actual skill is matching **trigger to pattern** — not recognizing names.

**Primary source:** [[archive/clippings/2026-06-04-eleven-agentic-patterns|The Eleven Patterns Behind Every Production Agentic System]]

---

## The Trigger → Pattern Map

| Trigger | Pattern |
|---------|---------|
| Steps knowable before run starts | Prompt Chaining |
| Distinct input categories | Routing |
| Independent subtasks running together | Parallelization (Sectioning) |
| Consensus through diverse perspectives | Parallelization (Voting) |
| Plan emerges from exploration | Orchestrator-Workers / Dynamic Decomposition |
| Output needs to be checked and revised | Evaluator-Optimizer (single) / Iterative Refinement (multi) |
| Coordinator + specialist subagents | Hub-and-Spoke (base) |
| Ordering must be guaranteed deterministically | Programmatic Prerequisites |
| Independent failure shouldn't waste good work | Graceful Degradation |
| Handing off to human or external system | Structured Handoff |

---

## The Five Canonical Patterns (Anthropic)

### 1. Prompt Chaining
Fixed sequence of LLM calls; each step's output feeds the next.

![[resources/media/809c5b04973e9ef33bb651551dc9a41e_MD5.webp]]

**Trigger:** Steps are known in advance. Each step has a simpler job than the whole. Latency is acceptable.

### 2. Routing
A classifier dispatches input to one of several specialized paths.

![[resources/media/129d60b226da713f7d5fa0aab3153a64_MD5.webp]]

**Trigger:** Inputs fall into distinct categories; one-size-fits-all degrades across them.

### 3. Parallelization
Multiple calls run concurrently; results aggregated programmatically.

*Sectioning* — different tasks at the same time:
![[resources/media/1736686ac91552ef6b415ba156af189a_MD5.webp]]

*Voting* — same task multiple ways, take consensus:
![[resources/media/646fe23f05e709641161f466ed1d72fd_MD5.webp]]

**Trigger:** Independent subtasks (sectioning) or consensus needed (voting). Graph independence is non-negotiable.

### 4. Orchestrator-Workers
Central orchestrator dynamically breaks task into subtasks at runtime; dispatches to workers; synthesizes.

![[resources/media/dc02a6edf6b6f8f79b853d7e809a476f_MD5.webp]]

**Trigger:** Open-ended problems where the subtask list cannot be predefined.

### 5. Evaluator-Optimizer
Generator produces output; evaluator scores against criteria; generator revises. Loop until evaluator passes.

![[resources/media/f482a9e111b046d9dea21e268a1f8dab_MD5.webp]]

**Trigger:** Clear evaluation criteria exist *and* iterative feedback measurably improves output.

---

## Six CCA Curriculum Patterns

### 6. Hub-and-Spoke
One coordinator; multiple specialist subagents; all inter-agent communication routed through coordinator. Subagents never talk directly.

![[resources/media/9db9e0c3bbb82dc68ac3f688868f175e_MD5.webp]]

**Trigger:** Multi-agent work where subagents must not share context. The **base architecture** for almost every multi-agent system. Context isolation is non-negotiable — subagents start with blank context windows.

### 7. Dynamic Adaptive Decomposition
Coordinator generates next subtask based on what prior subtasks revealed. Plans emerge during execution.

![[resources/media/6b186be725a46cfb1dc97e7f3c3d2b61_MD5.webp]]

**Trigger:** Unknown problem space; plan reveals itself as findings come in.

### 8. Sequential Prompt Chaining (Coordinator-Scoped)
Fixed, predictable sequence of subagent calls.

![[resources/media/9eb9e3a4c7fecbf9d599ad3a66c514b4_MD5.webp]]

**Trigger:** Multi-aspect work that decomposes into a known, predictable order.

### 9. Programmatic Prerequisites
Hooks block downstream tool calls until prerequisite steps complete. Prompt-based ordering has non-zero failure rate; programmatic gates have zero.

![[resources/media/12c8f5f88a035387e33067ad88ee360a_MD5.webp]]

**Trigger:** Ordering must be guaranteed. Cost of skipping the prerequisite is unacceptable.

### 10. Graceful Degradation with Partial Results
When a subagent fails, coordinator preserves completed work and annotates the gap rather than aborting.

![[resources/media/62db0763d7cc3b4628bc8e5f5b782773_MD5.webp]]

**Trigger:** Independent subtask failed but successful work should not be discarded.

### 11. Structured Handoff (Human Escalation)
When escalating, agent compiles structured context (typed fields) rather than dumping a transcript.

![[resources/media/97eeb06f18564de219f7516e9af76e33_MD5.webp]]

**Trigger:** Escalation to a human or downstream system that lacks conversation context.

---

## Composition Rules

- **Hub-and-spoke is the skeleton.** Evaluator-optimizer and iterative refinement are hub-and-spoke + evaluator + loop.
- **Orchestrator-workers vs. iterative refinement:** orchestrator-workers is driven by *exploration*; iterative refinement by *quality evaluation of a synthesized output*.
- **Parallelization composes inside orchestration:** the initial delegation wave in iterative refinement *is* parallelization.
- **Graceful degradation at loop boundary:** when iterative refinement exhausts `max_iterations`, it ships with gaps annotated — that's graceful degradation.

---

## Related

- [[JSON Schema Discipline]] — where structured outputs belong at pattern handoffs
- [[Harness Engineering]] — patterns are the actionable practice layer
