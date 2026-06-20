---
tags: [concept, transformers, llm]
cluster: transformers
aliases: ["BPE tokenization", "subword tokenization", "token vocabulary", "byte pair encoding"]
related: ["Transformer Architecture", "Embeddings", "Softmax"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms]]"
---

# Tokenization

Text is converted to integer token IDs before entering a transformer. Tokenization determines what a model "sees."

## Byte Pair Encoding (BPE)

GPT-style models use BPE: start with individual characters, iteratively merge the most frequent adjacent pairs. Result: common words become single tokens; rare words split into subwords.

GPT-3 vocabulary: 50,257 tokens. Common English words are one token; uncommon words, technical terms, and non-English text require multiple tokens.

## Why Context Limits Are in Tokens

The transformer's attention operates over token positions, not characters. A 128K token limit is ~96K English words or far fewer tokens of code or non-Latin-script text — tokenization efficiency varies dramatically by language and domain.

## Non-Intuitive Splits

- `"unhappiness"` → `["un", "happiness"]` or `["un", "happy", "ness"]` depending on training corpus
- Numbers: `"12345"` might be `["123", "45"]` — arithmetic on tokenized numbers is indirect
- Code: whitespace-heavy languages tokenize less efficiently than dense languages

## Connections

- [[Transformer Architecture]] — tokenization is the first stage of the pipeline
- [[Embeddings]] — each token ID maps to an embedding vector via W_E
- [[Attention Mechanism]] — attention operates over token positions; positional encoding is per-token
