---
title: "Diffused Responsibility at Scale"
question: "How does an organization maintain collective intent ownership over a codebase whose components were built with varying degrees of AI involvement and human oversight?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-0.md"
  - "archive/clippings/2026-05-29-idsd-method.md"
related:
  - "[[Vibe-Coding Anti-Pattern]]"
  - "[[IDSD]]"
tags:
  - methodology
  - organizational
  - ai-discipline
---

# Diffused Responsibility at Scale

*How does an organization maintain collective intent ownership over a codebase whose components were built with varying degrees of AI involvement and human oversight?*

The Vibe-Coding Anti-Pattern's stage 5 failure is diffused responsibility at the individual level: when AI generates code that no engineer fully reviewed, nobody owns it. At organizational scale, this compounds: different engineers use AI with different levels of ICE discipline, producing modules with radically different ownership profiles. A codebase where Module A has been carefully reviewed and Module B was accepted with minimal review is not uniformly safe to modify.

IDSD addresses individual-level intent ownership clearly but hasn't been developed for organizational governance. What structures make collective ownership viable? Candidates: ADR culture (every significant decision has an author and rationale); ownership mapping (explicit accountability for each module); AI involvement disclosure in code review (like "this was AI-generated, reviewed to level X"); mandatory review depth based on AI-generation percentage. None of these is specified or tested.

The problem compounds over time: a codebase with five years of mixed-ownership history becomes harder to reason about as the original authors' context fades and the AI-generated components outnumber the reviewed ones.

## See Also

- [[Vibe-Coding Anti-Pattern]] — diffused responsibility as stage 5 of the failure cascade
- [[IDSD]] — intent ownership at the individual level; the organizational extension is missing
