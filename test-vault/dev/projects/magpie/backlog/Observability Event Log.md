---
title: "Observability Event Log"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: medium
due: ""
created: 2026-06-02
updated: 2026-06-02
tags: [post-1.0, observability, events]
---

# Observability Event Log

Every vault operation writes to the human-readable `wiki/log.md`. That log is for humans — prose-heavy, unsuitable for machine analytics. This item adds a machine-readable sidecar that enables vault health checks, stats, and trend analysis without requiring LLM parsing of prose.

## Design

Every vault operation emits a structured JSON event to `.magpie/events/YYYY-MM-DD.jsonl` (one JSON object per line, append-only):

```json
{
  "ts": "2026-06-02T14:23:00Z",
  "op": "archive",
  "files": ["inbox/2026-06-01-some-note.md"],
  "dest": "archive/clippings/2026-06-01-some-note.md",
  "exit": 0,
  "duration_ms": 42
}
```

## Append-Only Discipline

Same discipline as `wiki/log.md` — events are never edited after being written. The JSONL format (one object per line) makes appending cheap and parsing line-by-line trivial.

## What This Enables

- `magpie stats` — counts, sizes, last activity, operations per day — all computed from the event log without LLM
- `magpie health` — detect anomalies (no archive ops in 30 days, inbox growing unbounded)
- Cron operations (see [[Autonomous Cron Operations]]) can use event history to avoid redundant work
- Audit trail for hook-triggered actions — exit code 2 blocks are recorded even though no wiki log entry is written

## Relationship to wiki/log.md

The event log and the wiki log are complementary, not redundant. `wiki/log.md` is human-authored synthesis (what the wiki learned). The event log is mechanical telemetry (what happened). Don't conflate them.

## Prerequisites

None — can be added to any operation in Phase 1 or later as an opt-in side effect. The file format is stable from day one: append a JSON line, never rewrite.
