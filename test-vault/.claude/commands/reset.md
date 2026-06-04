---
description: Reset the test vault to a clean baseline state
allowed-tools: ["Bash(find:*)", "Bash(rm:*)"]
---

Reset the test vault to a clean baseline state by clearing all test-generated content and restoring session state.

**Step 1 — Show the plan and ask for confirmation:**

Print exactly:
```
📋 Plan: Reset vault to baseline
  Delete: wiki/concepts/ (all files)
  Delete: wiki/log.md
  Delete: archive/ (all subdirs and files)
  Delete: resources/media/ (all files)
  Rewrite: context.md → session-count: 0, inbox-count: 0, focus restored

  Preserves: .claude/, .obsidian/, .fixtures/, dev/, inbox/, CLAUDE.md

Proceed? (y/n)
```

Wait for the user's response. If anything other than `y`, print "Reset cancelled." and stop.

**Step 2 — Execute on confirmation:**

1. Delete all files inside `${VAULT_PATH:-$(pwd)}/wiki/concepts/` (leave the directory itself)
2. Delete `${VAULT_PATH:-$(pwd)}/wiki/log.md` if it exists
3. Delete all files inside every subdirectory of `${VAULT_PATH:-$(pwd)}/archive/` (leave the subdirs)
4. Delete all files inside `${VAULT_PATH:-$(pwd)}/resources/media/` (leave the directory)
5. Rewrite `${VAULT_PATH:-$(pwd)}/context.md` with baseline values — set `last-session` to today's date in `YYYY-MM-DD` format, all other fields as shown:

```
---
last-session: "YYYY-MM-DD"
current-focus: Testing ingest pipeline
active-projects: []
next-actions: []
parked-ideas: []
inbox-count: 0
session-count: 0
last-nudged-command: ""
---
```

**Step 3 — Report:**

Print a one-line summary of what was cleared, e.g.:
```
✓ Reset complete. Cleared N wiki pages, N archive files, N media files. context.md restored.
```
