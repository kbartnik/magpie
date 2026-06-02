---
title: "Agentic RAG vs Advanced RAG Threshold"
question: "At what query complexity and quality threshold does the overhead of agentic RAG orchestration justify itself over a well-tuned advanced RAG pipeline?"
type: question
status: open
domain: agentic
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-ibm-rag-evolution-agentic-ai.md"
related:
  - "[[Retrieval-Augmented Generation]]"
  - "[[Agent Memory Architectures]]"
tags:
  - rag
  - agent-memory
---

# Agentic RAG vs Advanced RAG Threshold

*At what query complexity and quality threshold does the overhead of agentic RAG orchestration justify itself over a well-tuned advanced RAG pipeline?*

Agentic RAG adds an orchestration layer: the agent decides which retrieval strategies to use, decomposes queries, and iterates. Advanced RAG uses fixed, well-tuned pipelines without agent orchestration. Agentic RAG should win on complex multi-hop queries; Advanced RAG should win on latency-sensitive, well-defined queries.

The unresolved question is empirical: for what query distributions does agentic orchestration overhead (additional LLM calls, higher latency, higher cost) produce quality gains that justify it? No published benchmark fairly compares well-tuned Advanced RAG against well-tuned Agentic RAG on a realistic enterprise query distribution. Most comparisons favor whichever approach the paper authors implemented better.

For this vault: if a retrieval layer is added, this threshold question determines the architecture.

## See Also

- [[Retrieval-Augmented Generation]] — the 6-stage evolution arc; agentic RAG as stage 6
- [[Agent Memory Architectures]] — where RAG fits in the broader memory decision framework
