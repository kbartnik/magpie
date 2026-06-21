---
title: "magpie — Design Overview"
type: project-doc
project: magpie
created: 2026-06-02
updated: 2026-06-21
tags: [go, tooling, cli, design]
---

# magpie — Design Overview

A five-minute read. What magpie is, why it exists, and the principles that should govern every decision made while building it.

---

## What Magpie Is

Magpie is a **library-first, domain-specific LLM harness** for knowledge vaults.

Its core is a **stateless Go library** of vault operations that accept structured inputs and return `Result` value objects. **Thin front-ends** — the CLI (`cmd/magpie`) and the MCP server (`cmd/magpie-mcp`) — consume the library and own serialization, envelope formatting, and process lifecycle. "Harness" is the role these front-ends play when wiring the library to an LLM-driving context, with the MCP server being the most harness-like.

The LLM is magpie's primary consumer. Magpie returns complete, domain-aware facts in one round trip — using fixed, versioned schemas (the **oracle**), handling every purely mechanical operation the LLM shouldn't have to reason through (the **executor**), and reporting what changed about vault health after each write. The differentiator vs. a general tool server: *complete domain-aware facts in one round trip* — magpie delivers judgment-ready data, not suggestions about what to do next.

The founding split remains: if a task requires reasoning, it belongs to the LLM; if it's mechanical and testable without a model, it belongs to magpie. The harness framing adds a third role: if it requires *domain knowledge* to contextualize an LLM's action, that also belongs to magpie.

More formally: `Agent = Model + Harness`. Magpie is the harness — the compiled runtime that converts LLM decisions into reliable filesystem operations, feeds the LLM exactly what it needs to orient and decide, and uses vault domain knowledge to validate and enrich every interaction. It doesn't think. It orients, validates, executes, and reports — facts in, judgment out.

---

## What Magpie Is Not

- **Not an LLM — despite being LLM-first.** Designing *for* an LLM caller and depending *on* one are different things. The library has zero AI dependency and runs identically with any provider, or none; `magpie-claude` is what bridges it to Claude Code specifically.
- **Not a knowledge tool.** The vault and Claude do that. Magpie moves files, counts things, and validates structure.
- **Not a replacement for Claude Code.** The `magpie-claude` plugin bridges them. Magpie is indifferent to which LLM uses it.
- **Not an Obsidian tool.** The `magpie-obsidian` plugin handles Obsidian-specific features. The library only knows markdown.
- **Not a shell script wrapper.** Compiled binary, real test suite, TDD discipline throughout.

## Plugin Ecosystem

The library is deliberately narrow — it owns only the vault contract. Three plugin families surround it:

- **Harness plugins** (above the library): `magpie-claude` today; future provider plugins (`magpie-openai`, etc.) follow the same contract. They bridge magpie to whichever LLM is driving it and own the `model_tier` / `effort` / `token_budget` routing hints.
- **Domain plugins** (alongside the library): `magpie-stats`, `magpie-git`. They extend what the vault can do — statistics, git operations — without bridging to an LLM or a rendering surface. Pure vault-domain capability; no dependency beyond their own tooling.
- **Frontend plugins** (below the library): `magpie-obsidian` today, future rendering surfaces tomorrow. They translate vault state into something a human-facing tool can render.

| Plugin | Family | Scope | Depends on |
|--------|--------|-------|------------|
| `magpie-stats` | Domain | Vault statistics — bundled, read-only | magpie library |
| `magpie-git` | Domain | Git helpers scoped to vault | magpie library, git |
| `magpie-claude` | Harness | Claude Code: session, hooks, skills | magpie library, jq |
| `magpie-obsidian` | Frontend | Obsidian: canvas/base, embed, `.obsidian/` config | magpie library |

