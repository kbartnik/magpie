---
title: "Progressive Disclosure Architecture"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources:
  - "archive/clippings/2026-05-29-skills-conceptual-deep-dive.md"
  - "archive/clippings/2026-05-29-skills-deep-dive-part-1.md"
  - "archive/clippings/2026-05-29-skills-deep-dive-part-2.md"
related:
  - "Claude Code Skills"
  - "Context Rot"
  - "Harness Engineering"
tags:
  - claude-code
  - skills
  - context-engineering
  - cca-f
---

# Progressive Disclosure Architecture

PDA is the design pattern underlying [[Claude Code Skills]]: a skill should be an **orchestrator that loads what it needs**, not an encyclopedia that carries everything.

> "A skill should be an orchestrator, not an encyclopedia."

The alternative — loading all documentation upfront (the "documentation dump") — causes context rot: 50KB of PlantUML syntax loaded for every diagram request, 78–90% of it wasted.

## The Three Tiers

| Tier | What | Token cost | When loaded |
|------|------|-----------|-------------|
| 1: Metadata | YAML frontmatter (`name` + `description`) | ~100 tokens/skill | Session start — always |
| 2: Instructions | Full `SKILL.md` body | 1–5K tokens | When skill is selected for a task |
| 3: Resources | Reference files, scripts, templates | 0 (filesystem) until read | On-demand, directed by Tier 2 |

Tier 3 resources don't count against the context window until explicitly read by Tier 2 instructions — this is what makes the knowledge base "effectively unbounded."

## Orchestrator vs. Encyclopedia

**Encyclopedia approach (50KB monolith):**
- Loads all diagram syntax (sequence, class, ER, flowchart) every time
- User requests a flowchart: 45KB wasted, 5KB used
- Slower, more expensive, more noise for the LLM

**Orchestrator approach (3KB core + lazy-loaded refs):**
- SKILL.md contains routing logic: "If sequence diagram → load `references/sequence_diagrams.md`"
- User requests a flowchart: 3KB + 5KB = 8KB total
- 5 diagrams in a session: 44KB vs. 250KB traditional (82% savings)

## The Three Pillars

**1. Reference Files + Lazy Loading**

Heavy documentation lives in `references/` outside the skill. The skill loads only what's needed.

- Keep references focused: one topic per file
- Target 5–15KB per reference file
- Organize by use case (recommended), complexity level, or feature area
- Name files descriptively: `sequence_diagrams.md` not `ref1.md`

**2. Scripts for Mechanical Work**

API calls, data processing, and complex operations move to `scripts/`. The skill only specifies inputs/outputs.

- Scripts run outside Claude's context: their internal logic costs 0 tokens
- Scripts must emit structured, parseable output: `SUCCESS: <url>` or `ERROR: 404 - Database not found`
- Vague errors (`"Something went wrong"`) break the AI resilience layer
- Use environment variables for secrets, not hardcoded values

**3. AI Resilience Layer**

Claude provides the intelligence layer for edge cases, error interpretation, and UX. Scripts fail hard with structured codes; the AI interprets those codes into user guidance, alternative lookups, and retry logic.

Example error recovery flow:
```
Script: "ERROR: 404 - Database not found: abc123"
AI: Searches workspace for databases → presents options → user selects → retries
```

"Every error path leads to either user guidance or automated recovery — no dead ends."

## Token Math

| Scenario | Traditional | PDA | Savings |
|----------|------------|-----|---------|
| Single flowchart | 50KB | 8KB | 84% |
| 5 mixed diagrams | 250KB | 44KB | 82% |
| Notion uploader workflow | 150KB | 10KB | 93% |

## When to Apply PDA

Apply when your skill:
- Has >10KB of reference documentation
- Supports multiple distinct use cases (each needs different docs)
- Integrates with external APIs or complex tools
- Will grow over time

Stay basic when:
- Total size <5KB (overhead isn't worth it)
- All information is always needed (no conditional loading benefit)
- Simple, stable, focused task

The inflection point is ~10KB. Below that, basic skills are fine. Above that, PDA pays dividends quickly.

## SKILL.md Compliance

Per the official spec:
- `name`: hyphen-case, matches folder name (`^[a-z0-9-]+$`)
- `description`: third-person, specific trigger phrases ("This skill should be used when...")
- Body: imperative/infinitive form throughout ("Load the reference file. Validate inputs.")
- Target size: 1,500–2,000 words; max ~5,000 words
- Detailed docs → `references/`; reusable code → `scripts/`; output templates → `assets/`

## Declarative vs. Imperative

PDA skills are **declarative** (procedural manuals) vs. OpenAI Tools which are **imperative** (function schemas). Skills excel for multi-step, ambiguous workflows where reasoning is required; Tools excel for discrete, typed API calls. The hybrid architecture uses skills to orchestrate MCP tool calls.

## Related

- [[Claude Code Skills]] — the product that implements PDA; Skills 2.0 adds subagent execution and dynamic context injection
- [[Context Rot]] — the failure mode PDA prevents; unbounded documentation loading is its specific cause in the skills context
- [[Harness Engineering]] — PDA is a harness-level pattern: deterministic context management via three-tier loading
