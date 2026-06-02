---
description: Add a session note to a learning track
argument-hint: "[topic-slug]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/learning/PLUGIN.md` and `.claude/plugins/learning/skill.md` before proceeding.

**Route by argument:**
- **Argument is a topic slug**: Find or create the track at `dev/learning/<topic-slug>/`. Read `index.md`. Add a session note for today. Elicit: source, key insight, what surprised, connections, open questions. Update track index after writing note.
- **No argument**: List active tracks:
  ```bash
  find dev/learning -name "index.md" | xargs grep -l "status: active" 2>/dev/null
  ```
  Ask which to continue.

After writing:
`"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append learn "<topic>" "<files>" "<summary>" "<open questions>"`
