---
title: "Data Poisoning Legal Status"
question: "Is deliberately poisoning an AI company's training corpus or profiling system legally permissible when the company is itself engaging in unauthorized data collection?"
type: question
status: open
domain: security
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-benn-jordan-poison-pilling-music.md"
  - "archive/videos/2026-06-01-addie-lamarr-data-poisoning-surveillance.md"
related:
  - "[[AI Scraping Resistance]]"
tags:
  - ai-resistance
  - legal
  - data-poisoning
---

# Data Poisoning Legal Status

*Is deliberately poisoning an AI company's training corpus or profiling system legally permissible when the company is itself engaging in unauthorized data collection?*

Both training-time poisoning (injecting adversarial noise into published works) and inference-time behavioral poisoning (generating false engagement signals) are offensive actions. The defensive framing: they target systems engaging in unlawful behavior — unauthorized scraping, unauthorized data collection, potential GDPR violations.

The CFAA (Computer Fraud and Abuse Act) analysis is unresolved. Accessing a public website to post content (even adversarially structured content) is not obviously unauthorized access. Deliberately generating false behavioral signals to degrade an analytics system is not obviously computer fraud. But the legal question hasn't been tested.

GDPR's right-to-object provisions may provide some cover in the EU for actions that degrade profiling systems. Some copyright frameworks may support data poisoning as a form of technical protection measure. No court has ruled on offensive data poisoning as a response to unlawful scraping.

## See Also

- [[AI Scraping Resistance]] — the technical landscape; collective action problem; legality noted as open
