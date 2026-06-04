---
description: Open or update an investigation
argument-hint: "[question | existing slug fragment]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/investigation/PLUGIN.md` and `.claude/plugins/investigation/skill.md` before proceeding.

**Route by argument:**
- **No argument or new question**: Create a new investigation. Draft the file. Present for review. Write after approval. Add to wiki/index.md under ## Dev.
- **Argument matches existing investigation** (slug fragment or question keywords): Find the file with `grep -ril "<text>" dev/investigations/`. Open it, show current state, ask what to add (new evidence, updated confidence, conclusion).

After writing or updating:
`"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append investigate "<question title>" "<files>" "<summary>" "<open threads>"`
