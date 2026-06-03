---
title: "Plugin System"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-06-02
tags: [phase-1, plugin, syscall, manifest]
---

# Plugin System

**Dispatch wires in Phase 1, not Phase 4.** The `"plugin dispatch not yet implemented"` stub in `cmd/root.go` must be completed alongside Vault IO — before session management, lint, or init. Every subsequent phase depends on knowing that unknown subcommands route correctly.

Two contracts define what a plugin is. This phase implements both: the runtime dispatch contract (`syscall.Exec`) and the install-time manifest contract (`plugin.yaml` via `--manifest` flag). Everything from Phase 2 onward depends on this.

**session-start, session-end, context-update are NOT core commands.** These are provided by the `magpie-claude` plugin, which depends on magpie, not the other way around. The core binary never imports or calls Claude Code APIs. If you find yourself writing `session.go` or `context.go` in the core package, stop — that code belongs in magpie-claude.

## Learning Focus

- `syscall.Exec` — process replacement vs. subprocess spawning (`exec.Command`)
- `go:embed` — embedding files in a binary at compile time
- Schema versioning in YAML manifests — why an integer schema field matters
- `os.Environ()` for injecting env vars before exec

## Contracts

**Runtime:** unknown subcommand → look up in config registry → `syscall.Exec` with `MAGPIE_VAULT` injected. Current process *becomes* the plugin — zero overhead, correct signal handling.

**Install-time:** Plugin embeds `plugin.yaml` via `go:embed` and exposes it via `--manifest` flag. `magpie plugin install <path>` calls `<binary> --manifest` to retrieve the manifest — no co-location required, works with `go install` and Homebrew. Core handles tool checks, skill copying, binary registration. Plugin-specific setup (hooks, MCP) runs via `post_install` commands.

## Commands

```
magpie plugin install <path>   # call <binary> --manifest, process result, register
magpie plugin remove <name>    # unwind installation, deregister from config
magpie plugin status <name>    # check each declared dep/skill against live system
magpie plugin list             # list registered plugins
magpie <name> [args]           # dispatch to registered binary (via unknown command handler)
```

## You Drive

`execPlugin()` — look up the binary path from config (not PATH), build the env with `MAGPIE_VAULT` injected, call `syscall.Exec`. Key question: what do you pass as `argv[0]`? Convention is the binary name, not the full path.

Tool version checking — `check:` runs a shell command, `min_version:` is a semver lower bound. How do you extract a version string from arbitrary command output and compare it?

## Plugin Metadata as the Router

The resolver insight from agent memory architecture: **plugin descriptions are the router**. When Claude encounters an unknown task, it reads plugin descriptions and matches intent automatically — no explicit routing code needed on the LLM side.

This makes `--manifest` metadata quality a first-class design concern, not documentation. A plugin's `description` field in `plugin.yaml` should be specific enough that Claude can distinguish it from other plugins without reading source code.

The `--manifest` response should include a human-readable `description` alongside the tool/skill/hook declarations. This is what makes the plugin system self-describing to an LLM — and what lets plugin descriptions serve as the implicit dispatcher.

## Note on syscall.Exec

`syscall.Exec(path, args, env)` never returns on success — the current process *becomes* the plugin. On failure it returns an error. This is fundamentally different from `exec.Command(...).Run()` which spawns a child process.
