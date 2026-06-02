---
title: "Sonar"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, sonar, go-extract]
---

# Sonar

Graduate `go extract` from magpie into its own standalone project (`github.com/kbartnik/sonar`). It has no vault dependency and belongs on its own so it can grow independently and be used outside the vault context.

## Notes

- `magpie extract go` remains as a thin wrapper or alias pointing to `sonar` post-split
- `sonar` registers itself as a magpie plugin via the standard plugin contract
- The split validates that the plugin system handles a fully independent project cleanly
- Timeline: after magpie 1.0 ships and the plugin contract is stable
