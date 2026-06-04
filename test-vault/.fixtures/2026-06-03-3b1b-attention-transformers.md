---
title: "Attention in Transformers, Visually Explained | Deep Learning Chapter 6"
type: video
captured-date: 2026-06-03
source-url: "https://www.youtube.com/watch?v=eMlx5fFNoYc"
---

# Attention in Transformers, Visually Explained | Deep Learning Chapter 6

**Channel:** 3Blue1Brown
**Duration:** ~26 min
**Series:** Deep Learning / Essence of Neural Networks — Chapter 6

## Overview

The mechanism that makes transformers powerful: self-attention. Chapter 5 established that vectors flow through attention and MLP blocks; this chapter explains what attention *does* — how tokens at different positions exchange information, and why the query-key-value formulation works.

## The Core Idea

Every token in the sequence needs to update its vector based on context — the same word "bank" means something different in "river bank" vs. "bank account." Attention is the mechanism by which each token looks at every other token and decides how much to pull from it.

The intuition: each token *asks a question* (query), every other token *offers an answer* (key), and the answer content is a *value*. The dot product of a query with a key measures how well that key answers that question. High dot product = high attention weight = more of that value flows into the querying token's update.

## Query, Key, Value Matrices

Three learned weight matrices for each attention head:
- **W_Q** — projects each token's vector into a query
- **W_K** — projects each token's vector into a key  
- **W_V** — projects each token's vector into a value

For token i:
```
q_i = W_Q · x_i      (what am I looking for?)
k_j = W_K · x_j      (what do I offer?)
score_ij = q_i · k_j  (how relevant is j to i?)
```

The scores are passed through softmax (scaled by √d_k to prevent vanishing gradients) to produce attention weights. Each token's update is a weighted sum of values from all positions.

## Multi-Head Attention

Run the attention mechanism H times in parallel, each with its own W_Q, W_K, W_V. Each head learns to attend to different *types* of relationships simultaneously:
- One head might track grammatical subject-verb agreement
- Another might track coreference (which pronoun refers to which noun)
- Another might track positional proximity

The outputs of all heads are concatenated and projected back to the model dimension. GPT-3 uses 96 attention heads per layer across 96 layers.

## Attention Is Not Lookup

A common misconception: attention is not a soft database lookup where queries retrieve matching keys. The Q, K, V matrices are *learned* — what counts as a "matching" query-key pair is not hardcoded but discovered from training data. The model invents its own notion of relevance.

## Positional Encoding

Attention is inherently position-agnostic — the mechanism treats the sequence as a set, not an ordered list. Positional information is injected by adding a positional encoding vector to each token's embedding before the first layer. Original transformers used sinusoidal encodings; modern LLMs use learned or rotary positional embeddings (RoPE).

## Causal Masking

For autoregressive generation, each token may only attend to tokens *before* it in the sequence. This is enforced by setting future positions' attention scores to -∞ before softmax, so they receive zero weight. The result: each token's vector encodes only information from its past context.

## Deep Read

**Key Insight:** Attention is the part of the transformer that moves information *between* positions. Without it, each token's vector updates would be entirely local — the same computation applied independently at every position. Attention is what allows "bank" in position 7 to know that "river" appeared in position 2 and update accordingly.

**What Surprised Me:** The Q and K weight matrices are typically much smaller than the model dimension — GPT-3 uses d_k=128 per head despite a model dimension of 12,288. Each head operates in a compressed 128-dimensional subspace. This means the model has 96 × 128 = 12,288 dimensions of "question-asking" capacity per layer — exactly the model dimension, but distributed across specialized heads rather than one monolithic attention.

**Open Questions:**
- Attention heads in early layers appear to track syntactic relationships; later layers track more semantic ones. Is this a consistent finding across architectures, or an artifact of the models studied?
- Causal masking means each token only sees its past. Encoder models (BERT) use bidirectional attention — every token can attend to every other. Is there a clean way to characterize which tasks benefit from bidirectionality vs. which work fine with causal masking?
- RoPE (Rotary Position Embedding) has become the dominant positional encoding in new LLMs. What property does it have that learned or sinusoidal encodings lack, and why does it generalize better to longer contexts?

**Wikilink Candidates:**
- [[Attention Mechanism]] — the full QKV formulation, scaled dot-product attention, causal masking; primary source
- [[Multi-Head Attention]] — running H attention heads in parallel; concatenate and project; what each head specializes in
- [[Query Key Value]] — the Q, K, V matrix projections and their intuition; not yet a page

**Connections:**
- [[Transformer Architecture]] — attention blocks are the core of the transformer; this chapter fills in what Ch.5 described as "vectors talk to each other"
- [[Softmax]] — used inside attention to convert raw scores to attention weights; temperature scaling applies here too
- [[Tokenization]] — positional encoding is added per token; the attention pattern is over token positions, not characters
