---
title: "Onto AI Scalability and Auditability"
question: "Do Onto AI's scalability and auditability claims for graph-structured memory survive independent validation — and does the 800K neuron demonstration actually predict behavior at brain scale?"
type: question
status: open
domain: ai-theory
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-llm-wrong-problem.md"
related:
  - "[[Neuromorphic and Bio-Inspired AI]]"
  - "[[Autonomous Learning Architecture]]"
  - "[[LLM Mental Model]]"
tags:
  - ai-theory
  - neuromorphic
---

# Onto AI Scalability and Auditability

*Do Onto AI's scalability and auditability claims for graph-structured memory survive independent validation — and does the 800K neuron demonstration actually predict behavior at brain scale?*

Onto AI's four-layer software brain uses graph-structured knowledge representation rather than static weights. Claims: better auditability (you can inspect the knowledge graph), better scalability (adding knowledge doesn't require retraining). Both are plausible in principle; neither has been independently validated at scale.

The scalability concern: current demonstrations run on 800,000 neuron-equivalent units — approximately 0.008% of the human brain's 86 billion neurons. Graph-structured memory's scaling laws may differ fundamentally from transformer scaling laws. Whether the architecture retains its advantages at 10,000× scale is unknown; every architectural approach looks better in limited demonstrations than it does under production scale.

The auditability concern: "inspectable" doesn't mean "correct." A knowledge graph that stores factually incorrect relationships is auditable and wrong. The harder question — how do you validate that the knowledge graph is accurate? — isn't addressed. The same challenge faces rule-based systems generally.

## See Also

- [[Neuromorphic and Bio-Inspired AI]] — full coverage of Onto AI and the post-transformer landscape
- [[LLM Mental Model]] — frozen-after-training transformers as the baseline these approaches aim to improve on
