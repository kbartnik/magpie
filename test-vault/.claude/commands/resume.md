---
description: Restore session context mid-session
argument-hint: ""
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Run `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" session-start` and display the output.
2. Read last 5 log entries in full; surface the three most recent operations with their Open: fields.
3. Read any dev/ files referenced in `current-focus` or `next-actions`.
4. Display:
   ```
   SESSION RESUME

   <vault-tools session-start output>

   Recent activity:
   - <last 3 log entries with Open: fields>

   Parked: <parked-ideas list or "none">
   ```
5. State directly: "Ready to: <next-actions[0]>"
6. Do not ask a follow-up question.
