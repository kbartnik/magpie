---
title: "LLM Wiki Pattern"
type: concept
status: active
created: 2026-05-31
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-31-llm-wiki-compiler-deep-dive.md"
  - "archive/clippings/2026-05-31-karpathy-llm-wiki-built-twice.md"
  - "archive/clippings/2026-05-31-rag-llm-wiki-gbrain-comparison.md"
  - "archive/clippings/2026-06-01-karpathy-vibe-coding-to-agentic-engineering.md"
related:
  - "Harness Engineering"
  - "Agent Memory Architectures"
  - "Nexus Vault Template"
  - "magpie"
tags:
  - knowledge-management
  - agentic-systems
  - harness-engineering
---

# LLM Wiki Pattern

A paradigm inversion proposed by Andrej Karpathy (April 2026 gist) that replaces RAG retrieval with wiki compilation. Instead of having an LLM re-derive knowledge from raw sources on every query, the LLM incrementally builds and maintains a persistent, interlinked wiki — a structured collection of markdown files that sits *between* the user and the raw sources.

> "Obsidian is the IDE, the LLM is the programmer, the wiki is the codebase."
> — Andrej Karpathy

The key phrase: the wiki is a **persistent, compounding artifact**. Cross-references already exist at query time. Contradictions have already been flagged. Synthesis already reflects everything you've read.

## The Core Shift

![[resources/media/7f25c3d5773ffc59ba05523d284ab8b1_MD5.webp]]

RAG re-reads the same books for every exam, never actually learning the material. The LLM Wiki compiles once and improves incrementally. The cross-references that would take a human hours to build accumulate automatically as sources are ingested.

![[resources/media/d155a287dccef1983cccb95b9df9929b_MD5.webp]]

## Karpathy's Personal Practice (AI Ascent 2026)

At Sequoia's AI Ascent 2026, Karpathy described his own wiki practice unprompted:

> "I have my wiki that's being built up from these articles. I love asking questions about things... anytime I see a different projection onto information, I always feel like I gain insight. So it's really just a lot of prompts for me to do synthetic data generation over some fixed data."

He frames wikis as tools for *understanding*, not retrieval — and understanding as the irreplaceable human function that AI cannot take over:

> "You can outsource your thinking but you can't outsource your understanding. [...] You can't be a good director if you can't direct well, and the LLMs certainly don't excel at understanding. You're still uniquely in charge of that."

This is significant: the person who proposed the LLM Wiki Pattern is using it personally and articulates precisely *why* — not to retrieve facts faster, but to build the understanding needed to direct agentic work. The wiki serves the director, not the executor.

## The Division of Labor

- **Human:** curates sources, explores, asks good questions
- **LLM:** summarizes, cross-references, files, bookkeeping
- **Wiki:** the compounding artifact that makes every future query richer

## Implementations — A Design Space

The pattern is a *template*, not a tool. Implementations occupy different points on a spectrum from frozen-in-code to fully agentic:

| Implementation | Type | Frozen/Agentic | When to use |
|---|---|---|---|
| `llm-wiki-compiler` | TypeScript CLI + MCP server | Hybrid (pipeline + review queue) | Small-medium corpora; want auditable artifact with claim-level provenance |
| `wiki-llm` (Python) | Full pipeline: Pydantic models, LangGraph repair agent | Mostly frozen | Large/recurring pipelines, team use, downstream automation, audit trail required |
| AGENTS.md approach | Pure agent instruction file | Fully agentic | Personal/small wiki in discovery mode; using Claude Code/Cursor already; iterate without deploying |
| **This vault (Nexus)** | Binary (vault-tools) + CLAUDE.md schema + skills | **Hybrid: frozen mechanics, agent reasoning** | Agent-driven wiki management with deterministic file operations; deliberate middle ground |

The core design question (Leandro Bernardo): *how much implementation should be frozen in code vs. negotiated with the agent at runtime?* Deterministic code for reproducible pipelines; agentic instructions for discovery-mode wikis where the shape is still being found. Output quality is comparable — the difference is what surrounds the compilation loop.

