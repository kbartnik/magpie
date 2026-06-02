---
title: "magpie — Red Team Review"
type: project-doc
project: magpie
created: 2026-05-21
updated: 2026-06-02
tags: [review, red-team, architecture]
---

# Red Team Review — 2026-05-21

Adversarial review of [[magpie — Design Spec]] conducted at end of brainstorming session before implementation begins.

---

## Restatement

Magpie is a Go CLI with a vault sentinel (`.magpie/` directory), a two-tier config merge, and a two-contract plugin system (runtime via `syscall.Exec`, install-time via `plugin.yaml` manifest). Claude Code integration is entirely in a separate `magpie-claude` plugin. All file writes use same-directory temp-rename. Output format is TTY-detected: JSON when piped, human-readable when terminal.

*Note: the original design reviewed here included an event bus. That was deferred to a research backlog item after this review identified unresolved design problems (see Assumptions C/D and Failure Mode 3 below).*

---

## Hidden Assumptions

### A. magpie core processes manifest hooks/MCP sections without jq

The manifest contract allows plugins to declare `hooks` and `mcp_servers`. `magpie plugin install` is a core command that processes the full manifest. But patching Claude's `settings.json` requires JSON manipulation — and jq lives in `magpie-claude` only, not core.

**Status: unverified — internally contradictory.** Either core needs jq (contradicts the decision), core uses `encoding/json` for settings.json patching (fine, but not stated), or hooks/MCP sections in the manifest are not processed by core and must be delegated somewhere else.

### B. plugin.yaml lives in the same directory as the installed binary

`magpie plugin install <path>` reads `plugin.yaml` from the binary's directory. `go install` places binaries in `~/go/bin/` with no adjacent files. Homebrew installs to `Cellar/`, also with no adjacent data files.

**Status: unverified — assumption is false for the two most common Go distribution methods.**

### C. Sequential event dispatch is fast enough to be invisible

`magpie inbox capture` emits `inbox.captured`, which synchronously invokes every subscriber before returning. A normal git commit takes 200ms+. With multiple subscribers, latency compounds on the critical path of every command.

**Resolved (2026-06-02):** Event bus deferred to research backlog. This assumption no longer applies to v1.0.

### D. The subscriber stdin/terminal ambiguity is solvable

When invoked as an event subscriber, a plugin receives JSON on stdin. When invoked directly by the user, stdin is the terminal. The protocol for detecting which mode is unspecified. Affects every plugin that subscribes to events.

**Resolved (2026-06-02):** Event bus deferred to research backlog. This assumption no longer applies to v1.0.

---

## Failure Modes

### 1. Plugin install breaks for standard Go distribution

User runs `magpie plugin install ~/go/bin/magpie-git`. Magpie looks for `~/go/bin/plugin.yaml` — doesn't exist. Install fails. The install-time contract is unusable with `go install`, the standard Go toolchain distribution path.

### 2. The manifest contradiction blocks Phase 4

Phase 4 implements `plugin install` including the full manifest. When it reaches `hooks:` or `mcp_servers:`, core has no mechanism to act — no jq, no JSON patching logic. Either those sections are silently ignored or Phase 4 blocks until the contradiction is resolved.

### 3. Event bus unvalidated for four phases

Phase 5 builds the full event bus. Phase 6 (magpie-stats) doesn't subscribe. Phase 7 (magpie-git) is the first real subscriber. Phases 1–4 stub event emission as dead code. If the event bus design is wrong, it's discovered in Phase 7 after being "done" for two phases.

**Resolved (2026-06-02):** Event bus removed from v1.0 scope entirely. Moved to research backlog.

---

## Steelman: No Event Bus in v1.0

Every concrete v1.0 use case for the event bus can be satisfied by Claude Code hooks or shell composition:

- Claude Code hooks call `magpie-git commit` directly after `magpie inbox capture`
- Users chain commands in shell: `magpie inbox capture "idea" && magpie git commit`
- `magpie-git` invoked explicitly by the user when they want it

Removing the event bus from v1.0 eliminates Phase 5, simplifies the manifest (no `events` or `subscriptions` sections), removes the stdin/terminal ambiguity problem, and removes dispatch latency risk. It can be designed properly post-1.0 with multiple real plugins providing design pressure.

---

## Principles

**"Who executes each section of a shared contract?"**
Name the executor before naming the schema field. If no clear executor exists, the field is premature.
*Category: hidden coupling / wrong abstraction level*

**"Wait for the second consumer before abstracting."**
One data point isn't enough to validate an abstraction. The event bus has one concrete v1.0 subscriber. That's not sufficient design pressure to justify the protocol complexity.
*Category: premature abstraction / complexity debt*

**"Who controls the filesystem layout at install time?"**
If it's not the component author, co-location assumptions will fail. Verify install layout before designing file-adjacency contracts.
*Category: violated invariant / unverified deployment model*

---

## Verdict

> [!warning] Revisit — two blockers before Phase 4
>
> **Blocker 1 — Manifest/jq contradiction:** Decide how `hooks:` and `mcp_servers:` manifest sections are processed. Options: core uses `encoding/json` directly (no jq), or those sections are delegated to the plugin via a `<binary> --install-hooks` protocol. Must be explicit in the spec before Phase 4.
> **Resolved (2026-06-02):** Hooks/MCP sections delegated to the plugin via `post_install` commands. Core handles tool checks, skill copying, and binary registration only — no JSON patching in core.
>
> **Blocker 2 — plugin.yaml co-location:** The file co-location assumption breaks standard Go installs. Recommended fix: embed the manifest in the binary via `go:embed` and expose it via `<binary> --manifest` flag. Core calls the binary to get the YAML rather than reading a file adjacent to it.
> **Resolved (2026-06-02):** `go:embed` + `--manifest` flag approach adopted. Spec updated.

> [!tip] Strong recommendation
> Defer the event bus to post-1.0. Claude Code hooks satisfy every concrete v1.0 use case. The event bus should be designed when there are multiple real plugins providing design pressure — not speculatively before the first plugin ships.
> **Accepted (2026-06-02):** Event bus moved to research backlog. Phase 5 (previously "Event Bus") removed from v1.0 phase sequence.

---

## Obsidian Notes to Write

Two principles worth permanent notes in NeoCortex:

1. **"Who executes each section of a shared contract?"** — broadly applicable to any interface design, API schema, or plugin protocol
2. **"Wait for the second consumer before abstracting."** — the rule of three applied to infrastructure design
