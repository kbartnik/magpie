---
name: learning
version: "1.0.0"
zone: dev/learning/
commands: [/learn, /review]
skill: .claude/plugins/learning/skill.md
schema-version: "1"
---

# Plugin: Learning

## Zone Rules

Each topic has an index file and time-stamped session notes. Always read the track index before adding a session note. Session notes are append-only artifacts — never edit a past session note.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

- Topic index: `dev/learning/<topic-slug>/index.md`
- Session notes: `dev/learning/<topic-slug>/YYYY-MM-DD.md`

Example:
```
dev/learning/rag-chunking/index.md
dev/learning/rag-chunking/2026-05-19.md
```

## Frontmatter Schema

**Track index:**
```yaml
---
title: ""
type: learning
topic: ""
status: active             # active | paused | complete
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---
```

**Session note:**
```yaml
---
title: ""
type: learning-note
topic: ""
source: ""                 # what was read/watched/practiced this session
created: YYYY-MM-DD
---
```

## Status Lifecycle

- `active` → `paused`: user decision (taking a break from this topic)
- `active` | `paused` → `complete`: user decision (topic sufficiently understood)

Complete tracks are candidates for `wiki/syntheses/` pages.

## Index Convention

Do not list learning tracks in `wiki/index.md` — they are too granular. The track's own `index.md` is its index. Use `/review` to synthesize a track into a wiki page.

## Stale Threshold

30 days in `active` status without a new session note — flagged by `vault-tools lint-gather`.
