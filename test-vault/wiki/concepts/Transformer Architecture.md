---
title: "Transformer Architecture"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/videos/2026-06-01-3b1b-transformers-tech-behind-llms.md"
  - "archive/videos/2026-06-01-3b1b-attention-transformers-step-by-step.md"
  - "archive/videos/2026-06-01-3b1b-how-llms-store-facts.md"
related:
  - "Neural Network Mechanics"
  - "LLM Mental Model"
tags:
  - llm-fundamentals
  - deep-learning
  - transformers
---

# Transformer Architecture

The architecture underlying GPT, BERT, and nearly every modern large language model. Built from the 3Blue1Brown deep learning series — updated as each chapter is ingested.

**Series progress:** Ch. 5 ✓ · Ch. 6 ✓ · Ch. 7 ✓ · Ch. 8 (training, RLHF, scaling laws — not yet released)

---

## The Complete Data Flow

```
Input text
  → tokenize → token IDs
  → W_E (embedding matrix) → sequence of vectors
  → [attention block → MLP block] × N layers
  → W_U (unembedding matrix) → logits
  → softmax → probability distribution over next tokens
  → sample → append → repeat
```

This loop — predict next token, sample, append, repeat — is the entirety of what LLM chatbots do. GPT-2 on a laptop produces incoherent text; GPT-3 (same architecture, 100× parameters) produces coherent stories. Scale is the variable.

---

## Chapter 5: Foundations

*Source: "Transformers, the tech behind LLMs" (3B1B, 2024)*

### Tokenization

Input text is broken into **tokens**: words, word-pieces, punctuation, or other common character combinations. GPT-3 vocabulary: 50,257 tokens. For images or audio, tokens are patches or chunks of the signal.

### Embedding Matrix (W_E)

The first weight matrix maps each token to a high-dimensional vector.

| Property | GPT-3 value |
|----------|-------------|
| Shape | 12,288 × 50,257 |
| Parameters | ~617 million |
| Initialization | Random; learned during training |

Each column of W_E is one token's embedding vector. Looking up a token = selecting that column.

**Embedding geometry — the key insight:** Directions in the high-dimensional embedding space carry semantic meaning. This is learned, not designed — the model discovers it minimizes prediction error.

- King − Man + Woman ≈ Queen
- Italy − Germany + Hitler ≈ Mussolini  
- `(cats − cat)` defines a "plurality direction" — dot-producting against nouns ranks plural ones higher
- Directions for gender, nationality, number, and more emerge spontaneously

**Dot product as similarity:** `v · w > 0` when vectors align, `= 0` when perpendicular, `< 0` when opposed. Used throughout the attention mechanism to measure relevance.

### Vectors as Context Containers

Initially, each vector encodes only its token in isolation. The network's primary job is transforming these vectors so they encode rich, context-dependent meaning by the final layer. "Model" in "machine learning model" must become a different vector than "model" in "fashion model" — the attention blocks handle this.

### Context Window

GPT-3 processes 2,048 tokens simultaneously. The data array is always 2,048 columns × 12,288 dimensions. This limit is why early ChatGPT "lost the thread" in long conversations: tokens beyond 2,048 were simply invisible.

### Attention and MLP Blocks (Preview)

Two alternating block types (details in Ch. 6–7):

**Attention blocks:** Vectors communicate across positions — each vector can pull information from others to update its meaning based on context. See Ch. 6 below for full mechanics.

**MLP (feed-forward) blocks:** Each vector is processed independently through the same operation. Interpretable as asking a list of questions about each vector and updating based on the answers. No cross-position communication here.

All operations are matrix multiplications. GPT-3's 175B weights are organized into ~28,000 matrices across 8 categories.

### Unembedding Matrix (W_U)

The final weight matrix maps the last-layer vectors back to vocabulary logits.

| Property | GPT-3 value |
|----------|-------------|
| Shape | 50,257 × 12,288 |
| Parameters | ~617 million |
| Note | Transposed shape of W_E; sometimes weights are tied |

Training efficiency trick: every vector in the final layer simultaneously predicts what token comes *after its position* — not just the last one. This gives the model 2,048 training signal sources per forward pass instead of one.

**Running parameter count:** ~1.2B of 175B total after W_E + W_U. The remaining ~174B live in the attention and MLP blocks.

