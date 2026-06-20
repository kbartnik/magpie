---
title: "Embeddings: What They Are and Why They Matter"
author:
  - "Simon Willison"
url: "https://simonwillison.net/2023/Oct/23/embeddings/"
description: "A practical explanation of text embeddings: what they are, how they're produced, how cosine similarity works, and what they're actually good for in production applications."
tags:
  - "clippings"
  - "llm"
  - "embeddings"
  - "semantic-search"
published: 2023-10-23T00:00:00Z
created: 2026-06-05T09:30:00-04:00
---

# Embeddings: What They Are and Why They Matter

By Simon Willison
Published: 2023-10-23

>[!summary]
>Embeddings are dense vector representations of text where semantic similarity is encoded as geometric proximity. Produced by running text through a transformer encoder and extracting the final hidden state. Useful for semantic search, clustering, classification, and retrieval-augmented generation — but "semantic similarity" is actually "training distribution similarity."

## What Is an Embedding?

An embedding is a fixed-length vector of floating-point numbers that represents a piece of text. A text embedding model maps strings of arbitrary length to a vector of fixed dimension (e.g., 1536 dimensions for OpenAI's text-embedding-ada-002).

The vectors are organized so that texts with similar meaning are geometrically close — measured by **cosine similarity**: the cosine of the angle between two vectors. Cosine similarity of 1.0 means identical; 0.0 means orthogonal (unrelated); negative values indicate opposite meaning.

## How Embeddings Are Produced

The same transformer architecture used for language models can produce embeddings. Instead of the full autoregressive output, you extract the final hidden state for a special `[CLS]` token (in encoder models like BERT) or pool the last layer's token representations.

Embedding models are trained on large datasets of (text, text) pairs with contrastive loss: semantically similar pairs should have high cosine similarity; dissimilar pairs should have low similarity.

## What They're Good For

**Semantic search:** Store document embeddings in a vector database. Given a query, embed it and find the nearest neighbors. Unlike keyword search, this finds documents that mean the same thing even if they use different words.

**Clustering:** Group documents by topic without predefined categories. K-means on embeddings finds natural topic clusters.

**RAG (Retrieval-Augmented Generation):** Retrieve relevant documents by embedding similarity, then pass them to a language model. The embedding retrieval step is what "gives the model memory."

**Classification:** A linear classifier trained on embeddings often outperforms keyword-based classifiers with much less labeled data.

## The Limits of Embedding Similarity

Embedding similarity is not semantic similarity — it's training distribution similarity. Two concepts can be semantically related but have low cosine similarity if they rarely co-occur in the training data. Conversely, terms that co-occur frequently can have high cosine similarity even if they're conceptually distinct.

Practical consequences:
- Domain-specific jargon (medical, legal, code) often needs fine-tuned embedding models
- Cross-lingual similarity degrades without multilingual training
- Long documents need chunking — embedding a full page loses fine-grained information

## Deep Read

**Key Insight:** The geometric structure of embedding space is not arbitrary — it reflects statistical patterns in the training data. This is what makes embeddings useful (related concepts cluster) and unreliable (the clusters reflect training distribution, not ground truth semantics). Understanding this distinction is essential for building robust RAG systems.

**What Surprised Me:** Willison demonstrates that embeddings encode bias directly: occupational embeddings cluster by gender, nationality embeddings carry cultural associations. This is not a bug introduced by misuse — it's a direct consequence of training on human-generated text. Any system using embeddings for consequential decisions inherits these biases.

**Open Questions:**
- Cosine similarity in 1536 dimensions is a well-defined measure but may not correspond to human intuitions about "similar." Are there tasks where L2 distance or dot product outperform cosine similarity as a retrieval metric?
- Chunking documents for embedding is a known hard problem (semantic chunking vs fixed-size). How sensitive is RAG retrieval quality to chunking strategy, and is there a principled way to choose chunk size for a given domain?
- Embedding models have context limits (typically 512–8192 tokens). Long documents must be split. Does embedding the whole document (with a longer-context model) outperform pooling chunk embeddings, or does the pooling capture enough?

**Wikilink Candidates:**
- [[Embeddings]] — dense vector representations; cosine similarity; contrastive training; semantic search; training distribution ≠ semantic similarity; not yet a wiki page
- [[Vector Similarity]] — cosine similarity; L2 distance; dot product; approximate nearest neighbor search (FAISS, HNSW); not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — embeddings are produced by transformer encoders; the final hidden states are the representations used
- [[Tokenization]] — embedding a text requires tokenizing it first; the embedding is at the token level, not character level; context limit is in tokens
- [[Retrieval-Augmented Generation]] — embeddings are the retrieval mechanism in RAG; the quality of the embedding model sets the ceiling for retrieval quality
- [[Softmax]] — contrastive training uses a softmax over similarity scores; the same function that converts logits to probabilities is used to train embedding models
