---
title: "Dumb Zone Threshold Universality"
question: "Is the ~40% context fill 'dumb zone' degradation threshold specific to Claude, or does it generalize across models?"
type: question
status: open
domain: ai-ml
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-no-vibes-allowed-context-engineering.md"
related:
  - "[[Context Rot]]"
  - "[[Harness Engineering]]"
  - "[[Transformer Architecture]]"
tags:
  - context-management
  - harness-engineering
---

# Dumb Zone Threshold Universality

*Is the ~40% context fill 'dumb zone' degradation threshold specific to Claude, or does it generalize across models?*

Dex Horthy (No Vibes Allowed, AI Engineer 2025) claims model performance noticeably degrades around 40% context window fill — the "dumb zone." This is cited as a practical engineering constraint: structure agentic tasks to stay below this threshold. But the threshold was derived from Horthy's experience specifically with Claude.

Whether the threshold transfers to other models is unknown. It may reflect specific Claude architectural decisions around attention patterns, or RLHF training choices that affect performance at long contexts. Different models (GPT-4, Gemini) may have different degradation profiles. It may also be task-type-specific: factual retrieval from context may degrade at different fill levels than reasoning or code generation.

If it's Claude-specific and Claude's architecture changes, the threshold changes too. If it generalizes, it suggests a structural property of transformer attention that should inform harness design universally.

## See Also

- [[Context Rot]] — the broader failure mode; intentional compaction as the mitigation
- [[Harness Engineering]] — practical implications for harness design around context budgets
