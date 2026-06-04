---
name: debrief
version: "1.0.0"
zone: dev/debriefs/
commands: [/debrief]
skill: .claude/plugins/debrief/skill.md
schema-version: "1"
---

# Plugin: Debrief

## Testbed Role

This vault is a **testbed**, not a reference vault. Debriefs here serve two purposes:

1. **Real incidents from magpie development** — build failures, vault-tools bugs, incorrect command behavior, test regressions. These are preferred: real incidents produce the best learning.
2. **Workflow testing** — synthetic incident scenarios to validate the `/debrief` command: elicitation flow, blameless framing, wiki concept extraction.

The blameless format requirement is the same regardless of purpose — it's what we're testing.

**Debriefs here do not belong in Nexus.** They are either magpie-specific or synthetic; neither is general-purpose knowledge.

---

## Zone Rules

Blameless format required — the "What Didn't Work" section must not name individuals. Requires plan + approval before writing any file. Use the incident/event date for the filename, not today's date.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

`YYYY-MM-DD-short-slug.md` where the date is the **incident date** (not today's date).

Example: `2026-04-15-api-outage-cascade.md`

## Frontmatter Schema

```yaml
---
title: ""
type: debrief
status: draft              # draft | final
incident-date: YYYY-MM-DD
severity: low              # low | medium | high | critical
duration-minutes: 0        # optional — how long the incident lasted
tags: []
related-adrs: []           # ADR numbers related to this incident
---
```

## Status Lifecycle

- `draft` → `final`: only on explicit user declaration ("mark this final" or "this is done")

There is no expiry — debriefs do not go stale.

## Index Convention

List in `wiki/index.md` under `## Dev`:
```
[[Debrief: Title]] — <incident-date>, severity: <severity>
```

## Stale Threshold

None — debriefs do not go stale.
