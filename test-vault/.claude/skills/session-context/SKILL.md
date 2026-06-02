# Session Context Skill

## Purpose

This skill governs reading and writing `context.md`, running the session start/end protocol, and keeping the vault's session state accurate.

---

## context.md Field Guide

| Field | Type | Description |
|-------|------|-------------|
| `last-session` | string | Date of last session end (YYYY-MM-DD). Written by binary at session-end. |
| `current-focus` | string | 1-5 words, action-verb first. Written by LLM at session end. |
| `active-projects` | list | Free-form list of active projects. |
| `next-actions` | list | Up to 3 domain-tagged action strings (see format below). |
| `parked-ideas` | list | Append-only. Never remove entries. |
| `inbox-count` | int | Updated by binary automatically. Do not set manually. |
| `session-count` | int | Incremented by LLM at session start. |
| `last-nudged-command` | string | Last command surfaced in a discovery nudge. |

### next-actions Format

Each item is a string: `"[domain] Action verb + specific target"`

**Domain tags:** `[wiki]`, `[adr]`, `[investigate]`, `[learn]`, `[debrief]`, `[general]`

**Good examples:**
- `"[adr] Draft decision section for ADR-0003 caching strategy"`
- `"[wiki] Ingest the RAG paper currently in inbox"`
- `"[investigate] Add latency evidence from benchmark to open investigation"`

**Bad examples (never write these):**
- `"[wiki] Continue working on things"` — vague, no verb+target
- `"Think about the architecture"` — missing domain tag, no verb

### parked-ideas Format

Each entry: `"[YYYY-MM-DD] idea text"`

Append-only. Never remove entries. Never edit past entries.

---

## Session Start Protocol

Every session, before responding to the user's first message:

1. **Read context.md** — extract all fields into active memory.
2. **Plugin discovery** — `ls .claude/plugins/` and read each PLUGIN.md.
3. **Verify hook ran** — look for `vault-tools session-start` output in additionalContext. If absent, emit the briefing manually:
   ```
   Last focus: <current-focus or "none recorded">
   Next: <next-actions[0]> [+N more] or "none set"
   Inbox: <inbox-count> items
   Dev: <plugin zone summary> or omit if empty
   ```
4. **Increment session-count** in context.md.
5. **Command nudge** — if `session-count` is divisible by 5, emit one discovery nudge (see rotation below).
6. Wait for the user's first message. Take no further action.

---

## Session End — Who Writes What

**LLM writes** (before calling vault-tools session-end):
- `current-focus`: 1-5 words, action verb first. Derived from what was actually worked on this session.
- `next-actions`: list of up to 3 domain-tagged items, ordered by priority. Elicited from user during /wrap.
- `parked-ideas`: append any new items from this session. Never remove existing items.

**Binary writes** (via `vault-tools session-end`):
- `last-session`: today's date
- `inbox-count`: current count of non-hidden inbox files

**Never** touch `inbox-count` or `last-session` yourself — the binary owns those fields.

---

## Good current-focus Examples

Write `current-focus` as: action verb + brief object, 1-5 words.

- `"Writing vault build plan"` ✓
- `"Investigating RAG chunking"` ✓
- `"Debriefing prod incident"` ✓
- `"vault stuff"` ✗ — too vague
- `"Working on the system"` ✗ — no action verb, too long

---

## Command Discovery Nudge Rotation

Every 5th session, emit a nudge for the next command in the rotation. Track position via `last-nudged-command`.

**Rotation order (all 15 commands, repeat):**
1. `/ingest` — Process inbox items into archive and wiki
2. `/query` — Synthesize an answer from vault knowledge
3. `/lint` — Health-check the vault for broken links and orphans
4. `/capture` — Zero-friction inbox dump from terminal
5. `/park` — Shelve an idea and return to current focus
6. `/next` — Break a topic into 3 concrete first steps
7. `/focus` — Check whether this session has drifted from planned focus
8. `/resume` — Restore full session context mid-session
9. `/wrap` — End session, update context, log summary
10. `/inbox` — List inbox contents
11. `/adr` — Create or update an Architecture Decision Record
12. `/debrief` — Write a blameless post-mortem
13. `/investigate` — Open or update an investigation
14. `/learn` — Add a session note to a learning track
15. `/review` — Synthesize a learning track into a wiki synthesis

**Nudge format:**
```
Tip: /command — one sentence description.
```

To find the next nudge: look up `last-nudged-command` in context.md, find its position in the list above, advance one, and emit that command. If `last-nudged-command` is empty, start with `/ingest`. After emitting, write the command name to `last-nudged-command`.

---

## When to Update context.md Mid-Session

Update immediately (don't wait for /wrap) when:
- User runs `/park`: append the parked idea to `parked-ideas` right away
- User runs `/next` and chooses to add action 1: prepend it to `next-actions`, trim to 3 items
- User states a new focus explicitly: update `current-focus` immediately

---

## Writing next-actions at /wrap

Elicit from the user: "Next actions? Up to three, domain-tagged."

Example user answer: "[wiki] ingest the RAG paper, [adr] finish ADR-0002 context section"

Parse into a list:
```yaml
next-actions:
- "[wiki] Ingest the RAG paper currently in inbox"
- "[adr] Write context section for ADR-0002"
```

Rules:
- Maximum 3 items
- Each must start with a domain tag
- Action verb must be first word after the tag
- Trim to the most specific target possible
