---
title: "AI-Checking-AI Stop Hooks"
question: "Should the vault add a Stop hook that invokes a second AI model to verify the agent's final output before accepting it — and does the quality benefit justify the cost?"
type: question
status: open
domain: claude-code
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-hooks-production-patterns.md"
related:
  - "[[Claude Code Hooks]]"
  - "[[Harness Engineering]]"
  - "[[Vibe-Coding Anti-Pattern]]"
tags:
  - claude-code
  - vault-meta
---

# AI-Checking-AI Stop Hooks

*Should the vault add a Stop hook that invokes a second AI model to verify the agent's final output before accepting it — and does the quality benefit justify the cost?*

The hooks production patterns article describes a pattern: a Stop hook invokes a different AI model to verify the agent's proposed final output before the session ends. For this vault: ingest operations modify wiki pages that the agent reads in future sessions, so incorrect wiki modifications could propagate across sessions.

The case for adding it: ingest is the highest-stakes write operation in the vault; a second-model check could catch factual errors, structural issues, or pages that contradict existing knowledge before they're committed. The case against: additional API calls and latency on every ingest, added complexity, and uncertain quality benefit since the second model may have the same blind spots as the first.

A targeted implementation might run only on wiki page writes (not log appends or index updates), using a lightweight model for the check. Whether this is worth building depends on whether the vault has actually accumulated problematic ingest errors — an empirical question.

## See Also

- [[Claude Code Hooks]] — Stop hook pattern; exit code governance
- [[Harness Engineering]] — AI-checking-AI as a harness verification dimension
