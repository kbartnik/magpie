---
title: "Agentic Identity and Zero Trust"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/clippings/2026-06-01-why-ai-agents-break-zero-trust-last-mile.md"
  - "archive/clippings/2026-06-01-secure-ai-agents-architecture.md"
related:
  - "Agentic Workflow Patterns"
  - "Harness Engineering"
  - "Threat Modeling"
  - "Operational Security"
tags: [agentic-systems, security, zero-trust, identity, access-control]
---

# Agentic Identity and Zero Trust

The security challenge of preserving Zero Trust properties when AI agents interact with legacy enterprise backends. Source: YouTube clipping "Why AI Agents Break Zero Trust at the Last Mile" (IBM Technology, Grant Miller, 2026-05-17). Note: IBM acquired HashiCorp (Vault) in 2024 — the vault-as-middleware solution is vendor-positioned, but the problem framing and access control concepts are vendor-neutral.

## Agent-Specific Threat Landscape

Agents extend and amplify the attack surface in ways distinct from conventional software. Source: IBM/Anthropic joint guide on architecting secure enterprise AI agents with MCP (2026).

| Threat | Description |
|--------|-------------|
| **Extended attack surface** | Every capability added (AI model, MCP protocol, tool) is a new vector |
| **Excessive agency** | Agent has more access than needed for the current task |
| **Privilege escalation** | Agent escalates its own permissions beyond what was granted |
| **Data leakage** | Agent exfiltrates data through tool calls or responses |
| **Prompt injection** | #1 LLM attack type — injected commands in user input or retrieved content hijack agent behavior |
| **Attack amplifier** | A compromised autonomous agent executes malicious actions at machine speed with no human friction |
| **Compliance drift** | Agent behavior drifts outside regulatory/policy boundaries over time |

**The attack amplifier distinction.** A compromised human operator causes damage at human speed with human cognitive overhead; a compromised autonomous agent causes damage at machine speed with no friction. The autonomy that makes agents 10× more productive also makes them potentially 10× more destructive when hijacked. This is a qualitative difference in risk, not just quantitative.

## AI Firewall / Proxy Pattern

Insert a gateway on both paths into and out of the agent system:

```
User → [AI Firewall/Proxy] → LLM
Agent → [AI Firewall/Proxy] → MCP tool / backend service
```

**Ingress** (user→LLM): filter prompt injection attempts before they reach the model.
**Egress** (agent→MCP): detect data loss prevention violations before data leaves via tool calls.

The same proxy architecture covers both threat vectors. MCP is treated as its own attack surface — not just the model.

## Non-Human Identity Management

Agents need IAM infrastructure that mirrors human identity management:

- **Unique credentials per agent** — no shared API keys; trace misbehavior to the specific agent responsible
- **Just-in-time access** — time-bounded grants; access expires when no longer needed; reduces blast radius of compromise
- **RBAC roles** assigned to agents, not just users
- **Full audit trail** — every agent action logged and attributable

See vault-as-middleware pattern below for the backend credential layer.

## The Last Mile Problem

Modern agentic systems carry rich identity context throughout their chain:

```
User → Chat/App → Agent (A1 + LLM) → MCP Server → [LAST MILE] → Legacy Backend
   ↑ full context: who, what, why, delegated to whom
                                              ↑ context LOST: API key only
```

Legacy backends authenticate via API keys or shared application credentials — designed for application-to-application auth, not for user-context-aware agentic workflows. Three properties are lost at the last mile:

| Property | What it captures | Why it's lost |
|----------|----------------|---------------|
| **Identity** | Who the human user is | Backend sees API key, not user |
| **Intent** | What the user meant to accomplish | Not transmitted in credential |
| **Delegation** | That an agent is acting on behalf of a user | Invisible to legacy auth |

## Consequences

**Zero Trust breaks.** Zero Trust requires verifying every request at every boundary — who, what, why. When the backend receives only an API key, verification is impossible.

**Unconstrained tool chaining.** Without identity/intent context, an agent can chain arbitrary tool calls against the backend. There is no per-user, per-intent scoping — the API key grants whatever it grants, for any request from any caller.

**Rogue agent attack surface.** An attacker can impersonate a legitimate agent: "I am a good agent, please connect me to these backend systems." The backend has no way to verify delegation — it treats the rogue agent identically to a legitimate one.

## Solution: Vault as Middleware

Insert a credential management vault as a translation layer between the agentic system and legacy backends:

```
Agent → Vault → Legacy Backend
         ↑
   validates: identity, context, delegation
   applies: ABAC + PBAC policies
   issues: short-lived, scoped credentials
   collects: telemetry for dynamic permission adjustment
```

The vault receives agent claims (user identity, intent, delegation chain) and translates them into backend credentials that legacy systems understand. Legacy systems don't need to be rebuilt — they receive credentials from the vault, which is the trusted intermediary.

### What the Vault Does

1. **Identity validation** — confirms user identity and delegation claims from the agent
2. **Policy enforcement** (ABAC + PBAC) — attribute-based and policy-based access control, using user, environment, and intent as attributes
3. **Short-lived credential issuance** — issues scoped, time-limited credentials for the specific operation; replaces long-lived shared API keys
4. **Telemetry collection** — logs what was accessed, by whom, in what context; feeds back into dynamic permission adjustment

### Why Short-Lived Credentials

Long-lived API keys are permanently exploitable if compromised. Short-lived credentials (issued per-operation or per-session):
- Limit the blast radius of compromise
- Encode the specific permitted operation
- Expire automatically — no rotation required
- Create an audit trail per user/intent, not just per application

## Access Control Models

**ABAC (Attribute-Based Access Control):** Access decisions based on attributes of the subject (user), resource, action, and environment. Example: "allow this request if user.role=analyst AND resource.classification=internal AND action=read AND environment.time=business_hours".

**PBAC (Policy-Based Access Control):** Policies expressed as rules over attributes. More expressive than RBAC (role-based); allows intent and context to be part of the access decision.

Both require the backend (or the vault in front of it) to receive and evaluate contextual attributes — which is exactly what the vault-as-middleware pattern provides.

## Threat Model

| Threat | Mechanism | Countermeasure |
|--------|-----------|----------------|
| **Rogue agent impersonation** | Attacker claims to be a legitimate agent | Vault validates delegation chain; short-lived creds scoped to caller |
| **Credential theft** | Long-lived API key exfiltrated | Short-lived creds limit window; per-operation scoping limits impact |
| **Tool chaining attacks** | Agent chains arbitrary calls via shared credentials | Intent-scoped credentials restrict what each issued token can do |
| **Compromised agent** | Legitimate agent is hijacked | User identity + delegation chain; abnormal patterns caught by telemetry |

## Open Questions

- If the vault itself is compromised, does the attacker get backend access for all users? What's the vault threat model?
- Long-running agentic tasks need credential renewal mid-task — how does handoff work without interrupting the workflow?
- Legacy backends that can't support attribute-aware auth still exist — what's the migration path when you can't modify the target system?

## See Also

- [[Agentic Workflow Patterns]] — the last mile gap affects all tool-use patterns; this is the security layer those patterns require
- [[Harness Engineering]] — vault-as-middleware is the "tool contracts and validators" harness dimension extended to backend identity bridging
- [[Threat Modeling]] — rogue agent impersonation is a concrete agentic threat with defined adversary, asset, and countermeasure
- [[Operational Security]] — credential hygiene (short-lived, scoped) is OPSEC applied to machine credentials
- [[AI Firewall Attack Surface]], [[Config Drift vs Adaptive Learning]], [[Credential Renewal for Long-Running Tasks]] — open questions on agentic security
