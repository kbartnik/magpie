---
title: "Verbose Output Mode"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-06-02
updated: 2026-06-02
tags: [post-1.0, research, output, ux]
---

# Verbose Output Mode

Research item. The 1.0 output model uses TTY detection: JSON when piped, human-readable when terminal. A verbose flag is a natural extension but the design needs real usage before the shape is clear.

## Open Questions

- Is `--verbose` the right mechanism, or is it better modeled as a log level (`--log-level debug`)?
- Should verbosity be per-command or a global flag that propagates via `PersistentPreRun`?
- What does verbose output actually add that the default terminal output doesn't? Progress indicators? Intermediate steps? Decision traces?
- Does verbose mode break any downstream tooling that pipes magpie output?

## Design Pressure Needed

At least two concrete use cases where the default output is insufficient for debugging or understanding what magpie did. Build when those cases emerge naturally from usage — not speculatively.

## See Also

[[magpie — Design Overview]] — Terse by default principle
