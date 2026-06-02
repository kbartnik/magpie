---
title: "Context"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [phase-7, context]
---

# Context

Vault state display. Reads `context.md` frontmatter and body sections and presents them as a unified view. The `--json` flag makes it machine-readable for LLM consumption.

## Learning Focus

- Parsing markdown body sections (find heading, extract content until next heading)
- Dual output modes: human-readable vs. `--json` from the same data structure
- stdout/stderr discipline: all output goes to stdout, diagnostics to stderr

## Commands

```
magpie context status          # display vault state: inbox count, current focus, next actions
magpie context status --json   # same, as JSON to stdout
```

## What it reads

- `inbox-count`, `vault-name`, `schema` from frontmatter
- `## Current Focus` and `## Next Actions` body sections (user-maintained)
- File timestamps via `os.Stat()` on `context.md`

## You Drive

The section parser — given `## Current Focus` as the target heading, find it in the markdown body and extract everything until the next `##` heading (or end of file). What's the simplest implementation that handles empty sections, missing sections, and multi-paragraph content?

## Note

`context park` is deferred to [[magpie-claude — Phase B — Session + Park|magpie-claude Phase B]]. The distinction between parking and capturing isn't meaningful without session context.
