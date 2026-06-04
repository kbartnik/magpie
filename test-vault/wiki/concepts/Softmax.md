---
type: concept
tags: [ml, math, activation-function]
---

# Softmax

Converts a vector of arbitrary real-valued **logits** into a valid probability distribution (all values positive, sum to 1). Used at the output layer of [[Transformer Architecture]] to score the next token, and inside the attention mechanism to score which positions to attend to.

**Source:** [[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms|3Blue1Brown — Transformers Ch. 5]]

---

## The Operation

Given logits `z₁, z₂, ..., zₙ`:

```
softmax(zᵢ) = e^zᵢ / Σ e^zⱼ
```

1. Raise `e` to the power of each logit — makes all values positive
2. Divide each by the sum — normalizes to sum = 1
3. The largest inputs dominate exponentially; small inputs get near-zero weight

## Temperature Scaling

Temperature `T` is applied as `e^(zᵢ/T)`:

| T | Effect |
|---|--------|
| T → 0 | Argmax — always picks the highest logit (greedy decoding) |
| T = 1 | Standard softmax |
| T < 1 | Sharper distribution — concentrates probability on top tokens |
| T > 1 | Flatter distribution — spreads probability, more surprising outputs |

Temperature is a sampling-time hyperparameter, not a learned weight. Higher T increases creativity at the cost of coherence; lower T increases consistency at the cost of variety.

## In Attention

Inside attention blocks, softmax converts raw attention scores (dot products between query and key vectors) into **attention weights** — how much each position contributes to updating a given position's vector.

---

## Related

- [[Transformer Architecture]] — softmax at the output layer and inside attention
- [[Tokenization]] — the vocabulary over which softmax produces a distribution
