---
name: code
version: "1.0.0"
zone: dev/code/
commands: [/code]
skill: .claude/plugins/code/skill.md
schema-version: "1"
---

# Plugin: Code

## Testbed Role

This vault is a **testbed**, not a reference vault. The code plugin here has one primary purpose:

**Document the magpie repo itself.** Run `/code magpie` against `/Users/kurt/Source/magpie` to build and refresh the magpie code zone. This simultaneously:
- Tests the `/code` command workflow end-to-end
- Produces genuinely useful documentation for magpie development
- Exercises `vault-tools go-extract` against a real Go codebase

Other projects may be added for testing specific code plugin scenarios (multi-package layout, different directory structures, etc.), but magpie is the primary and persistent target.

**Code zones here are snapshots, not authoritative references.** If the magpie code zone is stale, it represents the code at a past point in time — this is expected behavior for a testbed.

---

## Zone Rules

Each project lives in its own subdirectory under `dev/code/`. Always read the project index before modifying any file in the zone. The zone has three layers per project — docs, packages, and functions — each with distinct frontmatter types.

Creating or modifying a code zone requires presenting a plan and receiving approval before writing.

If a plugin zone is touched before this file has been loaded, stop and load it first.

## File Naming

- Project index: `dev/code/<slug>/index.md`
- Imported doc: `dev/code/<slug>/docs/<name>.md` — name is kebab-case of the source filename without extension
- Package page: `dev/code/<slug>/packages/<pkg-slug>.md` — pkg-slug is the import path with `/` replaced by `-`
- Function note: `dev/code/<slug>/functions/<pkg-slug>-<FuncName>.md`
- Bases view: `dev/code/<slug>/functions.base`

Example:
```
dev/code/magpie/index.md
dev/code/magpie/docs/architecture.md
dev/code/magpie/packages/cmd.md
dev/code/magpie/functions/cmd-Execute.md
dev/code/magpie/functions.base
```

## Frontmatter Schemas

**Code index** (`type: code-index`):
```yaml
---
title: ""
type: code-index
project: ""
repo: ""
language: go
snapshot-date: YYYY-MM-DD
status: current
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---
```

**Code doc** (`type: code-doc`):
```yaml
---
title: ""
type: code-doc
project: ""
source-file: ""
snapshot-date: YYYY-MM-DD
tags: []
---
```

**Code package** (`type: code-package`):
```yaml
---
title: ""
type: code-package
project: ""
path: ""
language: go
confidence: extracted
snapshot-date: YYYY-MM-DD
imports: []
imported-by: []
key-types: []
tags: []
---
```

**Code function** (`type: code-function`):
```yaml
---
title: ""
type: code-function
project: ""
package: ""
signature: ""
receiver: ""
calls: []
called-by: []
confidence: extracted
snapshot-date: YYYY-MM-DD
tags: []
---
```

`called-by` is always written as `[]` — it is reserved for future use. Query call relationships via the `calls` field in `functions.base`.

## Status Lifecycle

- `current` → `stale`: lint flags after 60 days without refresh, or when `status` is explicitly set to `stale`
- `stale` → `current`: after a `/code` refresh completes

## Index Convention

Do not add code projects to `wiki/index.md`. Link from the existing project wiki page (`wiki/concepts/Project: <name>.md`) to `dev/code/<slug>/index.md`. Code documentation is an attribute of the project, not a top-level wiki entity.

## Stale Threshold

60 days in `current` status without a snapshot refresh — flagged by `vault-tools lint-gather` under `STALE_CODE_INDEXES`.
