---
title: magpie — Design Spec
type: project-doc
project: magpie
created: 2026-05-21
updated: 2026-06-03
downstream: "[[magpie — Red Team Review]]"
tags:
  - go
  - tooling
  - cli
---


# magpie — Design Spec

## What Magpie Is

Magpie is a Go CLI tool that provides a clean, testable, single-binary interface for vault operations. It replaces `vault-tools` — a collection of fragile shell scripts — with a compiled binary that is reliable, hookable, and ruthlessly modular via a plugin system.

Magpie is a **pure vault tool**. It has no dependency on Claude Code, Obsidian, or any LLM. Integration with those systems is handled by plugins that depend on magpie — not the other way around:

- **`magpie-claude`** — Claude Code integration: session management, hook installation, skill deployment
- **`magpie-obsidian`** — Obsidian-specific features: canvas/base file handling, `.obsidian/` config integration, Obsidian-aware vault health checks
- **`magpie-stats`** — read-only vault statistics (bundled; validates runtime contract)
- **`magpie-git`** — vault git helpers (bundled; validates manifest contract)

The core binary knows nothing about any of these systems. If you find yourself importing Claude Code APIs, Obsidian APIs, or LLM SDKs into core — stop. That code belongs in a plugin.

---

## Core Principles

- **Read operations never block.** Vault discovery failures return a clear error with a recovery hint. If data is partially missing, output what's available and report gaps to stderr. LLM-dependent operations especially must not block Claude Code from running.
- **Destructive write operations fail safely.** Archive moves and other irreversible operations use exit 2 when preconditions aren't met. See exit code contract in Error Handling.
- **Partial results over no results.** Output what's available; report what's missing to stderr.
- **stdout is for machines, stderr is for humans.** TTY detection (`os.Stdout.Stat()`) determines format: JSON when piped (hook or script context), human-readable when stdout is a terminal. Warnings and diagnostics always go to stderr.
- **Frontmatter is for automation, body is for users.** Frontmatter drives magpie, DataViews, Bases, and other tooling. Body content is primarily for the user — magpie writes it where useful but does not own it.
- **Don't store what you can derive.** Timestamps come from the filesystem. Frontmatter holds only state magpie uniquely owns.
- **All file writes use temp-rename.** Write to a temp file in the **same directory** as the target, verify, then rename over the original. Same-directory ensures the rename is an atomic inode swap (never cross-device) and triggers Obsidian's file history correctly via FSEvents path-level watching. Prevents partial writes regardless of what's being modified — frontmatter or body.
- **Frontmatter is written as a complete unit.** Parse the full frontmatter into a struct, construct a new complete struct with the desired changes, write the whole section. Never mutate a single field in place. Magpie owns the frontmatter structure entirely.
- **Body writes are surgical.** Read the whole file, find the target section by heading, make the targeted modification, write the whole file back via temp-rename. Never reconstruct body content programmatically — magpie modifies only the section it owns and leaves everything else byte-for-byte identical.

---

## Vault Structure

A vault is any directory containing a `.magpie/` directory. Its presence is the sentinel — no other markers required.

```
vault/
├── .magpie/
│   └── config.yaml          # vault-local config (overrides global)
├── context.md               # vault state file
├── wiki/
│   └── log.md
├── inbox/
├── archive/
│   ├── clippings/
│   ├── papers/
│   ├── books/
│   ├── daily/
│   ├── ideas/
│   └── docs/
└── dev/
    └── projects/
```

### Vault Discovery

Vaults are self-registering. Any directory with `.magpie/` is a vault — no registry required. `magpie vaults list` (post-1.0) finds all vaults by scanning for `.magpie/` directories.

### Vault Resolution (3-tier)

1. `MAGPIE_VAULT` env var — explicit override, always set by hooks
2. CWD walk upward — stop at first directory containing `.magpie/`
3. `default_vault` in global config — fallback when outside any vault

---

## Configuration

### Global config

`~/.config/magpie/config.yaml` — user-wide settings.

```yaml
default_vault: ~/Documents/Obsidian/Nexus   # optional fallback

plugins:
  claude: ~/.local/bin/magpie-claude
  stats:  ~/.local/bin/magpie-stats
  git:    ~/.local/bin/magpie-git
```

### Vault-local config

`<vault>/.magpie/config.yaml` — per-vault overrides. Merged over global config; vault wins on conflict.

```yaml
inbox_path: inbox/
archive_path: archive/
log_path: wiki/log.md
```

### Config format

All magpie config is **YAML**. Go's `yaml.v3` library handles all reading and writing — no external tool dependency in core.

---

## context.md Schema

`context.md` is the vault state file. Frontmatter holds machine-readable fields owned by magpie. The body is user-owned content.

### v1.0 Frontmatter

```markdown
---
schema: 1
vault-name: "Nexus"
inbox-count: 3
---

## Current Focus
...

## Next Actions
- [ ] item
```

