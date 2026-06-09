---
tags: [concept, transformers, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["text embeddings", "vector embeddings", "dense representations", "embedding space", "semantic embeddings"]
related: ["Vector Similarity", "Transformer Architecture", "Tokenization", "Retrieval-Augmented Generation"]
sources:
  - "[[archive/clippings/2026-06-04-embeddings-and-semantic-search]]"
---

# Embeddings

Dense fixed-length vector representations of text where semantic proximity is encoded as geometric proximity.

## Production

Text is passed through a transformer encoder; the final hidden state (or pooled representation) is the embedding. Embedding models are trained with **contrastive loss**: similar pairs → high cosine similarity; dissimilar pairs → low similarity.

## Cosine Similarity

`cos(θ) = (A · B) / (|A| |B|)` — the cosine of the angle between two vectors.
- 1.0: identical direction
- 0.0: orthogonal (unrelated)
- -1.0: opposite

## Embedding Similarity ≠ Semantic Similarity

Embedding similarity reflects **training distribution similarity**. Two semantically related concepts with low co-occurrence in training data will have low cosine similarity. Practical consequences:
- Domain-specific terms (medical, legal, code) need domain-fine-tuned models
- Cross-lingual similarity degrades without multilingual training

## Applications

- **Semantic search:** embed query + documents; retrieve nearest neighbors
- **Clustering:** group by topic without predefined categories
- **RAG retrieval:** find relevant chunks by embedding similarity
- **Classification:** linear classifier on embeddings

## Connections

- [[Vector Similarity]] — cosine similarity and ANN search are the operational tools for embedding-based retrieval
- [[Transformer Architecture]] — embeddings are produced by transformer encoders; W_E produces token embeddings; intermediate layers produce contextual embeddings
- [[Tokenization]] — the embedding context limit is in tokens; long documents must be chunked
- [[Retrieval-Augmented Generation]] — embeddings are the retrieval mechanism; their quality sets the retrieval ceiling
