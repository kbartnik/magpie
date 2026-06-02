---
title: "Lint"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [phase-2, lint, validation]
---

# Lint

Vault health checker. Validates structure, detects stale schema versions, and reports issues in both human-readable and machine-readable (`--json`) formats. The `--json` flag establishes the output pattern used by all subsequent machine-readable commands.

## Learning Focus

- Recursive directory traversal with `filepath.WalkDir`
- Regex compilation and reuse
- Designing `--json` output: when is JSON the right default vs. an opt-in flag?

## Commands

```
magpie lint            # validate vault structure, print human-readable report
magpie lint --json     # same, output as JSON to stdout
```

## What to validate

- `.magpie/` directory exists at vault root
- `context.md` present with valid frontmatter
- `context.md` schema version matches current (warn if stale, suggest `magpie init vault --upgrade`)
- Required directories present: `inbox/`, `archive/`, `wiki/`
- `wiki/log.md` present
- No orphaned files in archive (files without frontmatter)

## You Drive

The `--json` output design — what does the JSON structure look like? Should warnings and errors be separate arrays or a single list with a severity field? Think about what an LLM or script would want to parse.

The check composition — each check is a function. How do you accumulate results? Do checks bail early on a fatal error or always run all checks?
