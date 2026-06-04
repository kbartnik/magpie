---
description: Write a blameless post-mortem
argument-hint: "[incident description]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/debrief/PLUGIN.md` and `.claude/plugins/debrief/skill.md` before proceeding.

Elicit one question at a time (see skill):
1. What happened?
2. Timeline?
3. What worked?
4. What didn't work?
5. Root cause?

Draft the debrief. Present for review. Write after approval.

`"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append debrief "<title>" "<files>" "<summary>" "<open>"`
