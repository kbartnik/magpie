---
title: "magpie — Design Overview"
type: project-doc
project: magpie
created: 2026-06-02
updated: 2026-06-20
tags: [go, tooling, cli, design]
---

# magpie — Design Overview

A five-minute read. What magpie is, why it exists, and the principles that should govern every decision made while building it. Implementation details live in [[magpie — Design Spec]].

---

## What Magpie Is

Magpie is a **domain-specific LLM harness** for knowledge vaults.

Not a general tool server that happens to operate on vaults, but a harness that *understands* vaults well enough to co-pilot the LLM through vault operations. It contextualizes every operation with domain knowledge, validates intent before execution, and reports domain-relevant feedback after execution.

The LLM is magpie's primary consumer. Magpie returns information in a minimal, machine-friendly way using fixed, versioned schemas (the **oracle**), handles every purely mechanical operation the LLM shouldn't have to reason through (the **executor**), and actively uses domain knowledge to make each response more useful — reporting what changed about vault health, what attention items exist, and what to consider next.

The founding split remains: if a task requires reasoning, it belongs to the LLM; if it's mechanical and testable without a model, it belongs to magpie. The harness framing adds a third role: if it requires *domain knowledge* to contextualize an LLM's action, that also belongs to magpie.

More formally: `Agent = Model + Harness`. Magpie is the harness — the compiled runtime that converts LLM decisions into reliable filesystem operations, feeds the LLM exactly what it needs to orient and decide, and uses vault domain knowledge to validate and enrich every interaction. It doesn't think. It orients, validates, executes, and reports.

---

## What Magpie Is Not

- **Not an LLM — despite being LLM-first.** Designing *for* an LLM caller and depending *on* one are different things. Core has zero AI dependency and runs identically with any provider, or none; `magpie-claude` is what bridges it to Claude Code specifically.
- **Not a knowledge tool.** The vault and Claude do that. Magpie moves files, counts things, and validates structure.
- **Not a replacement for Claude Code.** The `magpie-claude` plugin bridges them. Magpie is indifferent to which LLM uses it.
- **Not an Obsidian tool.** The `magpie-obsidian` plugin handles Obsidian-specific features. Core only knows markdown.
- **Not a shell script wrapper.** Compiled binary, real test suite, TDD discipline throughout.

## Plugin Ecosystem

The core binary is deliberately narrow — it owns only the vault contract. Three plugin families surround it:

- **Harness plugins** (above core): `magpie-claude` today; future provider plugins (`magpie-openai`, etc.) follow the same contract. They bridge magpie to whichever LLM is driving it and own the `model_tier` / `effort` / `token_budget` routing hints.
- **Domain plugins** (alongside core): `magpie-stats`, `magpie-git`. They extend what the vault can do — statistics, git operations — without bridging to an LLM or a rendering surface. Pure vault-domain capability; no dependency beyond their own tooling.
- **Frontend plugins** (below core): `magpie-obsidian` today, future rendering surfaces tomorrow. They translate vault state into something a human-facing tool can render.

| Plugin | Family | Scope | Depends on |
|--------|--------|-------|------------|
| `magpie-stats` | Domain | Vault statistics — bundled, read-only | magpie core |
| `magpie-git` | Domain | Git helpers scoped to vault | magpie core, git |
| `magpie-claude` | Harness | Claude Code: session, hooks, skills | magpie core, jq |
| `magpie-obsidian` | Frontend | Obsidian: canvas/base, embed, `.obsidian/` config | magpie core |

