---
tags: [concept, ml-fundamentals, transformers]
cluster: ml-fundamentals
aliases: ["transfer learning", "domain adaptation", "pre-trained models", "feature transfer"]
related: ["Fine-Tuning vs Prompting", "Embeddings", "Transformer Architecture", "Backpropagation", "RLHF"]
sources:
  - "[[archive/clippings/2026-06-04-transfer-learning-fine-tuning]]"
---

# Transfer Learning

Using representations learned on a source task to improve performance on a target task. Dominant paradigm in NLP since BERT (2018).

## Why It Works: Representational Hierarchy

Deep networks learn features at multiple levels of abstraction. In language models:
- **Early layers:** Syntax, morphology, co-occurrence statistics — universal, highly transferable
- **Middle layers:** Sentence structure, semantic roles — moderately transferable
- **Late layers:** Task-specific patterns — less transferable, often replaced or fine-tuned

Earlier layers transfer better; later layers need task-specific adjustment.

## Fine-Tuning Approaches

| Approach | Trainable Params | Use When |
|---|---|---|
| Feature extraction | 0 (frozen backbone) | Target ≈ source task; data limited |
| LoRA | ~0.1% (low-rank adapters) | Standard approach for LLMs |
| Full fine-tuning | 100% | Large target dataset; distant task |

**LoRA:** Represent weight updates as low-rank matrices. 7B parameter model fine-tuned with ~4M parameters (0.06%). Updates are intrinsically low-rank for task adaptation.

## Negative Transfer

When source and target tasks are too dissimilar, transfer hurts. Domain-adaptive pre-training (continue pretraining on domain text before task fine-tuning) bridges the gap.

## Connections

- [[Fine-Tuning vs Prompting]] — transfer learning via fine-tuning vs transfer via prompting are the two main options
- [[Transformer Architecture]] — pre-trained transformers are the dominant source model; their representational hierarchy is what transfers
- [[Embeddings]] — feature extraction uses the transformer as an embedding model; the frozen backbone produces contextual embeddings
- [[Backpropagation]] — fine-tuning is backpropagation starting from pre-trained weights; lower learning rate prevents catastrophic forgetting
