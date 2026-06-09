---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["data contamination", "benchmark leakage", "test set contamination", "training-test overlap"]
related: ["Evaluation Metrics", "Transfer Learning", "RLHF"]
sources:
  - "[[archive/clippings/2026-06-04-llm-evaluation-benchmarks]]"
---

# Benchmark Contamination

The presence of benchmark test examples in a model's training corpus. Contaminated benchmarks overestimate capability — the model may have memorized answers rather than learned to reason.

## Why It's Hard to Detect

Detection requires knowing the exact training data. Large model providers often don't fully disclose training data composition. Even when disclosed, fuzzy matching is needed (paraphrased examples may still constitute contamination).

## Validity Decay

Published benchmarks decay in validity as models improve and training data expands. A benchmark released in 2020 is more likely to be contaminated for a 2024 model than for a 2020 model. Performance on old benchmarks is systematically over-estimated for new models.

## Mitigations

- **Held-out test sets:** Never released publicly until after a model has been trained and evaluated
- **Time-locked evaluation:** Test only on content published after the training cutoff
- **Private evaluation infrastructure:** Third-party evaluation without model access to test examples
- **Continuous benchmark rotation:** Replace saturated or contaminated benchmarks with new ones

None of these fully solve the problem; contamination remains a structural challenge for LLM evaluation.

## Connections

- [[Evaluation Metrics]] — contamination is why no single benchmark number is trustworthy; multi-dimensional evaluation is more robust
- [[Transfer Learning]] — contamination is more severe for models with internet-scraped training data; domain-fine-tuned models may be less contaminated on domain benchmarks
