---
title: "Init"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [phase-3, init, setup]
---

# Init

Vault and project setup commands. `init vault` bootstraps a new vault from scratch; `init project` scaffolds a project within an existing vault. Both are idempotent — safe to re-run.

## Learning Focus

- Idempotent setup: check before creating, don't clobber existing files
- Interactive prompts: reading user input cleanly without a library
- YAML generation: writing structured frontmatter programmatically

## Commands

```
magpie init vault              # create .magpie/, context.md, required dirs
magpie init vault --upgrade    # migrate context.md to current schema version
magpie init project            # scaffold a project in dev/projects/
```

## What `init vault` creates

```
.magpie/
  config.yaml
context.md          # schema 1 frontmatter + body section headers
wiki/log.md
inbox/
archive/
  clippings/ papers/ books/ daily/ ideas/ docs/
dev/projects/
```

## You Drive

The migration chain — `migrate1to2()`, `migrate2to3()` chained to bring any vault to current. How do you structure this so adding a new migration is a one-function addition? How do you handle a vault with no `schema` field at all (pre-magpie)?

Idempotency — when a file already exists, do you skip silently, print a notice, or diff and warn?
