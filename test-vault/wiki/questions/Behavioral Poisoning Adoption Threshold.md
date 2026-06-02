---
title: "Behavioral Poisoning Adoption Threshold"
question: "What adoption rate of deliberate false behavioral signals is required to meaningfully degrade inference-time surveillance profiler accuracy?"
type: question
status: open
domain: security
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-addie-lamarr-data-poisoning-surveillance.md"
related:
  - "[[AI Scraping Resistance]]"
  - "[[Digital Surveillance Resistance]]"
tags:
  - ai-resistance
  - surveillance
---

# Behavioral Poisoning Adoption Threshold

*What adoption rate of deliberate false behavioral signals is required to meaningfully degrade inference-time surveillance profiler accuracy?*

Addie LaMarr's behavioral poisoning strategy: generate deliberate false engagement signals (interact with content that doesn't reflect your actual preferences) to degrade the accuracy of behavioral profiling systems. The structural asymmetry LaMarr identifies: larger surveillance datasets are more vulnerable to noise because the signal-to-noise ratio shrinks as the corpus grows.

The Taleb minority rule suggests a committed minority consistently generating false signals could swamp authentic signals. But "committed" is load-bearing — the strategy requires deliberate, consistent unnatural behavior with personal inconvenience costs: degraded recommendations, potential fraud detection triggers, additional cognitive overhead.

The threshold for meaningful profiling degradation hasn't been empirically established. The feedback loop structure means the effect may be nonlinear: below a threshold, noise is filtered; above a threshold, the classifier degrades rapidly. Where that threshold is depends on the specific profiling system.

## See Also

- [[AI Scraping Resistance]] — behavioral poisoning alongside training-time techniques; collective action problem
- [[Digital Surveillance Resistance]] — the broader surveillance threat model
