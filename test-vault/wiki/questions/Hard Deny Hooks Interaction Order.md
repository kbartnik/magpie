---
title: "Hard Deny Hooks Interaction Order"
question: "When a command appears in both hard_deny and in a hook's trigger list, which takes precedence — and does the evaluation order create security gaps?"
type: question
status: open
domain: claude-code
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-05-20-claude-auto-mode.md"
related:
  - "[[Claude Code Hooks]]"
  - "[[Harness Engineering]]"
tags:
  - claude-code
  - security
---

# Hard Deny Hooks Interaction Order

*When a command appears in both hard_deny and in a hook's trigger list, which takes precedence — and does the evaluation order create security gaps?*

If a command is in `hard_deny` but a hook fires that would execute that command, the evaluation order determines whether the hook can bypass the deny or is itself blocked. Two failure modes: (1) if hooks run before deny evaluation, a hook could bypass a deny restriction; (2) if deny evaluation runs first and blocks the tool call entirely, the hook's security check never runs.

For a vault with a security-gate hook (PreToolUse that validates an operation before allowing it), the second failure mode is the dangerous one: a misconfigured deny list that matches the hook's trigger could silently prevent the security check from running, giving a false impression of protection.

Anthropic's documentation doesn't specify the evaluation order. Empirical testing would resolve this, but the test needs to be carefully designed to distinguish "hook ran but was bypassed" from "hook never ran."

## See Also

- [[Claude Code Hooks]] — lifecycle control points; exit code governance (exit 2 vs exit 1)
