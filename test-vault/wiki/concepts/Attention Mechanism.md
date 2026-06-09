---
tags: [concept, transformers, llm]
cluster: transformers
aliases: ["self-attention", "QKV attention", "scaled dot-product attention", "query key value"]
related: ["Transformer Architecture", "Multi-Head Attention", "Softmax", "Embeddings"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-attention-transformers]]"
---

# Attention Mechanism

The mechanism that moves information *between* token positions in a transformer. Without attention, each token's vector updates would be entirely local.

## Query-Key-Value

Three learned weight matrices per attention head:

```
q_i = W_Q · x_i    (what am I looking for?)
k_j = W_K · x_j    (what do I offer?)
score_ij = q_i · k_j / √d_k   (how relevant is j to i?)
attn_ij = softmax(score_ij)    (normalized weight)
out_i = Σ attn_ij · (W_V · x_j)  (weighted sum of values)
```

## Why Scaled?

Dot products grow in magnitude with dimension. Dividing by √d_k keeps scores in the range where softmax has useful gradients (not saturated near 0 or 1).

## Causal Masking

For autoregressive generation, each token may only attend to *prior* tokens. Future positions' scores are set to -∞ before softmax → zero weight. Each token's representation encodes only past context.

## Attention Is Learned

Q, K, V matrices are learned from data — the model discovers its own notion of "relevance." It is not a lookup table; it's a learned similarity function.

## Connections

- [[Transformer Architecture]] — attention is the inter-position communication layer; MLP is the per-position transformation
- [[Multi-Head Attention]] — multiple attention heads run in parallel; each learns different relationship types
- [[Softmax]] — converts raw attention scores to normalized weights
- [[Embeddings]] — attention operates on token embedding vectors, updating them in place
