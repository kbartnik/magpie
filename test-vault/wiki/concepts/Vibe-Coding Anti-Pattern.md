---
title: "Vibe-Coding Anti-Pattern"
type: concept
status: active
created: 2026-05-31
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-course-intro.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-0.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-4.md"
  - "[[archive/papers/2026-06-01-metr-ai-developer-productivity-rct]]"
  - "archive/clippings/2026-06-01-karpathy-vibe-coding-to-agentic-engineering.md"
  - "archive/clippings/2026-06-01-software-fundamentals-ai-age-pocock.md"
related:
  - "IDSD"
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
  - "LLM Mental Model"
  - "Prompting as Specification"
  - "AI Without Illusions (Series)"
tags:
  - methodology
  - ai-discipline
  - cca-f
---

# Vibe-Coding Anti-Pattern

A failure mode in AI-assisted programming where code is generated through conversational momentum and accepted based on surface appearance, without task specification, scope control, or verification. Named by [[Andrej Karpathy]] in 2025 for zero-stakes exploratory prototyping — then adopted by the industry as a general practice, with the playfulness stripped out.

## Karpathy's Own Framing (AI Ascent 2026)

Karpathy draws a sharp line between vibe coding and what he calls **agentic engineering**:

> "Vibe coding is about raising the floor for everyone in terms of what they can do in software. Agentic engineering is about preserving the quality bar of what existed before in professional software."

**Vibe coding** raises the floor — anyone can build anything, which is genuinely amazing. **Agentic engineering** is the discipline of coordinating fallible, stochastic agents to go faster *without sacrificing* your professional responsibility for the output. You're still accountable for your software. The engineering challenge is how to preserve that accountability while going much faster.

He also notes that "vibe coding" has itself undergone [[Semantic Diffusion]] — it was coined for zero-stakes prototyping with the playfulness intact, and became a general label for AI-assisted coding with the stakes stripped out.

**The MenuGen insight (spurious apps).** Karpathy built a full OCR + image generation pipeline app to show restaurant menu photos. Then saw the Software 3.0 version: give the photo to a multimodal model directly. His entire app "shouldn't exist" — it was architected in the old paradigm. The implication: when a new computing paradigm arrives, apps built in the previous paradigm don't just become slower comparatively — they become *structurally obsolete*, because the new paradigm eliminates the need for the intermediate architecture entirely.

> "The enemy is not AI. The enemy is sloppiness." — *AI Without Illusions*, Thomas Byern

The distinction: vibe-coding is not about *how much* AI you use. It's about whether usage is bounded, reviewable, verifiable, and owned.

## Specs-to-Code Is Vibe Coding by Another Name

The "specs-to-code" movement argues: write a specification, run AI to generate code, fix problems by editing the spec and running again — never look at the code. Matt Pocock (AI Engineer 2026, after 18 months teaching Claude Code to engineers) calls this vibe coding by another name.

Running the AI again and again without reading the code doesn't produce stable code — it compounds entropy. The Pragmatic Programmer: *software entropy* is what happens when each change is made without thinking about the system as a whole. Specs-to-code is divestment from system design, one iteration at a time.

**"Code is not cheap."** The counter to the common claim:

> Bad code is the most expensive it's ever been. A hard-to-change codebase blocks you from benefiting from AI. AI in a good codebase does really, really well. Good codebases matter more than ever, which means software fundamentals matter more than ever.

## Deep Modules: The AI-Specific Architecture Argument

John Ousterhout (*A Philosophy of Software Design*): a *deep module* has a simple interface hiding large, complex implementation. A *shallow module* has a complex interface with little hidden complexity.

AI creates shallow-module codebases by default. This is bad for AI for a **concrete architectural reason**: the model must explore more files and dependencies to understand the system, which burns context window faster. Shallow codebases trigger [[Context Rot]] structurally.

| | Deep modules | Shallow modules |
|--|---|---|
| Interface | Simple, well-designed | Complex, numerous |
| Implementation | Large, hidden | Small, minimal |
| AI navigability | Easy — few entry points | Hard — many files |
| Context cost | Low exploration overhead | High exploration overhead |
| Testability | Test at the interface boundary | Many mocked dependencies |

**"Design the interface, delegate the implementation."** The human designs the interface; the AI fills the implementation. Tests verify at the boundary. This is cognitively sustainable at high output velocity because the human only needs to hold the interface map in mind, not every implementation detail.

Kent Beck: "Invest in the design of the system every day." Specs-to-code is divestment from design.

## The Five-Stage Failure Cascade

Each stage enables the next:

1. **Ambiguous task specification** — "make the form better" instead of "add email validation using this regex, return 422 with this schema on failure." The AI complies. The output is plausible. Whether it did what was needed is a separate question vibe-coding never asks.

2. **Unconstrained scope** — because generation is fast, it's tempting to let the AI expand beyond the original task. A bug fix becomes a refactor; a refactor becomes an architecture change. Each step feels smooth. The cumulative result is a large diff nobody fully understands.

3. **Surface trust** — LLM output is fluent, well-formatted, and syntactically correct, which makes it easy to assume it's semantically correct. The code *looks* right. Vibe-coding exploits this bias.

