---
description: Health-check the vault — gather findings, reason over them, propose fixes
argument-hint: ""
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Run `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" lint-gather` — capture full output.
2. Parse by section (## BROKEN_WIKILINKS, ## ORPHAN_PAGES, etc.).
3. Read wiki pages referenced in BROKEN_WIKILINKS and ORPHAN_PAGES to understand context.
4. Read log entries referenced in OPEN_LOG_ITEMS.
5. Run alias-gap check: `grep -rL "^aliases:" wiki/concepts/ wiki/entities/ wiki/syntheses/ 2>/dev/null` — list any pages missing the `aliases:` frontmatter field.
6. Run orphan check: find wiki pages with no inbound wikilinks by scanning for `[[PageName]]` references across the wiki. A page is orphaned if its title appears in no other page's body or frontmatter.
7. Present findings as a numbered list: Critical → Important → Suggestions. Include file paths. Add alias gaps and orphans as separate sections.
8. Present OPEN_LOG_ITEMS as a separate "Unresolved from log" section.
9. When fixing: respect the causality chain — alias gaps before duplicate checks, merges before dead link repair.
10. Ask: "Which issues should I address? (all critical / item numbers / skip)"
11. For approved fixes: present a plan per fix, execute after approval.
12. `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append lint "Lint pass" "<N issues>" "<summary>" "<remaining open>"`
