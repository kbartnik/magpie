---
title: "No Vibes Allowed: Solving Hard Problems in Complex Codebases"
type: clipping
captured-date: 2026-06-01
source-url: "https://www.youtube.com/watch?v=rmvDxxNubIg"
---

# No Vibes Allowed: Solving Hard Problems in Complex Codebases

**Speaker:** Dex Horthy, HumanLayer
**Event:** AI Engineer 2025
**Published:** 2025-12-02
**Duration:** 20:31

## Summary

Dex Horthy argues that AI coding tools can work effectively in brownfield, complex codebases — not by waiting for smarter models, but by applying context engineering principles today. The talk introduces the "dumb zone" (context window degradation past ~40% capacity) and the Research → Plan → Implement (RPI) workflow as a structured discipline for staying in the "smart zone."

Core claims:
- Most AI coding failures in production codebases are context management failures, not model capability failures
- Sub-agents should be used to control context, not to anthropomorphize team roles
- Mental alignment (keeping the team synchronized on how the codebase is changing) is code review's highest-leverage purpose
- "Spec-driven development" is semantically diffused and already useless as a term (Martin Fowler's 2006 pattern)
- The ceiling on solvable complexity scales with how much context engineering you're willing to do

## Key Concepts

**The Dumb Zone:** Beyond ~40% context fill, performance begins to degrade. Too many MCPs means always working in the dumb zone.

**Intentional Compaction:** Deliberately compress the context window at phase boundaries. Research → compact to a research doc. Plan → compact to a plan file with code snippets. Implement → execute the plan with a small context.

**RPI Workflow (Research → Plan → Implement):**
1. Research: find exact files and line numbers, stay objective, output a compact research document
2. Plan: outline exact steps with file names, line snippets, and test verification points after every change
3. Implement: execute the plan, keep context low

**Sub-agents for context control:** Don't use sub-agents to model team roles (frontend agent, backend agent). Use them to isolate expensive codebase-reading work. A sub-agent forks a blank context, does all the searching/reading, and returns a single succinct finding. The parent agent never enters the dumb zone.

**Mental Alignment:** Code review's true purpose is keeping the team on the same page about how the codebase is changing and why — not just finding bugs. Sharing plan files (not just diffs) on PRs takes reviewers on the same journey as the implementer.

**Don't Outsource Thinking:** AI amplifies the thinking you bring. Bad research = 100x bad code. Human effort should concentrate at the highest-leverage stages: reviewing the research and approving the plan.

**Harness Engineering:** Explicitly named as the discipline of customizing integration points in Claude Code/Codex/Cursor. Distinct from generic "context engineering" — harness engineering is codebase-specific.

**Progressive Disclosure for CLAUDE.md:** Instead of a massive root-level CLAUDE.md that consumes the entire smart zone just to orient the agent, shard documentation down the stack. Each directory level has only the context needed for work at that level. On-demand sub-agents pull vertical slices of codebase context as needed.

**Semantic Diffusion (Martin Fowler, 2006):** Good terms with precise definitions get diluted when they become popular. "Agent," "spec-driven development" are current casualties. RPI itself is probably next.

## Transcript Excerpt (Key Passage)

> "Sub-agents are not for anthropomorphizing roles. They are for controlling context."

> "AI cannot replace thinking. It can only amplify the thinking you have done or the lack of thinking you have done."

> "A bad line of code is a bad line of code and a bad part of a plan could be a hundred bad lines of code and a bad line of research — a misunderstanding of how the system works — your whole thing is going to be hosed."

## Deep Read

**Key Insight:** The ceiling on what AI can solve in a complex codebase is almost entirely a context engineering problem, not a model capability problem — and the primary lever is *compaction at phase boundaries* (research → plan → implement), not smarter prompting.

**What Surprised Me:** Dex declares "spec-driven development" semantically dead using Martin Fowler's 2006 "semantic diffusion" concept — the pattern where a precise term becomes popular, gets applied to 100 different things, and loses all meaning. He applies the same diagnosis that killed "agent" as a useful term to "spec-driven dev," noting it now means anything from "a better prompt" to "markdown files while coding" to "documentation for an open source library."

**Open Questions:**
- The 40% "dumb zone" threshold is presented as empirical but model-agnostic — does it hold for Claude 3.5 vs. GPT-4o vs. Gemini 1.5, or is it highly model-specific?
- Mental alignment via plan files works for small teams reviewing each plan — how does this scale when shipping 3x volume, since there are 3x more plans to review?
- If RPI workflow is about to undergo semantic diffusion (as Dex acknowledges), what's the stable underlying principle that should survive the term's eventual dilution?

**Wikilink Candidates:**
- Context Compaction — does not exist; integrates into Context Rot as "Intentional Compaction" section
- RPI Workflow — does not exist; integrates into [[Agentic Workflow Patterns]]
- Semantic Diffusion — does not exist; the Martin Fowler concept worth a stub

**Connections:**
- [[Harness Engineering]] — "harness engineering" is explicitly named here as the codebase-specific layer of context engineering; customizing integration points (hooks, CLAUDE.md, skills) in Claude/Codex/Cursor
- Context Rot — "dumb zone" is a concrete threshold framing: ~40% context fill = measurable performance cliff; intentional compaction is the systematic harness response
- [[Agentic Workflow Patterns]] — RPI workflow is a composition of Prompt Chaining (research → plan → implement) with sub-agent Parallelization for the research phase; sub-agents used for context isolation, not role anthropomorphism
- Vibe-Coding Anti-Pattern — "no vibes allowed" / "no slop" is the direct thesis of this talk; the antidote is structured context engineering, not better vibes
- LLM Mental Model — statelessness + tokens-in/tokens-out framing is explicit; trajectory problem (bad conversation history biases next output) is named
- Progressive Disclosure Architecture — "shard CLAUDE.md down the stack" matches PDA exactly; on-demand sub-agents for vertical context slices
- AI Productivity Research — directly cites Eigor's 100k-developer survey finding: AI tools increase code volume but much of it is rework of AI-generated slop from the previous week

**Image Candidates:** none (video talk, no embedded images in archive file)
