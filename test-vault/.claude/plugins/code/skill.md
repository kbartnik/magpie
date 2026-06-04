# Code Plugin Skill

## Purpose

Populate and refresh `dev/code/<project>/` with three layers of content from a source repository:
1. **Docs** — verbatim imports from the repo's `doc/` folder
2. **Packages** — one authored page per package: architectural role, import relationships, key types
3. **Functions** — one frontmatter-heavy note per exported function with signature and cross-package calls

All writes follow the plan-before-write protocol from CLAUDE.md.

---

## Invocation

`/code <project-slug>`

Where `<project-slug>` matches the project's subdirectory name (e.g. `/code preflight-sync-go`).

---

## First Invocation (zone absent)

When `dev/code/<project>/` does not exist:

1. Ask for the local repo path if not already known from context.

2. Run extraction:
   ```bash
   vault-tools go-extract <repo-path>
   ```
   Capture the full JSON output — this is the extraction payload for all page authoring.

3. List the doc files in `<repo-path>/doc/` (if it exists):
   ```bash
   ls <repo-path>/doc/
   ```
   If `doc/` does not exist in the repo, skip this step and omit `dev/code/<project>/docs/` from the plan below.

4. Present the import plan before writing anything:
   ```
   📋 Plan
     Operation: Initial code zone import for <project>
     Files to create:
       dev/code/<project>/index.md
       dev/code/<project>/docs/<N files from doc/>
       dev/code/<project>/packages/<N package pages>
       dev/code/<project>/functions/<N function notes>
       dev/code/<project>/functions.base
     Files to move: none
     Files to update: wiki/concepts/Project: <project>.md (add link to code index)

   Proceed? (y/n)
   ```

5. On approval, author all files per the Authoring Guidelines below.

6. Add a link to `wiki/concepts/Project: <project>.md`:
   ```markdown
   ## Code Documentation
   [[dev/code/<project>/index.md]] — snapshot YYYY-MM-DD
   ```

7. Append log entry:
   ```bash
   vault-tools log-append "code" "<project> code zone import" \
     "dev/code/<project>/ (+N created)" \
     "Imported N doc pages, N package pages, N function notes from <repo-path>" \
     "<any gaps or open questions noticed during authoring>"
   ```

---

## Subsequent Invocations (zone exists)

When `dev/code/<project>/` exists:

1. Read `dev/code/<project>/index.md`. Report:
   ```
   Code zone: <project>
   Snapshot: <snapshot-date> (<N> days ago)
   Status: <current|stale>
   
   Refresh options:
     1. Full refresh (docs + packages + functions)
     2. Packages + functions only (skip docs — they change rarely)
     3. Docs only
   ```

2. Wait for the user to select an option.

3. Read the `repo` field from `dev/code/<project>/index.md` frontmatter to obtain the repo path for steps 3 and 4.

4. For options 1 and 2: run `vault-tools go-extract <repo-path>` and capture JSON.

5. For option 1: also list `<repo-path>/doc/` to check for new doc files.

6. Present plan, proceed on approval.

7. Overwrite the selected layer(s). No merging — clean snapshot. Update `snapshot-date` and `updated` in `index.md` after writing.

---

## Special Cases

### Missing doc/ directory

If `doc/` does not exist in the repo, skip the docs layer and omit `dev/code/<project>/docs/` from the plan.

---

## Authoring Guidelines

### index.md

```markdown
---
title: "<project>"
type: code-index
project: "<project>"
repo: "<repo-path>"
language: go
snapshot-date: <today YYYY-MM-DD>
status: current
created: <today YYYY-MM-DD>
updated: <today YYYY-MM-DD>
tags: []
---

# <project> Code Index

Snapshot taken <today>. Run `/code <project>` to refresh.

## Layers

- [Docs](docs/) — <N> imported documentation files
- [Packages](packages/) — <N> package pages
- [Function Inventory](functions.base) — <N> exported functions

## Module

`<module from extraction JSON>` — <N> packages: <comma-separated package names>
```

### docs/<name>.md

File naming: `<source-filename-without-extension>.md` (e.g. `doc/ARCHITECTURE.md` → `docs/architecture.md`).

Copy content verbatim. Add frontmatter at the top only — do not alter the body:

```markdown
---
title: "<first heading from source file>"
type: code-doc
project: "<project>"
source-file: "doc/<filename>"
snapshot-date: <today YYYY-MM-DD>
tags: []
---

<verbatim body of source file, starting from the first line after the heading>
```

### packages/<pkg-slug>.md

File naming: replace `/` with `-` in the import path (e.g. `internal/transfer` → `internal-transfer.md`).

> Strip the module prefix first (the `module` field from the extraction JSON root), then replace `/` with `-`. Example: module `github.com/kbartnik/preflight-sync-go`, path `github.com/kbartnik/preflight-sync-go/internal/transfer` → strip prefix → `internal/transfer` → replace slashes → `internal-transfer.md`.

Write a 2-3 sentence description of the package's architectural role — what it owns, what decisions it encapsulates, why it exists as a separate package. This is synthesis, not just a list of exports.

```markdown
---
title: "<import-path>"
type: code-package
project: "<project>"
path: "<import-path from extraction JSON>"
language: go
confidence: extracted
snapshot-date: <today YYYY-MM-DD>
imports: <imports array from extraction JSON>
imported-by: <imported_by array from extraction JSON>
key-types: <key_types array from extraction JSON>
tags: []
---

# <package name>

<2-3 sentences: what this package owns, what problem it solves, its role in the system.>

## Key Types

<one line per type in key-types: name — what it represents>

## Functions

<one line per function in extraction JSON: FuncName — one-sentence description>
```

### functions/<pkg-slug>-<FuncName>.md

File naming: `<pkg-slug>-<FuncName>.md` (e.g. `internal-transfer-Copy.md`).

> Strip the module prefix first (the `module` field from the extraction JSON root), then replace `/` with `-` for the pkg-slug. Example: module `github.com/kbartnik/preflight-sync-go`, path `github.com/kbartnik/preflight-sync-go/internal/transfer`, function `Copy` → strip prefix → `internal/transfer` → replace slashes → `internal-transfer-Copy.md`.

Body is empty — all data lives in frontmatter. The Bases view surfaces it.

```markdown
---
title: "<pkg-name>.<FuncName>"
type: code-function
project: "<project>"
package: "<import-path>"
signature: "<signature from extraction JSON>"
receiver: "<receiver from extraction JSON, empty string if none>"
calls: <calls_packages from extraction JSON, or [] if field is absent>
called-by: []
confidence: extracted
snapshot-date: <today YYYY-MM-DD>
tags: []
---
```

### functions.base

```yaml
filters:
  and:
    - type == "code-function"
    - file.inFolder("dev/code/<project>/functions")
properties:
  receiver:
    displayName: Receiver
  signature:
    displayName: Signature
  calls:
    displayName: Calls Packages
views:
  - type: table
    name: All Functions
    order:
      - file.name
      - receiver
      - signature
  - type: table
    name: By Package
    groupBy:
      property: package
      direction: ASC
    order:
      - file.name
      - receiver
      - calls
```

---

## After Completion

Update `snapshot-date` and `updated` in `dev/code/<project>/index.md` to today.

Set `status: current`.
