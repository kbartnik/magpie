---
title: "Additional Context Subcommands"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, context, magpie-claude]
---

# Additional Context Subcommands

Richer context management commands, most meaningful once magpie-claude adds session intelligence.

## Candidates

```
magpie context focus <text>     # set ## Current Focus in context.md
magpie context next <text>      # append to ## Next Actions
magpie context pop              # remove top item from ## Next Actions
magpie context clear-parked     # clear ## Parked Ideas (with confirmation)
```

## Notes

- `context focus` and `context next` are the write counterparts to what `context status` reads
- Without session context these are useful but not urgent — user can edit `context.md` directly
- `context park` ships with magpie-claude (Phase B) as a session-aware command; these are the standalone equivalents
