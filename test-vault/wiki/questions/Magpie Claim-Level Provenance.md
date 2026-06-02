---
title: "Magpie Claim-Level Provenance"
question: "Should magpie implement claim-level provenance (each assertion tagged to its source) or is source-level provenance sufficient for this vault's size and use patterns?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-llm-wiki-compiler-deep-dive.md"
related:
  - "[[magpie]]"
  - "[[LLM Wiki Pattern]]"
tags:
  - vault-meta
  - magpie
---

# Magpie Claim-Level Provenance

*Should magpie implement claim-level provenance (each assertion tagged to its source) or is source-level provenance sufficient for this vault's size and use patterns?*

The LLM Wiki compiler uses claim-level provenance: every assertion in every wiki page is tagged to its specific source. This enables precise attribution queries and makes contradictions between sources visible at the claim level. This vault uses source-level provenance: each page lists the sources it drew from, but individual claims are not tagged.

Claim-level provenance becomes load-bearing when: (a) contradictions between sources need to be surfaced automatically, (b) individual claims need to be challenged or updated without rewriting entire pages, or (c) the vault is used in contexts where attribution matters. At current vault size, an agent can re-derive claim origins from context. As the vault grows and more sources cover overlapping topics, this becomes harder.

The implementation cost is significant: claim-level provenance requires either structured markup in wiki pages or a separate provenance store. The question is whether the benefit justifies the cost given current and projected use patterns.

## See Also

- [[magpie]] — the project this decision belongs to
- [[LLM Wiki Pattern]] — the compiler architecture and claim-level provenance design
