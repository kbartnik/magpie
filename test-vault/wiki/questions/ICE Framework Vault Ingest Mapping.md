---
title: "ICE Framework Vault Ingest Mapping"
question: "Does the IDSD ICE framework (Intent, Context, Expectations) map onto the vault's ingest workflow — and would formalizing this mapping improve ingest consistency?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-29-idsd-method.md"
related:
  - "[[IDSD]]"
  - "[[Harness Engineering]]"
  - "[[Vibe-Coding Anti-Pattern]]"
tags:
  - methodology
  - vault-meta
---

# ICE Framework Vault Ingest Mapping

*Does the IDSD ICE framework (Intent, Context, Expectations) map onto the vault's ingest workflow — and would formalizing this mapping improve ingest consistency?*

The vault's ingest workflow already has structure: present a plan, get approval, execute. IDSD's ICE framework is: explicit intent statement (why are you doing this?), context (what existing knowledge is relevant?), expectations (what should the output look like?). The informal ingest plan covers "what" but not always "why" or "what success looks like."

An explicit ICE framing at ingest time might reduce scope creep: "intent: understand the security implications of this protocol; context: we already have Agentic Identity and Zero Trust; expectations: a concept page that links to both and surfaces the open questions" is more constrained than "ingest this article." The cost is friction on straightforward single-topic ingests.

The structural question: if the vault's ingest workflow is itself an AI-assisted task, IDSD's critique of vague task specification applies directly. The "angle question" in the ingest skill is already a weak form of intent clarification. Whether a more explicit ICE structure would improve output quality or just add ceremony is worth testing.

## See Also

- [[IDSD]] — the ICE framework; critique of spec-driven development
- [[Vibe-Coding Anti-Pattern]] — the scope creep failure mode the ICE structure prevents
