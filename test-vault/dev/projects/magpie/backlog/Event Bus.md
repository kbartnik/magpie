---
title: "Event Bus"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, event-bus, pubsub]
---

# Event Bus

Plugin-to-plugin pub/sub: every command emits a named JSON event; plugins subscribe in their manifest and receive payloads on stdin. Deferred from v1.0 because Claude Code hooks satisfy all immediate use cases and the design needs real multi-plugin pressure to get right.

## Design (from brainstorm — verify against live codebase before implementing)

- Every command emits a named event with JSON payload on completion
- Plugin manifests declare `subscriptions` (event name → command to run)
- Magpie invokes subscribers sequentially; failures never block the original command
- Payload piped to subscriber's stdin
- Plugin events namespaced by plugin name: `claude.session.ended`

## Why deferred

One concrete v1.0 subscriber (magpie-git auto-commit) is insufficient design pressure. The stdin/terminal ambiguity problem (how does a subscriber know it was invoked by the bus vs. the user?) needs solving with real use cases. Build when there are at least two independent consumers.

## See also

[[2026-05-21 — Red Team Review]] — failure mode 3 and steelman section explain the deferral reasoning.
