---
description: Synthesize a learning track into a wiki synthesis page
argument-hint: "[topic-slug]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/learning/PLUGIN.md` and `.claude/plugins/learning/skill.md` before proceeding.

1. Find the track at `dev/learning/<topic-slug>/`
2. Read all files: `index.md` and all `YYYY-MM-DD.md` session notes
3. Synthesize:
   - What I now know (confident understanding)
   - How key concepts relate to each other
   - What's still unclear or unresolved
4. Offer: "Want me to save this as `wiki/syntheses/<topic>.md`?"
5. If yes: present draft, wait for approval, write it, update wiki/index.md under ## Syntheses
6. `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append learn "<topic> review" "<files>" "<synthesis summary>" "<remaining open questions>"`
