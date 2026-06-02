---
title: "Semantic Diffusion"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/clippings/2026-06-01-no-vibes-allowed-context-engineering.md"
related:
  - "Agentic Workflow Patterns"
  - "Vibe-Coding Anti-Pattern"
tags:
  - epistemology
  - language
---

# Semantic Diffusion

[[Martin Fowler]]'s term (2006) for the process by which a precise technical term loses its meaning through popularity. A term starts with a clear, agreed-upon definition. It gains traction. More people use it. Each applies it to slightly different things. Eventually the term covers so many distinct practices that it conveys almost no information — it means whatever the speaker wants it to mean.

**The mechanism:** Precision is a barrier to entry. A vague term is easier to claim membership in. As a community grows, the incentive to broaden the term's scope increases, while the incentive to preserve precision decreases.

## Examples in AI/Software

**"Agent"** — once meant a specific pattern (tools in a loop, goal-directed behavior). Now applied to: microservices, chatbots, workflows, people, any software with an API. As Simon Willison noted: "An agent is just tools in a loop" — the return to precision comes after the term has been diluted.

**"Spec-driven development"** — started with a precise definition around verifiable feedback loops and code-as-assembly. Diluted to mean: a longer prompt, a PRD, a bunch of markdown files while coding, documentation for an open-source library. Dex Horthy (AI Engineer 2025) declared it fully diffused — the term is no longer useful for communication.

**"Microservices," "DevOps," "Agile"** — canonical earlier examples. Each started with a precise manifesto or definition; each was diluted within a few years of mainstream adoption.

## Why It Matters

Once a term is semantically diffused, conversations using it generate more confusion than clarity. Practitioners either:
- Adopt hyper-precise sub-terms ("RPI workflow" vs. "spec-driven dev")
- Return to first principles (define the specific practice, not the category)
- Accept that the term is now marketing language, not technical vocabulary

Dex Horthy's prediction: "RPI" (Research → Plan → Implement) will undergo the same diffusion once it becomes widely used. The stable underlying principle — compaction at phase boundaries — will outlast any term attached to it. See [[Agentic Workflow Patterns]].

## Countermeasures

- Anchor to observable behavior, not labels: describe what the practice *does*, not what it *is called*
- Use examples as definitions: "spec-driven dev means X, as demonstrated by Y" — the example resists drift better than the label
- Watch for the early diffusion signal: when a term starts appearing in contexts where its original definition wouldn't fit, the clock has started

## See Also

- [[Agentic Workflow Patterns]] — RPI workflow; Horthy's prediction that "RPI" will itself diffuse
- [[Vibe-Coding Anti-Pattern]] — "vibe-coding" risks the same diffusion; the anti-pattern is defined by specific failure modes, not by feeling
- [[IDSD]] — a term that attempts to resist diffusion by grounding itself in a specific framework (ICE: Intent, Context, Expectations)
