---
title: "Agentic Workflow Patterns"
type: concept
status: active
created: 2026-05-20
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-20-eleven-agentic-patterns.md"
  - "archive/clippings/2026-05-20-harness-engineering-memory-gap.md"
  - "archive/clippings/2026-06-01-no-vibes-allowed-context-engineering.md"
related:
  - "JSON Schema Discipline"
  - "Harness Engineering"
tags:
  - agentic-systems
  - patterns
---

# Agentic Workflow Patterns

A catalog of 11 patterns for multi-agent systems. The skill is matching the *trigger phrase* in a scenario to the correct pattern — not memorizing names.

## The Diagnostic Move

When you read a scenario description: find the trigger phrase first. The trigger picks the pattern.

## Pattern Reference

| Pattern | Trigger phrase | Core shape |
|---------|---------------|------------|
| **Prompt Chaining** | "Steps are known in advance" | Linear pipeline, fixed before execution |
| **Routing** | "Distinct input categories" / "cost optimization" | Classifier dispatches to specialized paths |
| **Parallelization (Sectioning)** | "Independent subtasks running together" | Fan-out, programmatic aggregation |
| **Parallelization (Voting)** | "Consensus through diverse perspectives" | Same task N ways, majority wins |
| **Orchestrator-Workers** | "Plan emerges from exploration" | Dynamic task discovery, open-ended |
| **Evaluator-Optimizer** | "Output needs to be checked and revised" (single agent) | Generate → evaluate → revise loop |
| **Hub-and-Spoke** | "Coordinator orchestrating specialist subagents" | Base architecture; subagents never talk to each other |
| **Dynamic Adaptive Decomposition** | "Unknown problem space" / "plan reveals itself" | Orchestrator-workers with emphasis on exploration trigger |
| **Sequential Prompt Chaining** | "Multi-aspect work, known predictable order" | Coordinator-level chaining of subagent calls |
| **Programmatic Prerequisites** | "Ordering must be guaranteed" / "cost of skipping is unacceptable" | Hook gates enforce ordering deterministically |
| **Graceful Degradation** | "Independent failure should not waste good work" | Coordinator preserves partial results, annotates gaps |
| **Structured Handoff** | "Handing off to a human or external system" | Typed payload, not transcript dump |

## Common Wrong-Answer Distractors

- "Fixed five-step pipeline" offered when the trigger says *exploration* → that's chaining applied to a dynamic decomposition problem
- "Spawn more subagents" offered when quality iteration is needed → that's parallelization applied to an evaluator-optimizer problem
- Routing vs. orchestration: routing makes a *one-shot decision* on a *finite closed set*; orchestration makes *repeated decisions* on an *open problem space*

## Iterative Refinement (Multi-Agent Evaluator-Optimizer)

Iterative refinement = hub-and-spoke base + evaluator-optimizer feedback loop:

```
Coordinator → Researcher(s) [parallel]
           → Synthesizer
           → Evaluator → if gaps: spawn targeted ResearchTask(s) → loop
                       → if pass: ship
```

The `if verdict.pass_: break` line is load-bearing. Without a structured evaluator output, the loop has no reliable exit condition. See JSON Schema Discipline.

## Two-Agent Architecture (Anthropic Pattern)

For long-running development tasks:
- **Initializer agent**: runs once, sets up `init.sh`, `feature_list.json` (JSON, not Markdown), `claude-progress.txt`, git repo
- **Coding agents**: each session reads progress log → checks feature list → runs `init.sh` to verify state → implements one feature → commits → updates `passes: true` on completed feature
- **JSON structural gravity**: feature list in JSON prevents agents from rewriting requirements or declaring victory prematurely; a `passes: false` boolean resists casual override in ways a Markdown checkbox does not

## RPI: Research → Plan → Implement

A workflow composition pattern for context-managed coding in complex codebases (Dex Horthy, AI Engineer 2025). Built on Sequential Prompt Chaining with compaction gates at each phase boundary.

**Core principle:** Every phase transition compacts the context. The phase's *output document* is the handoff — not the raw conversation.

| Phase | Output | What's in context at start |
|-------|--------|---------------------------|
| Research | Compact research doc: exact files, line numbers, relevant code excerpts | Blank or minimal — sub-agents do the heavy searching |
| Plan | Plan file: exact steps, file names, line snippets, post-change test checkpoints | Research doc only |
| Implement | Committed code | Plan file only |

