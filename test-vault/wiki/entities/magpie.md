---
title: "Project: magpie"
type: project
status: active
created: 2026-05-21
updated: 2026-05-29
sources: []
related: ["Nexus Vault Template", "Harness Engineering"]
tags: [project, go, vault-tools]
---

# Project: magpie

Migration of `vault-tools` into a standalone Go project. Magpie is the compiled interface layer between Claude Code and the vault — not a collection of shell scripts.

**GitHub (planned):** `github.com/kbartnik/magpie`
**Local:** `~/Source/magpie`

## What it replaces

The current `vault-tools` binary lives in `.claude/tools/` inside the vault. Magpie moves this out to its own module so it can be versioned, tested, and installed independently of any vault.

## sonar split

`go-extract` is being extracted as a separate tool named **sonar** (`github.com/kbartnik/sonar`, `~/Source/sonar`). Sonar serves as the reference implementation for the magpie plugin contract.

## Vault resolution hierarchy

When `magpie` runs, it locates the vault via:
1. `MAGPIE_VAULT` environment variable
2. `~/.config/magpie/config.yaml`
3. CWD walk upward — sentinel: `context.md` + `.claude/` present together

Single-vault for 1.0; multi-vault deferred.

## Plugin system

Plugins use a **registered model** — plugins are declared in config or discovered via a plugin registry, not by PATH-scanning for `magpie-<name>` binaries. The original design spec described git-style PATH discovery; that approach was superseded. The spec has not yet been updated to reflect this.

`MAGPIE_VAULT` env is passed to plugin processes. `--magpie-describe` flag on a plugin binary returns its metadata for discoverability.

## Command renames

| Old (`vault-tools`) | New (`magpie`) |
|---------------------|---------------|
| `inbox-list` | `inbox` |
| `archive-file` | `archive` |
| `log-append` | `log` |
| `lint-gather` | `lint` |

## magpie init

Scaffolds vault directories, rewrites `.claude/settings.json` hooks to clean `magpie <cmd>` calls, and writes `~/.config/magpie/config.yaml`. The old `VAULT_PATH` env var pattern is retired in favor of CWD detection.

## Implementation plan

See `/Users/kurt/.claude/plans/let-s-back-up-if-jolly-parasol.md` for the full task list.

## Implementation

Built with [[Cobra CLI]] (command tree, `PersistentPreRunE` for vault resolution). Config follows [[XDG Base Directory]] spec via `os.UserConfigDir()`. See [[Go Knowledge Map]] for the Go concept cluster this project builds on. [[SSH Config]] covers the GitHub SSH key setup required to push.

## Design

See [[magpie-design-signals-from-wiki]] for what the vault's accumulated knowledge implies for magpie's design.

## Open Questions

- [[Magpie Claim-Level Provenance]] — how to track which archive source supports each wiki claim
- [[Magpie Pipeline vs Hybrid]] — pipeline (strict stage separation) vs hybrid (stages share context) architecture decision
- How changes refined in this testbed vault flow back into `~/Source/nexus-template` is not yet documented
