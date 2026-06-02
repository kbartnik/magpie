---
title: "MCP Protocol"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-7.md"
related:
  - "Tool Use in AI Systems"
  - "Harness Engineering"
  - "Claude Code Skills"
  - "Claude Code Hooks"
  - "AI Without Illusions (Series)"
tags:
  - agentic-systems
  - protocols
  - tool-use
  - cca-f
---

# MCP Protocol

The **Model Context Protocol** is an open protocol introduced by Anthropic in late 2024 that standardizes how AI models discover and interact with external tools, data sources, and prompts. It addresses the integration fragmentation problem: before MCP, every tool integration was bespoke, every framework invented its own shape, and switching models or orchestrators meant rewriting glue code.

MCP is to model-tool interaction what REST was to web APIs, or what LSP was to language server integration — a common interface that decouples producers from consumers.

## Architecture

MCP defines a client-server model:

```
AI Application (MCP Client)
    ↕ MCP Protocol (JSON-RPC over stdio or HTTP+SSE)
MCP Server
    ↕
[Files · APIs · Databases · Services]
```

The AI application (Claude Code, an IDE assistant, a custom agent) is the **client**. The system that owns the tools or data is the **server**. They communicate over a standardized protocol — either local `stdio` for co-located servers or HTTP with server-sent events for remote ones.

## Three Primitives

MCP servers expose three types of capabilities:

| Primitive | What it is | Analogy |
|-----------|------------|---------|
| **Tools** | Executable functions the model can call | Function calling, but standardized |
| **Resources** | Data sources the model can read dynamically | Context injection, but on-demand |
| **Prompts** | Reusable prompt templates the server offers | Shared system prompt fragments |

**Tools** are the primary primitive — the same concept as function calling, described in a standardized schema rather than ad-hoc JSON. A server might expose `query_database`, `search_documents`, `create_ticket`.

**Resources** are data the model needs to reason with but that isn't known at prompt-construction time. A resource might be a specific file, a database table schema, or a live web page. The model requests the resource; the server fetches and returns it. This is the tool-vs-context distinction applied at protocol level: resources don't require execution, just retrieval.

**Prompts** are reusable instruction templates the server can offer to the client. Less commonly used in practice, but useful for standardizing how certain tool interactions are framed across different AI clients.

## What MCP Gets Right

**Discoverability.** A client can query a server for its full capability set — what tools, resources, and prompts are available — and get back structured descriptions. The integration layer can be partially automated: connect to a server, discover its tools, present them to the model.

**Standardized transport.** JSON-RPC over stdio (for local) or HTTP+SSE (for remote) — one protocol, not one per integration. Reduces glue code significantly.

**Separation of concerns.** The AI application manages conversation and model interaction; the MCP server manages access to specific tools and data. Each side can be developed, tested, and maintained independently.

**Composability.** A client can connect to multiple servers simultaneously — your database through one, the file system through another, your ticketing system through a third — all through the same protocol, all presenting a unified tool surface to the model.

**Cross-client portability.** An MCP server built for one AI application can serve any MCP-compatible client. The team that owns the ticketing system can ship an MCP server; any AI tool in the organization can consume it without custom integration on either side.

## When MCP Earns Its Complexity

**Use MCP when:**
- Building integrations that need to work across multiple AI clients (Claude, an IDE assistant, a custom agent framework)
- Connecting a model to many external systems — the consistent protocol pays for itself by the third or fourth integration
- Tool providers and consumers are different teams or organizations — standard interface eliminates tight coordination
- Dynamic tool discovery is needed — capabilities change based on context, permissions, or server state

**Skip MCP when:**
- Single model, single tool, simple integration — direct function calling is simpler, faster, easier to debug
- Fixed tool set in a tightly controlled environment — the flexibility MCP provides isn't buying much
- Latency-critical paths — the protocol layer adds overhead; negligible usually, but worth accounting for
- Prototyping — hardcode the tool definitions first, add MCP when the integration pattern stabilizes

The honest framing: MCP is most valuable at scale and in ecosystems. For a focused, single-purpose integration, the setup cost may not be justified yet.

## Relationship to Claude Code

Claude Code's skills system (see [[Claude Code Skills]]) is the application-layer analog to MCP servers — skills provide tools, context, and behavioral guidance that the agent can discover and load. The underlying tool execution in Claude Code runs through a hook system (see [[Claude Code Hooks]]) that sits at the execution boundary MCP's server layer would occupy in a standards-based deployment.

MCP and Claude Code's native systems aren't in conflict — MCP is the cross-application standard; Claude Code's skills/hooks are the vault-specific implementation within Claude Code's runtime.

## See Also

- [[Tool Use in AI Systems]] — the broader concept; function calling mechanics, tool design principles, approval boundaries
- [[Harness Engineering]] — MCP servers are one possible implementation of the harness's tool execution layer
- [[Claude Code Skills]] — application-layer analog: skill discovery and loading within Claude Code
- [[Claude Code Hooks]] — the execution boundary in Claude Code's runtime that corresponds to MCP's server-side execution
- [[AI Without Illusions (Series)]] — source course (Part 7)
