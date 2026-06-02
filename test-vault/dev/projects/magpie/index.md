---
title: "magpie"
type: project
slug: magpie
status: active
repo: "~/Source/magpie"
tech-stack: [Go]
created: 2026-05-21
updated: 2026-05-21
related-adrs: []
related-investigations: []
tags: [go, tooling, vault-tools, cli]
---

# magpie

## What It Is

A Go CLI that provides a clean, testable, single-binary interface for all Nexus vault operations. Replaces `vault-tools` — a collection of fragile shell scripts — with a compiled binary that is reliable, hookable, and extensible via plugins.

Magpie is a **pure vault tool** with no dependency on Claude Code or any LLM. Claude Code integration is handled by a separate plugin (`magpie-claude`) in its own repo.

## Goals

- Provide a clean, testable, single-binary interface for all vault operations
- Replace `vault-tools` shell scripts with reliable, deterministic Go commands
- Fix hook fragility: replace `VAULT_PATH` env var pattern with `.magpie/` sentinel-based vault resolution
- Extensible via a two-contract plugin system (runtime + install-time manifest)

## Architecture

**Vault sentinel:** `.magpie/` directory at vault root — its presence identifies a vault. No other markers required. Vaults are self-registering; no registry needed.

**Vault resolution (3-tier):**
1. `MAGPIE_VAULT` env var — explicit override, always set by hooks
2. CWD walk upward — stop at first directory containing `.magpie/`
3. `default_vault` in `~/.config/magpie/config.yaml` — fallback when outside any vault

**Vault-local config:** `.magpie/config.yaml` — per-vault overrides merged over global config.

**Plugin system — two contracts:**
- *Runtime:* unknown subcommand → `syscall.Exec` dispatch with `MAGPIE_VAULT` injected
- *Install-time:* `plugin.yaml` manifest declares tools, skills, hooks, MCP servers, subscriptions — `magpie plugin install <path>` processes it

**Event bus:** deferred to post-1.0. Claude Code hooks satisfy all v1.0 use cases. Needs multiple real plugins for proper design pressure before the protocol is fixed.

**Config format:** YAML for all magpie config. JSON used only for Claude Code `settings.json` interop (magpie-claude only).

Full design: [[magpie — Design Overview]] · [[magpie — Design Spec]] · [[magpie — Red Team Review|Red Team Review]]

## Commands

| Command | Description |
|---|---|
| `inbox capture [text]` | Append to inbox, increment `inbox-count` |
| `inbox list` | List inbox items with filesystem timestamps |
| `archive add <file>` | Move file to archive, inject frontmatter |
| `log append <text>` | Append timestamped entry to `wiki/log.md` |
| `lint [--json]` | Validate vault structure and schema versions |
| `init vault` | Create `.magpie/`, `context.md`, required dirs |
| `init project` | Scaffold a project in `dev/projects/` |
| `context status [--json]` | Display vault state |
| `plugin install/remove/status/list` | Manage plugins |

Bundled plugins (same repo): `magpie-stats` (vault statistics), `magpie-git` (vault git helpers).

## Learning Project

Same model as preflight-sync-go — work proceeds phase by phase. Each phase has a plan with "You Drive" sections for the meaningful implementation decisions. Claude scaffolds boilerplate and advises; the user drives design decisions and core logic.

## Phase Sequence

| Phase | Name |
|---|---|
| 0 | Foundation — cobra CLI, vault resolution, `.magpie/` sentinel, config merge |
| 1 | Vault I/O — inbox, archive, log |
| 2 | Lint |
| 3 | Init |
| 4 | Plugin system — runtime + manifest contracts |
| 5 | magpie-stats (bundled plugin) |
| 6 | magpie-git (bundled plugin) |
| 7 | Context — `context status` |
| 8 | Migration from vault-tools |

magpie-claude (`github.com/kbartnik/magpie-claude`) starts after Phase 4 — separate project and repo.

## Backlog

!backlog.base

## Plans

- [[Foundation]]

## Open

- [2026-06-02] Need a strategy for re-integrating vault content between `test-vault` (in this repo) and the Nexus vault — wiki pages and learning notes will diverge over time.

## Log

_Session notes below, newest first_

---

2026-05-21 — Brainstorming session. Redesigned from original spec: `.magpie/` sentinel replaces `context.md + .claude/`, Claude integration moved to separate `magpie-claude` plugin, two-contract plugin system (runtime + manifest), event bus added, YAML for magpie config / JSON for Claude interop only. Phase 0 plan written.
