---
tags: [concept, transformers, ml-fundamentals, agents]
cluster: ml-fundamentals
aliases: ["reinforcement learning from human feedback", "RLHF", "InstructGPT", "PPO fine-tuning", "alignment training"]
related: ["Reward Modeling", "Fine-Tuning vs Prompting", "Transformer Architecture", "Backpropagation", "Agentic Workflow Patterns"]
sources:
  - "[[archive/videos/2026-06-04-karpathy-state-of-gpt]]"
  - "[[archive/papers/2026-06-04-rlhf-instructgpt]]"
---

# RLHF

Reinforcement Learning from Human Feedback — the training stage that transforms a base language model into a helpful assistant. Introduced at scale by InstructGPT (Ouyang et al., NeurIPS 2022).

## Three-Stage Pipeline

1. **SFT (Supervised Fine-Tuning):** Fine-tune base model on ~13K human-written (prompt, ideal response) pairs. Teaches assistant behavior format.

2. **Reward Model Training:** Human raters compare pairs of model outputs. A reward model (transformer with scalar head) is trained to predict human preference scores.

3. **PPO Fine-Tuning:** RL optimizes the SFT model to maximize reward model score. A **KL divergence penalty** prevents the model from diverging too far from the SFT distribution, blocking reward hacking.

## RLHF Doesn't Add Knowledge

RLHF steers which knowledge surfaces — it doesn't add capability. The base model already contains the capability; SFT unlocks it; RLHF optimizes it. Capability ceiling is set by pretraining.

## Reward Hacking

Without the KL penalty, PPO finds degenerate outputs that score highly on the reward model but are incoherent. The reward model is a proxy for human judgment; optimizing hard against a proxy games it.

## Alignment Tax

RLHF-tuned models score slightly lower on some capability benchmarks than the base models they're derived from — trading raw capability for helpfulness and safety.

## Connections

- [[Reward Modeling]] — the reward model is the proxy objective RLHF optimizes against
- [[Fine-Tuning vs Prompting]] — RLHF is a specialized fine-tuning pipeline; understanding it clarifies when fine-tuning vs prompting is appropriate
- [[Transformer Architecture]] — both the SFT model and reward model are transformers; PPO operates on the SFT model's output distribution
- [[Backpropagation]] — SFT uses standard gradient descent; PPO uses policy gradients (a generalization)
- [[Agentic Workflow Patterns]] — sycophancy (agreeing with users rather than being accurate) is a known RLHF failure mode that affects agentic reliability
