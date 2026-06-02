---
title: "Conceptual Integrity Divergence Detection"
question: "How can teams detect slow conceptual integrity divergence in async distributed teams before it becomes expensive to reverse?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/books/2026-06-01-learning-systems-thinking.md"
related:
  - "[[Systems Thinking]]"
  - "[[Software Architecture Fundamentals]]"
  - "[[Semantic Diffusion]]"
tags:
  - software-architecture
  - systems-thinking
  - organizational
---

# Conceptual Integrity Divergence Detection

*How can teams detect slow conceptual integrity divergence in async distributed teams before it becomes expensive to reverse?*

Montalion's "conceptual integrity" — the property of a system whose parts cohere around a shared mental model — is hard to measure. In async distributed teams, divergence happens without anyone noticing: each contributor makes locally reasonable decisions that are globally incoherent. The divergence becomes visible only when two components need to work together and their assumptions clash, at which point remediation is expensive.

Proposed detection mechanisms all have significant limitations: architecture sessions (expensive, rare), design documents (rarely updated to reflect actual decisions), API contracts (capture interface but not mental model), code review with conceptual integrity criteria (subjective, dependent on reviewer quality). None reliably catches slow drift before it becomes a rework event.

The systems thinking framing: conceptual integrity divergence is a reinforcing feedback loop — each incoherent decision makes the next incoherent decision more likely, because new contributors learn the system from existing code that is already partly incoherent. Breaking the loop requires a balancing feedback loop, which requires someone or something to detect and flag divergence. That detector doesn't exist yet.

## See Also

- [[Systems Thinking]] — conceptual integrity; reinforcing and balancing feedback loops; leverage points
- [[Semantic Diffusion]] — related: technical vocabulary diverges under the same organizational pressure; the vocabulary drift is a leading indicator of conceptual integrity drift
