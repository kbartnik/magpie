# ADR Writing Skill

## Purpose

Guide the creation and maintenance of Architecture Decision Records in `dev/adr/`. ADRs capture why a decision was made, not just what was decided.

---

## Elicitation — One Question at a Time

Never ask multiple questions at once. Work through these in order, waiting for each answer:

1. "What decision needs to be made (or was made)?"
2. "What context led to this decision — what constraints, requirements, or prior state made this decision necessary?"
3. "What was decided?"
4. "What are the consequences — both positive and negative?"
5. "What alternatives were considered and rejected?"

After gathering all answers, draft the full ADR and present it for review.

---

## ADR Structure

```markdown
---
title: "Use X for Y"
type: adr
status: proposed
decision-date: ""
created: YYYY-MM-DD
updated: YYYY-MM-DD
supersedes: []
superseded-by: ""
tags: []
---

# ADR-NNNN: Use X for Y

## Context

[Why this decision was needed. What constraints, requirements, or prior state made it necessary. What problem this solves.]

## Decision

[What was decided. Written as a clear declarative statement.]

## Consequences

### Positive
- [benefit]

### Negative
- [cost or tradeoff]

## Alternatives Considered

### Option: [Alternative name]
[Why it was rejected]

## References

- [Related investigations, docs, or ADRs]
```

---

## File Naming

Before creating: `ls dev/adr/ | grep "^ADR-" | sort | tail -1` to find the next number.

Zero-pad to 4 digits: ADR-0001, ADR-0002, etc.

Slug: lowercase, hyphens, describes the decision. Example: `ADR-0003-use-postgres-for-sessions.md`

---

## Zone Rules (from PLUGIN.md)

- Never edit accepted ADR content — only `status` and `superseded-by` fields
- Always check for existing ADRs on the topic before creating
- Plan + approval required before any write

---

## After Creating

1. Check `dev/investigations/` for investigations that informed this decision:
   ```bash
   ls dev/investigations/*.md 2>/dev/null
   ```
   If any are related, add them to the References section.

2. Update `wiki/index.md` under `## Dev`:
   ```
   [[ADR-NNNN: Title]] — status: proposed
   ```

3. Call log-append:
   ```bash
   "${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append adr "ADR-NNNN: Title" "+1 created" "ADR drafted covering X" "Decision pending"
   ```

---

## Accepting an ADR

When user says "accept this" or "let's go with this":
1. Set `status: accepted` and `decision-date: YYYY-MM-DD`
2. Update `updated:` field
3. Update wiki/index.md entry to show `status: accepted`
4. If this supersedes an older ADR, set `superseded-by:` on the old one and `supersedes:` on this one

ADR content is now immutable. Do not edit the body again.
