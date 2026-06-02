---
description: Open or update a project
argument-hint: "[slug | new <name> | log <slug> | item <slug>]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

Read `.claude/plugins/projects/PLUGIN.md` and `.claude/plugins/projects/skill.md` before proceeding.

**Route by argument:**

- **No argument**: List all projects with `ls dev/projects/` and show their status.
- **`new <name>`**: Create a new project. Elicit repo, tech stack, goals. Draft `dev/projects/<slug>/index.md` and `dev/projects/<slug>/backlog.base`. Present for review. Write after approval. Add to `wiki/index.md` under `## Dev`.
- **`log <slug>`**: Add a session log. Read the project index first. Draft `dev/projects/<slug>/YYYY-MM-DD.md`. Present for review. Write after approval. Append link in index `## Log` section. Update `updated:` field.
- **`item <slug>`**: Add a backlog item. Read the project index first. Elicit title, description, priority, optional due date. Draft `dev/projects/<slug>/YYYY-MM-DD-<item-slug>.md`. Present for review. Write after approval.
- **`<slug>` (bare slug)**: Open and display the project index for `dev/projects/<slug>/index.md`.

After writing or updating:
`"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append project "<project title>" "<files changed>" "<summary>" "<open items>"`
