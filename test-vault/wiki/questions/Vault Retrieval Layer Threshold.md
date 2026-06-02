---
title: "Vault Retrieval Layer Threshold"
question: "At what vault corpus size and query complexity does the current compiled-wiki approach need augmentation with a retrieval layer?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-rag-llm-wiki-gbrain-comparison.md"
related:
  - "[[LLM Wiki Pattern]]"
  - "[[Agent Memory Architectures]]"
  - "[[Retrieval-Augmented Generation]]"
tags:
  - vault-meta
  - architecture
---

# Vault Retrieval Layer Threshold

*At what vault corpus size and query complexity does the current compiled-wiki approach need augmentation with a retrieval layer?*

The current architecture is a compiled wiki: the agent reads wiki pages directly rather than retrieving from raw sources. This works when total wiki content fits in a few context windows and the agent can identify relevant pages from their titles and the index. The failure mode is scale: as the wiki grows, loading all relevant pages becomes expensive and the index becomes harder to navigate.

The threshold depends on query type. Factual lookups degrade earlier (agent can't hold all relevant facts in context); synthesis queries degrade later (the agent can work from summaries). Semantic search matters more for poorly-titled pages. Current vault: ~90 concept/entity pages, plus questions and syntheses. The threshold is probably several hundred pages.

Tracking this threshold is primarily a measurement problem: when query answer quality noticeably degrades, that's the signal. But that requires a quality metric, which hasn't been established.

## See Also

- [[LLM Wiki Pattern]] — the compile-not-retrieve paradigm; frozen/agentic spectrum
- [[Agent Memory Architectures]] — the decision framework for when RAG is needed
