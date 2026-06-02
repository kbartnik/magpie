---
title: "Magpie Pipeline vs Hybrid"
question: "Should magpie evolve toward a batch pipeline architecture for ingest, or preserve the interactive hybrid model where an agent makes judgment calls about what's worth capturing?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-llm-wiki-compiler-deep-dive.md"
related:
  - "[[magpie]]"
  - "[[LLM Wiki Pattern]]"
tags:
  - vault-meta
  - magpie
---

# Magpie Pipeline vs Hybrid

*Should magpie evolve toward a batch pipeline architecture for ingest, or preserve the interactive hybrid model where an agent makes judgment calls about what's worth capturing?*

The LLM Wiki compiler implements a batch pipeline: ingest runs, claims are extracted, wiki is updated in batch. This vault uses a hybrid model: interactive ingest, one source at a time, with the agent deciding what's worth capturing based on what the wiki already contains.

The pipeline model offers reproducibility and auditability: the same sources always produce the same wiki state. The hybrid model offers context-sensitivity: the agent can notice that a new source connects to something the wiki already discusses and make that connection explicit, or decide a source is too thin to add. Batch pipelines require upfront specification of what to extract; the hybrid model defers this decision to runtime.

For a vault that grows organically, the hybrid model captures connections that a pipeline would miss. For a vault with well-defined, recurring ingestion requirements, a pipeline is more reliable and cheaper. Whether magpie's long-term use pattern will look more like the former or the latter is unclear.

## See Also

- [[magpie]] — current design and phase roadmap
- [[LLM Wiki Pattern]] — frozen/agentic spectrum; pipeline vs hybrid as a position on that spectrum
