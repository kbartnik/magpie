---
tags: [concept, ml-fundamentals, agents]
cluster: ml-fundamentals
aliases: ["LLM evaluation", "benchmarks", "MMLU", "HumanEval", "model evaluation", "benchmark contamination"]
related: ["Benchmark Contamination", "RLHF", "Transformer Architecture", "Transfer Learning"]
sources:
  - "[[archive/clippings/2026-06-04-llm-evaluation-benchmarks]]"
---

# Evaluation Metrics

A benchmark is a measurement instrument with a range, resolution, and systematic biases. No single number characterizes a model's quality.

## Common Benchmarks and Limits

**MMLU:** 57-subject knowledge test. Measures breadth; multiple-choice format masks confabulation. Contamination-prone.

**HumanEval:** Code generation. Pass@k metric. Tests producing working solutions, not readable/maintainable code.

**HellaSwag / WinoGrande:** Common-sense reasoning. Both saturated above human performance — models found statistical shortcuts without demonstrable common-sense reasoning.

## Evaluation Dimensions Beyond Accuracy

| Dimension | What It Measures |
|---|---|
| **Accuracy** | Correctness on defined tasks |
| **Calibration** | Confidence ↔ accuracy correspondence |
| **Robustness** | Stability under paraphrase/typo/distribution shift |
| **Fairness** | Performance consistency across demographic groups |
| **Efficiency** | Compute cost per accuracy unit |

## The Leaderboard Problem

Publishing benchmark scores creates incentives to optimize for the benchmark. Fine-tuning on benchmark-adjacent data inflates scores without improving real-world performance. Leaderboard position ≠ usefulness.

## Connections

- [[Benchmark Contamination]] — training data including test examples inflates apparent performance; validity decays as models improve
- [[RLHF]] — RLHF-tuned models score higher on human preference evaluations but may score lower on capability benchmarks (alignment tax)
- [[Transfer Learning]] — benchmark performance of fine-tuned vs base models measures transfer quality; contamination is more severe for internet-scraped data