**Sub-agents in RPI are context controllers, not role stand-ins.** Don't spawn a "frontend agent" and a "backend agent." Spawn a sub-agent to do expensive codebase-reading, have it return a single succinct finding, then discard its context. The parent agent stays in the smart zone throughout.

**Mental alignment:** Plan files serve a second purpose beyond execution. Sharing plans (not just diffs) on PRs gives reviewers the same journey as the implementer — the *why* and *order* of changes, not just the wall of green text. As shipping velocity increases, plan-level review is how technical leads maintain understanding without reading every line.

**The leverage hierarchy (from highest to lowest):**
1. Research quality — a wrong assumption here cascades into 100x bad code
2. Plan review — catching a bad step before implementation, not after
3. Implementation — by the time a good plan exists, execution is mostly mechanical

**Don't outsource the thinking.** AI amplifies the thinking you bring. There is no prompt that replaces reading and approving the plan.

> Note: Horthy acknowledges "RPI" will likely undergo Semantic Diffusion once widely adopted. The stable principle is *compaction at phase boundaries*, not the acronym.

## Pattern Composition

Patterns compose. A real system might:
- Route on intake (Routing)
- Use hub-and-spoke for the selected workflow (Hub-and-Spoke)
- Decompose dynamically within it (Orchestrator-Workers)
- Apply iterative refinement on the synthesized output (Evaluator-Optimizer)
- Fall back to graceful degradation when a subagent fails (Graceful Degradation)

## Pattern Diagrams

Visual reference for each pattern's shape. The shape is what makes patterns recognizable in scenarios — match the diagram to the trigger.

### Prompt Chaining
*Trigger: steps are known in advance*

![[resources/media/809c5b04973e9ef33bb651551dc9a41e_MD5.webp]]

### Routing
*Trigger: distinct input categories*

![[resources/media/129d60b226da713f7d5fa0aab3153a64_MD5.webp]]

### Parallelization — Sectioning
*Trigger: independent subtasks running together*

![[resources/media/1736686ac91552ef6b415ba156af189a_MD5.webp]]

### Parallelization — Voting
*Trigger: consensus through diverse perspectives*

![[resources/media/646fe23f05e709641161f466ed1d72fd_MD5.webp]]

### Orchestrator-Workers
*Trigger: plan emerges from exploration*

![[resources/media/dc02a6edf6b6f8f79b853d7e809a476f_MD5.webp]]

### Evaluator-Optimizer
*Trigger: output needs to be checked against criteria and revised (single-agent)*

![[resources/media/f482a9e111b046d9dea21e268a1f8dab_MD5.webp]]

### Hub-and-Spoke
*Trigger: coordinator orchestrating specialist subagents*

![[resources/media/9db9e0c3bbb82dc68ac3f688868f175e_MD5.webp]]

### Dynamic Adaptive Decomposition
*Trigger: unknown problem space*

![[resources/media/6b186be725a46cfb1dc97e7f3c3d2b61_MD5.webp]]

### Sequential Prompt Chaining (Coordinator-Scoped)
*Trigger: multi-aspect work, known predictable order*

![[resources/media/9eb9e3a4c7fecbf9d599ad3a66c514b4_MD5.webp]]

### Programmatic Prerequisites
*Trigger: ordering must be guaranteed deterministically*

![[resources/media/12c8f5f88a035387e33067ad88ee360a_MD5.webp]]

### Graceful Degradation
*Trigger: independent failure should not waste good work*

![[resources/media/62db0763d7cc3b4628bc8e5f5b782773_MD5.webp]]

### Structured Handoff
*Trigger: handing off to a human or external system*

![[resources/media/97eeb06f18564de219f7516e9af76e33_MD5.webp]]

---

## See Also

- JSON Schema Discipline — which patterns require schemas, which benefit, which are harmed
- [[Harness Engineering]] — the broader discipline these patterns live within
- Context Rot — failure mode that programmatic prerequisites and clean-state disciplines prevent
- IDSD — methodology for structuring what the human brings to these patterns: intent (what is wanted) and expectations (what done means), so the agent fills the implementation gap rather than the requirements gap
- Agentic Identity and Zero Trust — the security layer all tool-use patterns require; the last mile problem breaks Zero Trust when agents connect to legacy backends via API keys
