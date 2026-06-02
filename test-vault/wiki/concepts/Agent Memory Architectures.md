---
title: "Agent Memory Architectures"
type: concept
status: active
created: 2026-05-31
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-31-rag-llm-wiki-gbrain-comparison.md"
  - "archive/clippings/2026-06-01-four-types-agent-memory.md"
related:
  - "Harness Engineering"
  - "LLM Wiki Pattern"
  - "Claude Code Memory Architecture"
  - "Claude Code Skills"
tags:
  - agentic-systems
  - harness-engineering
  - knowledge-management
---

# Agent Memory Architectures

Three distinct architectures for how agents remember across sessions, each solving a different version of the same problem. Picking between them is a design decision, not a loyalty test.

> "The context window is not memory. It's a whiteboard that gets erased after every session."

## The Problem

Context windows (even 1M token ones) are ephemeral. Performance degrades starting around 300–400k tokens — roughly 30–40% of the ceiling. When the session ends, everything disappears. RAG was the first serious answer; LLM Wiki and Fat Skills are the next two.

## KOALA Framework: Memory by Cognitive Function

The KOALA framework (Cognitive Architectures for Language Agents, Princeton) classifies agent memory by *what cognitive job it does* — orthogonal to the RAG/LLM Wiki/Fat Skills taxonomy below, which classifies by *implementation mechanism*. Both axes are needed to design a complete memory system.

| KOALA Type | Human analog | Agent implementation | Always in context? |
|-----------|-------------|---------------------|--------------------|
| **Working** | Short-term / active thought | Context window | Yes (it *is* the context) |
| **Semantic** | Factual knowledge | CLAUDE.md, wiki, knowledge base | Yes — loaded at session start |
| **Procedural** | Learned skills | Skill files (skill.md), progressive disclosure | No — index always; full instructions on demand |
| **Episodic** | Personal experience | Distilled session notes, auto-memory | No — retrieved when relevant |

**Working memory** is volatile and size-limited. Performance degrades past ~40% fill (the "dumb zone" — see [[Context Rot]]). Every agent has it; it's just the context window.

**Semantic memory** is persistent factual knowledge that shapes every interaction. The production implementation is usually markdown files, not vector databases. CLAUDE.md is the canonical example. Without it, an agent repeats the same mistakes — it has no persistent knowledge to draw from. Semantic memory is *always present* in context, unlike procedural or episodic.

**Procedural memory** is how-to knowledge. Agent skills implement this via [[Progressive Disclosure Architecture]]: a lightweight index (name + description, ~100 tokens/skill) always in context; full instructions loaded only when a matching task arrives; referenced resources pulled in only during execution.

**Episodic memory** is distilled past experience — not raw transcripts, but compressed notes about what's worth remembering. "Last time we debugged the auth module, the issue was in the middleware layer." The hardest type to get right because of the **forgetting problem**: deciding what to discard when information becomes stale is an unsolved engineering challenge. Humans forget usefully; agents need explicit logic for it.

**Not every agent needs all four:**

| Agent type | Memory types needed |
|-----------|-------------------|
| Reflex agent (thermostat, routing bot) | Working only |
| Simple support bot (password reset) | Working + Procedural |
| Coding agent | All four |

> "Some memory is really what separates a chatbot from an agent." — Martin Keen, IBM Technology

## Architecture 1: RAG — The Retriever

**Formula:** embed → store → retrieve → generate  
**Wins at:** scale, freshness, shipping speed  
**Loses at:** depth, compounding, autonomous action

RAG handles corpus sizes the other architectures can't touch. 200,000 internal documents, indexed the same day. When a document changes, re-embed it. No wiki audit needed.

See [[Retrieval-Augmented Generation]] for the full internal evolution arc: keyword search → semantic search → hybrid → basic RAG → advanced RAG → agentic RAG.

**The three structural failure modes (2024 research, seven total):**
1. **Chunking problem** — a 30-page spec becomes 500-token fragments; the compliance requirement and the reason it exists land in different vectors; the retriever finds one, misses the other
2. **Re-derivation problem** — every query starts from scratch; the agent analyzed the same architecture doc yesterday and drew the same conclusions; it'll do it again tomorrow
3. **Passivity problem** — RAG waits to be asked; it never notices that today's document contradicts last Tuesday's; it never acts on what it knows

**Use RAG when:** corpus is 10,000+ documents, changes frequently, and the priority is shipping a production system with known trade-offs and compliance auditability. RAG has the most battle scars and the richest vendor ecosystem.

## Architecture 2: LLM Wiki — The Compiler

**Formula:** sources → compile → wiki → query → save → richer wiki  
**Wins at:** depth, compounding, synthesis quality  
**Loses at:** scale, freshness, autonomous action

See [[LLM Wiki Pattern]] for the full treatment. The core insight: compile knowledge once, maintain it incrementally. Every future query benefits from prior synthesis. The hundredth query is significantly better than the first — something RAG can never promise.

**Scale ceiling:** works well at ~100 sources / few hundred wiki pages. Navigation degrades past 10,000 sources. At that scale, you add a retrieval layer on top, which starts to look like RAG again.

**Use LLM Wiki when:** sources number in the hundreds, knowledge should compound, the goal is synthesis not just retrieval, and a human will stay in the loop via the review/lint workflow.

