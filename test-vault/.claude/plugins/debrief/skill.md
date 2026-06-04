# Debrief Writing Skill

## Purpose

Guide blameless post-mortem writing for incidents, failures, or significant surprises. The goal is learning, not blame.

---

## Elicitation — One Question at a Time

Work through these in order, waiting for each answer:

1. "What happened? Give me a brief description of the incident."
2. "Walk me through the timeline — what happened in what order?"
3. "What worked well during the response?"
4. "What didn't work? (Systems, processes, tooling — not people)"
5. "What was the root cause? If there were multiple contributing factors, list them."

After gathering all answers, draft the full debrief and present it for review.

---

## Debrief Structure

```markdown
---
title: ""
type: debrief
status: draft
incident-date: YYYY-MM-DD
severity: low
duration-minutes: 0
tags: []
related-adrs: []
---

# [Title]

## TL;DR

[2 sentences maximum: what happened and the key learning]

## Timeline

- **HH:MM** — [event]
- **HH:MM** — [event]
- **HH:MM** — [incident resolved]

## Root Cause

[The underlying systemic cause(s). Not what failed — why it was possible for it to fail.]

## What Worked

- [Things that helped — monitoring, communication, tooling, process]

## What Didn't Work

- [Systems, processes, or tooling that hindered — no individual names]

## Action Items

- [ ] [Concrete action] — owner: [team or role, not person name] — due: YYYY-MM-DD

## Generalizable Learning

[The principle this incident teaches that applies beyond this specific situation. This is the most important section — what would a future engineer need to know to prevent this class of failure?]
```

---

## Blameless Writing

The "What Didn't Work" and "Root Cause" sections must never name individuals. If the user names someone, reframe to the system:

- BAD: "Alice didn't check the deploy checklist"
- GOOD: "The deploy checklist was not verified before rollout"

If the user pushes back on this, explain: "Blameless post-mortems produce better learning by focusing on systemic causes. Individuals making mistakes is normal — the question is why the system allowed the mistake to propagate."

---

## After Writing

Always offer to create a `wiki/concepts/` page from the Generalizable Learning section:

> "The generalizable learning here is '[X]'. Want me to create a wiki concept page for this?"

If yes: draft the concept page, present it, create after approval, update wiki/index.md.

Then call log-append:
```bash
"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append debrief "<title>" "+1 created" "<brief summary>" "<action items open>"
```
