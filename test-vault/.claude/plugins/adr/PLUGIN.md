---
name: adr
version: "1.0.0"
zone: dev/adr/
commands: [/adr]
skill: .claude/plugins/adr/skill.md
schema-version: "1"
---

# Plugin: ADR

## Testbed Role

This vault is a **testbed**, not a reference vault. ADRs here serve two purposes:

1. **Workflow testing** — exercising the `/adr` command to validate creation, update, and supersession flows work correctly
2. **Magpie design decisions** — recording real architectural choices made during magpie CLI development (e.g. "Use vault-tools binary for file I/O", "Plugin discovery via PLUGIN.md manifest")

ADRs about magpie's own architecture are appropriate here — they're real decisions being made in this context. Generic synthetic scenarios are acceptable for workflow testing.

**Do not treat these ADRs as authoritative for anything outside magpie development.** Real project decisions belong in Nexus.

---

## Zone Rules

Never edit the body of an accepted ADR — only its `status` field can change, and only to `superseded`. Always check `dev/adr/` for existing ADRs on the same topic before creating a new one (grep by title keyword). Creating or editing any ADR file requires a plan presented and approved before writing.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

`ADR-NNNN-short-slug.md` where NNNN is zero-padded, sequential.

Before creating, list `dev/adr/` to find the next available number:
```bash
ls dev/adr/ | grep "^ADR-" | sort | tail -1
```

Example: `ADR-0003-use-postgres-for-sessions.md`

## Frontmatter Schema

```yaml
---
title: "Use X for Y"          # imperative sentence
type: adr
status: proposed              # proposed | accepted | superseded
decision-date: ""             # YYYY-MM-DD, set when accepted
created: YYYY-MM-DD
updated: YYYY-MM-DD
supersedes: []                # list of ADR numbers this replaces
superseded-by: ""             # ADR number that replaces this one
tags: []
---
```

## Status Lifecycle

- `proposed` → `accepted`: requires explicit user decision ("Let's accept this")
- `accepted` → `superseded`: only when a new ADR is created that replaces it; set `superseded-by` on the old ADR and `supersedes` on the new one

Accepted ADRs are immutable in content. Superseded ADRs are never deleted.

## Index Convention

List in `wiki/index.md` under `## Dev` section:
```
[[ADR-NNNN: Title]] — status: <status>
```

## Stale Threshold

14 days in `proposed` status without a decision — flagged by `vault-tools lint-gather`.
