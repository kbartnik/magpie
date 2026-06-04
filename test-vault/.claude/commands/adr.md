---
description: Create or update Architecture Decision Records
argument-hint: "[number | text | empty for new]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/adr/PLUGIN.md` and `.claude/plugins/adr/skill.md` before proceeding.

Route by argument:
- **No argument**: Start a new ADR. Elicit one question at a time (see skill). Draft. Present for review. Write after approval.
- **Numeric** (e.g. "7" or "0007"): Open ADR-0007 for update. Show current content. Ask what to change. Follow zone rules — accepted ADRs are content-immutable.
- **Text**: Search `dev/adr/` for matching title: `grep -ril "<text>" dev/adr/`

After writing:
1. Update `wiki/index.md` under `## Dev`
2. Run: `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append adr "<title>" "<files>" "<summary>" "<open>"`
