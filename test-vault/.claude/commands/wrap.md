---
description: End the session — update context, log summary
argument-hint: ""
allowed-tools: ["Bash(bash:*)"]
---

Load `.claude/skills/session-context/SKILL.md` before proceeding.

1. Run `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" session-end` (writes last-session and inbox-count).
1b. For any learning track session notes written today, append:
    ```
    *ended HH:MM*

    ## Where I Left Off
    [one sentence: what was actively in progress]
    ```
2. Ask: "One-line summary of what we accomplished?" Wait.
3. Ask: "Next actions? Up to three, domain-tagged. Example: [wiki] ingest the RAG paper, [adr] finish ADR-0002" Wait.
4. Ask: "Anything to park?" Wait (may be "nothing").
4b. Check for active learning tracks that had session notes written this session:
    ```bash
    find dev/learning -name "$(date +%Y-%m-%d).md" 2>/dev/null
    ```
    For each found: ask "Want to extract any concepts from today's [track] notes to wiki pages?" If yes, follow the concept extraction protocol in the learning plugin skill.
5. Update context.md: write `current-focus` (derive from session activity, 1-5 words, action verb first), `next-actions` (list from user's answer), append to `parked-ideas` if any.
6. `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append session-end "<summary>" "context.md updated" "<summary>" "<next-actions as open items>"`
7. Print:
   ```
   Done: <summary>
   Next: <next-actions[0]> [+N more]
   Parked: <parked ideas or "nothing">
   ```
8. Stop. No further output.
