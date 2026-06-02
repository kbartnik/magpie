---
title: "Vault Security Gate Exit Code"
question: "Do the vault's security-relevant hooks use exit 2 (block operation) where they should, or are any silently using exit 1 (warn only), creating false security?"
type: question
status: open
domain: claude-code
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-31-hooks-production-patterns.md"
related:
  - "[[Claude Code Hooks]]"
tags:
  - claude-code
  - security
  - vault-meta
---

# Vault Security Gate Exit Code

*Do the vault's security-relevant hooks use exit 2 (block operation) where they should, or are any silently using exit 1 (warn only), creating false security?*

The hooks production patterns article establishes a two-level governance model: exit 2 blocks the tool call; exit 1 signals a warning but allows the operation to proceed. For security gates, exit 2 is the correct choice — a hook that warns but doesn't block is not a security control.

The vault's hooks (SessionStart, PostToolUse audit, Stop) haven't been audited for correct exit code usage relative to their security intent. If any hook is classified as a security control but exits with 1 on a violation, it's logging the violation while allowing the operation — exactly what a security gate should not do.

This is a straightforward empirical question resolvable by reading `.claude/settings.json` and checking the hook scripts for their exit code logic.

## See Also

- [[Claude Code Hooks]] — exit code governance; the 3-tier ladder (automatic/confirm/plan-then-review)
