# Source Annotation Skill

## Purpose

After deep-reading an archived source, produce a `## Deep Read` section appended directly to the archived file. This section is the vault's memory of what was valuable in the source.

---

## When to Load This Skill

Load this skill before reading any source file during `/ingest`. Do not load it for other operations.

---

## Deep Read Section Structure

After reading the source fully, append this section to the archived file (do not create a separate file):

```markdown
## Deep Read

**Key Insight:** [The single most important takeaway — 1-2 sentences. If you can only remember one thing, what is it?]

**What Surprised Me:** [Something unexpected, counter-intuitive, or that challenges a prior assumption]

**Open Questions:** [2-3 questions this source raises but does not answer]
- 
- 
- 

**Wikilink Candidates:** [Concept names from this source that should become wiki pages]
- [[Concept Name]] — does this page exist? If not, create a stub.

**Connections:** [Links to existing wiki pages this source relates to]
- [[Page Name]] — how it connects

**Image Candidates:** [informational images from this source worth placing in wiki pages, or "none"]
- `![[resources/media/hash_MD5.webp]]` — [what it shows] → [[WikiPage]] § [Section]
```

---

## Image Scan

For clipping-type sources: after reading the body, check for embedded images:

```bash
grep '!\[\[resources/media' <archived-file-path>
```

For each image found:
1. View it with the Read tool
2. Classify as **informational** or **decorative**:
   - **Informational:** diagram, flowchart, architecture figure, labeled chart, data visualization, screenshot showing structure — anything that conveys meaning a reader can't get from prose alone
   - **Decorative:** photograph, stock illustration, author portrait, decorative art, hero image — visually appealing but adds no information not already in the text
3. Skip decorative images entirely
4. For each informational image: identify which wiki concept it illustrates and which section of that page it would replace or enhance — prioritize spots where prose currently describes what the image shows directly (ASCII art, text diagrams, dense process descriptions)

Record findings in the **Image Candidates** field of the Deep Read. When creating or updating wiki pages, place informational images at the identified locations using `![[resources/media/...]]` embeds rather than writing text approximations of what the image already shows.

---

## Wikilink Candidate Process

For each wikilink candidate identified:
1. Check if `wiki/concepts/<name>.md` or `wiki/entities/<name>.md` exists
2. If it does: add a wikilink to the concept in the Deep Read section
3. If it does not: offer to create a stub page now ("Want me to create a stub for [[Concept Name]]?")

---

## Investigation Connection Check

After deep-reading, always check `dev/investigations/` for open investigations this source might inform:

```bash
ls dev/investigations/*.md 2>/dev/null
```

For each open investigation file found, check if the investigation's question or topic is relevant to the source. If yes, surface it:

> "This source relates to [[Investigation: Your Question]]. Want me to add it as evidence?"

If the user says yes: open the investigation file, add the source to the Evidence section with a brief note.

---

## Archive Path

Archived sources live at: `archive/<type>/YYYY-MM-DD-slug.md`

Types: `clipping`, `paper`, `book`, `daily`, `idea`, `doc`

The Deep Read section is appended to this file — not to a separate note, not to the wiki.

---

## Quality Bar

A good Deep Read:
- Has a Key Insight that would be useful to someone who hasn't read the source
- Has Open Questions that are genuinely unresolved (not rhetorical)
- Has at least one Wikilink Candidate
- Has at least one Connection to an existing wiki page (or notes "no existing connections")
- Has run the Image Scan for clipping sources; Image Candidates says "none" only after actually checking

A bad Deep Read:
- Summarizes the source (summaries go in the archive file body)
- Has vague open questions ("What else is there to know?")
- Has wikilink candidates that are just keywords, not actual concepts
- Omits the Image Candidates field entirely for clipping sources
