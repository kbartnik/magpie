---
title: "Fitness Functions Political Viability"
question: "What organizational structures make fitness functions load-bearing rather than ignored — and why do they consistently fail in practice despite clear engineering rationale?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/books/2026-06-01-fundamentals-of-software-architecture.md"
related:
  - "[[Software Architecture Fundamentals]]"
  - "[[Systems Thinking]]"
tags:
  - software-architecture
  - organizational
---

# Fitness Functions Political Viability

*What organizational structures make fitness functions load-bearing rather than ignored — and why do they consistently fail in practice despite clear engineering rationale?*

Fitness functions are automated architectural checks: CI gates that fail when coupling exceeds a threshold, tests that verify response time SLAs, metrics that enforce module size limits. The engineering argument for them is clear. The political problem: violations require engineering time that competes with feature delivery. Product teams under pressure resist, delay, or deprioritize violations. A fitness function that consistently fails without triggering remediation normalizes ignoring violations and spreads to other quality checks.

The organizational structures that might make them load-bearing — architecture review boards, architectural debt budgets, fitness function champions — have inconsistent real-world records. The failure pattern seems to be: the function is added with management buy-in; violations begin appearing; the first few are addressed; later violations accumulate because the team is busy; eventually the function is disabled or ignored. The accountability mechanism needed to interrupt this pattern isn't specified in the architecture literature.

## See Also

- [[Software Architecture Fundamentals]] — Richards & Ford; fitness functions as the primary architectural governance tool
- [[Systems Thinking]] — leverage point theory: changing the rules (fitness functions) is high-leverage but only if enforcement follows
