---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["cosine similarity", "ANN search", "approximate nearest neighbor", "HNSW", "FAISS", "vector search"]
related: ["Embeddings", "Retrieval-Augmented Generation"]
sources:
  - "[[archive/clippings/2026-06-04-embeddings-and-semantic-search]]"
---

# Vector Similarity

The computational tools for finding nearest neighbors in high-dimensional embedding spaces.

## Distance Metrics

**Cosine similarity:** Angle between vectors. Scale-invariant — a short and long vector with the same direction are maximally similar. Standard for text embeddings.

**L2 (Euclidean) distance:** Magnitude-sensitive. Used when absolute distances matter (e.g., image embeddings).

**Dot product:** `A · B = |A||B|cos(θ)`. Equivalent to cosine similarity for normalized vectors. Used in attention mechanisms.

## Approximate Nearest Neighbor (ANN)

Exact nearest-neighbor search in 1536 dimensions across millions of vectors is too slow for production. ANN algorithms trade recall for speed:

**HNSW (Hierarchical Navigable Small World):** Graph-based. Build a multi-layer graph of connections; traverse it at query time. Fast and high-recall. Used by pgvector, Weaviate, Qdrant.

**FAISS (Facebook AI Similarity Search):** Index-based. Supports many index types (IVF, PQ, HNSW). Good for large-scale batch retrieval.

## Recall vs Latency Tradeoff

ANN parameters (`ef_construction`, `M`, `nprobe`) tune the recall-latency tradeoff. Higher recall → slower query. Set based on application requirements.

## Connections

- [[Embeddings]] — vector similarity is the operational layer on top of embeddings; the quality of the embedding sets the ceiling for similarity search
- [[Retrieval-Augmented Generation]] — ANN search over chunk embeddings is the retrieval step in RAG
