---
title: "Autonomous Cron Operations"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-06-02
updated: 2026-06-02
tags: [post-1.0, cron, autonomy, fat-skills]
---

# Autonomous Cron Operations

Fat Skills-style background operations that file results without being asked. The model: cron fires a magpie subcommand, it does deterministic work, it writes findings to `.magpie/reports/`, the next session-start briefing surfaces anything notable.

The LLM doesn't need to be running for this work to happen. Vault health maintains itself.

## Candidate Operations

| Operation | Trigger | Output |
|-----------|---------|--------|
| Orphan detection | Daily | List of wiki pages with no inbound links |
| Stale link sweep | Daily | Dead wikilinks (target page doesn't exist) |
| Inbox age warning | Daily | Items in inbox older than N days |
| Archive without wiki page | Weekly | Archived files with no corresponding wiki page |
| Entity linking | On demand | Unlinked mentions of known wiki pages in prose |

Each operation is a `magpie cron <name>` subcommand — deterministic, no LLM, testable in isolation.

## Output Convention

Results land in `.magpie/reports/<name>-YYYY-MM-DD.json`. Session-start hook reads reports newer than last session and appends notable findings to the briefing. Reports older than 30 days are pruned automatically.

## Relationship to Sonar

`sonar` (the extracted `go extract` plugin) is the first candidate for a cron-capable plugin — it can run extraction passes against new archive files without user prompting. This item defines the cron contract that `sonar` and future plugins implement.

## Prerequisites

- Plugin system (Phase 4) — cron operations are implemented as plugins or plugin subcommands
- Observability Event Log ([[Observability Event Log]]) — cron analytics need event history to avoid redundant work and to surface trends

## Design Pressure Needed

The cron protocol (how cron operations are registered, scheduled, and what they're allowed to touch) needs real plugins to design well — same reason the event bus was deferred. Start with one or two hardcoded `magpie cron` subcommands, validate the output convention, then generalize to a plugin-registered cron contract.
