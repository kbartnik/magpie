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

- [2026-06-02] Phase 0 plan + hands-on session — repo setup, module init, docs scaffold

## Key Concepts

- **Go modules**: `go.mod` declares the module path (import namespace) and Go version; `go.sum` locks dependency hashes
- **SSH config IdentityFile**: per-host key selection in `~/.ssh/config`; without a `Host github.com` block, `id_ed25519_github` is never offered to GitHub

## What I Now Understand

Repo is initialized and remote is wired. Go 1.26.3 is installed. Module is `github.com/kbartnik/magpie` with cobra and yaml.v3 dependencies. Session was mostly infrastructure — git, SSH, docs — not yet into the Go learning goals proper.

## Open Questions

- How far should the CWD walk go in vault resolution? (stopping at `$HOME` vs filesystem root)
- Should `MAGPIE_VAULT` pointing to a non-existent path be an error or trusted blindly?

## What to Explore Next

- Task 2: cobra root command — understand `PersistentPreRun` propagation and `RunE` vs `Run`
- Task 3: `Config` struct + `Load` — XDG spec in practice
