---
title: "magpie-git"
type: backlog-item
project: magpie
status: todo
milestone: "1.0"
priority: medium
due: ""
created: 2026-05-21
updated: 2026-05-21
tags: [phase-6, plugin, magpie-git]
---

# magpie-git

Second bundled plugin. Vault-scoped git helpers. Lives in `plugins/magpie-git/`. Validates the manifest contract: tool dependency check for `git`, `--manifest` flag, and `post_install` execution.

## Learning Focus

- `go:embed` for bundling `plugin.yaml` in the binary
- `plugin.yaml` manifest with tool dependency version check
- `exec.Command` for git subprocess calls scoped to vault root
- `post_install` pattern for plugin-specific setup

## Commands

```
magpie git status       # git status scoped to vault root
magpie git commit       # git commit scoped to vault root
magpie git log          # recent commits scoped to vault root
```

## plugin.yaml

```yaml
name: magpie-git
schema: "1"
requires_magpie: ">= 1.0.0"

tools:
  - name: git
    check: "git --version"
    min_version: "2.0"
    required: true
    hint: "install git from https://git-scm.com"

post_install:
  - "magpie-git configure"    # any git-specific setup
```

## You Drive

Implementing `--manifest` — the binary must embed `plugin.yaml` via `go:embed` and print it to stdout when called with `--manifest`. How does this interact with cobra's command parsing? Where does the flag live — on the root command or as a special case before cobra runs?
