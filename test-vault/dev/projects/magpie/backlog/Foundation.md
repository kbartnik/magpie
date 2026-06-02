---
title: "Foundation"
type: backlog-item
project: magpie
status: active
milestone: "1.0"
priority: high
due: ""
created: 2026-05-21
updated: 2026-06-02
downstream: "[[plans/Phase 0 — Foundation]]"
tags: [phase-0, foundation, cobra]
---

# Foundation

Project scaffolding, CLI framework, and vault resolution. No user-visible commands ship in this phase — it's the infrastructure every other phase builds on.

## Learning Focus

- Cobra's command/subcommand model: how `PersistentPreRun` propagates through subcommands
- XDG config path convention (`~/.config/magpie/config.yaml`)
- `filepath.WalkDir` sentinel detection (walk up, not down)
- Two-struct config merge: global defaults, vault-local overrides, vault wins on conflict

## Exit Code Contract

Three meaningful codes; nothing else:

| Code | Meaning | What Claude does |
|------|---------|-----------------|
| 0 | Success | Proceeds |
| 1 | Warning / user error | Logs, may proceed |
| 2 | Hard block | Cannot override — used by hooks for enforcement |

Exit 2 is the only real enforcement. A validation gate that exits 1 is logging, not blocking. Establish this in Phase 0 so every subsequent phase writes to the contract.

`stdout` carries results (machine-parseable). `stderr` carries diagnostics (human-readable). The `--json` flag pattern established in Phase 2 builds on this split.

## You Drive

`Resolve()` — the 3-tier vault resolution. Walk upward from CWD looking for a directory containing `.magpie/` (must be a directory). Key decisions: how far does the walk go? What happens if `MAGPIE_VAULT` is set but the path doesn't exist?

`Merge()` — two-tier config merge. Zero values in the local config must not override non-zero values in the global config. How do you handle the `Plugins` map (merge vs. replace)?

`configPath()` — XDG convention: `$XDG_CONFIG_HOME/magpie/config.yaml` with fallback to `~/.config/magpie/config.yaml`.
