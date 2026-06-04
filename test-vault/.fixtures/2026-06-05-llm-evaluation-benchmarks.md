---
title: "Benchmarking Large Language Models: What We Measure and What It Means"
author:
  - "Percy Liang"
  - "Rishi Bommasani"
url: "https://crfm.stanford.edu/2022/11/17/helm.html"
description: "The HELM (Holistic Evaluation of Language Models) framework from Stanford CRFM: what current benchmarks measure, their failure modes, and what a more comprehensive evaluation would look like."
tags:
  - "clippings"
  - "ml-fundamentals"
  - "evaluation"
  - "benchmarking"
published: 2022-11-17T00:00:00Z
created: 2026-06-05T11:00:00-04:00
---

# Benchmarking Large Language Models: What We Measure and What It Means

By Percy Liang, Rishi Bommasani (Stanford CRFM)
Published: 2022-11-17

>[!summary]
>Current LLM benchmarks measure a narrow slice of capability and suffer from contamination, metric gaming, and construct validity problems. HELM proposes a multi-dimensional evaluation framework: accuracy, calibration, robustness, fairness, toxicity, efficiency, and disentangled scenario coverage. No single number characterizes a model's quality.

## What Benchmarks Actually Measure

Most popular benchmarks measure one thing well and generalize poorly:

**MMLU (Massive Multitask Language Understanding):** 57 subjects from professional and academic domains. Tests factual recall and reasoning at the knowledge level of graduate students. Limitation: measures knowledge breadth, not reasoning depth; multiple-choice format masks confabulation.

**HumanEval:** Code generation. Pass@k (does at least one of k samples pass the test suite?) as the metric. Limitation: hand-written tests may not cover edge cases; tests the ability to produce a working solution, not readable or maintainable code.

**HellaSwag, WinoGrande:** Common-sense reasoning. Limitation: adversarially-filtered options inflate apparent difficulty; models have saturated both benchmarks above human performance without demonstrable common-sense reasoning.

## Benchmark Contamination

Models trained on internet-scale data have seen most benchmark test sets during pretraining. **Contamination** is the presence of benchmark examples in the training corpus. Contaminated benchmarks overestimate capability — the model may have memorized answers rather than learned to reason.

Contamination is hard to detect: it requires knowing the exact training data, which large model providers don't always disclose. Published benchmarks decay in validity as models improve and training data expands.

## Evaluation Dimensions Beyond Accuracy

**Calibration:** Does the model's confidence correlate with its accuracy? A calibrated model that says "I'm 80% confident" should be right ~80% of the time. LLMs are systematically overconfident.

**Robustness:** Does performance degrade under paraphrase, typo, or distribution shift? Most models are brittle to surface-level variation.

**Fairness:** Does performance differ across demographic groups? Models that perform well on average may perform poorly for underrepresented groups.

**Efficiency:** What computational cost (tokens, latency, API calls) is required for a given accuracy level? Important for production deployment but rarely reported.

## The Leaderboard Problem

Publishing benchmark scores creates incentives to optimize for the benchmark. Models fine-tuned specifically on benchmark-adjacent data inflate scores without improving real-world performance. Leaderboard position is not a reliable proxy for usefulness.

The evaluation community's response: held-out test sets, private evaluation infrastructure, and continuous benchmark rotation. None of these fully solves the problem.

## Deep Read

**Key Insight:** A benchmark is a measurement instrument, and like all instruments it has a range, a resolution, and systematic biases. MMLU measures knowledge breadth; HumanEval measures code generation; neither measures reasoning, calibration, robustness, or real-world usefulness. Choosing a benchmark is choosing what to optimize, and the field's habit of single-number rankings obscures this.

**What Surprised Me:** The saturation problem is worse than I expected. HellaSwag and WinoGrande — designed to be difficult common-sense tasks — are now solved above human-level performance by standard LLMs, but there's no evidence that the models actually have common-sense reasoning. They may have found statistical shortcuts in the adversarial filter used to construct the benchmarks. "Above human performance" on a saturated benchmark tells you almost nothing about capability.

**Open Questions:**
- If contamination makes benchmark performance unreliable, is there a principled way to construct benchmarks that are contamination-resistant? Time-locked evaluation (test only on events after the training cutoff) is one approach — what are its limits?
- Calibration can be measured, but improving calibration without training explicitly for it is hard. Is there a training objective that produces better calibration as a side effect, or does calibration require dedicated training?
- Human evaluation is the gold standard for quality, but it's expensive and doesn't scale. Can AI judges (using another LLM to evaluate outputs) approximate human evaluation reliably enough for production use, and what are the systematic failure modes?

**Wikilink Candidates:**
- [[Evaluation Metrics]] — accuracy, calibration, robustness, fairness, efficiency; multi-dimensional evaluation; no single number; not yet a wiki page
- [[Benchmark Contamination]] — training data includes test sets; validity decay over time; leaderboard gaming; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — all benchmarks are measuring transformer-based models; the architecture's capabilities and limitations (long context, factual recall, reasoning) are what benchmarks probe
- [[RLHF]] — RLHF-tuned models often score higher on human preference evaluations but may score lower on some capability benchmarks (the alignment tax); evaluation methodology affects which models appear best
- [[Transfer Learning]] — benchmark performance of fine-tuned models vs base models measures transfer quality; contamination is more severe for models fine-tuned on internet-scraped data
