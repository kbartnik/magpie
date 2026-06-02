---
title: "magpie-stats"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [phase-5, plugin, magpie-stats]
---

# magpie-stats

First bundled plugin. Read-only vault statistics — walks the vault and reports counts, sizes, and last activity. Lives in `plugins/magpie-stats/`. Primary purpose: validate the runtime plugin contract with a real implementation before magpie-claude depends on it.

## Learning Focus

- Building a cobra CLI that works as a magpie plugin (reads `MAGPIE_VAULT`, owns its own subcommand tree)
- `filepath.WalkDir` for counting and sizing
- Confirming `syscall.Exec` dispatch works end-to-end

## Commands

```
magpie stats             # vault summary: note counts by section, inbox depth, archive size, last log entry
magpie stats --json      # same, as JSON to stdout
```

## What it reports

- Notes by section (`wiki/`, `inbox/`, `archive/`, `dev/projects/`)
- Current `inbox-count` and actual inbox file count (flags drift)
- Archive size by subdirectory
- Last log entry timestamp and text
- Vault name and schema version from `context.md`

## You Drive

The cobra structure for a plugin — it's a standalone binary with its own root command, but `argv[0]` is `magpie-stats` when dispatched via magpie and `stats` when called as a subcommand. How does cobra handle this? What does `--help` look like in each case?
