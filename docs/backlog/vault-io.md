---
title: "Vault I/O"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: high
due: ""
created: 2026-05-21
updated: 2026-06-02
tags: [phase-1, vault-io, inbox, archive, log, plugin-dispatch]
---

# Vault I/O

The first user-visible commands. Three command groups covering the core vault operations: capturing to inbox, archiving files, and appending to the log. Each writes to the vault and updates `context.md` frontmatter where appropriate.

## Learning Focus

- YAML frontmatter read/write with `gopkg.in/yaml.v3`
- File append vs. overwrite patterns
- `os.Rename` cross-device fallback (archive moves files)
- stdout/stderr separation: results to stdout, diagnostics to stderr

## Commands

```
magpie inbox capture [text]   # append to today's inbox file, increment inbox-count
magpie inbox list             # list inbox items with filesystem timestamps
magpie archive add <file>     # move file to archive, inject frontmatter
magpie log append <text>      # append timestamped entry to wiki/log.md
magpie <unknown>              # dispatch to registered plugin via syscall.Exec
```

## Plugin Dispatch (ships this phase)

Wire the unknown-subcommand handler in `cmd/root.go` to do real `syscall.Exec` dispatch — not just print a placeholder. Steps: look up the subcommand name in merged config `plugins:` map → if found, `syscall.Exec(binaryPath, os.Args[1:], envWithVault)` → if not found, print "unknown subcommand" with a hint to `magpie plugin install`.

This is a small addition but it must ship in Phase 1. Phases 5 and 6 (magpie-stats, magpie-git) can't be integration-tested without it.

## You Drive

`inbox capture` — where does today's inbox file live and what does it look like? What's the naming convention for the file?

`archive add` — the safe move pattern: write frontmatter to a temp file, verify it's valid, then move. If the destination is on a different filesystem, `os.Rename` fails — what's the fallback?

`archive add --dry-run` — what output does dry-run produce? Is it a diff, a plan summary, or a structured JSON preview? Consider what's most useful to both a human running interactively and Claude parsing the result.

## Frontmatter Write Pattern

Frontmatter is always written atomically — parse the full struct, construct a new complete struct with changes applied, write the whole section as a unit:

```go
ctx, err := context.Read(vaultPath)   // parse full frontmatter into struct
ctx.InboxCount++                       // modify a copy
err = context.Write(vaultPath, ctx)   // write complete frontmatter as a unit
```

Never patch a single field in place. `Write` owns the full frontmatter section every time.

## Safety Levels for Destructive Operations

`archive add` moves a file — it's irreversible without git. The escalation ladder from the hooks design signal:

1. `--dry-run` — print what would happen, touch nothing
2. Diff preview — show the exact frontmatter that would be injected
3. Explicit approval gate — `y/n` before the move (interactive mode)
4. Hard block — exit 2, unconditional (hooks use this when preconditions aren't met)

`--dry-run` and diff preview are additions not in the original design. Level 3 is for interactive shell use. Level 4 is already handled by the exit code contract from Phase 0.