### Softmax

Converts raw logits (arbitrary real numbers) into a valid probability distribution.

**Mechanism:**
1. Raise `e` to the power of each logit → all positive
2. Divide each by the sum → normalizes to sum = 1
3. Largest logit dominates; others get near-zero weight

**Temperature T** (logit divided by T before exponentiation):

| T | Effect |
|---|--------|
| T → 0 | Always picks the most probable token (greedy) |
| T = 1 | Standard sampling |
| T > 1 | Flatter distribution; less-likely tokens get more weight |
| T < 1 | Sharper; model is more conservative |

Used both at the output layer (next-token distribution) and inside attention blocks.

### Weights vs. Data

A distinction 3B1B emphasizes throughout:

- **Weights** — learned during training, fixed at inference; "the brains of the model"
- **Data** — the specific input for one run; transformed layer by layer

Every computation is: data × weights → new data. Training = adjusting weights via backprop to minimize prediction error.

---

---

## Chapter 6: The Attention Mechanism

*Source: "Attention in transformers, step-by-step" (3B1B, 2024)*

### What Attention Does

Attention transforms context-free token embeddings into context-rich ones. Initial embeddings are lookup-table values — identical for "mole" regardless of whether the context is chemistry, biology, or espionage. After attention, each embedding encodes its token's meaning *given its surroundings*.

More concretely: attention moves information from one embedding to another, potentially across long distances in the context window. It is what allows the final vector in a sequence to summarize everything relevant from the full context when predicting the next token.

### Single Attention Head: Step by Step

**Step 1 — Query and Key vectors**

Every token generates two 128-dimensional vectors (GPT-3):

| Vector | Matrix | Meaning |
|--------|--------|---------|
| Query (q) | W_Q (128 × 12,288) | "What am I looking for?" |
| Key (k) | W_K (128 × 12,288) | "What do I offer?" |

Both matrices are learned parameters. Query and key live in the same 128-dim space so they can be directly compared.

**Step 2 — Attention scores**

Compute the dot product between every key-query pair → a grid of scores. Divide by √128 (the key-query dimension) for numerical stability — prevents scores from growing with dimension size and saturating softmax.

**Step 3 — Softmax → attention pattern**

Apply softmax column-by-column. Each column becomes a probability distribution: how much weight each token contributes when updating this position's embedding.

**Masking:** During training, every token simultaneously predicts what follows it — so later tokens must not "leak" their identity to earlier ones. Before softmax, set all future-position scores to −∞ → those become 0 after softmax, columns stay normalized. Always applied in GPT-style models.

**Step 4 — Value map (low-rank)**

The value map computes *what* to add to attending embeddings. It is factored into two matrices rather than one large one:

- **W_V↓** (128 × 12,288): projects embedding down to 128-dim space
- **W_V↑** (12,288 × 128): projects back up to embedding space

This low-rank factorization keeps the value map's parameter count equal to W_Q and W_K — avoiding a full 12,288×12,288 matrix (150M params). Each head can only propose updates that lie in a 128-dimensional subspace of the full embedding space; this constraint is a feature, not a limitation.

**Step 5 — Weighted sum → embedding update**

For each token position, sum all value vectors weighted by that column's attention pattern. Add the result (ΔE) to the original embedding. Output: refined, context-enriched embeddings.

### The Compact Formula

```
Attention(Q, K, V) = softmax(Q·Kᵀ / √d_k) · V
```

Q, K = full arrays of query/key vectors. V = value-down projection. The value-up (output) projection is applied separately. Softmax applied column-by-column.

### Multi-Headed Attention

A single head captures one type of contextual relationship. A full attention block runs many heads in parallel, each with independent W_Q, W_K, W_V↓, W_V↑.

**GPT-3 numbers:**

| Unit | Parameters |
|------|-----------|
| Per head (4 matrices) | ~6.3M |
| Per attention block (96 heads) | ~600M |
| All attention (96 layers) | ~58B ≈ ⅓ of 175B |

Each head proposes a ΔE per token position. All 96 are summed and added to the embedding. The value-up matrices for all 96 heads are concatenated into a single **output matrix (W_O)** for the block — this is what papers mean by "the output projection."

### Self-Attention vs. Cross-Attention

**Self-attention** (GPT): Q, K, V all come from the same sequence. Masking applied.

