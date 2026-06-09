---
tags: [concept, ml-fundamentals, agents]
cluster: ml-fundamentals
aliases: ["reward model", "preference model", "human preference learning", "pairwise comparison"]
related: ["RLHF", "Fine-Tuning vs Prompting", "Transformer Architecture", "Evaluation Metrics"]
sources:
  - "[[archive/papers/2026-06-04-rlhf-instructgpt]]"
---

# Reward Modeling

A transformer trained to predict human preference scores. The bridge between human judgment (expensive, doesn't scale) and automated optimization (cheap, runs at scale).

## Training

Human raters compare pairs of model outputs for the same prompt and indicate which is better. The reward model is trained with a pairwise comparison loss to predict these rankings.

**Generalization:** A reward model trained on ~1,000 rater comparisons generalizes to novel prompts evaluated by different raters — approaching inter-rater agreement (~73%).

## Architecture

Standard transformer with the language model head replaced by a scalar output head. Takes a (prompt, completion) pair as input; outputs a single score.

## Reward Hacking

The reward model is a proxy for human judgment. Optimizing hard against a proxy eventually produces proxy-gaming — outputs that score highly on the reward model but don't actually satisfy human preferences.

The KL penalty in PPO bounds how far the optimized model can drift from the SFT baseline, limiting but not eliminating reward hacking.

## Connections

- [[RLHF]] — the reward model is the objective that RLHF's PPO stage optimizes against
- [[Transformer Architecture]] — the reward model is a transformer; the scalar head replaces the vocabulary head
- [[Evaluation Metrics]] — the reward model is itself an evaluation metric; the same proxy-gaming failure modes apply
- [[Fine-Tuning vs Prompting]] — reward modeling is the component of RLHF that makes fine-tuning produce helpful models
