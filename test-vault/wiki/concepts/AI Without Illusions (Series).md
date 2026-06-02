---
title: "AI Without Illusions (Series)"
type: synthesis
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-0.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-1.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-2.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-3.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-4.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-5.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-6.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-7.md"
related:
  - "LLM Mental Model"
  - "Prompting as Specification"
  - "Vibe-Coding Anti-Pattern"
  - "Context Rot"
  - "JSON Schema Discipline"
  - "Tool Use in AI Systems"
  - "MCP Protocol"
  - "Harness Engineering"
tags:
  - ai-discipline
  - course
  - cca-f
---

# AI Without Illusions (Series)

A 20-part practitioner course by [[Thomas Byern]] on disciplined AI use in professional software work. The position: AI is powerful, imperfect, and worth understanding properly. The enemy is not AI — it's the sloppy, verification-free adoption that produces fragile systems and diffused responsibility.

> "This course is not anti-AI. It is anti-sloppiness."

Parts 0–7 ingested 2026-05-31. Parts 8–20 not yet captured.

## Course Structure

**Block I — Foundations** (Parts 1–4): what these systems are and how they actually work
**Block II — Control** (Parts 5–11): making them usable, reliable, and safe
**Block III — Engineering Use** (Parts 12–18): how AI fits into real programming work
**Block IV — Professional Operation** (Parts 19–20): team and system level

## Ingested Parts

| Part | Title | Archive slug | Key Concept |
|------|-------|-------------|-------------|
| 0/20 | How to use this course — and what "not vibe-coding" means | 2026-05-31-ai-without-illusions-part-0 | [[Vibe-Coding Anti-Pattern]], 7 principles |
| 1/20 | The generative AI landscape in 2026 | 2026-05-31-ai-without-illusions-part-1 | 5-layer model, agent spectrum |
| 2/20 | How LLMs work — without the math rabbit hole | 2026-05-31-ai-without-illusions-part-2 | [[LLM Mental Model]], fluency ≠ truth |
| 3/20 | Context windows, memory, and why models seem to forget | 2026-05-31-ai-without-illusions-part-3 | [[Context Rot]], repo-context gap |
| 4/20 | Reasoning, planning, and the illusion of intelligence | 2026-05-31-ai-without-illusions-part-4 | explanation quality ≠ answer quality |
| 5/20 | Prompting for professionals — specification, not spellcasting | 2026-05-31-ai-without-illusions-part-5 | [[Prompting as Specification]] |
| 6/20 | Structured outputs and schemas | 2026-05-31-ai-without-illusions-part-6 | [[JSON Schema Discipline]], 3-level taxonomy |
| 7/20 | Tool use, function calling, and MCP | 2026-05-31-ai-without-illusions-part-7 | [[Tool Use in AI Systems]], [[MCP Protocol]] |

## The Five-Layer Model (Part 1)

The structural backbone for reasoning about AI systems:

1. **Model** — LLM; takes text in, produces text out; powerful and unreliable in specific ways
2. **Context layer** — everything fed into the model's input; quality here determines output quality more than model choice
3. **Tool layer** — external capabilities the model can invoke (search, code execution, databases, APIs)
4. **Orchestration layer** — logic managing multi-step interactions; determines autonomy level
5. **Human + evaluation layer** — engineers, users, reviewers, and measurement; not optional in the current state of the technology

## The Seven Recurring Principles (Part 0)

Applied before accepting any AI output:

1. Clarity of task — can you state what you want before involving the AI?
2. Bounded scope — one well-defined thing, not an open directive
3. Inspectability — can you evaluate what was produced?
4. Evidence and grounding — is output based on something verifiable?
5. Validation — tested? meets actual requirement?
6. Human review — attention enough to catch the errors AI tends to make?
7. Responsibility — does someone own this in production?

## Key Insights Across Parts

**Alignment ≠ truthfulness** (Part 2): Models are trained to be helpful and polite, not truthful. A fully aligned model still hallucates. Grounding is the mechanism for truthfulness.

**Fluency heuristic** (Part 4): The better you are at recognizing good human reasoning, the more vulnerable you are to trusting polished AI output. Experienced engineers are more susceptible, not less.

**Context is workspace, not memory** (Part 3): Every call starts fresh. The repo-context gap means the model always works with a tool-selected subset of your codebase.

**Minimum viable schema** (Part 6): Over-specification degrades extraction quality. Start with 4–5 fields, expand when there's a concrete consumer.

## See Also

- [[LLM Mental Model]] — probabilistic pattern machine framing (Part 2)
- [[Prompting as Specification]] — 8-component anatomy, exploration vs execution (Part 5)
- [[Vibe-Coding Anti-Pattern]] — 5-stage failure cascade, 7 principles (Part 0)
- [[Context Rot]] — failure modes: pollution, truncation, drift (Part 3)
- [[JSON Schema Discipline]] — 3-level taxonomy, validation pipeline (Part 6)
- [[Harness Engineering]] — system-level complement to practitioner principles
- [[Stoicism and AI Disruption]] — same epistemic position ("you cannot outsource wisdom") argued from philosophy rather than software engineering
