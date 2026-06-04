# Investigation Tracking Skill

## Purpose

Guide the creation and ongoing maintenance of investigation files in `dev/investigations/`. Investigations are living documents that accumulate evidence until a conclusion is reached.

---

## Investigation Structure

```markdown
---
title: "Full question being investigated?"
type: investigation
question: "Full question being investigated?"
status: open
confidence: low
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---

# Investigation: [Full Question]

## Context

[Why this question matters. What prompted the investigation. What the current situation is.]

## Current Best Answer

[Updated each time new evidence arrives — not appended, replaced. This is the live hypothesis at the current confidence level. If confidence is low, say so explicitly: "Best guess: X, but this may be wrong because Y."]

## Evidence For

- [Source type, date] [Finding that supports the hypothesis]
- [Benchmark result, 2026-05-01] [Latency drops to 12ms with connection pooling]

## Evidence Against

- [Source type, date] [Finding that contradicts or complicates the hypothesis]

## Open Threads

- [Unresolved sub-question or missing evidence]
- [Something to follow up on]

## Conclusion

[Written only when status moves to `answered`. The final answer with supporting evidence cited. Once written, status → answered.]
```

---

## Evidence Standards

Informal evidence counts — the goal is honest epistemic accounting, not academic rigor:

- Benchmarks and profiling results
- Production data and logs
- Documentation and blog posts (note recency)
- Stack Overflow answers (note confidence)
- Direct experience ("I tried X and it behaved like Y")
- Conversations with domain experts

For each piece of evidence: note the source type and date. Old evidence may be stale.

---

## Updating an Existing Investigation

Always read the current file first. Then:

1. Update `Current Best Answer` if the new evidence shifts the picture
2. Add the new evidence to the appropriate section (For/Against)
3. Remove or strike through Open Threads that are now resolved
4. Update `confidence`: low → medium → high as evidence accumulates
5. Update `updated:` field to today

**Confidence levels:**
- `low`: hypothesis but little supporting evidence
- `medium`: multiple consistent data points, no strong contradictions
- `high`: strong evidence, contradictions explained, ready to conclude

---

## Investigation Connection During /ingest

When `/ingest` processes a source that connects to an open investigation, always surface it:

> "This source relates to [[Investigation: Your Question]]. Want me to add it as evidence?"

If yes: open the investigation file, add the source to Evidence For or Against with a brief note, update Current Best Answer if warranted, adjust confidence if appropriate.

---

## Closing an Investigation

When confidence reaches `high` and the user agrees it's answered:
1. Write the Conclusion section
2. Set `status: answered`
3. Remove from wiki/index.md Dev section
4. Offer: "Want me to create a `wiki/syntheses/` page from this investigation?"

When shelving:
1. Set `status: shelved`
2. Add a note at the top: "Shelved YYYY-MM-DD — [reason]"
3. Remove from wiki/index.md Dev section
