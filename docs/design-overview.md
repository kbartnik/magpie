---
title: "magpie — Design Overview"
type: project-doc
project: magpie
created: 2026-06-02
updated: 2026-06-02
tags: [go, tooling, cli, design]
---

# magpie — Design Overview

A five-minute read. What magpie is, why it exists, and the principles that should govern every decision made while building it. Implementation details live in [[magpie — Design Spec]].

---

## What Magpie Is

Magpie is the **deterministic half** of the Nexus vault system.

The LLM (Claude) handles synthesis, reasoning, and knowledge compilation. Magpie handles everything the LLM shouldn't have to: vault discovery, file operations, inbox counting, archiving, health checks. This split is the founding principle. Every design decision traces back to it.

More formally: magpie is a **harness component**. `Agent = Model + Harness`. Magpie is the harness — the compiled runtime that converts LLM decisions into reliable filesystem operations. It doesn't think. It executes.

---

## What Magpie Is Not

- **Not an LLM.** No AI dependency in core. Magpie runs the same with or without Claude.
- **Not a knowledge tool.** The vault and Claude do that. Magpie moves files, counts things, and validates structure.
- **Not a replacement for Claude Code.** The `magpie-claude` plugin bridges them. Magpie is indifferent to which LLM uses it.
- **Not a shell script wrapper.** Compiled binary, real test suite, TDD discipline throughout.

---

## Design Principles

**The deterministic split is the rule.**
If a task requires reasoning, it belongs to the LLM. If it's mechanical and testable without a model, it belongs to magpie. When in doubt: can this be unit-tested without mocking an LLM? If yes, it's magpie's job.

**Read operations never block.**
If the vault is partially missing data, output what's available and report gaps to stderr. Vault discovery failures get a clear error with a recovery hint. Never silently fail or hang.

**Destructive write operations fail safely.**
Archive moves and other irreversible operations use exit 2 when preconditions aren't met. This is intentional enforcement — hooks rely on it. A gate that exits 1 is logging, not enforcement.

**stdout carries results. stderr carries diagnostics.**
stdout is always machine-parseable. Warnings, hints, and progress messages go to stderr. Never mix them.

**TTY detection decides output format.**
When stdout is a pipe (hook or script context), output is JSON. When stdout is a terminal, output is human-readable. No `--json` flag required for normal use. The context determines the format.

**Terse by default.**
LLMs and hooks call magpie frequently. Output should be as compact as correct allows.

**All file writes use same-directory temp-rename.**
Write to a temp file in the same directory as the target, verify, then rename atomically. Never write directly to the target. Never write across filesystem boundaries.

**Frontmatter is written as a complete unit.**
Parse the full struct, modify a copy, write the whole section. Never patch a single field in place.

**Don't store what you can derive.**
Timestamps from the filesystem. Counts from directory walks. Frontmatter holds only state magpie uniquely owns.

---

## Exit Code Contract

Three meaningful codes; nothing else.

| Code | Meaning | Caller behavior |
|------|---------|----------------|
| 0 | Success | Proceeds |
| 1 | Warning / user error | Logs, may proceed |
| 2 | Hard block | Cannot override — hooks use this for enforcement |

Exit 2 is the only real enforcement. A validation gate that exits 1 is logging. A hook security check that exits 1 provides zero protection.

---

## Output Model

TTY detection via `os.Stdout.Stat()`:

- **Piped (hook, script):** JSON to stdout. Structured, parseable, no decoration.
- **Terminal (human):** Readable prose to stdout. Warnings and hints to stderr.

Structured error codes, not prose errors. The LLM interprets error codes into recovery guidance — "vault not found" as a JSON field lets Claude suggest `magpie init vault`; as prose it forces Claude to parse a sentence.

---

## Architecture at a Glance

**Vault sentinel:** `.magpie/` directory. Its presence is sufficient — no registry, no manifest.

