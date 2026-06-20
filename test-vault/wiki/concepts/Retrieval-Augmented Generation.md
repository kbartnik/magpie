---
tags: [concept, agents, llm]
cluster: agents
aliases: ["RAG", "retrieval augmented generation", "retrieval-augmented", "vector search retrieval"]
related: ["Embeddings", "Prompt Engineering", "Context Engineering", "Agentic Workflow Patterns", "Transformer Architecture"]
sources:
  - "[[archive/books/2026-06-04-prompt-engineering-for-llms]]"
---

# Retrieval-Augmented Generation

Architecture that grounds LLM responses in retrieved documents. Solves the "model only knows its training data" problem by fetching relevant context at query time.

## Basic Pipeline

1. **Index:** Chunk documents; embed each chunk; store in a vector database
2. **Retrieve:** Embed the query; find nearest-neighbor chunks by cosine similarity
3. **Generate:** Prepend retrieved chunks to the prompt; instruct the model to answer from them

## Chunking Is the Ceiling

Chunk quality sets the ceiling for all downstream stages. A retrieved chunk that doesn't contain the answer can't help the model. Strategies:
- **Fixed-size:** Simple; misses semantic boundaries
- **Semantic:** Split at paragraphs/sections; better but depends on document structure
- **Parent-document:** Store large chunks; retrieve child chunks; return parent for context

## Why "Deciding What to Look At" Is the Hard Problem

The retrieval step is where most RAG failures occur. Embedding similarity is not semantic similarity (see [[Embeddings]]). Domain-specific terminology, query-document vocabulary mismatch, and sparse training coverage all degrade retrieval.

## Connections

- [[Embeddings]] — the retrieval mechanism; embedding quality sets the ceiling for retrieval quality
- [[Context Engineering]] — RAG is context engineering applied to the information retrieval problem
- [[Prompt Engineering]] — retrieved documents need prompt instructions telling the model how to use them
- [[Agentic Workflow Patterns]] — agentic RAG adds routing, query rewriting, and iterative retrieval on top of basic RAG
