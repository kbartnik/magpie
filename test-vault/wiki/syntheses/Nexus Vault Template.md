---
title: "Nexus Vault Template"
type: synthesis
status: active
created: 2026-05-20
updated: 2026-05-20
sources:
  - "archive/docs/2026-05-20-vault-build-spec.md"
  - "archive/docs/2026-05-20-vault-build-plan.md"
related:
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
tags:
  - vault
  - project
  - architecture
---

# Nexus Vault Template

An open-source personal knowledge vault architecture powered by Claude Code.

**Project:** `nexus-vault-template`
**GitHub:** https://github.com/kbartnik/nexus-template
**Local:** `~/Source/nexus-template`
**Testbed:** This vault (`~/Documents/Obsidian/Nexus`) is the implementation and testbed for the project

**Spec:** `archive/docs/2026-05-20-vault-build-spec.md`
**Implementation plan:** `archive/docs/2026-05-20-vault-build-plan.md`

---

## What It Is

A two-axis knowledge system template:

- **Wiki axis** — reading, research, accumulated knowledge; articles and papers ingested and synthesized into an interlinked wiki
- **Dev axis** — engineering work; architectural decisions, post-mortems, investigations, learning tracks

The agent handles all maintenance. The human handles sourcing, direction, and approval of writes. This vault is both the reference implementation of the template and the working knowledge base used to develop and refine it.

---

## Core Architecture Principles

These are the non-negotiable design decisions from the spec. All generated files must reflect them.

### Go binary for mechanics, LLM for reasoning

The `vault-tools` binary at `.claude/tools/vault-tools` owns all deterministic operations:
- File counting, moving, parsing frontmatter, appending structured text
- All subcommands are unit-tested (`go test ./...`)
- The LLM never implements mechanical logic or acts as a shell wrapper

The division: **binary decides nothing; LLM decides everything; binary executes.**

This is [[Harness Engineering]] applied to a personal knowledge system — the same model/harness split that SWE-agent demonstrated: mechanical sympathy for the LLM substrate means keeping the model focused on reasoning, not file I/O.

### Plan-before-write

Every file-modifying operation presents a plan and requires explicit approval before the binary executes. Exceptions: log appends, inbox count updates, session-end timestamps.

### Plugin extensibility

Each knowledge domain is a self-contained plugin bundle in `.claude/plugins/<name>/`:
- `PLUGIN.md` — manifest, zone rules, frontmatter schema
- `skill.md` — domain reasoning skill

Plugin discovery is automatic at session start; no edits to CLAUDE.md needed to add a domain.

### Structured lint output

`vault-tools lint-gather` produces pipe-delimited structured output. The LLM reasons over findings; it does not gather them.

### Log `Open:` as live backlog

Every log entry has an `Open:` field. `/lint` and `/resume` surface unresolved Open: items from recent entries. Not buried history — an active follow-up queue.

---

## Component Map

```
vault/
├── inbox/           ← Zero-friction capture; agent reads, moves after approval
├── archive/         ← Write-once; clippings/, papers/, books/, daily/, ideas/, docs/
├── wiki/            ← Agent-owned; concepts/, entities/, syntheses/, questions/
├── dev/             ← Plugin zones; adr/, debriefs/, investigations/, learning/
└── .claude/
    ├── tools/vault-tools  ← Go binary (all mechanical ops)
    ├── skills/            ← session-context, source-annotation
    ├── commands/          ← 15 LLM command files (ingest, query, lint, wrap...)
    └── plugins/           ← 4 domain plugins (adr, debrief, investigation, learning)
```

---

## Installed Plugins

| Plugin | Zone | Commands | Stale threshold |
|--------|------|----------|----------------|
| adr | `dev/adr/` | `/adr` | 14 days in `proposed` |
| debrief | `dev/debriefs/` | `/debrief` | none |
| investigation | `dev/investigations/` | `/investigate` | 30 days in `open` |
| learning | `dev/learning/` | `/learn`, `/review` | 30 days in `active` |

---

## Binary Subcommands

| Subcommand | Purpose | Hook-safe? |
|-----------|---------|-----------|
| `capture <text>` | Append to today's inbox file | No |
| `session-start` | Emit session briefing | Yes (exit 0 always) |
| `session-end` | Update last-session, inbox-count | Yes |
| `update-inbox-count` | Sync context.md inbox count | Yes (async) |
| `inbox-list` | List inbox contents | No |
| `archive-file <src> <type> <slug>` | Move file to archive with frontmatter | No |
| `log-append <op> <title> <files> <summary> <open>` | Append to wiki/log.md | No |
| `lint-gather` | Structured vault health report | No |

---

## Backlog Structure

Backlog items live under `<project>/backlog/` in the source repo and are referenced in a backlog dashboard `.base` file (Obsidian Bases). Check `~/Source/nexus-template` for the authoritative backlog when planning template changes.

## Feature Ideas

- **Scopable slash commands** — `/capture` (and others) could accept a `@project` scope argument to route captures to a project's backlog rather than the vault-level inbox.

## Implementation Status

This vault (`~/Documents/Obsidian/Nexus`) serves as the reference implementation. The template source lives at `~/Source/nexus-template`. Implementation plan covers 22 tasks: directory structure → core static files → settings → Go binary (TDD: tests before source) → skills → commands → plugins → validation.

Open questions from implementation:
- Was the full 22-task plan completed? Which tasks remain?
- Were the final validation assertions (Task 21/22) all green?
- How does the testbed vault feed changes back into the template repo?

---

## See Also

- [[Harness Engineering]] — the architectural principle underlying the binary/LLM split
- [[Agentic Workflow Patterns]] — session hooks and plugin zones implement harness patterns
- [[Nexus Vault]] — this vault; the reference implementation and testbed
