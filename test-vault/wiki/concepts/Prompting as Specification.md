---
title: "Prompting as Specification"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-5.md"
related:
  - "LLM Mental Model"
  - "JSON Schema Discipline"
  - "Vibe-Coding Anti-Pattern"
  - "AI Without Illusions (Series)"
tags:
  - prompting
  - ai-discipline
  - cca-f
---

# Prompting as Specification

A prompt is not an incantation. It is a specification that shapes the probability distribution over possible outputs. Every useful detail eliminates a class of unwanted results — exactly as a function signature, API contract, or acceptance criterion does in conventional engineering.

The quality of results depends on how clearly the task is defined, how well the context is scoped, how precisely the expected output is described, and whether the model has enough information to do the job — not on knowing the right magical phrase.

## The Anatomy of a Professional Prompt

Eight components. Not all are needed for every prompt; the skill is knowing which this task requires.

| Component | Purpose | Most Often Omitted? |
|-----------|---------|---------------------|
| **Objective** | What the model should produce — specific, not "help me with this" | Rarely — but usually too vague |
| **Context** | Raw material the model works with; too little = guessing, too much = distraction | Common to have too much irrelevant context |
| **Role/perspective** | Sets implicit expectations for depth and vocabulary; weaker than explicit constraints | Often overused as substitute for constraints |
| **Constraints** | Boundaries of acceptable output — length, format, inclusions, exclusions, technical standards | **Most commonly missing; most powerful** |
| **Examples** | Concrete demonstrations; give the model a pattern to match, not an abstract instruction to interpret | Often skipped |
| **Counterexamples** | Demonstrations of what *not* to produce, with explanation; sharpens the decision boundary | Almost always skipped |
| **Output format** | The shape of the expected response — JSON, markdown, specific columns | Frequently missing or buried |
| **Failure instructions** | What to do when uncertain, when input is ambiguous, when task can't be completed | Nearly always missing |

Without failure instructions, the model defaults to confident improvisation. "If you are unsure, say so and explain what additional information would help" changes the output distribution in a genuinely useful way.

## Constraints Are the Most Valuable Component

Most prompts are all objective and no constraints. Constraints do three things:

1. **Reduce variance** — "summarize in 2–4 sentences, focusing only on technical claims" vs. "summarize this"
2. **Make output reviewable** — constraints create a checklist you can evaluate against (pass/fail, not subjective)
3. **Bridge the generalist gap** — LLMs produce generalist output by default; constraints shape that into something useful for a specific context

As you gain experience, most iteration happens in constraints, not the objective. The objective stays stable; constraints get refined as you learn what the model gets wrong.

## Exploration vs. Execution Prompts

These are fundamentally different modes requiring opposite approaches.

| Mode | When | Approach | Mistake |
|------|------|----------|---------|
| **Exploration** | Don't yet know what the right answer looks like; surveying a space | Open-ended; give the model room to roam | Applying execution constraints too early, closing off useful options |
| **Execution** | Know exactly what you want; need reliable, repeatable output | Tight specification; every degree of freedom you care about nailed down | Loose prompts when you need a specific target |

A practical workflow: explore to understand the problem space → identify the direction → write a tight execution prompt for the deliverable.

## The Bad Prompt vs. Good Prompt Pattern

**Vague:**
> "Review this PR description and give me feedback."

**Professional:**
> "You are reviewing a pull request description written for a team of backend engineers. The PR modifies the authentication middleware to support API key rotation.
>
> Evaluate against these criteria:
> 1. Does it clearly state what changed and why?
> 2. Does it describe how to test the change?
> 3. Does it mention migration steps or breaking changes?
> 4. Is it clear enough for a teammate unfamiliar with this subsystem?
>
> For each criterion: meets / partially meets / fails. For partial or failing: quote the specific passage and suggest a concrete revision. If a section is missing entirely, draft what it should contain.
>
> Respond in markdown, under 400 words."

The improvement comes entirely from making the task concrete, not from theatrical framing or magic phrases.

## Signs Your Prompt Is Too Vague

- You run it three times and get meaningfully different outputs
- You can't describe in one sentence what the model should produce
- No output format, length, or structure specified
- No constraints — anything the model produces technically follows instructions
- No failure instructions
- Someone else reads the prompt and asks "what exactly do you want here?"
- You're iterating on word choice rather than task structure
- You added "think step by step" but can't explain why that helps for this task

## Prompts as Engineering Artifacts

Once prompts are part of a workflow or product, treat them like code:

- **Version control** — who changed what and why; ability to diff and revert
- **Readability** — comments where intent isn't obvious; if a constraint was added because of a specific failure, document it
- **Review** — any prompt feeding a production system or customer-facing feature should be read by another person
- **Parameterization** — separate stable logic from variable inputs (prompt templates with clear placeholders)
- **Empirical iteration** — when you change a prompt, run it against your test cases; did it improve what you cared about, and did it regress anything?

## Connection to Output Schemas

The output format component of a prompt bridges directly into [[JSON Schema Discipline]]. For any output consumed by code, specifying the format in the prompt (Level 2) is a weak approximation of schema-constrained generation (Level 3). They address the same problem — output predictability — at different reliability levels.

## See Also

- [[LLM Mental Model]] — why prompts are probabilistic specifications, not commands
- [[JSON Schema Discipline]] — the output format component extended to machine-readable outputs
- [[Vibe-Coding Anti-Pattern]] — ambiguous task specification (stage 1) is the failure this concept prevents
- [[AI Without Illusions (Series)]] — source course (Part 5)
