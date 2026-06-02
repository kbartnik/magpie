---
title: "Bio-Inspired vs Gradient Descent"
question: "What are the genuine use cases where bio-inspired algorithms (evolutionary, swarm, simulated annealing) outperform gradient-based optimization — and which of the published claims hold up?"
type: question
status: open
domain: ai-theory
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/papers/2026-06-01-bongard-2009-bio-inspired-notes.md"
  - "archive/papers/2026-06-01-somvanshi-2025-bio-inspired-critique-notes.md"
related:
  - "[[Bio-Inspired Computing]]"
  - "[[Neuromorphic and Bio-Inspired AI]]"
tags:
  - ai-theory
  - bio-inspired
---

# Bio-Inspired vs Gradient Descent

*What are the genuine use cases where bio-inspired algorithms (evolutionary, swarm, simulated annealing) outperform gradient-based optimization — and which of the published claims hold up?*

The genuine advantages of bio-inspired methods over gradient descent are narrow but real: non-differentiable objectives (gradient descent requires a differentiable loss function; evolutionary methods don't), highly multimodal landscapes (population-based methods maintain diversity and escape local minima more effectively), and certain combinatorial problems (scheduling, routing, layout) that don't map naturally to continuous gradient-based formulations.

Somvanshi (2025) argues the bio-inspired field has a novelty-inflation problem: many recent papers demonstrate performance on toy benchmarks where gradient descent was already competitive, not on the cases where bio-inspired methods have genuine advantages. The published performance claims often don't survive contact with well-tuned baselines.

The practical question: when evaluating a bio-inspired approach for a specific problem, what criteria distinguish "this problem genuinely fits the bio-inspired advantage profile" from "this is a familiar tool being overapplied"?

## See Also

- [[Bio-Inspired Computing]] — evolutionary/swarm/physics-inspired algorithm families; Somvanshi critique of novelty inflation
- [[Neuromorphic and Bio-Inspired AI]] — hardware approaches; the energy argument against transformers