| Field | Type | Written by | Purpose |
|---|---|---|---|
| `schema` | int | init, migrations | format version for migration |
| `vault-name` | string | `init vault` | vault identification |
| `inbox-count` | int | `inbox capture` | unread inbox items |

Inbox filenames and capture timestamps are derived from the `inbox/` directory via `os.Stat()` — not stored in frontmatter. Other timestamps come from `os.Stat()` on the file itself.

### Schema versioning

Integer versions map to discrete migration functions: `migrate1to2()`, `migrate2to3()`. Chained to bring any vault to current. `magpie lint` warns on stale schema. `magpie init vault --upgrade` migrates.

- Adding an optional field with a clear default → no bump
- Adding a required field, removing, or renaming → bump

### Body content

Body sections belong to the user. Magpie reads them where useful and writes to them purposefully, without reformatting or reordering content it didn't create.

In v1.0:
- `## Current Focus` and `## Next Actions` — user-maintained. `context status` reads and displays them but nothing writes them until magpie-claude adds session intelligence.
- `## Parked Ideas` — deferred to magpie-claude. The distinction between parking and capturing isn't meaningful without session context. When magpie-claude ships, `context park` can snapshot active session state alongside the text.

---

## Plugin System

Two contracts define what a plugin is.

### Contract 1: Runtime

**Dispatch wires in Phase 1, not Phase 4.** The unknown-subcommand handler in `cmd/root.go` must route to registered plugins before Vault I/O commands ship. Every phase from 1 onward validates that the dispatch mechanism works.

When magpie receives an unknown subcommand, it looks up the subcommand name in the **`plugins:` map in merged config** (global `~/.config/magpie/config.yaml` merged with vault-local `.magpie/config.yaml`) and dispatches via `syscall.Exec` (process replacement — zero overhead). The current process becomes the plugin.

Plugins are **not** discovered by scanning PATH for `magpie-<name>` binaries. A plugin must be explicitly registered via `magpie plugin install` before it is available as a subcommand.

**Plugin obligations:**
- Be a binary
- Accept `MAGPIE_VAULT` env var (magpie injects it)
- Own your own cobra subcommand tree and help text

### Contract 2: Install-time (manifest)

Each plugin embeds a `plugin.yaml` manifest via `go:embed` and exposes it via a `--manifest` flag. `magpie plugin install <path>` calls `<binary> --manifest` to retrieve the YAML, then processes it. This works regardless of how or where the binary was installed — no co-location required.

```yaml
name: magpie-claude
description: "Claude Code integration — session management, park/capture, hook and skill installation"
version: "1.0.0"
schema: "1"
requires_magpie: ">= 1.0.0"

tools:
  - name: jq
    check: "jq --version"
    min_version: "1.6"
    required: true
    hint: "brew install jq"

skills:
  - source: skills/capture.md    # path within the embedded filesystem
    dest: ~/.claude/skills/

post_install:
  - "magpie-claude setup"        # plugin handles its own Claude-specific setup
```

Core handles: version check, tool dependency checks, skill file extraction and copying, binary registration. Everything else — hooks, MCP server registration, Claude-specific config — is handled by the plugin via `post_install` commands.

**Plugin obligations (both contracts):**
- Runtime: be a binary, accept `MAGPIE_VAULT` env var, own your cobra subcommand tree
- Install-time: implement `--manifest` flag that prints `plugin.yaml` to stdout

**Core commands:**
- `magpie plugin install <path>` — call `<binary> --manifest`, process result, register in config
- `magpie plugin status <name>` — check each declared dep and skill against live system
- `magpie plugin remove <name>` — unwind what install wrote, deregister from config

**Manifest schema versioning:** The `schema:` field lets core reject manifests it doesn't understand with a clear error. Breaking changes require a new schema version.

**Missing tools:** Core reports missing tools with their `hint` and exits non-zero. Does not auto-install.

---

## Commands

### magpie core

| Command | Description |
|---|---|
| `inbox capture [text]` | Append to today's inbox file, increment `inbox-count` |
| `inbox list` | List inbox items with timestamps from filesystem |
| `archive add <file>` | Move file to archive, inject frontmatter |
| `log append <text>` | Append timestamped entry to `wiki/log.md` |
| `lint` | Validate vault structure and schema versions |
| `init vault` | Create `.magpie/`, `context.md`, required dirs |
| `init project` | Scaffold a project in `dev/projects/` |
| `context status` | Display vault state (reads frontmatter + body sections) |
| `plugin install <path>` | Call `<binary> --manifest`, process result, register plugin |
| `plugin remove <name>` | Unwind plugin installation |
| `plugin status <name>` | Check plugin installation state |
| `plugin list` | List registered plugins |

### Bundled plugins (same repo)

**magpie-stats** — vault statistics. Walk vault, count notes by section, report inbox depth, archive size, last log entry. Read-only. Validates the runtime contract.

**magpie-git** — vault git helpers. `magpie git status/commit/log` scoped to vault root. Validates both contracts: runtime dispatch and manifest with tool dependency check and `post_install`.

