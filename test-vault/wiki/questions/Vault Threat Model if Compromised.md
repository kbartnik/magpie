---
title: "Vault Threat Model if Compromised"
question: "What is the attack surface if the vault itself is compromised — and do the current plan-before-write mitigations provide sufficient protection?"
type: question
status: open
domain: agentic
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-why-ai-agents-break-zero-trust-last-mile.md"
related:
  - "[[Agentic Identity and Zero Trust]]"
  - "[[Claude Code Hooks]]"
  - "[[Security Culture]]"
tags:
  - security
  - vault-meta
---

# Vault Threat Model if Compromised

*What is the attack surface if the vault itself is compromised — and do the current plan-before-write mitigations provide sufficient protection?*

The ingest workflow trusts the agent with write access to wiki pages, which the agent references in future sessions. An attacker who injects content into archive files (via a poisoned web clip, malicious PDF, or compromised source document) would be placing content in the agent's context window on future reads — a form of persistent prompt injection.

Current mitigations: vault is local (not network-accessible), archive files are reviewed before ingest (plan-before-write protocol), and the agent requires user approval for file writes. The question is whether these hold against a determined attacker with physical or network access, and whether the plan-before-write review would actually catch a subtly adversarial archive entry that looks legitimate but contains behavioral instructions embedded in body text.

The risk is low in current use but the architectural exposure is real: any system where an agent reads from a corpus it (or its environment) has previously written is a prompt injection surface.

## See Also

- [[Agentic Identity and Zero Trust]] — vault-as-middleware; the identity/intent/delegation gap
- [[Claude Code Hooks]] — hooks as a partial defense: can validate tool calls before execution
