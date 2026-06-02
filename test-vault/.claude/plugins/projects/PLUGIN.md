---
name: projects
version: "1.0.0"
zone: dev/projects/
commands: [/project]
skill: .claude/plugins/projects/skill.md
schema-version: "1"
---

# Plugin: Projects

## Zone Rules

Each project lives in its own subdirectory under `dev/projects/`. Always read the project index before adding session logs or backlog items. Backlog items are individual files — never inline them as list entries in `index.md`. The `backlog.base` file is auto-generated at project creation; do not hand-edit it.

Creating or modifying a project index requires showing the proposed change before writing.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

- Project index: `dev/projects/<slug>/index.md`
- Session log: `dev/projects/<slug>/YYYY-MM-DD.md` — date is the session date
- Backlog item: `dev/projects/<slug>/backlog/<Title>.md` — free-form title as filename
- Project doc: `dev/projects/<slug>/docs/<Title>.md` — free-form title as filename
- Backlog view: `dev/projects/<slug>/backlog.base`
- Dashboard: `dev/projects/dashboard.base`

Example:
```
dev/projects/nexus/index.md
dev/projects/nexus/2026-05-19.md
dev/projects/nexus/backlog/Add projects plugin.md
dev/projects/nexus/backlog.base
dev/projects/magpie/docs/Magpie Design.md
```

## Frontmatter Schemas

**Project index** (`type: project`):
```yaml
---
title: ""
type: project
slug: ""
status: active             # active | paused | completed | archived
repo: ""                   # optional URL
tech-stack: []
created: YYYY-MM-DD
updated: YYYY-MM-DD
related-adrs: []
related-investigations: []
tags: []
---
```

**Backlog item** (`type: backlog-item`):
```yaml
---
title: ""
type: backlog-item
project: ""                # slug, e.g. "nexus"
status: todo               # todo | in-progress | done | dropped
priority: medium           # low | medium | high
due: ""                    # optional YYYY-MM-DD
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---
```

**Session log** (`type: project-log`):
```yaml
---
title: ""
type: project-log
project: ""
created: YYYY-MM-DD
---
```

## Status Lifecycle

**Project:**
- `active` → `paused`: user decision (taking a break)
- `active` | `paused` → `completed`: user decision (work done)
- `completed` → `archived`: after a cooling-off period or explicit decision

**Backlog item:**
- `todo` → `in-progress` → `done`: normal flow
- `todo` | `in-progress` → `dropped`: item no longer relevant

Session logs are append-only artifacts — never edit a past session log.

## Index Convention

List **only active and paused** projects in `wiki/index.md` under `## Dev`:
```
[[Project: Name]] — status: active
```

Completed and archived projects are removed from the Dev index.

## Stale Threshold

30 days without a session log in `active` status — flagged by `vault-tools lint-gather`.
