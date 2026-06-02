---
title: "rtk"
type: tool
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-2026-05-30-rtk-repo.md"
  - "archive/clippings/2026-05-29-rust-token-killer.md"
related:
  - "Claude Code Hooks"
tags:
  - tool
  - claude-code
  - token-optimization
  - rust
---

# rtk

**rtk** is a Rust CLI proxy that intercepts Bash tool calls and rewrites common commands to filtered, token-efficient versions — reducing LLM context consumption by 60–90% per command.

- GitHub: [rtk-ai/rtk](https://github.com/rtk-ai/rtk)
- Install: `brew install rtk`

## How It Works

rtk installs as a `PreToolUse` hook in Claude Code's `settings.json`. When the agent issues a Bash call like `git status`, the hook rewrites it to `rtk git status` before execution. The agent receives compact output; the rewrite is transparent to the model.

```
Without rtk:  Claude → git status → ~2,000 tokens raw
With rtk:     Claude → rtk git status → ~200 tokens filtered
```

Four strategies applied per command type:
1. **Smart Filtering** — removes noise (comments, whitespace, boilerplate)
2. **Grouping** — aggregates similar items (files by directory, errors by type)
3. **Truncation** — keeps relevant context, cuts redundancy
4. **Deduplication** — collapses repeated log lines with counts

## Token Savings (30-min session estimates)

| Command | Savings |
|---------|---------|
| `ls` / `tree` | −80% |
| `cat` / read | −70% |
| `git diff` | −75% |
| `cargo test` / `go test` | −90% |

Estimated total across a session: ~80% reduction.

## Important Scope Limitation

**rtk only intercepts Bash tool calls.** Claude Code's built-in `Read`, `Grep`, and `Glob` tools bypass the hook entirely — they don't flow through Bash. For filtered reads of those workflows, use shell equivalents (`cat`/`rg`/`find`) or explicit `rtk read`/`rtk grep`/`rtk find` calls.

## Failure Recovery

When a filtered command fails, rtk saves the full unfiltered output and reports the path:

```
FAILED: 2/15 tests
[full output: ~/.local/share/rtk/tee/1707753600_cargo_test.log]
```

The LLM can read the full output without re-executing.

## Multi-Tool Support

Supports 13 AI coding tools. For Claude Code, integration is via `PreToolUse` hook in `settings.json`. Other tools use their own hook mechanisms.

```bash
rtk init -g          # Claude Code / GitHub Copilot
rtk init -g --gemini # Gemini CLI
```

## See Also

- [[Claude Code Hooks]] — rtk is implemented as a PreToolUse hook; hook architecture explains why and how it works
