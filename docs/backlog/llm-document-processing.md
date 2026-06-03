---
title: "LLM Document Processing Patterns"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: medium
due: ""
created: 2026-06-03
updated: 2026-06-03
tags: [post-1.0, magpie-claude, llm, ingest, context]
---

# LLM Document Processing Patterns

Patterns for magpie-claude (and other LLM plugins) when processing large documents. The core
binary does not call LLMs — these patterns live in plugins that do.

## Iterative Batch Extraction (Truncation-Retry)

When a document exceeds the LLM's context window, naive extraction silently truncates output.
The correct pattern: detect truncation → double the token limit → retry once.

**Detection signal:** `stop_reason: max_tokens` (Anthropic) or `finish_reason: length` (OpenAI).
These indicate the model hit its output limit, not that the task completed.

**Pattern:**
```
extract(doc, max_tokens=8000)
  if stop_reason == max_tokens:
    retry with max_tokens=min(16000, PROVIDER_CAP)
    if still truncated: surface warning to user, return partial result
```

**Cap behavior:** double once, then stop. Do not loop. A document that truncates at 16K tokens
needs chunking, not infinite retries. Partial results are better than silent incompleteness —
always surface a warning when extraction was truncated.

**Implementation note:** extract the truncation-retry logic into a shared pure function
(`truncationRetry(fn, maxTokens, cap)`) rather than duplicating across LLM clients.
The Karpathy LLM Wiki Plugin (obsidian-llm-wiki) extracts this to `src/core/truncation-retry.ts`
with 7 tests — cap behavior, error propagation, warning logging.

## Causality-Ordered Batch Fix

When running maintenance operations (alias completion, duplicate merges, dead link repair),
order matters: you cannot detect duplicates without aliases, and you cannot repair dead links
until merges have settled paths.

Fixed order:
1. Alias completion
2. Duplicate detection and merge
3. Dead link repair
4. Orphan linking
5. Empty page expansion

Each step should use `Promise.allSettled` (or Go's `errgroup`) so a single page failure
doesn't abort the batch — partial results are better than no results.

## Cancellation at Batch Boundaries

Long-running operations (folder ingestion, full lint) should be cancellable at logical batch
boundaries (not mid-page). Use a context cancellation signal checked at each batch iteration.
On cancellation: flush completed work, log partial-result count, exit cleanly.

This is cheaper to design for early than to retrofit. The magpie-claude plugin's ingest loop
should check `ctx.Done()` at each page boundary from the start.

## References

- Karpathy LLM Wiki Plugin: `[[2026-06-03-karpathy-llm-wiki-plugin]]` — reference
  implementation with SSE parser extraction, truncation-retry extraction, cancellation
- [[LLM Wiki Pattern]] — the compile-not-retrieve paradigm these patterns serve
