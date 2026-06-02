---
description: Check for topic drift relative to planned focus
argument-hint: ""
allowed-tools: ["Bash(grep:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Read context.md — get `current-focus` and `next-actions`.
2. Read last 10 log entries — identify topics this session has actually touched.
3. Compare planned focus vs. actual activity.
4. Report:
   ```
   FOCUS CHECK

   Planned: <current-focus>
   Next actions: <next-actions list>

   This session touched: <topics from log>

   Assessment: On track | Drifted into <topic>
   ```
5. If drifted, offer three options (no judgment):
   - Park the drift topic with /park and return to current-focus
   - Switch official focus to <drift topic> — update context.md
   - Continue — this drift is deliberate — update next-actions to reflect reality
6. Wait for choice. Execute it.