### External plugins

**magpie-claude** (`github.com/kbartnik/magpie-claude`) — Claude Code integration. Session start/end, `context park`, hook installation, skill deployment. Depends on magpie; core has zero knowledge of Claude Code.

**magpie-obsidian** — Obsidian-specific vault tooling. Canvas and Base file management, `.obsidian/` config integration, Obsidian-aware health checks (wikilink format, callout syntax, embed resolution). Depends on magpie; core has zero knowledge of Obsidian's APIs or file formats beyond plain markdown.

---

## Format Conventions

| Context | Format | Tool |
|---|---|---|
| Magpie config, manifests | YAML | Go `yaml.v3` |
| Claude Code `settings.json` | JSON | `jq` (magpie-claude only) |
| Machine output (TTY-detected, piped) | JSON | Go `encoding/json` |

---

## Error Handling

- **Exit code contract.** Three meaningful codes: 0 (success), 1 (warning or user error — caller may proceed), 2 (hard block — destructive operation with unmet precondition; hooks rely on this for enforcement). A gate that exits 1 is logging, not enforcement.
- **stdout/stderr split.** stdout carries results (machine-parseable, TTY-detected format). Warnings and diagnostics go to stderr.
- **Partial results.** Output what's available even when some data is missing or schema is stale.
- **File operations.** Write to temp file first, verify, then move. If post-move steps fail, lint surfaces the inconsistency — no rollback needed.

---

## Phase Sequence

### magpie (this repo)

| Phase | Name | Key deliverables |
|---|---|---|
| 0 | Foundation | Cobra CLI, vault resolution, `.magpie/` sentinel, two-tier config merge |
| 1 | Vault I/O | `inbox`, `archive`, `log` commands, `context.md` frontmatter writes, **plugin dispatch** (unknown subcommand → `syscall.Exec`) |
| 2 | Lint | Vault structure validation, alias-gap check, orphan detection, schema version check, TTY-detected output |
| 3 | Init | `init vault` (`.magpie/`, `context.md`, dirs), `init project` |
| 4 | Plugin management | Manifest contract (`--manifest` flag), `plugin install/remove/status/list` |
| 5 | magpie-stats | Bundled plugin — validates runtime contract end-to-end |
| 6 | magpie-git | Bundled plugin — validates manifest contract with tool dep check and `post_install` |
| 7 | Context | `context status`, body section read |
| 8 | Migration | Cutover from vault-tools, all existing tests pass |

**Phase 1 note:** plugin dispatch ships here — not Phase 4. Phase 4 delivers the management
commands (`plugin install/remove/status`). The dispatch mechanism (`syscall.Exec` on unknown
subcommands) is cheap to wire and validates the architecture immediately. Phases 5 and 6 would
be untestable without it.

### magpie-claude (separate repo — starts after Phase 4)

| Phase | Name | Key deliverables |
|---|---|---|
| A | Foundation | Plugin scaffold, cobra tree, `plugin.yaml`, jq dependency |
| B | Session + Park | `session start/end`, `context park` with session snapshot, schema → 2 |
| C | Sync | `sync inbox-count`, notification hook |
| D | Init | Install hooks, skills, MCP servers — full bootstrap with user permission |

**session-start, session-end, context-update are not core commands.** They live in Phase B of
magpie-claude. The core binary never writes session state — it does not know what a session is.

### magpie-obsidian (separate repo — post-1.0)

| Phase | Name | Key deliverables |
|---|---|---|
| A | Foundation | Plugin scaffold, `.obsidian/` config reader, canvas/base file types |
| B | Lint extensions | Wikilink format check, callout syntax validation, embed resolution |
| C | Vault ops | Obsidian-aware archive (preserve embeds), Base file scaffolding |

**Post-1.0:** `go extract` graduates to `sonar` (`github.com/kbartnik/sonar`).

---

## Testing Strategy

- **TDD unit tests** — written first, drive implementation, live alongside code in `_test.go` files
- **Integration tests** — behind `//go:build integration`, run with `go test -tags=integration ./tests/integration/...`
- **Fixture vault** — `testdata/vault/` with `.magpie/`, `context.md` (schema 1), required dirs. Tests that write copy it to `t.TempDir()` first.
- **Plugin tests** — build test binaries in `TestMain`; `magpie-stats` and `magpie-git` are real integration test subjects

---

## Post-1.0

- Multi-vault: `magpie vaults list` by scanning for `.magpie/` directories
- Additional `context` subcommands: `focus`, `next`, `pop`, `clear-parked`
- `magpie init vault --upgrade` migration UX
- Observability event log: `.magpie/events/YYYY-MM-DD.jsonl` structured sidecar alongside `wiki/log.md`; enables `magpie stats` and cron analytics
- Cron autonomy: background vault health checks, orphan detection, stale link sweeping; results filed to `.magpie/reports/`
- Event bus — plugin-to-plugin pub/sub (under research; needs multiple real plugins for design pressure)
- `sonar` project scope and interface with magpie-claude
