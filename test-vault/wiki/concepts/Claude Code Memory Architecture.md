---
title: "Claude Code Memory Architecture"
type: concept
status: active
created: 2026-05-20
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-20-claude-permanent-memory.md"
  - "archive/clippings/2026-06-01-four-types-agent-memory.md"
related:
  - "Context Rot"
  - "Harness Engineering"
  - "Nexus Vault Template"
tags:
  - claude-code
  - memory
  - harness-engineering
  - cca-f
---

# Claude Code Memory Architecture

Claude's memory system is four distinct layers, each trading control for convenience at a different point on the spectrum. Understanding which layer to use is the practical skill.

## The Four Layers

| Layer | Who it's for | Control | Persistence |
|-------|-------------|---------|-------------|
| **Chat Memory Synthesis** | All users | Low — Claude decides what to keep | Auto-updated ~24h; immediate via explicit instruction |
| **Project Memory Spaces** | Pro/Max/Team/Enterprise | Medium — scoped to project | Isolated per project; supports past-chat search |
| **CLAUDE.md File-Based** | Claude Code users | High — you write it | Loaded unconditionally every session |
| **API Memory Tool** | API/developers | Full — you own the storage | R/W filesystem at `/memories`; persists across sessions |

## Layer 1: Chat Memory Synthesis

Claude runs extractive summarization over conversation history, identifying patterns and preferences worth persisting. Routine Q&A is discarded; professional context and stated preferences are retained.

**Key behavior:** Synthesis errs toward *useful professional and personal context* over *ephemeral task content*. It won't remember that you asked about photosynthesis. It will remember that you're an ML engineer using PyTorch.

**Trigger immediate update:** "Remember that I prefer concise technical explanations with code examples." No settings menu needed.

## Layer 2: Project Memory Spaces

Each Claude Project maintains isolated memory — your product launch planning doesn't bleed into your client work. Supports retrieving specific decisions from weeks ago on demand.

**Limit:** Project memory and standalone chat memory are siloed. Choose one home for a given workflow.

## Layer 3: CLAUDE.md (Most Important for Claude Code Users)

CLAUDE.md files sidestep the retrieval problem entirely — they are loaded **unconditionally** at every session start, no synthesis delay, no retrieval call.

**Hierarchy:** global (`~/.claude/CLAUDE.md`) → project root → subdirectory → gitignored personal override. More specific files win on conflicts.

**Practical limit:** Keep under 200 lines. Beyond that, adherence degrades meaningfully. Use `.claude/rules/` for modular overflow.

**The high-ROI pattern:** End every meaningful session with "Update CLAUDE.md with everything important you learned today." Over a month, the file becomes the onboarding document you always meant to write.

**Share it.** Don't add it to `.gitignore`. Your teammates' sessions improve when the file is shared.

## Layer 4: API Memory Tool

Client-side R/W filesystem at `/memories`. Claude makes tool calls (`create`, `read`, `str_replace`, `insert`, `delete`) and your application executes them locally. You own the storage, schema, and what persists.

**Designed to pair with compaction:** context compaction summarizes conversations server-side when the context window fills. The API Memory Tool ensures nothing critical is lost across that compaction boundary.

**Use case:** long-running agents that genuinely pick up where they left off — architecture decisions from session 1 are available in session 20.

## What Memory Captures (and What It Doesn't)

Memory captures **preferences and facts**, not **decisions and reasoning chains**.

Claude will remember: you prefer PostgreSQL.
Claude will NOT automatically remember: *why* you chose PostgreSQL, what the schema looked like, which trade-offs you weighed.

For decisions of that weight: explicit documentation in CLAUDE.md or the API Memory Tool is essential.

## Limits

- Memory doesn't travel — it lives on Anthropic's servers and doesn't follow you to ChatGPT or Gemini
- Free users get synthesis but not past-chat search (search requires Pro+)
- Compaction (auto-summarization) is a band-aid, not a substitute for Memory Tool in long-running agents (see [[Context Rot]])

## How This Vault Uses Memory

This vault implements its own memory system on top of these layers:
- **CLAUDE.md** — loads the vault operating schema unconditionally
- **context.md** — agent-written session state (focus, next-actions, parked ideas)
- **wiki/** — accumulated knowledge (equivalent to API Memory Tool but in Markdown)
- **auto memory** — Claude Code's built-in auto-memory writes to `~/.claude/projects/*/memory/`

See [[Nexus Vault Template]] for the full architecture.

## KOALA Mapping

The KOALA framework (see [[Agent Memory Architectures]]) classifies agent memory by cognitive function. Claude Code's four layers map directly:

| KOALA type | Claude Code layer |
|-----------|------------------|
| **Working** | Context window (the session itself) |
| **Semantic** | CLAUDE.md — always loaded, persistent factual/convention knowledge |
| **Procedural** | Skills — progressive disclosure; index always present, instructions on demand |
| **Episodic** | Auto-memory + API Memory Tool — distilled cross-session experience |

This mapping is useful for diagnosing gaps: if your agent repeats mistakes it should know, the semantic layer needs work. If it can't execute familiar workflows reliably, the procedural layer needs work. If it can't build on past sessions, the episodic layer is missing.

## CCA-F Connection

Memory architecture is the **Memory dimension** of the six-dimensional agent model from [[Harness Engineering]]. The harness manages memory so the model can focus on reasoning. The four layers map to different points on the control/convenience tradeoff that harness designers must navigate.

## See Also

- [[Context Rot]] — the failure mode that good memory management prevents
- [[Harness Engineering]] — memory is one of the six harness dimensions
- [[Nexus Vault Template]] — this vault's memory implementation