The boundary rule: *if a feature requires importing or parsing something system-specific (Claude Code settings.json, Obsidian canvas format, a provider's SDK), it belongs in that system's plugin.* The library should be compilable and testable with zero knowledge of these systems.

**The library-first stack:** `internal/` packages (the library) → front-ends (`cmd/magpie`, `cmd/magpie-mcp`). The library returns `Result` value objects; each front-end owns its serialization and envelope format. **Plugin dispatch is a front-end concern:** the CLI uses `syscall.Exec`; the MCP server uses `exec.Command`. **MCP exposes core only for 1.0** — plugins remain CLI-reachable; plugins-as-own-MCP-servers is reserved post-1.0; the MCP-proxy approach was rejected (see [[ADR-0001: Adopt a library-first architecture for magpie]]).

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
The response envelope carries `schema_version` (passive drift detection), `status` (ok/warning/blocked), `dry_run` (preview flag), `data` (the operation's payload), `effects` (what changed — files created, modified, deleted), and `delta` (vault health changes caused by this operation). The LLM gets complete, domain-aware facts in one round trip — no follow-up calls needed to learn what happened.

**The envelope lives in the front-ends, not the library.**
The library returns `Result` value objects — plain Go structs with json tags as a convenience, carrying `Status`, `Effects`, `Delta`, and operation-specific data. Each front-end serializes its own envelope and maps `Status` its own way: the CLI maps it to exit codes (0/1/2); the MCP server maps it to `isError`.

**Graph-optionality is a parameter, never a parallel implementation.**
`checkHealth(graph, scope)` and `Index(vault, graph)` are single functions where the graph is an optional input. When graph data is available, results are richer; when it's absent, the operation still succeeds with reduced coverage. No separate "with-graph" vs. "without-graph" code paths.

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

Three plugin families surround the library — harness plugins above, domain plugins alongside, frontend plugins below — and the library is consumed through thin front-ends that own serialization:

| Layer | Purpose | Library, Front-end, or Plugin |
|---|---|---|
| CLI front-end | CLI flags → library; envelope + exit codes; `syscall.Exec` plugin dispatch | Front-end (`cmd/magpie`) |
| MCP front-end | JSON-RPC → library; envelope + `isError` mapping; `exec.Command` plugin dispatch | Front-end (`cmd/magpie-mcp`) |
| Oracle | Domain-aware LLM orientation (`context status`, `magpie schema`) | Library |
| Executor | Atomic writes with exit-code flow control and effects reporting (`inbox capture`, `archive add`, `log append`) | Library |
| Delta validation | Post-write vault health checks (broken links, orphans, alias mismatches) | Library |
| Content extraction | Domain-aware query + graph-optional traversal (`query`, `context gather`) | Library |
| Relevance signals | `page-confidence` — magpie-computed, never LLM-computed | Library |
| Model routing | `model_tier` / `effort` / `token_budget` hints declared in plugin manifest | Harness plugin |
| Frontend | Obsidian rendering, `.obsidian/` config | Frontend plugin |

The library lives under `internal/` (publish-on-demand) and returns `Result` value objects — plain Go structs. Each front-end owns its envelope formatting and serialization.

**Canonical package layout:**

```
internal/  result/ vault/ config/ graph/ health/ index/ query/ inbox/ archive/ log/
cmd/       magpie/  magpie-mcp/
```

`internal/result/` holds the `Status` enum + `Effects`/`Delta`/data structs (plain Go, json tags as convenience, no envelope). The envelope + `Status`→exit-code mapping live in `cmd/magpie`; `cmd/magpie-mcp` maps `Status`→`isError`.

**Vault sentinel:** `.magpie/` directory. Its presence is sufficient — no registry, no manifest.

**Vault resolution (3-tier):**
1. `MAGPIE_VAULT` env var — explicit override, set by hooks
2. CWD walk upward — stops at first `.magpie/` ancestor
3. `default_vault` in global config — fallback

**Response envelope** (front-end concern — the library returns `Result` value objects; each front-end wraps them):

```json
{
  "schema_version": 0,
  "status": "ok | warning | blocked",
  "dry_run": false,
  "data": { },
  "effects": { "created": [], "modified": [], "deleted": [] },
  "delta": { }
}
```

- `schema_version` — passive drift detection; version 0 = unstable until 1.0, field changes are free, no migration.
- `status` — `ok` (success), `warning` (user error, non-fatal), `blocked` (hard failure, cannot proceed).
- `dry_run` — `true` when previewing; same shape as a live run, no state touched.
- `data` — operation-specific payload.
- `effects` — filesystem changes: files created, modified, deleted.
- `delta` — vault health changes caused by this operation.

**Config merge:** global (`~/.config/magpie/config.yaml`) + vault-local (`.magpie/config.yaml`), vault wins, zero-values never override non-zero globals. YAML via `yaml.v3`.

**Plugin system — two contracts:**
- *Runtime:* unknown subcommand → config lookup → plugin dispatch with `MAGPIE_VAULT` injected. The CLI front-end uses `syscall.Exec` (zero overhead; current process becomes the plugin); the MCP front-end uses `exec.Command`.
- *Install-time:* plugin embeds `plugin.yaml` via `go:embed`, exposes via `--manifest` flag. The front-end calls the binary to retrieve it — no co-location assumption.

**Plugin dispatch wires in Phase 1.**
The unknown-subcommand handler ships alongside the first user-visible commands (`inbox`, `archive`, `log`). Deferring it to Phase 4 would mean Phases 2–3 can't be tested against real plugin binaries. Dispatch is cheap — wire it early.

**Plugin metadata is the dispatcher.**
A plugin's `description` field in `plugin.yaml` is how the LLM routes intent to plugin. Rich, specific descriptions mean no explicit routing code is needed on the Claude side. Metadata quality is a first-class design concern, not documentation.

**Provider strategy — narrow scope.**
The magpie library has no LLM dependency. `magpie-claude` is the Claude Code harness plugin. Future providers (`magpie-openai`, etc.) follow the same plugin contract. No abstraction layer needed in the library — the plugin contract *is* the abstraction.

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
- **Pull model.** No fixed phase sequence. Backlog items have dependency links and are pulled just-in-time based on readiness and interest. The walking skeleton — library (vault resolve + config + one write + graph-optional `Index` returning `Result`) → `cmd/magpie` front-end (serialize, exit codes, `syscall.Exec` plugin dispatch); MCP is the second front-end, the reward. VaultGraph + Lint are the natural second pull. Plans are written at the start of each sprint, not upfront.
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
| [[ADR-0001: Adopt a library-first architecture for magpie]] | Architecture decision: library-first reframe, front-end split, envelope ownership |
| [[magpie — Red Team Review]] | Adversarial review: hidden assumptions, failure modes, resolved blockers |
| Backlog items | Dependency-linked items with `depends-on` fields and `milestone` (1.0 vs post-1.0). Items tagged `mvp` are the walking skeleton. Items tagged `research` are open questions awaiting investigation. Development uses a scrum-like pull model — items are pulled by dependency readiness, not fixed phase order. `backlog.base` is the live, sortable view. |
