---
type: concept
tags: [ai, agents, architecture, structured-output]
---

# JSON Schema Discipline

A single diagnostic question determines where JSON schemas belong in agentic systems:

> **Does code branch on this output, or does a model/human read it?**

If code consumes the output (a `for` loop, `if` statement, switch on category), you need a schema. If the output is material for another model to reason over, or for a human to read, schemas are optional at best and counterproductive at worst.

**Primary source:** [[archive/clippings/2026-06-04-eleven-agentic-patterns|The Eleven Patterns Behind Every Production Agentic System]]

---

## Where Schemas Are Required

The pattern does not work without them.

| Pattern | Schema location | Why |
|---------|----------------|-----|
| Routing | Classifier output | Category drives a switch; enum prevents `"Refund Request"` vs `"refund"` mismatches |
| Evaluator-Optimizer | Evaluator output | `if verdict.pass: ship else: revise` cannot exist without a typed verdict |
| Iterative Refinement | Evaluator output | Loop control (`if verdict.pass_: break`) is the load-bearing line of the entire pipeline |
| Programmatic Prerequisites | Hook check fields | `if customer_verified == True: allow` — the hook *is* a schema check |
| Structured Handoff | Handoff payload | Named fields (`customer_id`, `root_cause`) let a human act without parsing prose |

## Where Schemas Are Useful but Optional

| Pattern | Benefit |
|---------|---------|
| Parallelization (Sectioning) | Preserves attribution and provenance through merge |
| Parallelization (Voting) | Verdicts tally cleanly; `{"verdict": "violation", "confidence": 0.87}` vs prose |
| Hub-and-Spoke handoffs | Typed findings prevent provenance leakage between subagents |
| Graceful Degradation | `coverage_gap` flags are easier to compose into final output as typed fields |

## Where Schemas Are Counterproductive

| Situation | Why |
|-----------|-----|
| Prose-to-prose chaining (outline → draft → revise) | Strips natural paragraph structure; next model gets less to work with |
| Synthesizer output | Human/model reads it; forcing JSON means the consumer renders it back to prose |
| Orchestrator planning reasoning | Dynamic decisions over open problem space; schematizing constrains the flexibility the pattern requires |
| Dynamic adaptive decomposition | Same reason — schema the *outputs of subtasks*, never the planning decision |

---

## Dynamic Systems and Schemas Are Compatible

The misconception: "dynamic means no schemas." Reality: **schema the layer below the decision**.

Three compatible layers:

1. **Schema outputs, not decisions.** Orchestrator reasons in prose about what to do next; each worker produces typed output.
2. **Schema the action menu, not the plan.** Give the orchestrator a *menu of tools* with typed parameter schemas. It dynamically picks which tool; each call is validated against its schema. This is exactly what the Claude API `tools` parameter is for.
3. **Schema the state, not the transitions.** For orchestrators that maintain world models across iterations, the running state object can be typed. The decision about what to do next stays prose.

The tell that you've built routing disguised as orchestration: if you want a schema that *enumerates everything the orchestrator might decide to do next*, you have a closed decision set — that's routing.

---

## SDK Enforcement Points

### `output_format` — top-level query boundary
Validates the final output of a `query()` call. The SDK validates, retries on mismatch, and delivers a typed object on `ResultMessage.structured_output`. Right tool when the entire `query()` purpose is to produce a structured payload (e.g., evaluator as a standalone `query()` call).

### `PostToolUse` hooks — subagent boundaries
Fires in the parent process when a subagent returns. Parse and validate `tool_response.result`; return `decision: "block"` with a `reason` to trigger a retry. Right tool when the coordinator spawns subagents via Agent and wants typed results back without round-tripping through prose.

```
Pipeline enforcement map:
  Planner → output_format (top-level query → Outline schema)
  Researchers → PostToolUse hook (subagents → ResearchResult schema)
  Synthesizer → no schema (prose output, model/human consumer)
  Evaluator → output_format or PostToolUse depending on architecture
```

---

## The Pipeline Rule

**Schemas at boundaries where structured data crosses from model output to code input. Prose at boundaries where data flows model-to-model as reasoning material, or model-to-human as something to read.**

![[resources/media/56d9d7461b1130e79cb44b27c3c70197_MD5.webp]]

---

## Related

- [[Agentic Workflow Patterns]] — which patterns require schemas, which benefit, which don't
- [[Harness Engineering]] — schema discipline is the enforcement layer of harness engineering
