---
type: concept
tags: [ml, llm, architecture, deep-learning]
---

# Transformer Architecture

The core invention underlying the current AI boom. GPT = Generative Pre-trained Transformer. Every major LLM (GPT, Claude, Gemini) is a transformer.

**Primary source:** [[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms|3Blue1Brown — Transformers, the tech behind LLMs (Ch. 5)]]

---

## End-to-End Data Flow

```
Input text
  → tokenize → token IDs
  → embedding matrix (W_E) lookup → sequence of vectors
  → [attention block → MLP block] × N layers
  → unembedding matrix (W_U) → logits
  → softmax → probability distribution over next tokens
  → sample → append to input → repeat
```

At inference: sample a token from the distribution, append it to the input, repeat. This is what every LLM chatbot is doing — one token at a time.

## Embedding Matrix (W_E)

The first weight matrix. Converts token IDs to dense vectors.

- **GPT-3 shape:** 12,288 × 50,257 — one 12,288-dimensional column per token in the vocabulary
- **Parameters:** ~617 million

**Geometry is learned, not designed.** Directions in the high-dimensional space carry semantic meaning because that organization minimized prediction error during training. The model invented its own internal language:

- King − Man + Woman ≈ Queen
- Plural direction: `(cats − cat)` points toward other plurals
- These regularities emerge; they are not programmed

## Context Window

GPT-3 processes 2,048 tokens simultaneously, each as a 12,288-dimensional vector. The full array (2,048 × 12,288) flows through all layers. The context limit explains why early ChatGPT lost the thread of long conversations.

## Attention and MLP Blocks

Two alternating block types, applied N times:

**Attention blocks** — vectors "talk to each other." Each position can read from and update based on other positions. This is where context-dependent meaning is assembled. (See Ch. 6 for the full mechanism.)

**MLP (multi-layer perceptron) blocks** — each vector is processed independently through the same operation. Interpretable as the model "looking up" stored knowledge about each position's current representation.

All operations are matrix multiplications. GPT-3's 175B weights are ~28,000 distinct matrices across 8 categories.

## Unembedding Matrix (W_U)

The last weight matrix. Maps the final-layer vector to logits over the vocabulary.

- **GPT-3 shape:** 50,257 × 12,288 — the transpose of W_E's shape
- **Parameters:** ~617 million
- In some models, W_U and W_E literally share weights — the same geometric space encodes input meaning and scores output candidates

**Training efficiency:** every vector in the final layer simultaneously predicts what comes *after* its position. Not just the last vector.

## Weights vs. Data

A sharp distinction emphasized throughout:

- **Weights** — learned during training, fixed at inference. The "brains." GPT-3: 175B.
- **Data** — the specific input for this run, transformed layer by layer. Discarded after the run.

Embedding + unembedding account for ~1.2B of GPT-3's 175B parameters. The remaining ~174B live in the attention and MLP layers (covered in Ch. 6/7).

## Temperature

Controls sampling randomness via [[Softmax]]:

| Temperature | Effect |
|-------------|--------|
| T → 0 | Always picks most probable token (greedy/deterministic) |
| T = 1 | Standard sampling |
| T < 1 | Sharper distribution — more conservative |
| T > 1 | Flatter distribution — more varied/surprising |

---

## Related

- [[Tokenization]] — how text becomes token IDs
- [[Softmax]] — converts logits to probability distribution
- [[LLM Mental Model]] — next-token prediction loop, weights-vs-data distinction
