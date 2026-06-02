---
title: "Training Poisoning Degradation at Scale"
question: "Does training-time AI poisoning (Nightshade, Harmony Cloak) actually degrade model quality meaningfully at realistic adoption rates?"
type: question
status: open
domain: security
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-benn-jordan-poison-pilling-music.md"
related:
  - "[[AI Scraping Resistance]]"
  - "[[LLM Mental Model]]"
tags:
  - ai-resistance
  - data-poisoning
---

# Training Poisoning Degradation at Scale

*Does training-time AI poisoning (Nightshade, Harmony Cloak) actually degrade model quality meaningfully at realistic adoption rates?*

Lab demonstrations (Harmony Cloak, Nightshade) show measurable quality degradation at 100% poison rate on controlled test sets. The real-world question is the degradation curve at realistic adoption rates: what percentage of a training corpus would need to be poisoned before a model trained on it shows measurable quality reduction?

Benn Jordan's "Pareto plateau" argument is the strongest claim: poisoning the top 20% of artists (by quality) creates a quality ceiling no compute can overcome, because model training requires high-quality examples to learn from. This is theoretically elegant but unverified at scale. AI companies build diverse corpora across many domains; no single artistic domain represents a training bottleneck they can't route around.

For image poisoning (Nightshade), adoption rates among visual artists are non-trivial but far below 20% of training-relevant content. The degradation signal, if present, has not been publicly confirmed in deployed models.

## See Also

- [[AI Scraping Resistance]] — full coverage; Pareto plateau argument; arms race structure
