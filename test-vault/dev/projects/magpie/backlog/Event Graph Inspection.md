---
title: "Event Graph Inspection"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, event-bus]
---

# Event Graph Inspection

Full visibility into the event subscription graph across all installed plugins — who emits what, who subscribes to what.

## Commands

```
magpie events list             # all core events + plugin-declared events
magpie events list --json      # machine-readable graph
magpie events subscribers <event>   # which plugins subscribe to a given event
```

## Notes

- All data is statically derivable from installed plugin manifests — no runtime needed
- `magpie plugin status <name>` already shows per-plugin subscriptions; this is the global view
- Useful for debugging unexpected subscriber behavior and auditing the event graph
