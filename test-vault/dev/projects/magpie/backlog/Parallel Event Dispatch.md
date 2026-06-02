---
title: "Parallel Event Dispatch"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, event-bus, performance]
---

# Parallel Event Dispatch

Optional parallel dispatch for event subscribers when order doesn't matter. v1.0 dispatches sequentially — predictable, easy to debug. Parallel is an opt-in once the event bus is proven in production.

## Design

```yaml
subscriptions:
  - event: inbox.captured
    run: "magpie-git commit"
    parallel: true    # opt-in per subscription
```

## Notes

- Default remains sequential — parallel is a conscious opt-in by the plugin author
- Parallel subscribers must treat failures as independent (no shared state assumptions)
- Requires a timeout mechanism to prevent slow subscribers from blocking indefinitely
- Only worth adding once there's a demonstrated performance need
