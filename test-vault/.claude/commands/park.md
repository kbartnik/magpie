---
description: Shelve an idea and return to current focus
argument-hint: "idea to park"
allowed-tools: ["Bash(grep:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Read context.md — get `current-focus`.
2. Evaluate $ARGUMENTS: is this a genuine question worth a wiki/questions/ stub?
3. If yes: present plan, create stub at `wiki/questions/<slug>.md` with minimal frontmatter and one-line description, add to wiki/index.md under ## Open Questions.
4. Append to context.md `parked-ideas`: "[YYYY-MM-DD] $ARGUMENTS"
5. Report:
   ```
   Parked: "<$ARGUMENTS>"
   Back to: <current-focus>
   ```
6. Stop. No elaboration on the parked idea.
