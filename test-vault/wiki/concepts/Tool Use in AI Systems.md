---
title: "Tool Use in AI Systems"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-7.md"
related:
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
  - "JSON Schema Discipline"
  - "Vibe-Coding Anti-Pattern"
  - "MCP Protocol"
  - "AI Without Illusions (Series)"
tags:
  - agentic-systems
  - ai-discipline
  - tool-use
  - cca-f
---

# Tool Use in AI Systems

The mechanism by which a language model moves from generating text to participating in workflows — calling functions, querying systems, reading files, and triggering actions in external services. Tool use is the inflection point where AI systems become system components rather than conversational partners.

> "The model proposes. Your system disposes. Build the system accordingly."

## The Core Mechanism: Function Calling

The model does not execute functions. It generates a structured request — tool name plus arguments — and your orchestration layer decides whether and how to act. The loop:

```
User prompt
  → Model reasons about available tools
    → Model emits function call (name + arguments)
      → Orchestration layer validates and executes
        → Result returned to model as context
          → Model generates final response
```

Every safety boundary lives in the gap between the model's proposal and the system's execution. The model's "call" is a structured output, not runtime execution. It can get the arguments wrong, select the wrong tool, or call when it shouldn't — your system is what catches those failures.

## Tool vs. Context

A critical distinction practitioners often conflate:

- **Tools** — for actions: things the model should *do* (query a database, send a notification, create a record)
- **Context** — for information: things the model should *reason with* (a config file, API documentation, a user's preferences)

If a model needs access to a configuration file, include the file contents in the prompt — don't build a `read_config` tool. Tools introduce execution complexity and failure modes; context is just tokens. Reach for tools only when the model needs to take an action or retrieve information that isn't known at prompt-construction time.

## Principles of Good Tool Design

**Narrow scope.** One tool, one thing. `manage_user` (create/update/delete/suspend) is a liability. Four tools with distinct names, clear descriptions, and separate authorization are safer and easier for the model to reason about.

**Explicit, precise descriptions.** The model selects tools based on their descriptions. "Does stuff with users" leads to incorrect selections. "Returns the email address and signup date for a user given their user ID" is precise enough to reason about. If two tools have overlapping descriptions, the model will confuse them.

**Typed, constrained parameters.** Every parameter needs a type, a description, and where possible, an enum. Free-text fields passed directly to backend queries are a security risk. If a status field can only be `"active"`, `"suspended"`, or `"closed"`, declare it as an enum — don't trust the model to guess valid values.

**Minimal authority.** Least privilege at the tool level: a read-only tool that returns order information is far safer than a read-write tool that can also modify orders. Separate the read and write surfaces; require different authorization for each.

**Idempotent where possible.** If a tool can be safely called twice with the same arguments without causing damage, that's a significant safety advantage. Models retry, loops re-execute, orchestration duplicates calls. Idempotent tools absorb these errors gracefully.

**Predictable, scoped output.** Return exactly what the model needs for the current task — not the full 200-field database record when it needs two fields. Scoped outputs reduce context waste, reduce misinterpretation, and reduce data leakage into the context window.

**Structured error returns.** Tools fail. Return interpretable error information:
```json
{"error": true, "error_type": "not_found", "message": "No account found with ID acct_0000."}
```
Not raw stack traces. Not empty responses. The model needs enough signal to tell the user something useful rather than hallucinating around the gap.

## The Three Approval Models

Match control structure to risk:

| Model | When to use | Mechanism |
|-------|-------------|-----------|
| **Automatic execution** | Read-only, safe, idempotent tools | Model calls → system executes → result returned |
| **Confirm-then-execute** | Write operations, anything with side effects | Model proposes → user sees and approves → then executes |
| **Plan-then-review** | Complex multi-step sequences | Model generates full plan → user reviews whole sequence → then executes |

The worst pattern — and the one most demos use — is automatic execution of everything. This works in demos because the stakes are zero. In production, the stakes are never zero.

## Input Validation Is Not Optional

Treat the model's function call arguments exactly as you'd treat input from an untrusted client, because that's effectively what they are. Validate before execution:

- Account ID doesn't match your format? Reject it.
- Date string in wrong format? Reject it.
- Extra parameters not in schema? Strip them.
- Argument outside enum range? Reject it.

Read-only operations with bad arguments return an error and life goes on. Write operations with malformed or hallucinated parameters can cause real damage. Validate at the boundary, always.

## The Ten-Question Checklist

Before exposing a system capability as a model-accessible tool:

1. **Well-defined interface?** Clean function signature, typed parameters, documented return shape — if not, the capability isn't ready.
2. **Can the model determine when to use it?** If even a human would struggle to decide when this tool applies vs. another, the model will struggle too.
3. **Worst case if called incorrectly?** Failed lookup = fine. Deleted production data = approval controls required, not just logging.
4. **Parameters constrained enough?** Enums, validated formats, bounded ranges — the more constrained the input, the less surface for damage.
5. **Output appropriately scoped?** Don't leak a full DB record when the model needs one field.
6. **Audit trail?** Every invocation logged: which conversation, what arguments, what result. If you can't log it, don't expose it.
7. **Human in the loop for high-risk ops?** Default assumption: write = approval required.
8. **Testable independently?** Should be callable with known inputs outside the AI context. If it only works when the model calls it, your testing has a critical gap.
9. **Better as a tool or as context?** If the model just needs to reason with some data, inject it into context — no tool needed.
10. **Appropriate at scale?** A tool called 10×/day by one user behaves differently at 1,000 calls/day by 1,000 users. Rate limits, cost, backend capacity.

## Observability Requirements

Tool-using systems require trace-level logging, not just prompt-in/response-out. Every interaction is a multi-step sequence, and every step is a potential failure point:

1. User message received
2. Model reasoning
3. Tool selection + argument construction
4. Execution request (including validation outcome)
5. Execution result
6. Model incorporates result
7. Final response

**What to log:** tool name, arguments (pre- and post-validation), execution result, errors, latency, correlation ID tying back to user and conversation.

**What to monitor:** tool call failure rate, argument validation failure rate (signals unclear tool definitions), tool selection accuracy (wrong tool for context), per-tool latency (4 sequential 3-second calls = 12 seconds of user wait).

Without trace-level observability, failures in tool-using systems are nearly impossible to diagnose. The model's final response gives you no information about which step in the tool chain went wrong.

## See Also

- [[MCP Protocol]] — the open standard for model-tool interaction; standardizes discovery, transport, and tool/resource/prompt primitives
- [[Harness Engineering]] — the execution boundary is harness engineering at the tool layer; tool use is one of the six harness dimensions
- [[Agentic Workflow Patterns]] — tool use is the mechanical foundation for every agentic pattern
- [[JSON Schema Discipline]] — tool parameter schemas are a specific application; the 3-level taxonomy applies here
- [[Vibe-Coding Anti-Pattern]] — automatic execution without approval is vibe-coding at the action level
- [[Claude Code Hooks]] — hooks are a specific harness-layer implementation of tool execution controls
- [[AI Without Illusions (Series)]] — source course (Part 7)