## Architecture 3: Fat Skills (GBrain) — The Operator

**Formula:** thin harness + fat skill documents → autonomous scheduled execution  
**Wins at:** autonomy, compounding action, background operation  
**Loses at:** accessibility, setup cost, organizational scale

Introduced by Garry Tan's GBrain project (open source). The inversion: instead of a thick runtime with tool definitions consuming context, keep the harness to ~200 lines and put all intelligence in skill files — fat markdown documents that declare when to fire, what to check, what to write.

**Key properties:**
- **Always-on skills** fire on every message (signal detector, entity linker) — "an unlinked mention is a broken brain"
- **Cron skills** run autonomously on schedule, filing results to `reports/` — agent works while you sleep
- **Deterministic split** — latent work (synthesis, pattern recognition) stays with the LLM; deterministic work (SQL, file ops, calculations) routes to code
- **Fat = fewer, more comprehensive** — Tan's direction: fewer skills with branching parameters, not a library of narrow ones; resolvers stay short

**The resolver insight:** skill descriptions *are* the router. The model reads descriptions and matches intent automatically. No explicit routing code needed.

**Use Fat Skills when:** knowledge needs to trigger autonomous actions, you'll invest in engineering the skill layer, and one power user can define the workflows. Not suitable for broad organizational deployment — the skill-authoring barrier is high.

## Decision Framework

**Start here: what is your agent's job?**

```
Your agent retrieves answers from a large corpus
  → Use RAG
  
Your agent builds expertise that should compound over time
  → Use LLM Wiki (corpus <1k, knowledge should grow)
  
Your agent needs to act on what it knows without being asked
  → Use Fat Skills (accept engineering cost)
```

| | RAG | LLM Wiki | Fat Skills |
|---|---|---|---|
| **Wins at** | Scale (10k+ docs) | Depth, synthesis | Autonomy, action |
| **Loses at** | Depth, compounding | Scale, freshness | Accessibility, setup |
| **Corpus size** | Unlimited | Hundreds | Hundreds–thousands |
| **Freshness** | Re-embed on change | Recompile on change | Re-skill on change |
| **Action** | Passive | Passive | Autonomous |
| **Engineering** | Low | Medium | High |

## Architecture 4: Hebbian Knowledge Graph (Onto AI / "Software Brain")

**Formula:** ingestion (Sensorium) → persistent relational memory with self-learning weights (Hippocampus) → deliberate reasoning (Cortex) → monitoring (Meta-Cognition) → LLM as replaceable language interface  
**Wins at:** continuous learning, auditability, institutional knowledge persistence  
**Loses at:** high-frequency unstructured content, the tacit knowledge that can't be expressed as graph edges

The insight: separate what the brain separates. Don't ask one transformer to do ingestion, memory, reasoning, and monitoring simultaneously. Give each layer the mechanism appropriate for its job.

**Key differentiator:** Hebbian self-learning on the graph. Edge weights strengthen when relationships are repeatedly confirmed, weaken when contradicted. Learning happens in the graph, not in model weights — no retraining, no fine-tuning, no catastrophic forgetting. Every weight change is traceable to the evidence that caused it.

**LLM as interface, not core:** when a better model releases, swap the API endpoint. The institutional knowledge graph persists.

Source: [[archive/clippings/2026-06-01-llm-wrong-problem|TheLLMSeeker (2026)]], [[Neuromorphic and Bio-Inspired AI]]

## Relationship to Autonomous Learning

[[Autonomous Learning Architecture]] (Dupoux/LeCun/Malik 2026) proposes a related but orthogonal decomposition: System A (learning from observation), System B (learning from action), System M (meta-control coordinating both). System M is a more principled formulation of what Architecture 4's Hebbian graph is trying to accomplish — moving adaptation out of frozen model weights and into a persistent, auditable structure. The key difference: Dupoux et al. target *post-deployment learning*, while the Hebbian graph primarily targets *cross-session memory persistence*. Both identify the same structural gap in current systems.

## Convergence

The three-way split is already dissolving. Production systems will combine all three:
- RAG for the retrieval layer (finding relevant content at scale)
- Wiki for the synthesis layer (compiling retrieved content into persistent knowledge)
- Skills for the action layer (operationalizing knowledge into autonomous workflows)

**Claude Code already shows this convergence:** CLAUDE.md functions as a mini-wiki (persistent context that compounds across sessions); auto-memory as compounding knowledge; skills as the action layer. The same pressures produced the same solutions independently.

## Relationship to This Vault

The Nexus vault operates in the LLM Wiki quadrant — small, high-signal corpus, knowledge that compounds, human curation via the ingest/lint workflow. The [[magpie]] project may push toward the Fat Skills direction (cron-based autonomous operations, sonar split).

## See Also

- [[LLM Wiki Pattern]] — full treatment of the compiler architecture
- [[Harness Engineering]] — the memory dimension is one of six harness concerns; this page is the decision framework for that dimension
- [[Claude Code Memory Architecture]] — Claude Code's specific 4-layer memory system
- [[Claude Code Skills]] — the fat skills architecture as implemented in this vault's skill system
- [[Agentic RAG vs Advanced RAG Threshold]], [[Episodic Memory Forgetting Problem]], [[Multimodal Quality Gap for Enterprise]] — open questions on agent memory frontiers
