# Project Tracking Skill

## Purpose

Guide the creation and maintenance of project files in `dev/projects/`. Each project has an index (current state), dated session logs (what happened when), and backlog items (individual files queryable by Obsidian Bases).

---

## Project Index Structure

```markdown
---
title: ""
type: project
slug: ""
status: active
repo: ""
tech-stack: []
created: YYYY-MM-DD
updated: YYYY-MM-DD
related-adrs: []
related-investigations: []
tags: []
---

# [Project Name]

## What It Is

[One paragraph: purpose, problem it solves, current state.]

## Goals

- [Concrete goal]
- [Concrete goal]

## Decisions & Context

[Key decisions made, constraints, why things are the way they are. Link to ADRs where relevant: [[ADR-NNNN: Title]]]

## Open Questions

- [Question or blocker]

## Backlog

![[backlog.base]]

## Log

_Session notes below, newest first (links only — files are in this folder)_

- [[YYYY-MM-DD]] — one-line summary
```

---

## Session Log Structure

```markdown
---
title: "Session YYYY-MM-DD — [Project Name]"
type: project-log
project: ""
created: YYYY-MM-DD
---

# Session YYYY-MM-DD

## What Was Worked On

[What actually happened this session.]

## Decisions Made

[Any choices made and why. If a decision warrants a full ADR, note it here.]

## Blockers / Open Issues

[What is blocking progress, if anything.]

## Next

[What to pick up next session.]
```

---

## Backlog Item Structure

```markdown
---
title: ""
type: backlog-item
project: ""
status: todo
priority: medium
due: ""
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---

# [Item Title]

## Description

[What needs to be done and why.]

## Acceptance Criteria

- [ ] [Specific, testable condition]

## Notes

[Additional context, links, related items.]
```

---

## Creating a New Project

1. Ask for: project name, repo URL (optional), tech stack, initial goals
2. Derive the slug (lowercase, hyphens)
3. Draft the index file — present before writing
4. Create `backlog.base` using the standard template for this project's folder path
5. Add to `wiki/index.md` under `## Dev`
6. Log with `vault-tools log-append`

## Adding a Session Log

1. Read the project index first
2. Ask what was worked on (or accept inline from the user's message)
3. Draft the session log — present before writing
4. Append a link to the log list in `index.md` under `## Log`
5. Update `updated:` field in the index frontmatter

## Adding a Backlog Item

1. Read the project index first
2. Elicit: title, description, priority, due date (optional)
3. Draft the file — present before writing
4. No changes to `index.md` needed — `backlog.base` picks it up automatically

## Updating a Project's Status

When status changes (pause, complete, archive):
1. Update `status:` in the index frontmatter
2. Update `updated:` to today
3. Update `wiki/index.md` — remove if completed/archived, keep if paused

---

## Connection to Other Plugins

- When a project needs a formal architectural decision → suggest `/adr`
- When a project has an open question worth investigating → suggest `/investigate`
- When a project incident occurs → suggest `/debrief`
- When ingesting a source related to a project → surface the connection
