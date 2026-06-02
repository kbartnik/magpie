---
title: "CLAUDE.md Degradation Mechanism"
question: "What is the mechanism by which CLAUDE.md files above ~200 lines produce degraded instruction-following, and how should CLAUDE.md files be structured to avoid it?"
type: question
status: open
domain: claude-code
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-29-karpathy-claude-md-8-rules.md"
related:
  - "[[Claude Code Memory Architecture]]"
  - "[[CLAUDE.md Configuration Patterns]]"
  - "[[Harness Engineering]]"
tags:
  - claude-code
  - cca-f
---

# CLAUDE.md Degradation Mechanism

*What is the mechanism by which CLAUDE.md files above ~200 lines produce degraded instruction-following, and how should CLAUDE.md files be structured to avoid it?*

Multiple practitioners report that CLAUDE.md files above approximately 200 lines produce degraded instruction-following — the agent begins inconsistently applying rules from the lower portions of the file. The mechanism is undocumented by Anthropic.

Candidate hypotheses: (1) attention decay — the instruction-following training objective may not generalize well to very long instruction files, causing later instructions to receive systematically less attention weight; (2) context budget — CLAUDE.md loads early and displaces other context; at 200+ lines this may crowd out working memory for the current task; (3) training distribution — instruction files in training data may have been shorter on average, so long files are out-of-distribution.

The practical implication is architectural: rather than one long CLAUDE.md, the correct structure may be a short CLAUDE.md that imports or references domain-specific instruction files — similar to the Progressive Disclosure Architecture pattern for skills. The current vault CLAUDE.md is long; whether it's hitting the degradation threshold is unknown.

## See Also

- [[CLAUDE.md Configuration Patterns]] — the 8-rule architecture and structural principles
- [[Claude Code Memory Architecture]] — how CLAUDE.md fits into the 4-layer memory hierarchy
