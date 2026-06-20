---
tags: [concept, agents, llm]
cluster: agents
aliases: ["tool use", "function calling", "tool calling", "LLM tools", "function tools"]
related: ["Agentic Workflow Patterns", "JSON Schema Discipline", "Prompt Engineering", "Go HTTP Client Patterns"]
sources:
  - "[[archive/clippings/2026-06-04-llm-tool-use-in-go]]"
  - "[[archive/books/2026-06-04-prompt-engineering-for-llms]]"
---

# LLM Tool Calling

The mechanism by which LLMs request execution of external functions. The model outputs a structured tool call; the harness executes it and returns the result; the model continues with the result in context.

## Protocol

1. Tools defined in the API call (name, description, JSON Schema for parameters)
2. Model outputs a tool call (structured JSON matching the schema)
3. Harness executes the function
4. Result appended to context as a tool result message
5. Model continues generation

## Design Principles

**Tool descriptions matter more than tool names.** The model reads the description to decide when to use the tool. Ambiguous descriptions cause incorrect selection.

**Keep tools focused.** One tool, one capability. Compound tools (do A and B) make it harder for the model to select correctly.

**JSON Schema is the contract.** The schema validated at execution time; a schema violation means the model called the tool incorrectly. See [[JSON Schema Discipline]].

## Connections

- [[Agentic Workflow Patterns]] — tool calling is the mechanism behind orchestrator-workers and most multi-agent patterns
- [[JSON Schema Discipline]] — tool parameters are the canonical schema handoff boundary
- [[Prompt Engineering]] — tool descriptions are part of the prompt; quality of descriptions determines selection accuracy
- [[Go HTTP Client Patterns]] — Go implementation of tool execution uses HTTP clients to call external services
