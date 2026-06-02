---
title: "Typed Wikilinks Tractability"
question: "Is there a tractable lightweight approach to typed wikilinks — expressing edge semantics like 'causes', 'implements', 'contradicts' — without requiring new tooling?"
type: question
status: open
domain: vault
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-llm-wiki-compiler-deep-dive.md"
related:
  - "[[LLM Wiki Pattern]]"
  - "[[magpie]]"
tags:
  - vault-meta
  - knowledge-graph
---

# Typed Wikilinks Tractability

*Is there a tractable lightweight approach to typed wikilinks — expressing edge semantics like 'causes', 'implements', 'contradicts' — without requiring new tooling?*

Typed edges would make the vault a proper knowledge graph: instead of a plain wikilink, a typed edge like `causes::Concept` or `implements::Pattern` carries semantic meaning that enables structured queries. The obstacle is tooling: Obsidian doesn't natively support typed wikilinks; the lint tool doesn't validate types; the agent would need to reason about edge types when creating pages.

A lightweight approach that doesn't require tooling changes: use a controlled vocabulary in the `related` frontmatter field, adding a `relation` key alongside each link reference. The YAML structure captures the type; Obsidian renders it as-is; a future query tool could parse the relation. This doesn't give typed links in the body text, but it captures the most important relationship semantics in frontmatter.

Whether this is worth doing at current vault scale is unclear. The benefit grows as the vault gets larger and as cross-domain connections become harder to discover through reading alone.

## See Also

- [[LLM Wiki Pattern]] — the relationship failure mode typed wikilinks would address
