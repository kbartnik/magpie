---
title: "Vault LLM Wiki Failure Modes"
question: "Does this vault exhibit the LLM Wiki failure modes — identity, level, and relationship failures — and if so, where?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-llm-wiki-compiler-deep-dive.md"
related:
  - "[[LLM Wiki Pattern]]"
tags:
  - vault-meta
---

# Vault LLM Wiki Failure Modes

*Does this vault exhibit the LLM Wiki failure modes — identity, level, and relationship failures — and if so, where?*

The LLM Wiki compiler identifies three structural failure modes: **identity failure** (two pages that should be the same concept treated as separate, causing duplication and inconsistency), **level failure** (concepts described at the wrong level of abstraction — too general to be actionable, too specific to be cross-referenced), and **relationship failure** (correct pages but missing or wrong links between them).

This vault has never been systematically audited against these criteria. Some likely candidates for identity failures: Vibe-Coding Anti-Pattern and IDSD cover overlapping territory. Agentic Workflow Patterns and Harness Engineering may partially overlap. Level failures: some Go concept pages may be too implementation-specific to connect to conceptual pages. Relationship failures: probably the most common, given the vault's growth rate.

Whether a systematic audit is worth running, and whether vault-tools could be extended to automate it, hasn't been evaluated.

## See Also

- [[LLM Wiki Pattern]] — full description of the three failure modes
