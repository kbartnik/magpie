---
title: "JSON Schema Discipline"
type: concept
status: active
created: 2026-05-20
updated: 2026-05-20
sources:
  - "archive/clippings/2026-05-20-eleven-agentic-patterns.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-6.md"
related:
  - "Agentic Workflow Patterns"
  - "Harness Engineering"
  - "Prompting as Specification"
  - "AI Without Illusions (Series)"
tags:
  - agentic-systems
  - schemas
  - structured-outputs
---

# JSON Schema Discipline

When to use JSON schemas at handoff boundaries in multi-agent systems — and when not to.

## The One Diagnostic Question

> **Does code branch on this output, or does a model/human read it?**

- **Code branches on it** → schema required
- **Model reads it as material to reason over** → prose (schema optional, often harmful)
- **Human reads it** → prose

Apply this question at every handoff boundary. No other rules needed.

## Where Schemas Are Required

The pattern doesn't work without a schema here.

| Handoff | Why |
|---------|-----|
| **Routing classifier output** | `switch(category)` needs an enum — "Refund Request" vs "refund" breaks dispatch |
| **Evaluator output** (evaluator-optimizer) | `if pass: ship else: revise` is the load-bearing line — prose breaks the loop |
| **Evaluator output** (iterative refinement) | Same: `if verdict.pass_: break` needs a typed boolean |
| **Programmatic prerequisites** | The hook *is* a schema check: `if customer_verified == True: allow` |
| **Structured handoff payload** | The human escalation payload *is* the schema — typed fields, not a transcript |

## Where Schemas Help but Aren't Required

| Handoff | Benefit |
|---------|---------|
| Parallelization with sectioning | Preserves attribution, source URLs, claim text through aggregation |
| Parallelization with voting | `{"verdict": "violation", "confidence": 0.87}` tallies cleanly |
| Hub-and-spoke researcher → coordinator | Structured findings prevent provenance leaks; evaluator can score subsections programmatically |
| Graceful degradation | `coverage_gap` flag is easier to compose into final output as a typed field |

## Where Schemas Are Counterproductive

Don't add schemas here — they damage the pattern.

- **Prose-to-prose chaining** (outline → draft → revise): each step's output is material for the next, not data for code; JSON strips the paragraph structure
- **Synthesizer output**: the synthesizer's job is fluent prose; `{"paragraph_1": "..."}` forces consumers to render it back to text
- **Orchestrator planning reasoning**: the orchestrator's next-step decision is fluid; schematizing it constrains the flexibility that makes the pattern work
- **Dynamic adaptive decomposition planning step**: same reason — the planning decision's shape is unknown in advance

## Dynamic Orchestrators + Schemas (Compatible)

The common mistake: "dynamic means no schemas anywhere." It doesn't. Three places to schema without constraining dynamism:

1. **Schema the outputs, not the decisions** — orchestrator reasons in prose; each worker it spawns produces typed output
2. **Schema the action menu, not the plan** — give the orchestrator a *tool menu* with typed parameter schemas; it dynamically picks which tool to call; each call is validated
3. **Schema the state, not the transitions** — the running world model can be typed; the decision about what to do next stays prose

## SDK Enforcement Points

Two places to enforce schemas at runtime:

**`output_format` on `query()`** — validates the top-level query result, re-prompts on mismatch, delivers typed object on `ResultMessage.structured_output`. Right for: top-level evaluator calls, top-level planner calls.

**`PostToolUse` hook on Agent tool** — fires when a subagent returns, receives the raw result string, can parse/validate and `return {"decision": "block", "reason": "..."}` to trigger retry with correction. Right for: subagent boundaries inside a `query()`.

The key distinction: `output_format` validates the *top* of the pipeline; PostToolUse hooks validate the *seams inside* it.

## Worked Example: Research Pipeline Schema Map

![[resources/media/56d9d7461b1130e79cb44b27c3c70197_MD5.webp]]

```
Topic prompt → Planner → Outline (JSON: schema required — code iterates sections)
Outline → Researcher × N → ResearchResult (JSON: schema required — coordinator joins by subsection_number)
ResearchResult × N → Synthesizer → Draft prose (NO schema — evaluator reads prose as text)
Draft → Evaluator → Verdict (JSON: schema required — coordinator branches on verdict.pass_)
Verdict → Coordinator → branch on typed fields
```

The echo key pattern: `ResearchResult.subsection_number` echoes `ResearchTask.subsection_number` — gives the coordinator a clean join key without LLM parsing.

## The Three-Level Output Taxonomy (from AI Without Illusions, Part 6)

A practitioner framing that complements the agentic-systems schema rules above:

| Level | Name | Mechanism | Reliability | Use when |
|-------|------|-----------|-------------|----------|
| 1 | **Free text** | Model generates prose | Variable | Human consumer only |
| 2 | **Prompt-formatted** | Ask nicely for JSON in prompt | Mostly works; degrades under pressure | Prototyping, low-stakes automation |
| 3 | **Schema-constrained** | API-level decoder enforcement | Guaranteed to parse; types guaranteed | Any production pipeline |

Level 2 and Level 3 are categorically different, not different in degree. At Level 2, you rely on the model's willingness to follow formatting instructions — which degrades under long outputs, complex schemas, and adversarial inputs. At Level 3, the constraint is structural: the decoder cannot produce output that violates the schema.

**Decision rule** (three questions):
1. Will code consume this output? → lean toward structure
2. Is the output shape known and stable? → a schema is appropriate
3. Is field-level reliability more important than expressive nuance? → enforce the schema

All three yes → Level 3. First only → Level 2 with validation. None → Level 1.

### Validation Layers

Beyond schema enforcement, structured outputs need a validation stack:

1. **Structural** — does it parse and match the schema? (Level 3 makes this almost always pass; Level 2 will occasionally fail here)
2. **Semantic** — are values plausible? (date not in future when it shouldn't be; amount positive; name not suspiciously short)
3. **Cross-field** — do fields agree? (if status is "completed", is there a completion date?)
4. **Confidence-based** — if log probabilities are available, flag low-confidence fields for human review

When validation fails: retry with the specific error appended to the prompt ("field `amount_usd` was negative; please re-extract as positive") rather than blind retry.

### Minimum Viable Schema First

Over-specification is as dangerous as under-specification. A 20-field schema produces worse extraction quality than a 5-field schema — the model must fill every field, accumulating errors. Start with the fields your downstream code actually needs. Expand when there's a concrete consumer.

## See Also

- [[Agentic Workflow Patterns]] — which patterns require schemas (full table)
- [[Harness Engineering]] — schema discipline is part of harness design
- [[Prompting as Specification]] — the output format component bridges prompting into schema design
- [[AI Without Illusions (Series)]] — Part 6 is the primary source for the 3-level taxonomy and validation pipeline