**Cross-attention** (translation): Q from the target sequence; K and V from the source. Describes which source tokens each target token should attend to. No masking.

### Why Attention Scales

Attention's architectural advantage is **parallelizability** — every dot product in the attention grid computes simultaneously on GPU. RNNs process tokens sequentially; Transformers process the entire context at once. One decade's lesson: scale produces qualitative improvements. Parallelizability enables scale.

---

---

## Chapter 7: MLP Blocks and Fact Storage

*Source: "How might LLMs store facts" (3B1B, 2024)*

### MLP Structure

Each MLP block applies the same operation to every token vector independently — no cross-token communication. The structure:

```
E_out = E_in + W_down · ReLU(W_up · E_in + b_up) + b_down
```

| Component | Shape (GPT-3) | Role |
|-----------|--------------|------|
| W_up | 49,152 × 12,288 | Expands to "neuron space" (~50K neurons) |
| ReLU / GELU | — | Element-wise nonlinearity |
| W_down | 12,288 × 49,152 | Projects back to embedding space |

The 4× expansion (49,152 = 4 × 12,288) is a convention. The entire structure is architecturally identical to the basic neural network from [[Neural Network Mechanics]] Chs. 1–4 — just operating on context-enriched vectors inside a Transformer.

**"Neurons" in a Transformer** refers specifically to these ~50K intermediate values after the up projection and nonlinearity. Active when positive; inactive (clipped to 0) when negative.

### MLP as AND-Gate Fact Lookup

Toy example: storing "Michael Jordan plays basketball."

1. **W_up row** = (Michael direction + Jordan direction) in embedding space
   → dot product = 2 if E encodes full name, < 2 otherwise
2. **b_up** = −1 for this neuron → positive *only* when both names present
3. **ReLU** clips negatives → neuron fires iff "Michael Jordan"
4. **W_down column** for this neuron = basketball direction
   → when active, adds basketball direction to output

The residual connection (`E_in + ...`) means the block *proposes an addition* rather than replacing the vector. The input embedding is enriched, not overwritten.

This pattern generalizes: each neuron can implement a "query + threshold + response" circuit. Multiple facts coexist in different neurons within the same block.

### Parameter Count (Full GPT-3 Tally)

| Component | Parameters |
|-----------|-----------|
| Embedding W_E | 617M |
| Attention blocks (96 layers) | ~58B |
| MLP blocks (96 layers) | ~116B |
| Unembedding W_U | 617M |
| **Total** | **~175B** ✓ |

MLP blocks account for **~⅔ of all parameters** — the numerical majority, despite attention receiving most of the conceptual focus.

### The Superposition Hypothesis

Individual neurons rarely represent single clean features. Most appear *polysemantic* — responding to unrelated inputs. The likely reason is geometric.

**The constraint:** Features represented by strictly perpendicular directions → max features = dimensions (12,288 for GPT-3). That's far too few.

**The relaxation:** Allow *nearly* perpendicular directions (e.g., 89–91°). The Johnson-Lindenstrauss lemma implies the number of nearly-orthogonal vectors in n-dimensional space grows **exponentially** with n. A 12,288-dim space can encode vastly more than 12,288 independent features — at the cost of small cross-feature noise.

**Consequences:**
- Explains polysemantic neurons: each neuron participates in many superimposed features
- Explains why scaling yields disproportionate capability gains: doubling dimensions exponentially expands feature capacity
- Makes mechanistic interpretability very hard: features don't align with individual neurons

**Tool:** Sparse autoencoders — attempt to find the true sparse feature basis hidden in superimposed neuron activations. See [[Superposition Hypothesis]] for full treatment.

---

## See Also

- [[Neural Network Mechanics]] — Chs. 1–4: the foundational math (neurons, gradient descent, backpropagation) that Transformer training runs on
- [[LLM Mental Model]] — high-level mental model of how LLMs work; Transformer architecture is the "how" behind that model
- [[Superposition Hypothesis]] — the geometric explanation for polysemantic neurons and why LLMs scale so well
- [[Tokenization]] — dedicated page for BPE subword tokenization and its non-intuitive properties
- [[Softmax]] — dedicated page covering temperature, attention use, and output layer use
- [[Sparse Autoencoder]] — the interpretability tool for extracting superimposed features from MLP activations