## Technical Properties (llm-wiki-compiler reference)

**Two-phase pipeline:** Extract all concepts from all sources first, then generate pages. Eliminates order-dependence (first source no longer defines vocabulary), catches failures before dirtying the wiki, merges synonymous concepts from multiple sources into one page.

**Incremental compilation:** SHA-256 hash per source. Subsequent compile runs only process changed sources. Makes the tool usable day-to-day rather than a one-time expensive operation.

**Compounding queries:** Answers saved via `--save` flag become new wiki pages. Today's synthesis becomes context for tomorrow's query. Exploration doesn't evaporate into chat history.

**Claim-level provenance:** Citations link paragraphs to specific source files and line ranges. Linter validates citations by reading the rendered body — metadata can't be falsified. "Memory you can inspect" rather than a black box.

**Review queue:** `--review` mode writes candidates to a staging area for human approval before wiki pages are committed. Per-source locking prevents partial compilation. Human-in-the-loop without friction.

## Failure Modes (Known, Unsolved)

**Identity problem:** Synonymous concepts extracted from related sources become duplicate pages ("Cognitive Dissonance Marketing" and "Cognitive Dissonance and Urgency" from the same book). Two-phase concept merging and hash deduplication help, but are partial solutions.

**Level problem:** Macro themes and tactical findings end up on the same plane. "Harness Engineering" and "Go Empty Struct Pattern" coexist at the same level. Schema-based page typing (`overview`, `concept`, `entity`, `comparison`) is the partial answer, but requires explicit configuration.

**Relationship problem:** All connections reduce to "related." "Similar to," "contains," "contradicts" collapse to one word. The graph is useful for navigation but poor for reasoning. `contradictedBy` in frontmatter is the only typed edge most implementations have today.

**Scale ceiling:** Works well at 100 sources / a few hundred wiki pages. Navigation degrades at 10,000+ sources. At that point, a retrieval layer on top starts looking like RAG again.

## This Vault's Position

The Nexus vault is an LLM Wiki implementation in the sense Karpathy defined: the agent maintains a persistent, interlinked wiki between the user and raw sources, with the wiki as the primary knowledge artifact. The vault's specific design choices:

- `vault-tools` binary handles all file operations deterministically (frozen mechanics)
- CLAUDE.md provides the schema: frontmatter format, zone rules, link conventions (frozen schema)
- Ingest skill handles compilation reasoning (agentic, but within frozen schema constraints)
- Skills system enables capability acquisition without harness changes (agentic extension point)

This places the vault firmly in the "hybrid" quadrant — more frozen than pure AGENTS.md (the binary doesn't guess), more agentic than a full Python pipeline (no deterministic extraction stage). The trade-off: more flexible for a single-user knowledge vault, less reproducible than a production pipeline.

Open question: the [[magpie]] project (vault-tools → magpie migration) is the moment to decide whether to move further toward the pipeline end of the spectrum.

## See Also

- [[Agent Memory Architectures]] — where LLM Wiki sits in the RAG/Wiki/Fat Skills decision framework
- [[Harness Engineering]] — the memory layer dimension of harness engineering; knowledge compilation is how agents maintain auditable memory
- [[Nexus Vault Template]] — architecture documentation for this vault as an LLM Wiki implementation
- [[magpie]] — active project evolving this vault's implementation
- [[Claude Code Memory Architecture]] — Claude Code's own 4-layer memory system (CLAUDE.md as mini-wiki is one layer)
- [[Vault Retrieval Layer Threshold]], [[Vault LLM Wiki Failure Modes]], [[Vault Frozen-Agentic Position]], [[Typed Wikilinks Tractability]] — open questions on vault architecture
- [[Magpie Claim-Level Provenance]], [[Magpie Pipeline vs Hybrid]] — open questions on magpie design
