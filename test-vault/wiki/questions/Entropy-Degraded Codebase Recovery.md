---
title: "Entropy-Degraded Codebase Recovery"
question: "Is there a systematic recovery path from a codebase that has been entropy-degraded by AI-assisted development without deep review — and what does it look like?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-software-fundamentals-ai-age-pocock.md"
related:
  - "[[Ubiquitous Language Ownership]]"
  - "[[Interface Design Responsibility at Scale]]"
  - "[[Vibe-Coding Anti-Pattern]]"
  - "[[Systems Thinking]]"
tags:
  - software-architecture
  - ai-discipline
---

# Entropy-Degraded Codebase Recovery

*Is there a systematic recovery path from a codebase that has been entropy-degraded by AI-assisted development without deep review — and what does it look like?*

Specs-to-code without reading the code compounds entropy: each iteration adds code that no one fully understands, and the understanding debt compounds as subsequent iterations build on the previous layer. At some point, the codebase becomes effectively unreadable — no engineer has a reliable mental model of the whole, and AI tools perform worse because the codebase is shallow-module and hard to navigate.

Pocock names this as a risk but doesn't prescribe a recovery path. Options: (1) incremental refactoring — requires understanding the code you're refactoring, which is the thing that's broken; (2) full rewrite — loses history, introduces new bugs, and requires the domain knowledge that the entropy-degraded codebase obscured; (3) structured decomposition — identify the stable interfaces, extract them, refactor behind each interface boundary iteratively. Option 3 is most principled but requires the interfaces to already be coherent enough to serve as boundaries.

The Pragmatic Programmer's framing: software entropy is a ratchet. Preventing it requires continuous intervention; reversing it requires concentrated effort that is hard to justify under normal delivery conditions.

## See Also

- [[Ubiquitous Language Ownership]], [[Interface Design Responsibility at Scale]] — related questions in the same cluster
- [[Vibe-Coding Anti-Pattern]] — the 5-stage failure cascade that produces entropy-degraded codebases
