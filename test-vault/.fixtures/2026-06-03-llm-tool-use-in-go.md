---
title: "Implementing LLM Tool Use in Go: A Practical Guide"
author:
  - "Marcos Henrique"
url: "https://dev.to/mhenrique/implementing-llm-tool-use-in-go-a-practical-guide-4k2f"
description: "How to implement structured tool-calling with Claude and OpenAI APIs in idiomatic Go, using interfaces for tool definitions and JSON Schema for parameter validation."
tags:
  - "clippings"
  - "go"
  - "llm"
  - "agents"
published: 2026-04-15T09:00:00Z
created: 2026-06-03T08:14:22-04:00
---

# Implementing LLM Tool Use in Go: A Practical Guide

By Marcos Henrique
Published: 2026-04-15

>[!summary]
>A hands-on walkthrough of building LLM tool-calling in Go: defining tools as interfaces, generating JSON Schema from Go types, wiring up the request/response loop, and handling multi-turn tool conversations cleanly.

## Why Tool Use Changes the Agent Architecture

Tool use (also called function calling) is what turns a language model from a text predictor into an agent that can act. The model doesn't execute tools itself — it outputs a structured JSON object saying "call this tool with these arguments." Your host program executes the tool and feeds the result back. The model then continues reasoning with that result.

This is the **orchestrator-worker pattern** in practice: the LLM is the orchestrator, your Go program is the executor. The transformer's ability to produce structured, schema-conforming JSON output is what makes this reliable. Without that, you'd be parsing free text — fragile and frustrating.

## Defining Tools as Go Interfaces

The cleanest Go approach: define a `Tool` interface and let each tool implement it.

```go
type Tool interface {
    Name() string
    Description() string
    InputSchema() json.RawMessage  // JSON Schema for parameters
    Execute(ctx context.Context, input json.RawMessage) (string, error)
}
```

This gives you:
- A registry pattern: `map[string]Tool`
- Easy dispatch: look up by name, call `Execute`
- Schema generation per-tool without coupling to the API client

## JSON Schema for Tool Parameters

The LLM needs a JSON Schema describing valid inputs. You can write these by hand, or use a struct-tag-based generator like `invopop/jsonschema`:

```go
type SearchInput struct {
    Query   string `json:"query" jsonschema:"description=The search query"`
    MaxResults int `json:"max_results,omitempty" jsonschema:"default=5"`
}
```

The schema enforces that the model produces parseable input. This is the same discipline as using JSON Schema for API validation — the schema is the contract between the orchestrator (LLM) and the worker (your tool).

## The Tool Loop

The conversation loop for tool use is a while-not-done pattern:

```go
for {
    resp, err := client.Messages(ctx, req)
    if resp.StopReason == "end_turn" {
        break
    }
    if resp.StopReason == "tool_use" {
        for _, block := range resp.Content {
            if block.Type == "tool_use" {
                result := registry[block.Name].Execute(ctx, block.Input)
                req.Messages = append(req.Messages, toolResultMessage(block.ID, result))
            }
        }
    }
}
```

The loop is stateless between iterations — each pass appends to `Messages` and re-calls the API. This is the **prompt chaining** pattern: each tool result becomes part of the context for the next model call.

## HTTP Client Considerations

The Anthropic and OpenAI Go SDKs are thin wrappers around HTTP. For production use:
- Set explicit timeouts via `context.WithTimeout`
- Use `http.Client` with connection pooling (the default transport reuses connections)
- Handle rate limits with exponential backoff — wrap the client, don't scatter retry logic

The `net/http` package's interface-based design means you can inject a mock transport in tests, keeping tool-use logic testable without hitting the real API.

## Deep Read

**Key Insight:** The `Tool` interface pattern is idiomatic Go applied to a new domain. Accept interfaces, return structs — the tool registry accepts `Tool`, each implementation is a concrete struct. The LLM's JSON output becomes the glue between the model's reasoning and Go's type system, mediated by JSON Schema as the shared contract.

**What Surprised Me:** The tool loop accumulates the full conversation in `Messages` on the client side. There's no server-side session — every call is stateless from the API's perspective. This means the "agent's memory" between tool calls is just a Go slice that grows with each turn. For long-running agents, this is where context-window limits become a real engineering constraint, not just a theoretical one.

**Open Questions:**
- The article uses `json.RawMessage` for tool inputs to defer parsing to the tool itself. Is there a better pattern using generics — `Execute[T any](ctx context.Context, input T) (string, error)` — or does the registry pattern require the dynamic dispatch that `json.RawMessage` enables?
- When multiple tools are available, the model chooses which to call. Is there a way to bias tool selection without prompt engineering — e.g., providing per-tool usage examples in the schema description?
- The HTTP client section recommends connection pooling via the default transport. How does this interact with context cancellation — if the context is cancelled mid-stream, does the connection return cleanly to the pool?

**Wikilink Candidates:**
- [[LLM Tool Calling]] — the request/response loop, stop reasons, multi-turn tool conversations; primary source for this pattern
- [[Go HTTP Client Patterns]] — timeouts, connection pooling, mock transports for testing; not yet a dedicated page

**Connections:**
- [[Go Interfaces]] — the `Tool` interface is a direct application of Go's accept-interfaces/return-structs rule; the registry pattern is interface-based dispatch
- [[Agentic Workflow Patterns]] — tool use implements the orchestrator-workers pattern; the LLM is the orchestrator, Go tools are the workers
- [[JSON Schema Discipline]] — tool input schemas are the same JSON Schema discipline applied to LLM-facing APIs rather than HTTP APIs; the contract enforcement purpose is identical
- [[Transformer Architecture]] — the model's ability to produce schema-conforming JSON output (structured generation) is a direct consequence of next-token prediction over the vocabulary; without reliable structured output, the tool loop degrades to fragile text parsing
