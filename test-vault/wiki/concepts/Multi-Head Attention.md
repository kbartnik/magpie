---
tags: [concept, transformers, llm]
cluster: transformers
aliases: ["multi-head attention", "MHA", "attention heads", "parallel attention"]
related: ["Attention Mechanism", "Transformer Architecture", "Softmax"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-attention-transformers]]"
---

# Multi-Head Attention

Run the attention mechanism H times in parallel, each with its own W_Q, W_K, W_V matrices. Each head learns to attend to different relationship types.

## Why Multiple Heads?

A single attention head produces one weighted combination of values. Multiple heads allow the model to simultaneously track different relationship types:
- One head: grammatical subject-verb agreement
- Another: coreference (pronoun → antecedent)
- Another: positional proximity

GPT-3: 96 heads per layer × 128 dimensions per head = 12,288 total — exactly the model dimension, distributed across specialized heads.

## Concatenate and Project

After computing all H heads' outputs, concatenate and project back to model dimension:

```
MultiHead(Q,K,V) = Concat(head_1,...,head_H) · W_O
```

`W_O` projects the concatenated output back to the model dimension.

## Compressed Subspaces

Each head operates in a compressed subspace (d_k = 128 for GPT-3, vs model dimension 12,288). The model has 96 × 128 = 12,288 dimensions of "question-asking" capacity per layer — equal to model dimension, but distributed.

## Connections

- [[Attention Mechanism]] — multi-head attention runs the base mechanism H times; the heads specialize
- [[Transformer Architecture]] — multi-head attention is the primary component of each transformer block
- [[Softmax]] — each head applies its own softmax to attention scores
