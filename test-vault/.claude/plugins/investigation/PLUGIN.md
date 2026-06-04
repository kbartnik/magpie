---
name: investigation
version: "1.0.0"
zone: dev/investigations/
commands: [/investigate]
skill: .claude/plugins/investigation/skill.md
schema-version: "1"
---

# Plugin: Investigation

## Testbed Role

This vault is a **testbed**, not a reference vault. Investigations here serve two purposes:

1. **Real questions about magpie** — "Does vault-tools handle concurrent writes?", "What happens when VAULT_PATH is unset?", "Is the session-start hook idempotent?" These are genuinely useful and appropriate for this vault.
2. **Workflow testing** — synthetic investigations to exercise the `/investigate` command: evidence accumulation, confidence tracking, conclusion writing, wiki synthesis.

Real questions about magpie's behavior are strongly preferred — they generate genuinely useful knowledge while testing the command.

**If an investigation produces a broadly applicable insight** (not magpie-specific), surface it for Nexus rather than recording it here.

---

## Zone Rules

Investigations accumulate evidence over time — always read the existing file before adding to it. The `confidence` field must be updated honestly when evidence changes. Adding evidence or updating confidence requires showing the proposed change before writing.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

`YYYY-MM-DD-question-slug.md` — date is when the investigation was **opened**.

Example: `2026-04-10-why-does-p99-spike-at-noon.md`

## Frontmatter Schema

```yaml
---
title: "Full question being investigated?"
type: investigation
question: "Full question being investigated?"
status: open              # open | answered | shelved
confidence: low           # low | medium | high
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---
```

## Status Lifecycle

- `open` → `answered`: when confidence reaches `high` and a Conclusion section is written
- `open` → `shelved`: when stalled > 30 days or explicitly deprioritized by the user

Answered and shelved are terminal. To reopen a question: create a new investigation file and link back to the old one in its Context section.

## Index Convention

List **only open** investigations in `wiki/index.md` under `## Dev`:
```
[[Investigation: Question]] — status: open
```

Answered investigations move to `wiki/syntheses/` or are noted inline in related wiki pages — they are no longer listed in the Dev index.

## Stale Threshold

30 days in `open` status without an update to the file.
