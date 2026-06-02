---
title: "Compiling knowledge, not retrieving it: a hands-on deep dive into llm-wiki-compiler"
author:
  - "Roan Brasil Monteiro"
url: "https://medium.com/@roanmonteiro/compiling-knowledge-not-retrieving-it-a-hands-on-deep-dive-into-llm-wiki-compiler-a495523a8085"
description: "Compiling knowledge, not retrieving it: a hands-on deep dive into llm-wiki-compiler The thesis of this piece is simple and uncomfortable: the problem of making an LLM “remember” what you’ve …"
tags:
  - "clippings"
  - "medium"
published: 2026-05-28T05:16:24-04:00
created: 2026-05-31T10:36:35-04:00
---
# Compiling knowledge, not retrieving it: a hands-on deep dive into llm-wiki-compiler

By Roan Brasil Monteiro
Published: 2026-05-28T09:16:24Z

>[!summary]
>Compiling knowledge, not retrieving it: a hands-on deep dive into llm-wiki-compiler The thesis of this piece is simple and uncomfortable: the problem of making an LLM “remember” what you’ve …

> The thesis of this piece is simple and uncomfortable: the problem of making an LLM “remember” what you’ve read isn’t solved with more retrieval. It’s solved with compilation. The `llm-wiki-compiler` is the first implementation I've seen that treats this as a genuine compiler problem — two-phase pipeline, incremental change detection by hash, claim-level provenance, and a review queue — instead of yet another RAG layer with a prettier name. This is a hands-on deep dive into that tool: what it does, how it does it, where it breaks, and why it matters to anyone already treating Obsidian as a persistent memory layer for agents.

