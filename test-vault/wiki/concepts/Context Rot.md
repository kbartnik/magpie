---
title: "Context Rot"
type: concept
status: active
created: 2026-05-20
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-20-what-is-harness-engineering.md"
  - "archive/clippings/2026-05-29-rust-token-killer.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-3.md"
  - "archive/clippings/2026-06-01-no-vibes-allowed-context-engineering.md"
  - "archive/clippings/2026-06-01-software-fundamentals-ai-age-pocock.md"
related:
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
  - "Progressive Disclosure Architecture"
  - "LLM Mental Model"
  - "AI Without Illusions (Series)"
tags:
  - agentic-systems
  - failure-modes
---

# Context Rot

A documented failure mode in long-running agents: as the context window fills with stale, redundant, or low-signal tokens, model performance measurably decays — even though the window isn't exhausted.

Named in harness engineering literature as one of four LLM failure modes that harness design must compensate for. The others: context panic (skipping steps under pressure), lost-in-the-middle (under-attention to mid-context content), and U-shaped attention (generalization of lost-in-the-middle).

## Why It Happens

Context windows process tokens with attention mechanisms that don't distinguish "this token is still relevant" from "this token was relevant three hours ago." Old tool outputs, superseded plans, and redundant observations accumulate and compete with current working content for attention budget.

## The Dumb Zone

Dex Horthy (AI Engineer 2025) named the concrete threshold: beyond **~40% context fill**, performance degrades measurably — even before the window is exhausted. He calls this region the "dumb zone." Every tool call, every MCP response, every file read pushes toward it.

Corollary: if your agent has too many MCPs loaded, it does all its meaningful work inside the dumb zone and can never produce reliable results. The fix isn't better prompts — it's reducing what's in context.

## Intentional Compaction

The systematic response to the dumb zone: deliberately compress context at phase boundaries before starting the next phase.

**Research → Plan → Implement (RPI)** is a workflow built around this principle:
1. **Research phase**: sub-agents find exact files and line numbers; output is a compact research document — not the raw conversation
2. **Plan phase**: outline exact steps with file names, line snippets, and post-change test checkpoints; output is a plan file with code snippets — not the raw conversation
3. **Implement phase**: execute the plan with a small fresh context; the plan itself is the entire context handoff

At each phase transition the context is compacted. The parent agent never accumulates the exploratory search cost of the research phase.

## Harness Compensations

- **Working-memory discipline**: compress or drop old observations before they dilute the active context
- **Sub-agent isolation**: each subagent starts with a blank context, receiving only the typed payload relevant to its task — context control, not role anthropomorphism (see [[Agentic Workflow Patterns]])
- **Intentional compaction at phase boundaries**: RPI workflow enforces this structurally
- **Structured note-taking**: instead of raw tool outputs accumulating, agents write compressed structured summaries to external state
- **Context compression gates**: periodically replace low-signal context sections with structured summaries before rot sets in

## Architectural Source: Shallow-Module Codebases

Context rot has a structural cause that operates at the codebase level, not just the runtime level: **shallow-module codebases force exploration**.

John Ousterhout (*A Philosophy of Software Design*) distinguishes deep modules (simple interface, complex hidden implementation) from shallow modules (complex interface, little hidden complexity). AI coding tools create shallow-module codebases by default — many small files, many entry points, many cross-dependencies.

When an AI agent navigates a shallow-module codebase, it must read more files to understand the system. Each file read is context consumed. The exploration overhead fills the context window faster, pushing the agent into the dumb zone sooner.

The mechanism: `shallow modules → more files to read → more context consumed → faster rot onset → worse code output → more corrections needed → more context consumed`.

This is a self-reinforcing degradation loop. The remedy is architectural: invest in deep modules with simple interfaces. The AI then needs to read only the interface to reason about the module, not all of its dependencies.

See [[Vibe-Coding Anti-Pattern]] for the full argument that codebase architecture is a harness optimization.

## Tool-Layer Mitigations

Beyond harness design, context rot caused by verbose shell output can be addressed at the tool layer:

**RTK (Rust Token Killer)** — A Rust binary proxy that intercepts shell command output before it reaches the LLM. Four strategies: filtering (remove comments and blank lines), grouping (consolidate similar items), truncation (keep relevant portions of long output), deduplication (collapse repeated log lines). Benchmarked at 60–90% token reduction; `cargo test` with 262 passing tests: ~4,800 tokens → 11 tokens.

Limitation: RTK only applies to Bash shell commands. Claude Code's built-in tools (Read, Grep, Glob) bypass the hook and are not filtered.

**`/compact` command** — Triggers summarization of the current conversation, replacing history with a condensed summary. Recommended proactively at ~60% context fill (estimated 85% payload reduction). `/clear` resets entirely for a new task.

**Documentation loading (PDA)** — [[Progressive Disclosure Architecture]] addresses the most common source of context rot in skills: loading 50KB of documentation for a task that needs 5KB. Three-tier loading ensures only relevant documentation is ever in context.

## Three Failure Modes (from AI Without Illusions, Part 3)

A richer taxonomy of how context goes wrong, each with different causes and remedies:

**Context pollution** — the context window contains irrelevant, contradictory, or outdated information. The model doesn't know which parts matter and which are noise; it tries to use everything. Conflicting information produces averaged or arbitrarily biased output. Insidious because it doesn't produce obvious errors — it produces subtly wrong output.

**Context truncation** — when context exceeds the window limit, the tool silently drops or compresses earlier material. The model isn't ignoring a requirement you stated in turn 3; it literally cannot see it by turn 30. Most chat tools don't surface when this is happening.

**Context drift** — even without truncation, long conversations gradually shift focus. Each exchange builds on the previous one; the model's attention gravitates toward recent turns. The original specification is still present but functionally drowned out. Conversational scope creep.

**Lost in the middle** — the structural attention pattern where models attend most strongly to content at the beginning and end of context. A critical constraint buried in the middle of 50,000 tokens receives measurably less attention than the same constraint at position 1.

**The repo-context gap** — in coding workflows: the model never sees your whole codebase, only a tool-selected subset. "The model has read and understood my repository" is almost always false. It has read what the tool's retrieval logic chose to include. Quality of that selection determines quality of output.

### Practical Rules (from Part 3)

- Start new conversations for new tasks — clean context, free from accumulated noise
- Put critical information near the top or bottom — not buried in the middle
- Restate key constraints in long conversations — don't assume tracking
- Curate retrieved and file context aggressively — include the relevant, withhold the tangential
- Treat context as a budget — everything in it costs attention, not just tokens

## See Also

- [[Harness Engineering]] — the discipline built around preventing and compensating for this failure mode
- [[Agentic Workflow Patterns]] — hub-and-spoke isolation is a structural defense against context rot propagating across subagents
- [[Progressive Disclosure Architecture]] — the PDA pattern prevents documentation-loading as a source of context rot in skills
- [[LLM Mental Model]] — the underlying mechanics of why context management matters
- [[AI Without Illusions (Series)]] — Part 3 is the richest treatment of this topic
