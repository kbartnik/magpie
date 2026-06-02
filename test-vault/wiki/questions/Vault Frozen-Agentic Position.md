---
title: "Vault Frozen-Agentic Position"
question: "Is the vault's current hybrid position on the frozen/agentic spectrum — human-driven ingest, agent-maintained wiki — still correctly calibrated as vault size and agent capability grow?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-karpathy-llm-wiki-built-twice.md"
related:
  - "[[LLM Wiki Pattern]]"
  - "[[magpie]]"
tags:
  - vault-meta
  - architecture
---

# Vault Frozen-Agentic Position

*Is the vault's current hybrid position on the frozen/agentic spectrum — human-driven ingest, agent-maintained wiki — still correctly calibrated as vault size and agent capability grow?*

The vault was designed as a deliberate hybrid: humans drive ingest (choosing what sources to add), the agent maintains the wiki (synthesizing, cross-linking, updating), and the archive is immutable. This positions it between frozen (ingest-only, static pages) and fully agentic (LLM continuously rewrites based on new observations).

As the vault grows and the agent becomes more capable, two recalibrations become possible: (1) shift toward more autonomous maintenance — the agent proposes wiki updates based on pattern detection without explicit ingest events; (2) shift toward more frozen — user-written pages, agent reads only. Both would change the trust/control balance.

The calibration question hasn't been revisited since initial design. The right answer depends on how much the agent is trusted for wiki judgment and how the vault is primarily used (queries vs. ingest/lookup). These have both shifted since design.

## See Also

- [[LLM Wiki Pattern]] — the frozen/agentic spectrum; this vault's position as deliberate hybrid
- [[magpie]] — the project that may shift this position through architectural evolution
