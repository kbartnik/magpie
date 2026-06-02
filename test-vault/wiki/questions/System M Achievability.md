---
title: "System M Achievability"
question: "Is Dupoux/LeCun/Malik's System M (meta-control for post-deployment learning) tractable at frontier scale, or does it face a deeper architectural obstacle?"
type: question
status: open
domain: ai-theory
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/papers/2026-06-01-dupoux-lecun-malik-autonomous-learning.md"
related:
  - "[[Autonomous Learning Architecture]]"
  - "[[LLM Mental Model]]"
  - "[[Neuromorphic and Bio-Inspired AI]]"
tags:
  - ai-theory
  - autonomous-learning
---

# System M Achievability

*Is Dupoux/LeCun/Malik's System M (meta-control for post-deployment learning) tractable at frontier scale, or does it face a deeper architectural obstacle?*

The Dupoux/LeCun/Malik (2026) framework proposes System M — a meta-control layer that directs what System A (observation) should observe and when System B (action) should act. System M is the piece that enables post-deployment learning without catastrophic forgetting: it modulates which experiences get consolidated into long-term memory and which actions get executed.

No System M implementation exists at frontier scale. The paper is a research agenda. The question is whether the obstacle is primarily engineering (we know how to build it but haven't yet) or architectural (the inductive biases of current transformer architectures actively resist the kind of online learning System M requires). The history of AI suggests that "we know how in principle, just needs engineering" often conceals decades of hard work — the same claim was made about neural networks in the 1970s.

## See Also

- [[Autonomous Learning Architecture]] — full System A/B/M framework; the three roadblocks
- [[LLM Mental Model]] — the frozen-after-training property System M is designed to overcome
