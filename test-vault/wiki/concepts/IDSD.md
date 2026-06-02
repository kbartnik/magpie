---
title: "IDSD"
type: concept
status: active
created: 2026-05-29
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-29-idsd-method.md"
  - "[[archive/papers/2026-06-01-metr-ai-developer-productivity-rct]]"
  - "archive/clippings/2026-06-01-software-fundamentals-ai-age-pocock.md"
related:
  - "Agentic Workflow Patterns"
  - "Harness Engineering"
  - "Context Rot"
tags:
  - agentic-development
  - methodology
  - cca-f
---

# IDSD

**Intent-Driven Software Development** — a post-SDD methodology proposed by Kapil Viren Ahuja that argues Spec-Driven Development fails because humans cannot write complete specifications before the software exists.

> "SDD breaks because we asked humans to do the one thing humans cannot do: specify everything before it exists."

## The Core Critique of SDD

The OpenAI Symphony spec — 2,169 lines, 18 sections, formal RFC-grade language — was reverse-engineered from software that was already running. OpenAI built the internal tool first under one rule (no human-written code), then distilled the spec from the running system, then used that spec to generate reference implementations in six languages.

**The industry is selling this output as if it were the method.** A complete spec works when you can write it upfront. But the organization that produced the best-known complete spec produced it last, not first.

SDD also fails for a second reason: the agent fills the holes. Every engineer writes the spec however they want that morning, and the gaps are filled by agent discretion — not because the spec was poorly written, but because a complete specification is structurally impossible before the problem is fully understood.

## ICE: The Three Crafts

IDSD separates what was previously crammed into a single SPEC.md into three distinct, owned artifacts:

**Intent** — *What is wanted, owned by the human who wants it*

Five required parts:
1. Description of what is wanted (outcome, not implementation)
2. Constraints around it
3. Failure scenarios (what would make this wrong?)
4. Success scenarios (what does done actually look like?)
5. Connections — links to other intents this one touches, so changes are traceable

Miss any of the five, and you've left a hole the agent will fill by discretion.

**Expectations** — *The boundary, owned by the same person who owns the intent*

The scenarios under which the result counts as done, and the scenarios under which it has failed, written in terms the user would recognize — not implementation language.

This is "what would be called the spec," but IDSD deliberately doesn't call it that. The key constraint: **the definition of done must stay with the person who wanted the outcome.** The moment it drifts to whoever held the keyboard, the agent starts deciding "done" for them.

**Context** — *The how, owned by the harness*

The tech stack, existing system constraints, and codebase-specific knowledge. Context should be fed progressively as needed, not dumped as a wall at the start. This is [[Context Rot]] prevention at the methodology level.

## The Design Concept Gap (Brooks)

Frederick P. Brooks (*The Design of Design*): a **design concept** is the invisible shared theory of what you're building — not a markdown file, not a spec, but the ephemeral understanding that exists between collaborators. When multiple people (or a person and an AI) design something together, this design concept floats between them. If they don't share it, they'll produce something misaligned — each acting consistently from their own model.

Matt Pocock (AI Engineer 2026) applies this directly: "The AI didn't do what I wanted" is almost always a design concept gap, not a prompt quality gap.

**The "Grill Me" skill** operationalizes reaching a shared design concept:

> "Interview me relentlessly about every aspect of this plan until we reach a shared understanding. Walk down each branch of the design tree, resolving dependencies between decisions one by one."

The AI asks 40–100 questions before it's satisfied. The resulting conversation becomes a PRD or issue list.

Pocock explicitly contrasts this with Claude Code's plan mode: *"Plan mode is extremely eager to create an asset. I think it's nicer to reach a shared design concept first."* The insight: plan mode optimizes for shipping a plan; Grill Me optimizes for convergence on intent. These are different goals. When intent isn't shared, a well-formatted plan is still wrong.

The "Grill Me" skill connects to IDSD's Intent craft: both require that the human owns and articulates intent before the agent begins. The difference is framing — IDSD is methodological, Grill Me is tactical. Use Grill Me to surface the five Intent components before committing to an implementation plan.

## The Loop

![[resources/media/2ddab1c27d1f16b323d8338bd8703460_MD5.webp]]

The human never leaves the intent and expectations. The harness never invents what the human wanted.

## Why This Matters

**The METR controlled trial finding:** Experienced developers were measurably *slower* with AI assistance but walked away *certain* they had been faster. Being wrong while feeling fast is the core failure mode of AI-assisted development — not the AI's errors, but the developer's inability to detect them at speed and volume.

At 150–200 million tokens per day (Opus rates), three days of rework from a bad loop costs ~$985 of real money. The economics of agentic development mean that confident wrong code is expensive — not because per-token cost is high, but because agents burn many more tokens per wrong outcome than per right one.

## IDSD vs. Harnesses

Spec-kit, BMAD, and prompt-and-workflow tools are *harnesses* — useful, but only harnesses. IDSD is the *method* that decides what the work is before the harness handles it.

Adopting the harness without the method (the current default, because the harness has a download button) produces the same failure as SDD: the agent fills the intent gap with discretion.

## Relationship to Vault Workflow

IDSD's ICE framework maps loosely onto the vault's own knowledge workflow:
- **Intent** = what we want to understand from a source
- **Expectations** = what a well-formed wiki page looks like (frontmatter, connections, open questions)
- **Context** = the vault's existing wiki pages, fed via the ingest process

The vault's plan-before-write protocol and the ingest skill's angle-question step are informal implementations of the intent/expectations pattern.

## Related

- [[Agentic Workflow Patterns]] — catalog of patterns for multi-agent systems; IDSD provides the methodology for structuring inputs to those patterns
- [[Harness Engineering]] — IDSD's "harness owns the loop" is a specific application of harness engineering philosophy
- [[Context Rot]] — IDSD's progressive context feeding is an explicit anti-context-rot strategy at the methodology level
- [[Vibe-Coding Anti-Pattern]] — parallel framework naming the same failure mode from a practitioner angle; ICE maps to the 7 recurring principles
- [[ICE Framework Vault Ingest Mapping]], [[Diffused Responsibility at Scale]] — open questions on IDSD at scale
