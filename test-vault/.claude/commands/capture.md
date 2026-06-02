---
description: Zero-friction inbox dump — Go binary only, no LLM
argument-hint: "text to capture"
allowed-tools: ["Bash(bash:*)"]
---

Run `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" capture "$ARGUMENTS"` and output the one-line result. Stop. No analysis, no follow-up.
