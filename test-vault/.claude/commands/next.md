---
description: Break a topic into three concrete, time-bounded first steps
argument-hint: "topic or question"
allowed-tools: ["Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Search vault: grep wiki/ and dev/ for terms from $ARGUMENTS.
2. Read relevant open investigations, ADRs, learning tracks, project notes.
3. Identify current state: what's known, decided, or already done on this topic?
4. Output exactly three next actions:
   ```
   NEXT for: <topic>

   1. <Action verb> <specific concrete step> — ~X min
   2. <Action verb> <specific concrete step> — ~X min
   3. <Action verb> <specific concrete step> — ~X min

   Smallest first step: action 1.
   ```
   Rules: action verb first; step 1 completable in under 5 minutes; no vague goals.
5. Ask: "Add action 1 to next-actions in context.md? (y/n)"
6. If yes: prepend `[<domain>] <action 1>` to the `next-actions` list (trim to 3 items).