🐦 **X/Twitter:** [@roanbrasil](https://x.com/roanbrasil)

💼 **LinkedIn:** [in/roanbrasil](https://www.linkedin.com/in/roanbrasil/)

📰 **Substack:** [roanbrasil.substack.com](https://roanbrasil.substack.com/)

## The problem nobody wants to admit

Everyone who builds a serious “chat with your documents” workflow goes through the same slow-motion disappointment.

It starts magical. You drop twenty PDFs into NotebookLM, or you wire up a homegrown RAG setup with embeddings and a vector store, and you ask things. The answers come. You get excited.

Then you ask the real question — the one that requires cross-referencing five sources, noticing that the March document contradicts the January one, and synthesizing a conclusion that none of the sources states explicitly. And the system… starts over from scratch. It searches chunks, finds plausible fragments, and stitches an answer together on the spot. On the next similar question, it does it all again. Nothing accumulated. None of the connections it “discovered” in the previous answer were recorded anywhere.

Andrej Karpathy described this malaise with precision in a gist published in April 2026, the [“LLM Wiki”](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f). His diagnosis is that the dominant LLM-with-documents experience — RAG, ChatGPT file uploads, NotebookLM — makes the model *rediscover knowledge from scratch on every question*. There’s no accumulation. The cross-references don’t exist until the moment you ask, and they vanish as soon as the answer ends. Knowledge is re-derived, never *maintained*.

His proposal is a paradigm inversion: instead of retrieving from raw documents at query time, the LLM **incrementally builds and maintains a persistent wiki** — a structured, interlinked collection of markdown files that sits *between* you and the raw sources. When a new source arrives, the model doesn’t index it for later. It reads it, extracts what matters, integrates it into the existing wiki, updates entity pages, revises syntheses, notes where the new data contradicts the old. Knowledge is compiled once and kept current.

Karpathy’s key phrase is that the wiki is “a persistent, compounding artifact” — the cross-references are already there, the contradictions have already been flagged, the synthesis already reflects everything you’ve read. And the division of labor is clear: you handle curating sources, exploring, and asking good questions; the LLM does all the tedious work — summarizing, cross-referencing, filing, bookkeeping. In the metaphor he uses, *Obsidian is the IDE, the LLM is the programmer, the wiki is the codebase*.

Karpathy’s gist, however, is deliberately abstract. He says so himself: it’s an idea file, meant to be pasted into your agent so you can build the specifics together. There’s no code. No CLI. It’s a *pattern*, not a tool.

That’s where `llm-wiki-compiler` comes in.

## What llm-wiki-compiler is (and what it isn’t)

The `llm-wiki-compiler`, published under the `atomicstrata` account, is a TypeScript CLI that takes Karpathy's abstract pattern and turns it into an executable tool. The repo's tagline says it well: *"Raw sources in, interlinked wiki out."*

The vocabulary difference here isn’t cosmetic, and it’s where the thesis lives. Karpathy talks about “building and maintaining” a wiki. The `llm-wiki-compiler` talks about **compiling**. The word choice is intentional — and the README is explicit about positioning the tool against RAG in a framing worth quoting for its clarity:

```c
RAG:     query → search chunks → answer → forget
llmwiki: sources → compile → wiki → query → save → richer wiki → better answers
```

The honest read is the one the project itself makes: this **doesn’t replace RAG, it complements it**. RAG remains great for ad-hoc retrieval over huge corpora. The `llm-wiki-compiler` gives you a persistent, structured artifact *to retrieve from*. One solves "find the relevant chunk now." The other solves "maintain a body of knowledge that improves itself."

The minimal usage fits in three lines:

```c
npm install -g llm-wiki-compiler
export ANTHROPIC_API_KEY=sk-...
```
```c
llmwiki ingest https://some-article.com
llmwiki compile
llmwiki query "what is X?"
```

And it’s deliberately provider-agnostic. The default is Anthropic, but it reads from OpenAI, from local OpenAI-compatible servers (like `llama-server`), and from Ollama — which matters a lot for anyone running local inference. For an Ollama setup, for example:

```c
export LLMWIKI_PROVIDER=ollama
export LLMWIKI_MODEL=llama3.1
export LLMWIKI_EMBEDDING_MODEL=nomic-embed-text
export OLLAMA_HOST=http://ollama_host:11434/v1
```

This detail isn’t trivial. It means the entire compile pipeline can run without a single call to an external API — something that matters again later when I get to the Personal Harness.

## How it compiles: the two-phase pipeline

Here’s the heart of the thing, and the point where the “compiler” metaphor stops being marketing and becomes architecture.

At a very high level, the flow is:

```c
sources/  →  SHA-256 hash check  →  concept extraction (LLM)  →  page generation  →  wikilink resolution  →  index.md
```

The design decision that most deserves attention is what the README calls the **two-phase pipeline**. Phase 1 extracts *all* concepts from *all* sources. Only afterward, in Phase 2, are the pages generated. It looks like an implementation detail, but it has three practical consequences that change the quality of the output:

The first is that it **eliminates order-dependence**. In a naive approach, where you process source by source and write pages as you go, the source that arrives first defines the vocabulary, and the following ones contort themselves to fit. Separating extraction from generation breaks that.

The second is that it **catches failures before writing anything**. If extraction of one source fails, you find out before you’ve dirtied the wiki with half-written pages.

The third, and most important for Karpathy’s problem, is that it **merges concepts shared across sources into a single page**. If five different articles talk about “knowledge compilation,” you don’t end up with five near-identical pages — you end up with one page that cites all five sources.

The second essential property is that compilation is **incremental**. Each source has its SHA-256 hash computed and recorded in state. On a subsequent `compile`, only the sources that changed go through the LLM; the rest are skipped. For anyone paying per token (or waiting on local inference), this is the difference between a tool usable day-to-day and an expensive toy.

And the third property closes Karpathy’s loop: **compounding queries**. When you run `llmwiki query "question" --save`, the answer is written as a new wiki page, the index is rebuilt immediately, and that answer then shows up as context in future queries. Your exploration doesn't evaporate into chat history — it becomes a source. This is exactly the "file answers back into the wiki" that the gist describes as the central insight, but here it's implemented as a flag.

## What it actually produces

A deep dive without seeing the output is just theory. So let’s look at what actually comes out the other end.

A raw source — say, the Wikipedia article on *knowledge compilation* — becomes a structured page with YAML frontmatter. Per the repository’s own example, something like this:

```c
---
title: Knowledge Compilation
summary: Techniques for converting knowledge representations into forms that support efficient reasoning.
kind: concept
sources:
  - knowledge-compilation.md
createdAt: "2026-04-05T12:00:00Z"
updatedAt: "2026-04-05T12:00:00Z"
---
Knowledge compilation refers to a family of techniques for pre-processing
a knowledge base into a target language that supports efficient queries.
Related concepts: Propositional Logic, Model Counting
```

Notice three things. The frontmatter carries source attribution (`sources:`). The body ends with Obsidian-style `wikilinks` — which resolve to concept titles. And the `kind: concept` points to a schema layer I'll detail shortly.

The on-disk output structure is deliberately flat and legible:

```c
wiki/
  concepts/         one .md per concept, with YAML frontmatter
  queries/          saved answers, included in index and retrieval
  index.md          auto-generated table of contents
.llmwiki/
  schema.json       optional page-kind and cross-link policy
  candidates/       candidates pending review
```

The fact that the output is **Obsidian-compatible markdown** — real `wikilinks`, frontmatter that Dataview can read — is no coincidence. It's the same bet Karpathy makes: the wiki is just a git repo of markdown files, so you inherit version history, branching, and collaboration for free, and you can open it all in Obsidian to navigate the graph.

## Provenance: the part that separates a wiki from organized hallucination

If there’s one feature I’d single out as the most mature in the project, it’s **claim-level provenance**. It’s also the feature that most answers the obvious objection any serious person raises about this kind of system: *“how do I trust the page didn’t make this up?”*

The system has two levels of traceability.

At the paragraph level, each paragraph can carry a marker pointing back to the source file that contributed that content:

```c
This paragraph is grounded in the source. ^[source.md]
```

At the claim level, for assertions that need tighter verification, a page can pin a statement to a *specific line range* of the ingested source:

```c
The system uses a two-phase compile pipeline. ^[architecture-notes.md:42-58]
```

And `llmwiki lint` validates both forms — it reports missing source files, malformed citations, impossible ranges (line `0`, or `8-3`), and ranges that run past the end of the file. This is the kind of rigor that usually only shows up in academic tooling, not in "PKM with AI" projects.

On top of that, pages can carry **epistemic metadata** in the frontmatter, all optional:

```c
---
title: Knowledge Compilation
confidence: 0.82           # 0–1, LLM self-reported confidence in the synthesis
provenanceState: merged    # extracted | merged | inferred | ambiguous
contradictedBy:
  - slug: probabilistic-reasoning
---
```

When multiple sources merge into one slug, the metadata is reconciled conservatively: confidence becomes the `min`, `provenanceState` becomes `merged`, and the `contradictedBy` entries are unioned. The linter gains three corresponding rules — `low-confidence` (flags pages below a confidence threshold), `contradicted-page` (flags pages with recorded contradictions), and `excess-inferred-paragraphs` (flags pages with too much uncited prose).

That last one is subtle and clever: it counts uncited paragraphs *directly from the rendered text*, not from a frontmatter field. The body is the single source of truth. You can’t lie to the linter by editing metadata — it reads what’s actually written.

## The review queue: human in the loop, without friction

By default, `compile` writes pages straight into `wiki/`. But there's a `--review` mode that instead writes candidates to `.llmwiki/candidates/` for you to inspect before anything lands in the wiki:

```c
llmwiki compile --review     # produces candidates, leaves wiki/ untouched
llmwiki review list          # see what's pending
llmwiki review show <id>     # inspect a candidate
llmwiki review approve <id>  # promote to wiki/ + refresh index/MOC/embeddings
llmwiki review reject <id>   # archive to candidates/archive/
```

The concurrency details here reveal that someone thought about the real use case. `approve` and `reject` acquire a lock (`.llmwiki/lock`) so they serialize cleanly against each other and against any concurrent `compile`. Source state is deferred per source — when one source produces multiple candidates, it isn't marked compiled until the last candidate is approved, so unresolved siblings stay re-detectable on the next `compile --review`.

This maps directly to the workflow Karpathy describes preferring: ingest one source at a time, read the summaries, check the updates, guide what to emphasize. The difference is that here “checking” is a command with a lock, not a discipline you have to remember to maintain.

## The schema layer: when “concept” isn’t enough

For wikis that grow beyond flat concept pages, there’s an optional `.llmwiki/schema.json`. Existing projects don't need it — a missing or invalid `kind` falls back to `concept`. But when you define it, it supports four page kinds:

- `concept` — a standalone idea or pattern
- `entity` — a specific person, product, organization, or named artifact
- `comparison` — side-by-side analysis across concepts or entities
- `overview` — a map page that connects several concepts in a domain

Schema rules can set per-kind `minWikilinks` and optional `seedPages`. `compile` can materialize seed pages (like overviews), `lint` enforces per-kind cross-link minimums, and review candidates surface schema violations before approval.

This is exactly the kind of feature that answers one of the structural failures I’ll discuss below — the “level” problem, where macro themes and tactical findings end up flattened onto the same plane. Typing the pages is a first step toward returning hierarchy to the graph.

## The MCP server: the wiki as an agent tool

For me, this is the feature that connects `llm-wiki-compiler` to the topic that interests me most. The project ships an MCP (Model Context Protocol) server that exposes the entire pipeline to agents — Claude Desktop, Cursor, Claude Code:

```c
llmwiki serve --root /path/to/your/wiki-project
```

And the client registration is straightforward:

```c
{
  "mcpServers": {
    "llmwiki": {
      "command": "npx",
      "args": ["llm-wiki-compiler", "serve", "--root", "/path/to/wiki-project"],
      "env": { "ANTHROPIC_API_KEY": "sk-ant-..." }
    }
  }
}
```

The exposed tools include `ingest_source`, `compile_wiki`, `query_wiki`, `search_pages`, `read_page`, `lint_wiki`, and `wiki_status`. There's a nicely considered security detail: the tools that need an LLM check for credentials on every call, while the read-only ones (`read_page`, `lint_wiki`, `wiki_status`) and `ingest_source` work without any credentials.

The README draws a distinction worth highlighting, because it’s honest about positioning. Compared to `llm-wiki-kit`, which gives agents raw CRUD against pages, the `llm-wiki-compiler` exposes the **automated pipelines** — intelligent compilation, incremental change detection, and semantic query routing built in. One gives you bricks; the other gives you the factory.

## Connecting to the Personal Harness

This is where, for me, the whole thing clicks together. I’ve been describing what I call the **Personal Harness** for a while now — an architecture where Obsidian works as a persistent memory layer for AI agents, with five layers: Memory (Obsidian), Reference, Local Compute (Ollama), API+MCP, and Showcase (GitHub).

The `llm-wiki-compiler` isn't the Personal Harness, but it slots into nearly all of its layers at once, and that's why it caught my attention so strongly:

The **Memory** layer is literally the tool’s output — Obsidian-compatible markdown, with resolving `wikilinks`, frontmatter Dataview reads, a navigable graph. The wiki *is* the memory, and it lives in your vault.

The **Local Compute** layer is covered by native support for Ollama and OpenAI-compatible servers. The entire compilation can run without leaving your machine.

The **API+MCP** layer is the MCP server, which turns the wiki into a tool that named agents can drive.

The philosophical difference that matters is this: most attempts to give an agent “memory” treat memory as a side effect of retrieval — you store embeddings and hope the search brings back the right chunk. The `llm-wiki-compiler` treats memory as a *compiled, auditable artifact* that you can open, read, criticize, and version. For anyone building a serious personal harness, that's the property you want. Memory you can't inspect isn't memory; it's a black box that got lucky.

## Where this fails (the honest part)

No analysis of mine is complete without this, and in this case the material is abundant — partly because the project itself is transparent about its limits, partly because the community that formed around Karpathy’s pattern has already documented the structural failures.

**The project itself admits it’s early software.** The README is direct: best for small, high-signal corpora (a few dozen sources). Query routing is index-based. Don’t expect this to scale to tens of thousands of documents without pain — and the tool doesn’t promise that.

**The identity problem.** This is the most documented failure of the pattern, and it comes from an unexpected place: the author of `nohmitaina` (a competing desktop-editor implementation) recorded, in the thread on Karpathy's gist, that after a month feeding the system his own notes, the same concept was being extracted under slightly different names from related sources. The result is duplicate pages — in his example, "Cognitive Dissonance Marketing" and "Cognitive Dissonance and Urgency" came out of the same book. The `llm-wiki-compiler` attacks this partially with the two-phase pipeline's concept merging and hash deduplication, but hashing resolves identical sources, not synonymous concepts. This is an open frontier.

**The level problem.** The same account points out that life-scale themes (“Personal AGI”) end up at the same level as tactical findings (“Urgency Trigger”). When everything is flat, importance disappears. The `llm-wiki-compiler` 's schema layer (with its `overview`, `concept`, `entity` kinds) is a partial answer, but it's an answer *you* have to configure — it doesn't emerge on its own from the material.

**The relationship problem.** Concepts get linked as “related,” but the *type* of the relationship is lost. “Similar to,” “contains,” “contradicts” — all collapse into one word, which makes the graph useful for navigating but poor for *thinking*. The `contradictedBy` in the frontmatter is the only typed edge the project has today. It's a start, not a solution.

**Self-reported confidence is exactly that.** The `confidence: 0.82` field is the confidence the LLM *says* it has in the synthesis. Models are notoriously poorly calibrated at this. I'd treat that number as a loose signal for prioritizing review, never as a quality guarantee. The fact that the linter counts uncited paragraphs straight from the text is, ironically, a far more reliable quality signal than the confidence field.

**Truncation.** The project is honest about this — sources that exceed the character limit are truncated on ingest, with `truncated: true` and the original count recorded in frontmatter. It's transparent, but it means that for long sources you're working with partial content, and you need to know that.

**Token cost is real and hidden.** A single ingest can touch many pages. The per-concept budget (`LLMWIKI_PROMPT_BUDGET_CHARS`, default ~200k characters / ~50k tokens) exists precisely because popular concepts shared across many sources can blow past the context window. Incremental compilation helps enormously, but the first compile of a large corpus is not cheap.

## Who this makes sense for — and who it doesn’t

After all this analysis, the honest recommendation is narrow, and that’s a virtue.

It makes sense if you have a small, high-signal corpus — a few dozen sources you want to turn into a living body of knowledge: a thesis literature review, a months-long deep dive into a topic, the notes from a dense book you want to dissect. It makes a lot of sense if you already live in Obsidian and want an agent to do the bookkeeping you abandon. And it makes total sense if you’re building agent memory and want it to be *auditable* rather than a black box of embeddings.

It doesn’t make sense if your problem is ad-hoc retrieval over a huge, constantly changing corpus — that’s RAG’s job, and the project says so. It doesn’t make sense if you need scale of tens of thousands of documents today. And it doesn’t make sense if you’re not willing to stay in the loop — the review queue and the linter are where the real value shows up, and both presuppose a human curating.

The `llm-wiki-compiler` is, in my reading, the cleanest proof of concept that *knowledge compilation* is executable and not merely an elegant metaphor of Karpathy's. It doesn't solve the hard problems of identity, level, and relationship — nobody has yet, and the gist thread is a living graveyard of implementations trying. But it gets the problem that matters right: the wiki is a persistent, auditable, compounding artifact, and the maintenance that would kill a human wiki costs almost nothing when it's an LLM disciplined by a schema doing the bookkeeping.

The question worth taking away from here isn’t “should I use this tool?” It’s the deeper one Karpathy planted: why on earth did we accept, for so long, that our knowledge should be rediscovered from scratch on every question — instead of compiled once and kept alive?

*Primary sources: Andrej Karpathy’s* [*LLM Wiki*](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f) *gist (April 2026) and the thread of implementations that formed around it; the* `*atomicstrata/llm-wiki-compiler*` *repository (README and CLAUDE.md, v0.6.0). The accounts of the identity, level, and relationship problems come from the public comment by the author of the* `*nohmitaina*` *project on the gist thread.*

## Deep Read

**Key Insight:** The two-phase pipeline (extract all concepts from all sources first, then generate pages) isn’t just an implementation detail — it eliminates order-dependence (first source no longer defines vocabulary), catches extraction failures before dirtying the wiki, and merges synonymous concepts across sources into single pages. This is why compilation outperforms iterative per-source processing.

**What Surprised Me:** `llmwiki lint` validates claim-level citations by reading paragraph text directly from the rendered body — not from frontmatter metadata. You cannot lie to the linter by editing metadata. The body is the single source of truth. Unusually rigorous for a PKM tool.

**Open Questions:**
- How does the review queue’s locking model (source not marked compiled until last candidate approved) compare to vault-tools’ archive-file approach? Is per-source locking a missing feature?
- Does the magpie project need claim-level provenance (source:line-range citations) or is source-level attribution sufficient for this vault’s use case?
- The three unsolved problems (identity, level, relationship) — does this vault’s wiki already have these failure modes? Worth checking for near-duplicate concept pages and macro/tactical concept flattening.

**Wikilink Candidates:**
- LLM Wiki Pattern — new page; this article is one of three primary sources for the concept
- [[Harness Engineering]] — the memory layer is precisely what knowledge compilation addresses

**Connections:**
- LLM Wiki Pattern — this article provides the deepest implementation detail for the pattern (two-phase pipeline, provenance, review queue)
- [[Harness Engineering]] — "memory you can’t inspect isn’t memory; it’s a black box that got lucky" maps directly to the harness memory dimension
- [[magpie]] — review queue, locking model, and claim provenance are directly relevant to magpie’s design decisions
- Nexus Vault Template — this vault is an informal implementation of the same pattern; MCP server integration worth evaluating for magpie