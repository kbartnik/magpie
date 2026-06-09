---
tags: [concept, agents, llm, transformers]
cluster: agents
aliases: ["context management", "context window management", "context design"]
related: ["Agentic Workflow Patterns", "Prompt Engineering", "Retrieval-Augmented Generation", "Working Memory", "Transformer Architecture"]
sources:
  - "[[archive/clippings/2026-06-04-context-engineering-guide]]"
---

# Context Engineering

The discipline of deciding what information appears in the context window — distinct from prompt engineering (how you phrase a request).

## The Distinction

**Prompt engineering:** Optimizes *within* a fixed context — phrasing, instructions, examples.
**Context engineering:** Shapes the *context itself* — what information is present at all.

Prompt engineering cannot exceed the ceiling set by context engineering. Most reliability failures in production AI systems are context failures, not prompt failures.

## Failure Modes

| Failure | Cause | Fix |
|---|---|---|
| **Stale context** | Outdated information | Freshness checks, re-retrieval |
| **Missing context** | Relevant document not retrieved | Better retrieval, retrieval verification |
| **Noisy context** | Too much irrelevant information | Reranking, filtering |
| **Context overflow** | Exceeds window limit | Compression, progressive loading |

## Context Window as Working Memory

The context window is a functional analogue to human working memory: fixed capacity, active manipulation, bottleneck for complex reasoning. Context engineering is working memory management for AI systems.

Larger context windows don't eliminate context engineering — they shift the failure mode from truncation to *dilution* (the model attends to wrong parts of a large context).

## Patterns

- **Progressive context loading:** Start minimal; add information as the task progresses
- **Context compression:** Summarize previous turns before including them
- **Scratchpad separation:** Keep reasoning traces separate from information context
- **RAG:** Pull relevant documents from a vector store at query time

## Connections

- [[Agentic Workflow Patterns]] — context management is a cross-cutting concern for all patterns; hub-and-spoke explicitly isolates context
- [[Prompt Engineering]] — prompt engineering operates within the context; context engineering shapes the context
- [[Retrieval-Augmented Generation]] — RAG is context engineering applied to information retrieval
- [[Working Memory]] — the context window is the AI functional analogue to working memory; context engineering ≈ central executive function
- [[Transformer Architecture]] — attention is quadratic in context length; larger contexts increase compute cost
