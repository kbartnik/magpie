---
title: "Retrieval-Augmented Generation"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/videos/2026-06-01-ibm-rag-evolution-agentic-ai.md"
related:
  - "Agent Memory Architectures"
  - "LLM Mental Model"
  - "Agentic Workflow Patterns"
  - "Tool Use in AI Systems"
tags:
  - agentic-systems
  - retrieval
  - llm-fundamentals
---

# Retrieval-Augmented Generation

RAG gives LLMs external memory without retraining — grounding generation in retrieved context rather than relying on frozen training-time knowledge. It emerged as the direct answer to the LLM knowledge boundary problem and has since evolved from a simple linear pipeline into a fully agentic reasoning component.

> "The hardest part of AI isn't generation. It's deciding what to look at." — Sam Anthony, IBM Technology

See [[Agent Memory Architectures]] for how RAG compares to LLM Wiki and Fat Skills as a memory strategy.

## The Problem RAG Solves

LLMs are trained once on a large corpus and then frozen. Their knowledge is locked to what existed in that corpus before a cutoff date. They don't know today's information, and they certainly don't know your organization's specific documents. Retraining is expensive and slow.

RAG separates the knowledge store from the model: retrieve relevant documents at query time, inject them into the prompt, generate against grounded context. Model weights stay frozen; the knowledge layer stays current.

## The Six-Stage Evolution

### Stage 1 — Keyword Search

Inverted indices map terms to documents. TF-IDF and BM25 rank by term frequency and importance. Still powers much of the internet today.

**Limit:** treats words as symbols, not meaning. Synonyms, ambiguity, and complex intent are invisible. The user must find exactly the right words.

### Stage 2 — Semantic Search

Text represented as dense vectors — high-dimensional embeddings learned by neural networks on large corpora. Similar concepts cluster together in vector space regardless of surface form: "espresso" ends up near "coffee" without any explicit synonym rule.

Enables intent understanding without exact keyword match. The user's meaning matters more than their exact words.

### Stage 3 — Hybrid Retrieval

Semantic search complemented rather than replaced keyword search. Hybrid systems bridge keyword *precision* (exact term matching) with semantic *recall* (meaning-based matching). For the first time, search could approximate understanding rather than just matching symbols.

### Stage 4 — LLMs Arrive, Create a New Problem

Large language models trained on massive corpora can generate fluent, useful answers — but only from patterns learned during training. Powerful, but knowledge-locked. The solution to the knowledge boundary problem is retrieval.

### Stage 5 — Basic RAG

**Architecture:** documents embedded offline into a vector database → user query triggers retrieval → top-k documents injected into LLM prompt → answer generated.

**What it solved:** gave LLMs external memory; dramatically reduced hallucinations; enabled domain specialization without retraining; sources became citable.

**Structural limitations** (see [[Agent Memory Architectures]] for the full failure mode catalog):
- **Chunking problem** — long documents split into fragments; related context lands in different vectors
- **Re-derivation problem** — every query starts fresh; prior reasoning doesn't compound
- **Passivity problem** — retrieval only happens when asked; the system never notices contradictions or acts on what it knows
- **Single-retrieval ceiling** — one search step means the answer quality is bounded by the quality of that one query

### Stage 6 — Advanced RAG

Addressed basic RAG's single-retrieval ceiling with a richer pipeline:

| Addition | What it does |
|----------|-------------|
| **Rerankers** | Reorder retrieved results by relevance after initial retrieval |
| **Query rewriting** | Expand or rephrase the user's query to improve recall |
| **Query decomposition** | Break complex questions into sub-queries |
| **Hybrid retrieval** | Combine keyword and semantic search in the retrieval step |

Far more accurate than basic RAG. Still fundamentally *static* — the pipeline is predetermined by the designer. "Retrieval was smarter, but still not intelligent."

### Stage 7 — Agentic RAG

The fixed pipeline becomes a reasoning loop. An AI agent with tools (LLMs, memory, planners, critics, retrievers) decides at runtime:

- **Whether** retrieval is needed at all
- **Where** to search (which knowledge base, which API)
- **What** to ask (query formulation as reasoning, not template)
- **When** enough information has been gathered
- **How** to synthesize across sources, validate claims, resolve contradictions

Retrieval is no longer a pipeline step — it is a *tool invoked as part of reasoning*. The agent can compare sources, iterate on queries, invoke external APIs, handle multimodal data, and perform multi-step research across documents that no single retrieval query would surface.

**What agentic RAG enables that prior stages could not:**
- Multi-step research (retrieve → reason → refine → retrieve again)
- Cross-document synthesis with conflict resolution
- Adaptive behavior when initial retrieval is insufficient
- Tool orchestration across heterogeneous knowledge sources

## The Core Insight

The evolution from keyword search to agentic RAG is not a story of better answers — it is a story of better *question-asking*. Each stage improved the system's ability to decide what to look at, not just how to generate from what it found.

Basic RAG: the pipeline designer decides what to look at.
Agentic RAG: the agent decides what to look at, dynamically, as part of reasoning.

This is the same shift that distinguishes disciplined AI use from vibe-coding: the human (or agent) stays in the reasoning loop rather than accepting whatever the pipeline produces. See [[Vibe-Coding Anti-Pattern]].

## Relationship to Other Architectures

**[[Agent Memory Architectures]]** — RAG is one of four memory strategies. Its structural failure modes (chunking, re-derivation, passivity) explain why the LLM Wiki and Fat Skills architectures exist: they solve the depth and compounding problems RAG cannot.

**[[Agentic Workflow Patterns]]** — agentic RAG is a specific instantiation of the orchestrator + tool-invoking agent pattern; retrieval is a tool among many.

**[[Tool Use in AI Systems]]** — the agentic RAG shift is precisely the move from retrieval as a fixed pipeline step to retrieval as a callable tool with a schema, invocation policy, and observability requirements.

**[[LLM Mental Model]]** — the frozen-after-training property is what makes RAG necessary. Without it, you could just ask the model directly.
