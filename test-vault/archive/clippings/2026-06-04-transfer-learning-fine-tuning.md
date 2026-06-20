---
title: "Transfer Learning and Fine-Tuning: The Practical Guide"
author:
  - "Sebastian Ruder"
url: "https://ruder.io/transfer-learning/"
description: "A comprehensive survey of transfer learning in NLP and deep learning: when and why representations learned on one task transfer to another, how fine-tuning works, and what determines whether transfer is positive or negative."
tags:
  - "clippings"
  - "ml-fundamentals"
  - "transfer-learning"
  - "fine-tuning"
published: 2019-03-21T00:00:00Z
created: 2026-06-05T10:30:00-04:00
---

# Transfer Learning and Fine-Tuning: The Practical Guide

By Sebastian Ruder
Published: 2019-03-21 (updated through 2024)

>[!summary]
>Transfer learning uses representations learned on a source task to improve performance on a target task. In deep learning, this works because neural networks learn a hierarchy of features: early layers learn general representations that transfer across tasks; later layers learn task-specific representations that must be fine-tuned or replaced. Pre-trained language models have made transfer learning the dominant paradigm in NLP.

## The Representational Hierarchy

Deep networks learn features at multiple levels of abstraction. In a vision network:
- **Early layers:** edges, textures, color gradients (highly transferable)
- **Middle layers:** object parts, patterns (moderately transferable)
- **Late layers:** object identities specific to the training task (less transferable)

In language models:
- **Early layers:** syntax, morphology, low-level co-occurrence statistics
- **Middle layers:** sentence structure, semantic roles
- **Late layers:** task-specific patterns (for BERT: NSP, MLM; for GPT: next-token prediction)

This hierarchy is why transfer works: the lower layers capture universal knowledge; only the upper layers need task-specific adjustment.

## Fine-Tuning Approaches

**Full fine-tuning:** Update all parameters of the pre-trained model on the target task. Computationally expensive but achieves best performance when the target dataset is large.

**Feature extraction (frozen backbone):** Freeze the pre-trained weights; train only the final task-specific head. Fast and memory-efficient; works well when target and source tasks are similar and target data is limited.

**Adapter layers:** Insert small trainable modules between pre-trained layers; keep original weights frozen. Enables multi-task fine-tuning without interference.

**LoRA (Low-Rank Adaptation):** Reparametrize weight updates as low-rank matrices. Achieves full fine-tuning quality with 0.1–1% of the trainable parameters. The dominant approach for large model fine-tuning.

## Negative Transfer

Transfer can hurt if source and target tasks are too dissimilar. A model pre-trained on English may transfer poorly to a highly specialized domain (clinical notes, legal contracts, code) where the vocabulary and sentence structure diverge significantly. Domain-adaptive pre-training — continuing pretraining on domain-specific text before task fine-tuning — bridges this gap.

## Why Large Language Models Changed the Paradigm

Before large pre-trained models, NLP required task-specific architectures and large labeled datasets per task. After BERT (2018) and GPT-2 (2019), the paradigm inverted: one large pre-trained model fine-tuned with a small labeled dataset outperforms task-specific models trained from scratch on large datasets.

The transfer is so strong that, in many cases, *prompting* without any fine-tuning (zero-shot or few-shot) outperforms previous state-of-the-art fine-tuned models.

## Deep Read

**Key Insight:** The representational hierarchy in deep networks is what makes transfer learning work — and this is not a coincidence of architecture but a consequence of training on diverse data. A network that learns to predict text tokens on internet-scale data must develop general representations because the task itself is general. The generality of pretraining is the source of transferability.

**What Surprised Me:** LoRA's claim — that fine-tuning a 7B parameter model requires only ~4M trainable parameters (0.06%) — is validated empirically across many benchmarks. The interpretation: the weight updates needed for task adaptation are intrinsically low-rank. The full parameter space is redundant for fine-tuning even if it's necessary for pretraining.

**Open Questions:**
- LoRA represents weight updates as low-rank matrices. Why would task-specific adaptations be low-rank? Is this a fundamental property of the loss landscape near a pre-trained model, or an empirical observation that might break for more distant tasks?
- Negative transfer from domain mismatch is well-documented in NLP. In multimodal models (vision + language), is negative transfer more or less of a problem when source and target modalities differ?
- Continued pretraining on domain data before task fine-tuning helps with domain mismatch. But continued pretraining can cause catastrophic forgetting of general knowledge. Is there a principled way to balance domain adaptation against forgetting?

**Wikilink Candidates:**
- [[Transfer Learning]] — representational hierarchy; early layers transfer, late layers need fine-tuning; pre-trained LM as universal starting point; not yet a wiki page
- [[Fine-Tuning vs Prompting]] — full fine-tuning vs LoRA vs feature extraction vs prompting; when each approach is appropriate; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — pre-trained transformers are the dominant source model for transfer learning; BERT and GPT architectures instantiate the hierarchy described here
- [[Backpropagation]] — fine-tuning is backpropagation on the target task, starting from pre-trained weights; the learning rate is typically much lower than during pretraining to avoid catastrophic forgetting
- [[Embeddings]] — the "frozen backbone" approach to fine-tuning uses the transformer as an embedding model; the late-layer representations are task-specific embeddings
- [[RLHF]] — RLHF is a form of fine-tuning; SFT (supervised fine-tuning) is full fine-tuning on human-labeled examples; the pre-trained base model provides the capability that SFT adapts
