---
title: "Transformers, the tech behind LLMs | Deep Learning Chapter 5"
type: video
captured-date: 2026-06-01
source-url: "https://www.youtube.com/watch?v=wjZofJX0v4M"
---

# Transformers, the tech behind LLMs | Deep Learning Chapter 5

**Channel:** 3Blue1Brown
**Duration:** ~27 min
**Series:** Deep Learning / Essence of Neural Networks — Chapter 5

## Overview

GPT = Generative Pre-trained Transformer. The Transformer is the core invention underlying the current AI boom. This chapter covers the complete end-to-end data flow and foundational background (embeddings, softmax, dot products as similarity) needed before diving into attention in Ch. 6.

## The Data Flow (High Level)

```
Input text
  → tokenize → token IDs
  → embedding matrix lookup → sequence of vectors
  → [attention block → MLP block] × N layers
  → unembedding matrix → logits
  → softmax → probability distribution over next tokens
```

At inference, sample from this distribution, append to input, repeat. This is what every LLM chatbot is doing token by token.

## Tokenization

Input is broken into **tokens** — words, word-pieces, punctuation, or other common character combinations. GPT-3 vocabulary: 50,257 tokens. For images/audio, tokens are patches or chunks.

## Embedding Matrix

The first weight matrix, `W_E`:
- Shape: 12,288 × 50,257 (GPT-3)
- Each column = the embedding vector for one token
- Parameters: ~617 million
- Learned during training from random initialization

**Embedding geometry:** Directions in the high-dimensional embedding space carry semantic meaning — this emerges from training, not design.

- King − Man + Woman ≈ Queen
- Italy − Germany + Hitler ≈ Mussolini
- Plural direction: words closest to (cats − cat) are other plurals
- The model finds it useful to organize the space this way

**Dot product as similarity:** `v · w` is positive when vectors align, zero when perpendicular, negative when opposite. This property is used throughout the attention mechanism.

## Vectors as Context Containers

Initially each vector is just the embedding of one token — no context. The network's job is to transform these vectors so they encode rich, context-dependent meaning by the final layer. Example: "model" in "machine learning model" vs. "fashion model" — the same token embedding needs to become different vectors after attention.

## Context Window

GPT-3 processes 2,048 tokens at once, each with 12,288 dimensions. The entire array of 2,048 × 12,288 values flows through all layers. The context limit explains why early ChatGPT lost the thread of long conversations.

## Attention and MLP Blocks

Two alternating operation types (details in Ch. 6+):

**Attention blocks:** Vectors "talk to each other" — pass information between positions to update their values based on context. The heart of the Transformer.

**MLP (multi-layer perceptron) blocks:** Each vector is processed independently and in parallel through the same operation. Interpretable as asking a long list of questions about each vector and updating based on answers.

All operations are matrix multiplications. GPT-3's 175B weights are organized into ~28,000 distinct matrices across 8 categories.

## Unembedding Matrix

The last weight matrix, `W_U`:
- Shape: 50,257 × 12,288 (GPT-3) — transpose of embedding matrix shape
- Maps the final-layer vector(s) to logits over the vocabulary
- Parameters: ~617 million
- Training efficiency: every vector in the final layer simultaneously predicts what comes after its position (not just the last one)

Running parameter count after embedding + unembedding: ~1.2B of 175B total.

## Softmax

Converts arbitrary real-valued logits to a valid probability distribution:

1. Raise `e` to the power of each logit (makes all values positive)
2. Divide each by the sum (normalizes to sum = 1)
3. Largest inputs dominate; smaller inputs get near-zero weight

**Temperature T** (added as `e^(logit/T)`):
- T → 0: always picks the most probable token (greedy/deterministic)
- T = 1: standard sampling
- T > 1: flatter distribution, more varied/surprising output
- T < 1: sharper distribution, more conservative output

## Weights vs. Data

Sharp distinction emphasized throughout:
- **Weights** (blue/red in 3B1B animations): learned during training, fixed at inference, "the brains"
- **Data** (gray): the specific input for this run, transformed layer by layer

The entire network is a function that transforms data using weights. Weights are what training optimizes.

## Deep Read

**Key Insight:** The embedding matrix is the model's entire vocabulary of concepts, and its geometry is learned — not designed. King−Man+Woman≈Queen emerges because that organization minimized prediction error across billions of examples. Directions in embedding space are an emergent language the model invented for itself.

**What Surprised Me:** The unembedding matrix is almost the transpose of the embedding matrix (same shape, reversed dimension order). In some models they're literally tied — the same weights shared. This suggests the model uses one geometric space both to represent input meaning and to score output candidates: the same "directions" that encode meaning on the way in also score relevance on the way out.

**Open Questions:**
- GPT-3 has 175B parameters but only ~1.2B in embedding + unembedding. Where are the other ~174B? (Ch. 6/7 on attention and MLP blocks should answer this.)
- Temperature above T=2 is blocked by the API as "arbitrary." Is there a principled reason higher temperatures degenerate so sharply, or is it purely empirical?
- The context window of 2048 is a training-time choice. What actually breaks with larger context — compute cost, positional encoding, or something else?

**Wikilink Candidates:**
- [[Transformer Architecture]] — new page created from this source; home for Ch. 5+
- [[Tokenization]] — not yet a wiki page; the mechanics of breaking text into tokens
- [[Softmax]] — not yet a wiki page; used at the output layer and inside attention

**Connections:**
- [[Neural Network Mechanics]] — Chs. 1–4 cover the foundations this chapter builds on; series progress updated
- [[LLM Mental Model]] — next-token prediction loop, context window, and weights-vs-data distinction all map directly to this page

**Image Candidates:** none (transcript-only source)
