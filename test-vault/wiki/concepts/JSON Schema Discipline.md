---
tags: [concept, agents, llm]
cluster: agents
aliases: ["JSON schema", "structured output", "schema discipline", "output schemas"]
related: ["Agentic Workflow Patterns", "LLM Tool Calling", "Prompt Engineering"]
sources:
  - "[[archive/clippings/2026-06-04-eleven-agentic-patterns]]"
---

# JSON Schema Discipline

The diagnostic question for schema use at agent handoff boundaries: **does code branch on this output?**

## The Rule

- **If yes** — code branches on the output (if/switch on field values, routing logic, downstream system consumption) → enforce schema. The consumer needs a contract.
- **If no** — output is passed to another LLM prompt or displayed to a human → schema is optional. Don't add schema friction for aesthetic reasons.

Dynamic orchestrators and schemas are compatible: a schema for the *structure* of a task decomposition is fine; the *content* of the tasks can remain dynamic.

## Connections

- [[Agentic Workflow Patterns]] — schema enforcement is a cross-cutting concern for all handoff-heavy patterns (orchestrator-workers, hub-and-spoke)
- [[LLM Tool Calling]] — tool call responses are the canonical schema handoff; function signatures are contracts
- [[Prompt Engineering]] — structured output section in prompt engineering literature covers the same ground
