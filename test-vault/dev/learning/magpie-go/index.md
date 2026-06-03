---
title: "Learning: Building magpie in Go"
type: learning
topic: "magpie-go"
status: active
created: 2026-06-02
updated: 2026-06-02
tags:
  - go
  - cli
  - cobra
  - tooling
---

# Learning: Building magpie in Go

## What I Want to Understand

- How to structure a real Go CLI project from scratch (modules, packages, cobra)
- XDG Base Directory spec and why it matters for config file placement
- Vault resolution via filesystem walk — how to walk up a directory tree reliably
- Two-tier config merge: global defaults overridden by vault-local settings
- How `PersistentPreRun` propagates context through cobra's command tree
- What "idiomatic Go" looks like in practice across error handling, package layout, and testing

Done looks like: confidently writing and reasoning about these patterns without reference.

## Sources Consumed

- [[2026-06-02]] Phase 0 plan + hands-on session — repo setup, module init, docs scaffold
- [[2026-06-02]] (continued) — Go testing patterns, testify setup, first config test red phase

## Key Concepts

- **Go modules**: `go.mod` declares the module path (import namespace) and Go version; `go.sum` locks dependency hashes — see [[Go Knowledge Map]] for the full concept map
- **SSH config IdentityFile**: per-host key selection in `~/.ssh/config`; without a `Host github.com` block, `id_ed25519_github` is never offered to GitHub — see [[SSH Config]]
- **Cobra CLI**: `PersistentPreRunE` propagates to all subcommands; `RunE` vs `Run` for error handling — see [[Cobra CLI]]
- **XDG Base Directory**: `os.UserConfigDir()` + manual `XDG_CONFIG_HOME` check for cross-platform config paths — see [[XDG Base Directory]]
- **testify**: `assert` continues on failure, `require` stops; `require.NoError` is the standard setup guard; `suite` provides struct-based setup/teardown — see [[Go Testing Patterns]]
- **runtime.Caller**: returns call-stack file path baked in at compile time; `skip=1` in shared helpers anchors to the caller's package, not the helper's — see [[Go Testing Patterns]]

## What I Now Understand

Repo is initialized and remote is wired. Go 1.26.3 is installed. Module is `github.com/kbartnik/magpie` with cobra, yaml.v3, and testify dependencies. First config test is scaffolded and red — `TestLoad/valid_YAML_file_returns_correct_fields` fails on `require.NotNil` because `Load` returns nil. TDD red phase complete for that case.

## Open Questions

- How far should the CWD walk go in vault resolution? (stopping at `$HOME` vs filesystem root)
- Should `MAGPIE_VAULT` pointing to a non-existent path be an error or trusted blindly?

## What to Explore Next

- Implement `Load` — YAML parsing, missing-file handling, error on malformed input
- Fill in remaining `TestLoad` and `TestConfigPath` stubs once `Load` is green
