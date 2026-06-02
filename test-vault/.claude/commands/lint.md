---
description: Health-check the vault — gather findings, reason over them, propose fixes
argument-hint: ""
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Run `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" lint-gather` — capture full output.
2. Parse by section (## BROKEN_WIKILINKS, ## ORPHAN_PAGES, etc.).
3. Read wiki pages referenced in BROKEN_WIKILINKS and ORPHAN_PAGES to understand context.
4. Read log entries referenced in OPEN_LOG_ITEMS.
5. Present findings as a numbered list: Critical → Important → Suggestions. Include file paths.
6. Present OPEN_LOG_ITEMS as a separate "Unresolved from log" section.
7. Ask: "Which issues should I address? (all critical / item numbers / skip)"
8. For approved fixes: present a plan per fix, execute after approval.
9. `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append lint "Lint pass" "<N issues>" "<summary>" "<remaining open>"`
