---
title: "State of GPT | Microsoft Build 2023"
type: video
captured-date: 2026-06-05
source-url: "https://www.youtube.com/watch?v=bZQun8Y4L2A"
---

# State of GPT | Microsoft Build 2023

**Speaker:** Andrej Karpathy
**Event:** Microsoft Build 2023
**Duration:** ~43 min

## Overview

Karpathy walks through the full training pipeline for GPT-class models: pretraining, supervised fine-tuning (SFT), reward modeling, and reinforcement learning from human feedback (RLHF). The talk is the clearest accessible explanation of how a base language model becomes a useful assistant.

## Stage 1: Pretraining

The base model is trained on internet-scale text to predict the next token. The model learns to complete documents. It has no concept of "being helpful" — given the prompt "Tell me about Paris", it might continue with "Tell me about London, Tell me about Berlin" because the training data contains lists of city questions.

The base model is a document simulator, not an assistant.

## Stage 2: Supervised Fine-Tuning (SFT)

Human contractors write (prompt, ideal response) pairs demonstrating assistant behavior. The model is fine-tuned on this much smaller dataset. The model learns to follow the prompt-response format and produce helpful, coherent answers.

SFT data is expensive to produce (requires human labor) but essential for the behavior shift from document completion to assistant. Typically 10K–100K examples.

## Stage 3: Reward Modeling

Human raters compare pairs of model outputs and indicate which is better. A **reward model** is trained to predict human preference scores. The reward model is itself a language model with a scalar output head.

The reward model is a proxy for human judgment — it generalizes from the rated pairs to unseen outputs.

## Stage 4: RLHF (Reinforcement Learning from Human Feedback)

The SFT model is optimized to produce outputs that maximize the reward model's score. The RL algorithm used is PPO (Proximal Policy Optimization). A KL penalty term prevents the RL-optimized model from drifting too far from the SFT model's output distribution (preventing degenerate reward hacking).

The result is a model that produces outputs humans prefer — more helpful, less harmful, more honest.

## Key Insight: Base vs. RLHF Models

Base models and RLHF models are fundamentally different objects. A base model knows everything in the training data but doesn't "want" to help. An RLHF model has been shaped to be a helpful assistant — but it can only express capabilities that were latent in the base model. RLHF doesn't add knowledge; it steers which knowledge surfaces and how.

## Prompting RLHF Models

RLHF models are prompt-sensitive in ways base models are not. The system prompt matters because it activates the "assistant mode" the model was fine-tuned on. Few-shot examples in the prompt can shift behavior significantly. The model is not doing retrieval — it's using the prompt as a distribution shift signal.

## Deep Read

**Key Insight:** The training pipeline distinction — pretraining → SFT → RLHF — maps directly onto a behavioral distinction. Base models are document completers that happen to be useful for generation tasks. RLHF models are assistants that have been shaped to be helpful. Treating them as the same object leads to wrong intuitions about why prompting works.

**What Surprised Me:** Karpathy's estimate that SFT requires only ~10K examples to produce a large behavioral shift. The pretraining data is billions of tokens; the fine-tuning data is thousands of examples. The ratio is extreme. The implication: the base model already contains the capability to be an assistant — SFT just unlocks it by showing the model the prompt-response format.

**Open Questions:**
- RLHF requires human raters to compare outputs. The reward model generalizes from these ratings. How well does the reward model generalize to novel capability areas not covered in the rating data? Is there evidence of systematic failure modes?
- The KL penalty prevents the RL-optimized model from diverging too far from the SFT model. How is the penalty coefficient chosen? Is there a principled way to set it, or is it empirically tuned?
- Constitutional AI (Anthropic's method) uses AI feedback instead of human feedback for some stages. How does the quality of AI-labeled reward data compare to human-labeled data, and which failure modes does it introduce?

**Wikilink Candidates:**
- [[RLHF]] — reinforcement learning from human feedback; reward modeling; PPO training; KL penalty; base vs aligned models; not yet a wiki page
- [[Fine-Tuning vs Prompting]] — when to fine-tune vs prompt; SFT as behavior unlocking; prompting as distribution shift; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — the base model being fine-tuned is a transformer; the reward model is also a transformer with a scalar head; both stages build on the same architecture
- [[Agentic Workflow Patterns]] — RLHF-trained models behave differently in agentic contexts than base models; the helpfulness training can conflict with agentic reliability (over-confidence, sycophancy)
- [[Prompt Engineering]] — understanding the SFT and RLHF stages explains why prompt engineering works: the model has been trained on prompt-response pairs, and the prompt activates that distribution
- [[Backpropagation]] — both SFT and reward model training use standard gradient descent via backpropagation; the RL stage (PPO) uses policy gradient methods that are more complex