4. **Skipped verification** — generated code is not tested, reviewed, or validated against the actual requirement. "If it runs, it ships." This is where real damage accumulates, because LLM errors are exactly the kind that survive a casual glance: correct structure, wrong logic; valid syntax, invented API.

5. **Diffused responsibility** — when code is generated conversationally and accepted without deep review, nobody actually understands or owns it. The engineer didn't write it; the AI doesn't "understand" it. Code that belongs to nobody is a maintenance problem, a security problem, and an accountability problem.

## The Seven Recurring Principles (Antidote)

Applied before accepting any AI-generated output:

| # | Principle | The question to ask |
|---|-----------|---------------------|
| 1 | **Clarity of task** | Can you state, before involving the AI, what you want, what the output should look like, and how you'll know if it succeeded? |
| 2 | **Bounded scope** | Is the AI asked to do one well-defined thing, or given a vague directive to interpret freely? |
| 3 | **Inspectability** | Can you read, understand, and evaluate what the AI produced? If not, the workflow is broken regardless of output quality. |
| 4 | **Evidence and grounding** | Is the output based on something verifiable — documentation, test results, explicit context — or bare generation? |
| 5 | **Validation** | Is the output tested? Does it pass checks you'd apply to human-written code? Does it meet the actual requirement? |
| 6 | **Human review** | Has a competent person looked with enough attention to catch the errors AI tends to make? |
| 7 | **Responsibility** | Is there a person who will own this output in production, maintain it, debug it, and answer for it? |

These are not a checklist to run once — they're a habit of thinking applied to every AI interaction.

## Relationship to Other Frameworks

**[[IDSD]]** — Both frameworks arrive at the same place from different angles. IDSD's ICE framework (Intent, Context, Expectations) is a structural solution to the same problem vibe-coding names as a pattern: when the agent fills intent gaps with discretion, the output belongs to no one. "Diffused responsibility" (stage 5) maps directly to IDSD's requirement that intent ownership stays with the person who wanted the outcome.

**[[Harness Engineering]]** — The 7 principles describe what a disciplined *practitioner* does. Harness engineering describes what the *system* enforces. Both are necessary: the harness can enforce scope boundaries and validation gates, but a practitioner without principles will find ways around them.

## The Experienced-Engineer Trap

Counterintuitively, experienced engineers are *more* susceptible to vibe-coding, not less. AI-generated code resembles their own style — familiar structure, sensible variable names, clean formatting. This mimicry triggers the same pattern-match shortcuts that let experienced engineers read code fast. The code looks like something they would have written, so they approve it faster than they should.

The rigorous verification that comes naturally when reviewing a junior engineer's unfamiliar code gets short-circuited by the fluency of AI output.

## The Fluency Heuristic and Experienced Engineers

From Part 4 of [[AI Without Illusions (Series)]]: explanation quality and answer quality are *separate variables*. A beautifully structured, logically sequenced explanation can accompany a completely wrong answer. LLMs produce fluent, well-structured text because that was the training objective — not because they verified the claims.

This is why "surface trust" (stage 3 above) is especially dangerous: experienced engineers are *more* susceptible to it, not less. AI-generated code mimics their style — familiar structure, sensible names, clean formatting — triggering the same pattern-match shortcuts that let them read code efficiently. The rigorous review that kicks in for unfamiliar junior-engineer code gets short-circuited.

The three-level trust framework from that course:
- **Plausible** — sounds reasonable; the model always achieves this
- **Evidence-backed** — grounded in verifiable sources
- **Verified** — independently checked; the only level carrying operational trust

See [[LLM Mental Model]] for the underlying mechanics.

## Empirical Evidence

The METR RCT (Becker et al. 2025) — see [[AI Productivity Research]] — provides direct quantification of this anti-pattern in action:

- Experienced developers using Cursor Pro + Claude 3.5/3.7 Sonnet on real OSS tasks were **19% slower** than without AI
- After the study concluded, those same developers estimated they had been **20% faster** — a 39-point gap
- Developers accepted <44% of AI generations; 56% frequently made major changes to accepted code; 9% of total task time went to review/cleanup

The post-hoc overestimate is the experimental confirmation of the **experienced-engineer trap** (Stage 3: surface trust) and the **expectation-reality gap** that vibe-coding creates. The feeling of productivity is decoupled from actual productivity.

## See Also

- [[IDSD]] — parallel methodology for the same problem; ICE framework maps to the 7 principles
- [[Harness Engineering]] — system-level enforcement of the boundaries these principles describe
- [[Agentic Workflow Patterns]] — bounded scope and human review apply directly to agentic orchestration
- [[Claude Code Hooks]] — hooks implement "bounded scope" and "validation" at the runtime level deterministically
- [[LLM Mental Model]] — the probabilistic pattern machine framing that explains why fluency ≠ correctness
- [[Prompting as Specification]] — the positive counterpart: ambiguous task specification (stage 1) is exactly what specification prevents
- [[AI Without Illusions (Series)]] — source course
- [[Stoicism and AI Disruption]] — philosophical convergence: Stoicism's "you cannot outsource wisdom" arrives at the same place as stage 5 (diffused responsibility) from the identity direction
- [[Ubiquitous Language Ownership]], [[Interface Design Responsibility at Scale]], [[Entropy-Degraded Codebase Recovery]] — open questions from Pocock's argument
