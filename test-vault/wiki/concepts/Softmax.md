---
tags: [concept, transformers, ml-fundamentals]
cluster: transformers
aliases: ["softmax function", "temperature scaling", "logits to probabilities"]
related: ["Transformer Architecture", "Attention Mechanism", "Loss Functions"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms]]"
---

# Softmax

Converts a vector of real-valued logits into a probability distribution: all values positive, sum to 1.

## Formula

`softmax(x_i) = e^x_i / Σ e^x_j`

The exponential amplifies differences — a small lead in raw scores becomes a large difference in probabilities.

## Temperature Scaling

Dividing logits by temperature T before softmax:
- **T < 1** — sharpens the distribution (more deterministic, top token dominates)
- **T > 1** — flattens the distribution (more random, lower-probability tokens get more weight)
- **T → 0** — argmax (always pick the highest-scoring token)

## Uses in Transformers

1. **Output layer** — converts final logits to next-token probabilities for sampling
2. **Inside attention** — converts attention scores (QK dot products) to attention weights; scaled by 1/√d_k to prevent saturation

## Connections

- [[Transformer Architecture]] — final stage of the pipeline; also used inside every attention head
- [[Attention Mechanism]] — softmax converts raw QK dot products to attention weights
- [[Loss Functions]] — cross-entropy loss is computed from softmax output probabilities during training