**Vault resolution (3-tier):**
1. `MAGPIE_VAULT` env var — explicit override, set by hooks
2. CWD walk upward — stops at first `.magpie/` ancestor
3. `default_vault` in global config — fallback

**Plugin system — two contracts:**
- *Runtime:* unknown subcommand → config lookup → `syscall.Exec` with `MAGPIE_VAULT` injected. Zero overhead; current process becomes the plugin.
- *Install-time:* plugin embeds `plugin.yaml` via `go:embed`, exposes via `--manifest` flag. Core calls the binary to retrieve it — no co-location assumption.

**Plugin metadata is the dispatcher.**
A plugin's `description` field in `plugin.yaml` is how the LLM routes intent to plugin. Rich, specific descriptions mean no explicit routing code is needed on the Claude side. Metadata quality is a first-class design concern, not documentation.

**Provider strategy — narrow scope.**
Magpie core has no LLM dependency. `magpie-claude` is the Claude Code adapter. Future providers (`magpie-openai`, etc.) follow the same plugin contract. No abstraction layer needed in core — the plugin contract *is* the abstraction.

---

## Knowledge Architecture

The vault's wiki structure follows two hub conventions that magpie init and lint enforce:

**Domain hubs:** every `wiki/` subdirectory has a corresponding hub page at the same level. `wiki/concepts/` → `wiki/concepts.md`. `magpie init vault` scaffolds these. `magpie lint` flags missing ones. Domain hubs are the compiled view of a folder — MOCs in the LLM Wiki sense.

**Cross-domain hubs:** author-created pages that aggregate across domains (e.g., a Go learning MOC spanning concepts, entities, and books). Not scaffolded; not lint-checked. Created when an author decides a cross-cutting view is valuable.

**wiki/index.md** is the hub-of-hubs — it links to domain hubs, not to individual pages.

---

## Learning Project

Same model as [[preflight-sync-go]], with an explicit Go learning layer added:

- **TDD throughout.** Tests are written before implementation. Every phase has a failing test suite before a line of production code is written.
- **Tutorial material before implementation.** Where relevant, each phase opens with Go language concepts and idiomatic patterns — not generic documentation, but the specific idioms the phase will use. Each tutorial block is labeled and skippable: say "I know this" and we move straight to the You Drive. If a gap surfaces during implementation, it gets filled in context rather than by backtracking.
- **You Drive sections.** Each phase identifies the design decisions that matter — the ones with real trade-offs. Claude scaffolds boilerplate and advises; the user implements the meaningful choices.
- **Phase by phase.** No phase starts until the previous one passes `go test ./...`. The detailed implementation plan is written at the start of each phase — not upfront — so it can adapt to what emerged during the previous one. Backlog items are the stable reference; plans are just-in-time scaffolding.
- **Test vault in repo.** `testdata/vault/` is a hermetic fixture vault embedded in the repo. Tests that write copy it to `t.TempDir()` first. No external vault dependency.

---

## Post-1.0 Direction

**Observability event log.** Every vault operation emits a structured JSON event to `.magpie/events/YYYY-MM-DD.jsonl`. Enables `magpie stats`, vault health analytics, and cron operations without LLM parsing.

**Cron autonomy.** Background operations that file results without being asked: orphan detection, stale link sweeping, inbox age warnings. Results land in `.magpie/reports/`. Session-start surfaces notable findings.

**Multi-vault.** `magpie vaults list` by scanning for `.magpie/` directories. Single-vault is the 1.0 scope.

**Event bus.** Plugin-to-plugin pub/sub. Under research — the design needs real multi-plugin pressure before the protocol is fixed. See backlog item.

---

## Document Map

| Document | Purpose |
|----------|---------|
| **This doc** | Design philosophy, principles, architecture vision |
| [[magpie — Design Spec]] | Implementation reference: schemas, config formats, command contracts, phase sequence |
| [[magpie — Red Team Review]] | Adversarial review: hidden assumptions, failure modes, resolved blockers |
| Backlog items | Phase-by-phase execution plan with You Drive sections |
