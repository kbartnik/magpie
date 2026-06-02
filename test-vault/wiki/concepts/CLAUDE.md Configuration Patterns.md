---
title: "CLAUDE.md Configuration Patterns"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources:
  - "archive/clippings/2026-05-29-claude-md-10-rules.md"
  - "archive/clippings/2026-05-29-karpathy-claude-md-8-rules.md"
related:
  - "Harness Engineering"
  - "Context Rot"
  - "Claude Code Memory Architecture"
tags:
  - claude-code
  - harness-engineering
  - cca-f
---

# CLAUDE.md Configuration Patterns

`CLAUDE.md` is the primary mechanism for locking persistent behavior into a Claude Code project. Because every session starts with zero memory, CLAUDE.md is the only way to preserve architectural decisions, coding standards, and operational constraints across sessions.

> "You create a plain text file named CLAUDE.md in your project root. Claude Code reads this file automatically at the start of every session. It becomes the permanent DNA of your project."

## The 8-Rule Architecture

The community-validated architecture uses exactly 8 rules — no more, because past ~200 lines, compliance reportedly plummets (the model pattern-matches "rules exist" rather than reading them).

### Foundation Rules (Karpathy's 4 — single-prompt failures)

**1. Think Before Coding**
State assumptions explicitly. Ask rather than guess. Push back when a simpler approach exists. Stop when confused.

**2. Simplicity First**
Minimum code that solves the problem. Nothing speculative. No abstractions for single-use code.

**3. Surgical Changes**
Touch only what you must. Don't "improve" adjacent code, comments, or formatting. Match existing style.

**4. Goal-Driven Execution**
Define success criteria and loop until verified, rather than blindly following rigid steps.

### Agentic Extension Rules (4 new — multi-step pipeline failures)

**5. Hard Token Budgets**
Establish a per-task token budget (e.g., 4,000 tokens) and per-session budget (30,000 tokens). If breached, force the agent to summarize and restart. Surfacing a breach is always better than silently overrunning API limits.

**6. Read Before You Write**
Before adding code to a file, read the file's exports, immediate callers, and shared utilities. "Looks orthogonal" is a dangerous assumption that leads to duplicate functions 30 lines apart.

**7. Checkpoint Multi-Step Operations**
After every significant step, the agent must summarize what was verified and what remains. It cannot continue from a state it cannot describe back to you.

**8. Fail Loud**
Default to surfacing uncertainty. If the agent skipped anything, or a test passed for the wrong reasons, it must fail loud and alert immediately. "Migration completed" is a lie if 30 records were silently skipped.

## The Persistent Memory Pattern

In addition to behavioral rules, CLAUDE.md can configure external memory files that the agent is instructed to maintain:

**MEMORY.md** — Architectural state log
After any significant decision regarding architecture, logic, or format, add an entry: date, what was decided, reasoning, alternatives rejected. Read this file at session start.

**ERRORS.md** — Hallucination loop breaker
When an approach fails more than 2 attempts, log the failed approaches and the final working solution. Check before attempting complex logic to avoid repeating known failures. This guards against the pattern where an agent tries the same failing approach on Friday that failed on Tuesday.

**Session Checkpoints**
On session end, write to MEMORY.md: what was completed, what is in progress, and the exact next step required.

## Stack Lock

Without a defined stack, Claude defaults to whatever framework appears most in training data. An architecture lock instruction:

```
Tech stack: always use these, never suggest alternatives unless I ask.
[Language, Framework, Package Manager, Database, Linter]
If a suggested tool conflicts with this stack, flag it but use the defined stack anyway unless explicitly overridden.
```

## Compliance Constraint

CLAUDE.md over ~200 lines loses effectiveness — the model pattern-matches "rules exist" rather than actually reading them. Keep the file to 8 high-impact rules; project-specific stack details can follow below a separator.

## Relationship to Other Mechanisms

CLAUDE.md configures **LLM behavior** (what it says and doesn't say). Hooks configure **runtime behavior** (what runs when). Together they form the harness layer:

- CLAUDE.md → shapes the model's reasoning and output style
- [[Claude Code Hooks]] → deterministic side-effects (auto-format, auto-test, auto-approve)
- External memory files (MEMORY.md, ERRORS.md) → cross-session state persistence

## Related

- [[Harness Engineering]] — CLAUDE.md is a harness artifact: behavioral constraints externalized into configuration
- [[Claude Code Memory Architecture]] — CLAUDE.md is one of the four memory layers; MEMORY.md/ERRORS.md are external files that extend it
- [[Context Rot]] — CLAUDE.md rules 5–8 directly address context rot symptoms (token overflow, silent failures, unbounded loops)
- [[Claude Code Hooks]] — complementary mechanism: hooks handle deterministic runtime actions while CLAUDE.md handles LLM behavior