The boundary rule: *if a feature requires importing or parsing something system-specific (Claude Code settings.json, Obsidian canvas format, a provider's SDK), it belongs in that system's plugin.* Core should be compilable and testable with zero knowledge of these systems.

**Transport adapters are not plugins.** CLI and MCP determine *how* a caller reaches core — plugins determine *what* the vault does. Core is transport-agnostic: it returns Go structs, and each adapter owns its own serialization and envelope format. The MCP adapter is a thin wrapper around the same core, not a redesign.

---

## Design Principles

**The deterministic split is the rule.**
If a task requires reasoning, it belongs to the LLM. If it's mechanical and testable without a model, it belongs to magpie. When in doubt: can this be unit-tested without mocking an LLM? If yes, it's magpie's job.

**Read operations never block.**
If the vault is partially missing data, output what's available and report gaps to stderr. Vault discovery failures get a clear error with a recovery hint. Never silently fail or hang.

**Destructive write operations fail safely.**
Archive moves and other irreversible operations use exit 2 when preconditions aren't met. This is intentional enforcement — hooks rely on it. A gate that exits 1 is logging, not enforcement.

**Writes preview before they commit.**
Every write operation accepts `--dry-run`: same JSON schema as the live run, `dry_run: true` in the envelope, no state touched. The LLM previews an irreversible action before it spends it.

**stdout carries results. stderr carries diagnostics.**
stdout is always machine-parseable. Warnings, hints, and progress messages go to stderr. Never mix them.

**TTY detection decides output format.**
When stdout is a pipe (hook or script context), output is JSON. When stdout is a terminal, output is human-readable. No `--json` flag required for normal use. The context determines the format.

**Every response is a harness contract.**
The response envelope carries `schema_version` (passive drift detection), `status` (ok/warning/blocked), `effects` (what changed — files created, modified, deleted), `delta` (vault health changes caused by this operation), and `hint` (suggested next action). The LLM gets complete feedback in one round trip — no follow-up calls needed to learn what happened.

**Terse by default.**
LLMs and hooks call magpie frequently. Output should be as compact as correct allows. Read commands support `--limit` and `--summary` to control output size — the LLM requests "top 5 results" rather than "everything."

**All file writes use same-directory temp-rename.**
Write to a temp file in the same directory as the target, verify, then rename atomically. Never write directly to the target. Never write across filesystem boundaries.

**Frontmatter writes are field-scoped, not document-scoped.**
Frontmatter is shared state — the LLM, magpie, and the human author all edit it. A full parse-and-reserialize round trip is a liability: naive YAML serializers normalize quote styles, mangle `HH:MM` as sexagesimal integers, and scramble fields the caller never touched. AST-aware write primitives parse the structure, modify only the target node, and emit everything else verbatim — a hard invariant that touching `page-confidence` cannot corrupt `sources`, `aliases`, or `created`.

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

### The Layer Stack

Three plugin families surround core — harness plugins above, domain plugins alongside, frontend plugins below — and core is reached only through transport adapters, which own serialization and aren't plugins themselves:

| Layer | Purpose | Core or Plugin |
|---|---|---|
| Transport adapter | CLI flags → core; MCP JSON-RPC → core (both 1.0) | Adapter, not plugin |
| Oracle | Domain-aware LLM orientation with attention signals (`context status`, `magpie schema`) | Core |
| Executor | Atomic writes with exit-code flow control and effects reporting (`inbox capture`, `archive add`, `log append`) | Core |
| Delta validation | Post-write vault health checks (broken links, orphans, alias mismatches) | Core |
| Content extraction | Domain-aware query + bidirectional graph traversal (`query`, `context gather`) | Core |
| Relevance signals | `page-confidence` — magpie-computed, never LLM-computed | Core |
| Model routing | `model_tier` / `effort` / `token_budget` hints declared in plugin manifest | Harness plugin |
| Frontend | Obsidian rendering, `.obsidian/` config | Frontend plugin |

Core is **transport-agnostic**: it accepts structured requests and returns Go structs. Each adapter — CLI and MCP — owns its own envelope formatting and serialization.

**Vault sentinel:** `.magpie/` directory. Its presence is sufficient — no registry, no manifest.

**Vault resolution (3-tier):**
1. `MAGPIE_VAULT` env var — explicit override, set by hooks
2. CWD walk upward — stops at first `.magpie/` ancestor
3. `default_vault` in global config — fallback

**Plugin system — two contracts:**
- *Runtime:* unknown subcommand → config lookup → `syscall.Exec` with `MAGPIE_VAULT` injected. Zero overhead; current process becomes the plugin.
- *Install-time:* plugin embeds `plugin.yaml` via `go:embed`, exposes via `--manifest` flag. Core calls the binary to retrieve it — no co-location assumption.

**Plugin dispatch wires in Phase 1.**
The unknown-subcommand handler ships alongside the first user-visible commands (`inbox`, `archive`, `log`). Deferring it to Phase 4 would mean Phases 2–3 can't be tested against real plugin binaries. Dispatch is cheap — wire it early.

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
- **Pull model.** No fixed phase sequence. Backlog items have dependency links and are pulled just-in-time based on readiness and interest. The walking skeleton (Foundation + Write Primitives + Index) ships first; VaultGraph + Lint are the natural second pull. Plans are written at the start of each sprint, not upfront.
- **Test vault in repo.** `testdata/vault/` is a hermetic fixture vault embedded in the repo. Tests that write copy it to `t.TempDir()` first. No external vault dependency.

---

## Post-1.0 Direction

**Observability event log.** Every vault operation emits a structured JSON event to `.magpie/events/YYYY-MM-DD.jsonl`. Enables `magpie stats`, vault health analytics, and cron operations without LLM parsing.

**Multi-vault.** `magpie vaults list` by scanning for `.magpie/` directories. Single-vault is the 1.0 scope.

**Event bus.** Plugin-to-plugin pub/sub. Under research — the design needs real multi-plugin pressure before the protocol is fixed. See backlog item.

**Dream-gather.** Semantic maintenance signals (long pages, hub promotion, tag clusters, purge candidates). Excluded from 1.0 migration scope — `/dream` skill is disabled until this ships. Strict separation: lint = correctness, dream = optimization.

---

## Document Map

| Document | Purpose |
|----------|---------|
| **This doc** | Design philosophy, principles, architecture vision |
| [[magpie — Design Spec]] | Implementation reference: schemas, config formats, command contracts, phase sequence |
| [[magpie — Red Team Review]] | Adversarial review: hidden assumptions, failure modes, resolved blockers |
| Backlog items | Dependency-linked items with `depends-on` fields and `milestone` (1.0 vs post-1.0). Items tagged `mvp` are the walking skeleton. Items tagged `research` are open questions awaiting investigation. Development uses a scrum-like pull model — items are pulled by dependency readiness, not fixed phase order. `backlog.base` is the live, sortable view. |
