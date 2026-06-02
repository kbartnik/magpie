---
title: "Multi-Vault"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [post-1.0, multi-vault]
---

# Multi-Vault

Vault discovery and switching across multiple vaults. Because vaults are self-registering (any directory with `.magpie/` is a vault), no registry is needed — just a scan.

## Commands

```
magpie vaults list             # scan known locations for .magpie/ directories
magpie vaults list --all       # deeper scan from home dir
```

## Notes

- Resolution tier 3 (`default_vault` in global config) becomes the fallback when outside all vaults
- `magpie vaults list` is essentially `find ~ -name .magpie -type d -maxdepth 5` expressed in Go
- No `magpie vaults add` needed — `magpie init vault` in a directory makes it discoverable
