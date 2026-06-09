---
tags: [concept, transformers, llm]
cluster: transformers
aliases: ["transformers", "LLM architecture", "attention architecture", "GPT architecture"]
related: ["Tokenization", "Softmax", "Attention Mechanism", "Backpropagation", "RLHF"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms]]"
  - "[[archive/videos/2026-06-04-3b1b-attention-transformers]]"
---

# Transformer Architecture

End-to-end data flow: **tokenize → embed → [attention + MLP] × N → unembed → softmax → sample**

## Data Flow

1. **Tokenization** — text split into tokens (subwords); each token gets an integer ID
2. **Embedding** — token IDs mapped to high-dimensional vectors (W_E matrix)
3. **Positional encoding** — position information added to embeddings
4. **Transformer blocks** × N — alternating attention layers (tokens exchange information) and MLP layers (per-token transformation)
5. **Unembedding** — final vector projected back to vocabulary space (W_U matrix)
6. **Softmax** — logits converted to probability distribution over next token

## Embedding Geometry

The embedding space is not arbitrary — it encodes semantic relationships. Famous example: `king - man + woman ≈ queen`. These relationships are emergent from training, not designed.

W_E (embedding matrix) and W_U (unembedding matrix) are approximate transposes. Some models tie these weights entirely, halving parameter count for those matrices.

## Scale

GPT-3: 175B parameters. Most are in the MLP layers (~2/3), not the attention mechanism. Each attention block uses 96 heads × 128 dimensions = 12,288 total attention dimensions per layer.

## Connections

- [[Tokenization]] — first stage of the pipeline; context limits are in tokens, not characters
- [[Softmax]] — final stage; also used inside attention to convert scores to weights
- [[Attention Mechanism]] — the mechanism that lets tokens exchange information across positions
- [[Backpropagation]] — transformers are trained via gradient descent through backpropagation
- [[Transfer Learning]] — pre-trained transformer weights are the source model for fine-tuning
- [[Embeddings]] — the W_E matrix produces token embeddings; intermediate representations are contextual embeddings
