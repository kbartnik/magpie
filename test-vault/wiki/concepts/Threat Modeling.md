---
title: "Threat Modeling"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/clippings/2026-06-01-secure-ai-agents-architecture.md"
related:
  - "[[Digital Surveillance Resistance]]"
  - "[[Operational Security]]"
  - "[[Security Culture]]"
tags:
  - security
  - privacy
  - risk
---

# Threat Modeling

A structured process for identifying security threats, their likelihood, and appropriate countermeasures — calibrated to a specific context and adversary. Threat modeling prevents both under-protection (missing real risks) and over-protection (wasting resources on negligible ones).

The core questions:
1. **What are you protecting?** (assets: communications, identity, location, contacts, documents)
2. **From whom?** (adversary: local police, federal agencies, corporate surveillance, doxxers, stalkers)
3. **How likely is each threat?** (probability × impact)
4. **What can you afford to lose?** (acceptable risk)
5. **What measures are proportionate?**

## Why It Matters for [[Digital Surveillance Resistance]]

The EFF and WIRED's surveillance guides emphasize threat modeling as the prerequisite for good security choices. Tor and Signal are not appropriate for every user in every context — they are the right tool when the adversary has the capability and motivation to monitor standard channels. Most people most of the time have different adversaries and different risk profiles.

Threat modeling prevents *security theater* (visible measures that don't address real threats) and *security friction* (measures so burdensome they're abandoned).

## AI Agent Threat Modeling

Agents require an extended threat model beyond conventional software. The core complication: **autonomy scales consequences**. A compromised tool causes damage proportional to the attacker's bandwidth; a compromised autonomous agent causes damage at machine speed, continuously, until detected.

**Assets to protect:**
- Data the agent can read/write/exfiltrate via tool calls
- Backend systems the agent can affect via tool calls
- The agent's own credentials and permissions
- The agent's behavior (against prompt injection / hijacking)

**Agent-specific adversary capabilities:**
- Prompt injection: injecting commands into the agent's input (via user messages, retrieved documents, tool outputs) to redirect behavior
- Rogue agent impersonation: attacker claims to be a legitimate agent to access backend systems
- Privilege escalation: exploiting agent autonomy to accumulate permissions beyond what was granted

**Agent-specific threat taxonomy** (IBM/Anthropic joint guide, 2026):

| Threat | Why it's agent-specific |
|--------|------------------------|
| Excessive agency | Agent scope is harder to define than function scope |
| Privilege escalation | Autonomous agents can take actions that expand their own access |
| Prompt injection | LLMs are uniquely vulnerable; input can come from many untrusted sources |
| Attack amplifier | Compromise → machine-speed execution with no human friction |
| Compliance drift | Adaptive behavior can drift outside policy boundaries silently |

**Countermeasure stack:**
1. Principle of least privilege — per-task scoped access, JIT grants
2. Non-human identity management — unique credentials, audit trail per agent
3. AI firewall/proxy — filter injection on ingress, DLP on egress
4. Human-in-the-loop gates for high-risk operations
5. Real-time monitoring + anomaly detection + configuration drift detection

See [[Agentic Identity and Zero Trust]] for the full treatment.

## In Organizing Contexts

[[Security Culture]] implicitly applies threat modeling: the primary adversary is law enforcement and informants; the assets are identities, plans, and relationships; the countermeasures (compartmentalization, encrypted comms, meeting protocols) are calibrated to that specific threat landscape.

## See Also

- [[Digital Surveillance Resistance]] — practical countermeasures; threat modeling is the prerequisite for choosing among them
- [[Operational Security]] — the practice of implementing the measures threat modeling identifies
- [[Security Culture]] — community-level threat modeling applied to activist/organizing contexts
- [[Agentic Identity and Zero Trust]] — concrete threat model for AI agent systems: rogue agent impersonation, tool chaining, and credential theft, with defined countermeasures
- [[Vault Threat Model if Compromised]] — open question: what the blast radius of an LLM wiki vault compromise looks like in practice
