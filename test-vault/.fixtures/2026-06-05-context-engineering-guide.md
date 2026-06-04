---
title: "Context Engineering: The Hidden Skill Behind Effective AI Systems"
author:
  - "Shreya Shankar"
url: "https://www.shreya-shankar.com/context-engineering"
description: "Context engineering is the discipline of deciding what information to put in a model's context window — not just how to phrase it. Argues that this is the primary lever for improving AI system reliability in production."
tags:
  - "clippings"
  - "llm"
  - "agents"
  - "context-management"
published: 2025-09-15T00:00:00Z
created: 2026-06-05T10:00:00-04:00
---

# Context Engineering: The Hidden Skill Behind Effective AI Systems

By Shreya Shankar
Published: 2025-09-15

>[!summary]
>Context engineering is the discipline of selecting, structuring, and managing what information appears in the model's context window. Distinct from prompt engineering (how you phrase a request), context engineering is about what information is available at all — and this determines a ceiling that prompt engineering cannot exceed.

## The Distinction

**Prompt engineering:** How you phrase your request. "Be concise" vs "Answer in one sentence." "You are an expert." Few-shot examples. Chain-of-thought instructions.

**Context engineering:** What information is present in the context window at all. Retrieved documents, conversation history, tool results, system state, schemas, examples — the composition and management of the total context.

Prompt engineering optimizes within a fixed context. Context engineering shapes the context itself.

## Why Context Is the Ceiling

A model cannot use information it doesn't have. No amount of instruction will make a model answer correctly about a document that isn't in its context. This seems obvious but has a non-obvious implication: most reliability failures in production AI systems are context failures, not prompt failures.

Common context failures:
- **Stale context:** The model is working with outdated information (old document versions, previous conversation state)
- **Missing context:** A relevant document wasn't retrieved; a prior tool result wasn't included
- **Noisy context:** Too much irrelevant information; the model attends to the wrong parts
- **Context overflow:** The context exceeds the window limit; earlier content is lost or degraded

## Context Window as Working Memory

The context window functions as working memory for the model — but with important differences. Unlike human working memory, the model's context window is:
- **Non-persistent:** Each call starts fresh (unless state is explicitly managed)
- **Flat:** All tokens are nominally equal; there is no automatic salience hierarchy
- **Token-limited:** Hard capacity limit; overflow causes degradation, not graceful forgetting

Context engineering is essentially working memory management for AI systems.

## Practical Patterns

**Retrieval-augmented context:** Pull relevant documents from a vector store based on the current query. The quality of this retrieval step determines the quality of the answer.

**Progressive context loading:** Start with a minimal context; load more information as the task progresses. Prevents context overflow on long tasks.

**Context compression:** Summarize previous conversation turns or retrieved documents before including them. Reduces token count while preserving key information.

**Scratchpad separation:** Keep the model's working notes in a structured scratchpad separate from the task context. Prevents reasoning traces from contaminating the information context.

## Deep Read

**Key Insight:** Context engineering reframes "why is this AI system unreliable?" from "the prompt is wrong" to "the context is wrong." This is a more productive failure analysis — context failures are usually findable (missing document, stale state) whereas prompt failures are often subjective (this instruction isn't clear enough).

**What Surprised Me:** The article argues that context window size increases (100K+ tokens) don't eliminate the need for context engineering — they change the failure mode. Small windows fail by truncation; large windows fail by *dilution*: the model attends to the wrong parts of a large context, producing correct-sounding but poorly-grounded answers. Context management remains necessary even when the window could theoretically fit everything.

**Open Questions:**
- "Noisy context" is identified as a failure mode but measuring it is hard — how do you detect that the model attended to irrelevant parts vs. the correct parts? Is there a practical instrumentation pattern?
- Context compression (summarizing history) loses information. Is there a principled way to decide what to compress vs. retain, given that you don't know which parts of the history will matter for future queries?
- The scratchpad pattern (separate reasoning from information context) is advocated in several frameworks. Does empirical evidence support that separated scratchpads improve reliability vs. mixed reasoning, or is this a heuristic without strong evidence?

**Wikilink Candidates:**
- [[Context Engineering]] — context window as working memory; context failures vs prompt failures; stale/missing/noisy/overflow failure modes; progressive loading, compression, scratchpad; not yet a wiki page

**Connections:**
- [[Agentic Workflow Patterns]] — context management is a cross-cutting concern for all agentic patterns; hub-and-spoke specifically isolates context to prevent cross-contamination
- [[Transformer Architecture]] — the context window is the model's entire "view" of the world for a single call; attention operates over the full context; longer contexts increase compute quadratically
- [[Working Memory]] — the context window is a functional analogue to Baddeley's working memory; context engineering is what the central executive does for AI systems
- [[Retrieval-Augmented Generation]] — RAG is context engineering applied to the information retrieval problem; what gets retrieved determines what context the model has
