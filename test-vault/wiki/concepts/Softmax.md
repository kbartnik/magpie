---
title: "Softmax"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-3b1b-transformers-tech-behind-llms.md"
related:
  - "[[Transformer Architecture]]"
  - "[[Neural Network Mechanics]]"
  - "[[Tokenization]]"
tags:
  - llm-fundamentals
  - deep-learning
  - mathematics
---

# Softmax

A function that converts an arbitrary vector of real numbers (logits) into a valid probability distribution — all values positive, summing to 1. Used at the output of language models and inside attention mechanisms.

## The Operation

Given logits `z = [z₁, z₂, ..., zₙ]`:

```
softmax(zᵢ) = e^zᵢ / Σⱼ e^zⱼ
```

1. Raise `e` to the power of each logit → all positive
2. Divide each by the total → normalizes to sum = 1

The largest logit dominates because exponentiation amplifies differences — a logit of 3 gives `e³ ≈ 20` vs. a logit of 1 giving `e¹ ≈ 2.7`. Near-zero logits contribute negligibly.

## Temperature

Dividing all logits by a temperature parameter T before applying softmax controls the sharpness of the distribution:

| T | Effect |
|---|--------|
| T → 0 | Collapses to argmax (always picks the highest logit) |
| T = 1 | Standard softmax |
| T > 1 | Flatter — lower-probability tokens get more weight; more "creative" |
| T < 1 | Sharper — model becomes more conservative and repetitive |

Temperature is a runtime sampling parameter, not a learned weight.

## Two Uses in Transformers

**Output layer:** Applied to the final unembedding logits (one per vocabulary token) to produce the next-token probability distribution. The model samples from this distribution to generate the next token.

**Attention mechanism:** Applied column-by-column to the attention score grid `(Q·Kᵀ / √d_k)` to produce attention weights — the probability distribution over which tokens each position should attend to. Uses the same scaled division to prevent saturation. See [[Transformer Architecture]] § Ch. 6.

## See Also

- [[Transformer Architecture]] — both uses of softmax in the full architecture
- [[Tokenization]] — the vocabulary over which the output softmax distributes probability
- [[Neural Network Mechanics]] — the training objective (cross-entropy loss) that measures how well the softmax distribution matches the correct next token
