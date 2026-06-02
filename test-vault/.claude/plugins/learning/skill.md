# Learning Notes Skill

## Purpose

Guide the creation of learning tracks and session notes in `dev/learning/`. Learning tracks build cumulative understanding over time through multiple sessions with a topic.

---

## Track Index Structure

The index file lives at `dev/learning/<topic-slug>/index.md` and is updated after every session:

```markdown
---
title: "Learning: [Topic Name]"
type: learning
topic: "<topic-slug>"
status: active
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---

# Learning: [Topic Name]

## What I Want to Understand

[The specific questions this track is trying to answer. What does "done" look like?]

## Sources Consumed

- [YYYY-MM-DD] [Source title and type — book chapter, blog post, video, experiment]

## Key Concepts

- **[Concept]**: [Brief definition in my own words]

## What I Now Understand

[Updated each session — not appended. The current state of understanding. Be specific and honest about gaps.]

## Open Questions

- [Question that came up but isn't answered yet]
- [Something to explore next]

## What to Explore Next

- [Specific next source or experiment]
```

---

## Session Note Structure

Each session note lives at `dev/learning/<topic-slug>/YYYY-MM-DD.md` and is never edited after writing:

```markdown
---
title: "[Topic] — YYYY-MM-DD"
type: learning-note
topic: "<topic-slug>"
source: "[What was read/watched/practiced]"
created: YYYY-MM-DD
---

# [Topic] — YYYY-MM-DD

## Key Insight

[The single most important thing learned this session — 1-3 sentences. If you can only remember one thing, what is it?]

## What Surprised Me

[Something unexpected, counter-intuitive, or that challenged a prior assumption]

## Connections

- [[Wiki Page]] — how it connects
- [Other learning track] — how it relates

## Open Questions

- [Question this session raised but didn't answer]
```

---

## Session Timestamps

Daily files (`YYYY-MM-DD.md`) may contain multiple sessions. Mark boundaries with:

```markdown
---

### HH:MM

[content added during session]

*ended HH:MM*
```

- **Session start:** append `---\n\n### HH:MM\n` when first writing to the file in a new session
- **Session end:** append `\n*ended HH:MM*` at `/wrap`, followed by a `## Where I Left Off` block
The `## Where I Left Off` block is one sentence written at `/wrap`: what was actively in progress when the session ended. It is the re-entry point for the next session.

```markdown
*ended HH:MM*

## Where I Left Off

[One sentence: what was in progress — e.g. "Writing failing tests for Task 3 config loading."]
```

- Multiple sessions in a day stack naturally beneath each other
- Do not apply retroactively to existing notes

---

## During a Session

Append to the day's session note in real time — don't hold content for /wrap:
- Tutorial explanations and concept discussions as they happen
- Design decisions made and why
- Surprises or things that didn't work as expected
- Open questions as they arise

/wrap extracts from a complete record. It does not reconstruct from memory.

---

## After Each Session Note

1. Update the track `index.md`:
   - Add the source to "Sources Consumed"
   - Update "What I Now Understand" (replace, don't append — this is current understanding)
   - Add new open questions, remove resolved ones
   - Update "What to Explore Next"
   - Update `updated:` field

2. Check if any open questions warrant a new investigation:
   > "This open question — '[X]' — seems like it could become a proper investigation. Want me to open one?"

---

## Concept Extraction (at /wrap)

When `/wrap` triggers concept extraction for a track:

1. Read the session note(s) written today for this track.
2. Identify concepts that would stand alone as wiki pages — a concept qualifies if:
   - It has a clear definition independent of the session context
   - It would be useful to link to from future notes or wiki pages
   - It isn't already a wiki page (check `wiki/index.md`)
3. Propose the list with a one-line rationale for each candidate:
   > "Candidates for wiki pages:
   > - **[Concept 1]** — [why it qualifies: e.g. 'reusable mental model', 'referenced in 3 places', 'non-obvious']
   > - **[Concept 2]** — [why it qualifies]
   > Which should I create? (all / pick / skip)"
4. For each approved concept:
   - Create `wiki/concepts/<concept-slug>.md` with standard wiki frontmatter (`type: concept`)
   - Write a concise reference page: definition, why it matters, key properties, links to related concepts
   - Add a wikilink back to the session note under a `## Sources` section
   - Add to `wiki/index.md` under `## Concepts`
   - Scan recent session notes and wiki pages for mentions of this concept — add retroactive wikilinks
5. After all concepts are written, update the session note's `## Connections` section to link the new pages.

---

## When a Track Reaches `complete`

When the user declares the track complete or "What I Want to Understand" is fully answered:
1. Set `status: complete` in index.md frontmatter
2. Always offer synthesis:
   > "This track is complete. Want me to synthesize it into a `wiki/syntheses/<topic>.md` page?"
3. If yes: read all session notes and the track index, draft the synthesis, present for review, write after approval, update wiki/index.md.
