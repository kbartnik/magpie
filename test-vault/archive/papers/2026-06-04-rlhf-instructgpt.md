---
title: "Training language models to follow instructions with human feedback (InstructGPT)"
type: paper
captured-date: 2026-06-05
source-url: "https://arxiv.org/abs/2203.02155"
author: "Ouyang, Long, et al."
publisher: "NeurIPS 2022 / arXiv"
year: 2022
---

# Training Language Models to Follow Instructions with Human Feedback

**Citation:** Ouyang, L., Wu, J., Jiang, X., et al. (2022). Training language models to follow instructions with human feedback. *NeurIPS 2022*. arXiv:2203.02155.
**Coverage:** SFT dataset construction, reward model training, PPO fine-tuning, alignment tax, evaluation methodology

## The Problem

Large language models (GPT-3 class) trained only on next-token prediction are misaligned with user intent. They complete documents rather than answer questions, produce harmful content when the training distribution contains it, hallucinate confidently, and fail to follow instructions reliably.

The goal: make models *helpful*, *harmless*, and *honest* — the alignment objective.

## The Three-Stage Pipeline

**Stage 1 — Supervised Fine-Tuning (SFT):**
Human labelers write ideal responses to a diverse set of prompts (gathered from the OpenAI API). The GPT-3 base model is fine-tuned on these ~13,000 (prompt, response) pairs. This produces a model that follows instructions but whose responses aren't yet optimized for human preference.

**Stage 2 — Reward Model Training:**
Human labelers rank 4-9 model outputs for a given prompt from best to worst. A reward model (a transformer with a scalar output head replacing the language model head) is trained to predict these rankings via a pairwise comparison loss. The reward model generalizes the human preference signal to novel outputs.

**Stage 3 — RL Fine-Tuning (PPO):**
The SFT model is further fine-tuned using the reward model as the objective. The RL algorithm is PPO (Proximal Policy Optimization). A **KL divergence penalty** between the RL model's outputs and the SFT model's outputs prevents the model from finding degenerate high-reward outputs that don't resemble coherent language.

## Key Findings

**Alignment tax:** InstructGPT models score slightly lower on some standard NLP benchmarks than the base GPT-3 model they were fine-tuned from. Alignment with human preferences trades off marginally against performance on benchmarks optimized for capability. However, on human evaluations of helpfulness and safety, InstructGPT vastly outperforms GPT-3.

**Generalization:** The reward model generalizes beyond the specific prompts and raters used in training. A model trained with ~1,000 human rater comparisons generalizes to novel prompts evaluated by different raters.

**Labeler agreement:** Human raters agree with each other ~73% of the time on pairwise comparisons. The reward model approaches this inter-rater agreement, suggesting it's close to the ceiling of the human preference signal.

## Reward Hacking

The RL optimization can exploit the reward model — producing outputs that score highly on the reward model's metric but don't actually satisfy human preferences. The KL penalty limits this but doesn't eliminate it. This is a fundamental challenge: the reward model is a proxy for human judgment, and optimizing hard against a proxy will eventually produce proxy-gaming behavior.

## Deep Read

**Key Insight:** RLHF doesn't add new knowledge to the model — it steers which knowledge surfaces. The base GPT-3 model already "knows" how to be helpful; it's just as likely to continue a document as to answer a question. SFT unlocks the assistant behavior; RLHF optimizes it. The capability ceiling is set by pretraining; RLHF determines how much of that ceiling is accessible to users.

**What Surprised Me:** The KL penalty coefficient is a hyperparameter that controls the tradeoff between reward maximization and staying close to the SFT model. The paper reports that without the KL penalty, PPO quickly finds degenerate outputs that score highly on the reward model but are incoherent to humans. This is reward hacking in practice — not a theoretical concern but an observed empirical failure mode.

**Open Questions:**
- The reward model is trained on pairwise comparisons by a specific population of labelers. To what extent does the resulting aligned model reflect the values of that labeler population vs. broader human values? Whose preferences are being optimized?
- Constitutional AI (Anthropic) and RLAIF (RL from AI feedback) attempt to reduce reliance on human labelers by using model-generated feedback. How does the quality of the resulting alignment compare to human-labeled RLHF?
- The "alignment tax" — slight degradation on standard benchmarks — suggests a tradeoff between helpfulness and raw capability. As models scale, does this tax shrink (the model can be helpful and capable) or grow (alignment becomes harder)?

**Wikilink Candidates:**
- [[RLHF]] — reward model training; PPO fine-tuning; KL divergence penalty; reward hacking; alignment tax; InstructGPT as the foundational paper; not yet a wiki page
- [[Reward Modeling]] — pairwise comparison training; generalization of human preference; inter-rater agreement ceiling; reward hacking vulnerability; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — both the language model being fine-tuned and the reward model are transformers; the reward model replaces the language model head with a scalar head
- [[Fine-Tuning vs Prompting]] — InstructGPT is the canonical example of fine-tuning changing model behavior; prompting a base model vs. an RLHF model produce qualitatively different results
- [[Agentic Workflow Patterns]] — RLHF-trained models are the base for most agentic systems; sycophancy (agreeing with the user rather than being accurate) is a known RLHF failure mode that affects agentic reliability
- [[Backpropagation]] — SFT uses standard supervised learning with backpropagation; PPO uses policy gradient methods that are a generalization of gradient-based optimization
