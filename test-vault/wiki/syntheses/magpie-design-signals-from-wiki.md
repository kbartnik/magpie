---
title: "Magpie Design Signals from the Wiki"
type: synthesis
status: active
created: 2026-06-02
updated: 2026-06-02
sources: []
related:
  - "magpie"
  - "Harness Engineering"
  - "Agent Memory Architectures"
  - "Agentic Workflow Patterns"
  - "Claude Code Hooks"
  - "Progressive Disclosure Architecture"
tags:
  - magpie
  - harness-engineering
  - design
---

# Magpie Design Signals from the Wiki

What the vault's accumulated knowledge implies about magpie's design. Magpie is a **harness component** — the compiled interface layer between Claude Code and the vault. The [[Harness Engineering]] framework names what harnesses need to do, and most of those dimensions are currently unaddressed in magpie.

## The Six Harness Dimensions Applied

From [[Harness Engineering]]: `Agent = Model + Harness`. The harness does six things the model cannot do alone:

| Dimension | Current magpie state | What's missing |
|-----------|---------------------|----------------|
| **Context assembly** | `session-start` briefing | Signal/noise ratio not validated; format evolved empirically |
| **Tool contracts / validators** | Exit 2 on `archive-file` slug conflicts | Not systematized — only one subcommand has a hard block |
| **Memory / durable state** | `log-append` (append-only prose) | No machine-queryable log; no structured fields for analytics |
| **Observability** | None | No drift detection, no evaluation gates, no structured output |
| **Recovery** | None | No rollback for archive operations; no dry-run mode at binary level |
| **Orchestration** | n/a (single-vault, 1.0) | Deferred; will require design when multi-vault arrives |

## Exit Code Design Should Be Systematic

The wiki distinguishes two exit semantics (see Claude Code Hooks):

- **Exit 2** = hard block — the action does not run; Claude cannot override
- **Exit 1** = logged warning — Claude proceeds; you have logging, not enforcement

Magpie already uses this for `archive-file`. It should be a **design rule across all subcommands**, not an accident of one command's history.

The four-safety-levels escalation pattern (also from Claude Code Hooks):
1. Dry-run — show what would change
2. Diff preview — exact diff before applying
3. Explicit approval gate — require `y/n`
4. Hard block — exit 2, unconditional

Magpie's `--force` flag is essentially jumping from 2 to 4. The intermediate levels (a `--dry-run` flag and a diff-preview mode) are worth adding before 1.0.

## Structured Output Is an Interface Contract

[[Agentic Workflow Patterns]] is explicit: handoffs should be "typed payloads, not transcript dumps." Progressive Disclosure Architecture adds: "Scripts fail hard with structured codes; the AI interprets those codes into user guidance, alternative lookups, and retry logic."

Magpie's current CLI output mixes human prose and structured fields. The clean design:
- **stdout**: machine-parseable (structured text or `--json` flag)
- **stderr**: human-readable error messages
- **exit codes**: semantic, documented, consistent

Structured error codes let Claude reason about recovery without parsing prose. "Every error path leads to either user guidance or automated recovery — no dead ends."

## The Deterministic Split Is Already the Architecture — Name It

From Agent Memory Architectures (Fat Skills section):

> **Deterministic split**: latent work (synthesis, pattern recognition) stays with the LLM; deterministic work (SQL, file ops, calculations) routes to code.

Magpie *is* this principle compiled into a binary. Making the split explicit in documentation and plugin contracts clarifies what belongs in magpie vs. what belongs in a skill. If a task requires reasoning, it's the LLM's job. If it's mechanical, it's magpie's job. This should be a stated design principle, not just emergent behavior.

## Plugin Metadata Is the Router

From Agent Memory Architectures (Fat Skills / resolver insight):

> Skill descriptions *are* the router. The model reads descriptions and matches intent automatically. No explicit routing code needed.

Applied to magpie's registered plugin model: the `--magpie-describe` metadata *is* how Claude discovers and dispatches to plugins. Investing in rich, specific plugin descriptions is more valuable than investing in routing logic. Plugin metadata quality is a first-class concern.

## Fat Skills / Cron Autonomy Direction

The wiki notes explicitly: "the magpie project may push toward the Fat Skills direction (cron-based autonomous operations, sonar split)." The Fat Skills model from Agent Memory Architectures offers three concrete mechanisms worth evaluating:

- **Always-on skills** — fire on every message; e.g., passive entity linker, unlinked-mention detector
- **Cron skills** — run autonomously on schedule, filing results to `reports/`; e.g., vault health checks, orphan detection, stale link sweeping
- **Scheduled background work** — agent works while you sleep; results surface at next session start

This is the autonomy direction that sonar (extracted from `go-extract`) enables. Sonar as a plugin becomes the reference for what cron-style magpie plugins look like.

## Observability Is a Harness Responsibility

Magpie could emit structured log entries (not just human-readable prose to `wiki/log.md`) that a future `magpie health` or `magpie stats` subcommand could aggregate. The [[Harness Engineering]] observability dimension includes drift detection and evaluation gates — things that require machine-readable history to compute.

Concrete: `log-append` could accept a `--structured` flag that also writes a JSON sidecar to `.magpie/events/YYYY-MM-DD.jsonl`. This costs nothing at write time and unlocks analytics later.

## What the Wiki Is Thin On

- **Go CLI output design for LLM parseability** — general advice exists but no Go-specific research
- **Plugin contract versioning** — the registered model is documented but no schema discipline page exists for plugin metadata format
- **Multi-vault scenarios** — deferred to post-1.0 but no research captured for when to revisit

## See Also

- [[magpie]] — project entity with implementation plan and current status
- [[Harness Engineering]] — the six-dimension framework; magpie is one harness component
- Agent Memory Architectures — decision framework for memory; Fat Skills as the autonomy direction
- [[Agentic Workflow Patterns]] — Programmatic Prerequisites and Structured Handoff patterns
- Claude Code Hooks — exit code semantics, safety levels, AI-checking-AI pattern
- Progressive Disclosure Architecture — three-tier output design; AI resilience layer
- rtk — production-scale PreToolUse hook; concrete harness implementation to study
- Graphify — PreToolUse injection pattern; another harness component in the same ecosystem
