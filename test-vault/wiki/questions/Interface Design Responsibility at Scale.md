---
title: "Interface Design Responsibility at Scale"
question: "When 50 engineers are each designing module interfaces and delegating implementations to AI, what process ensures the interfaces compose coherently at the system level?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-software-fundamentals-ai-age-pocock.md"
related:
  - "[[Ubiquitous Language Ownership]]"
  - "[[Entropy-Degraded Codebase Recovery]]"
  - "[[Systems Thinking]]"
  - "[[IDSD]]"
tags:
  - software-architecture
  - organizational
  - ai-discipline
---

# Interface Design Responsibility at Scale

*When 50 engineers are each designing module interfaces and delegating implementations to AI, what process ensures the interfaces compose coherently at the system level?*

Pocock's argument: design the interface, delegate the implementation to AI. This works for individual developers. At scale — when 50 engineers are each designing interfaces for their modules — there is no guarantee the interfaces compose coherently. Each engineer's locally reasonable interface decision may be globally incoherent.

Interface design requires domain judgment that can't be fully delegated to AI, but it also can't remain purely individual — a 50-engineer team needs global coherence. The middle path isn't specified: team-level interface reviews? Architecture council with veto power? Automated consistency checks (type compatibility, naming convention enforcement)? Contract testing as system-level interface validation? Each of these has organizational overhead and coverage gaps.

This connects to Montalion's conceptual integrity problem: interface incoherence is a symptom of conceptual integrity divergence, but interface-level fixes (enforcing consistency) don't address the underlying mental model divergence that causes it.

## See Also

- [[Ubiquitous Language Ownership]], [[Entropy-Degraded Codebase Recovery]] — related questions in the same cluster
- [[Systems Thinking]] — conceptual integrity as a system property; why individual-level discipline doesn't produce system-level coherence
