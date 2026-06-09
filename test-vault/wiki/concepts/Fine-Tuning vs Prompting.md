---
tags: [concept, transformers, ml-fundamentals, agents]
cluster: ml-fundamentals
aliases: ["fine-tuning", "when to fine-tune", "SFT", "supervised fine-tuning", "prompting vs fine-tuning"]
related: ["RLHF", "Transfer Learning", "Prompt Engineering", "Transformer Architecture"]
sources:
  - "[[archive/videos/2026-06-04-karpathy-state-of-gpt]]"
  - "[[archive/clippings/2026-06-04-transfer-learning-fine-tuning]]"
---

# Fine-Tuning vs Prompting

## When to Prompt

- Task is well-covered in the base model's training distribution
- Few-shot examples achieve adequate quality
- No labeled dataset available
- Rapid iteration needed
- Behavioral change is temporary or context-specific

## When to Fine-Tune

- Consistent format/style requirements (structured output, brand voice)
- Domain-specific terminology not in the base model
- Efficiency: distill a complex prompt into model weights
- Privacy: sensitive examples shouldn't be in the context window
- Quality ceiling: prompting has plateaued

## SFT Unlocks, Doesn't Add

Supervised fine-tuning on ~10K examples produces a large behavioral shift because the base model already has the capability. SFT teaches the prompt-response format, not new knowledge.

## LoRA for Efficient Fine-Tuning

Fine-tuning 7B parameters requires only ~4M trainable LoRA parameters (0.06%). The weight updates for task adaptation are intrinsically low-rank. Full fine-tuning is rarely necessary.

## Connections

- [[RLHF]] — RLHF adds a reward-optimization stage on top of SFT; the SFT model is the starting point
- [[Transfer Learning]] — fine-tuning is transfer learning; LoRA exploits the low-rank structure of task adaptation
- [[Prompt Engineering]] — prompting and fine-tuning are complementary; understanding why prompting works informs when to fine-tune instead
- [[Transformer Architecture]] — the base model being fine-tuned is a transformer; fine-tuning only updates upper layers for most task-specific adaptations
