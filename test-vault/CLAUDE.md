# CLAUDE.md — magpie test vault

## This Vault Is a Testbed

This vault exists to exercise vault command workflows and validate magpie CLI behavior. It is **not** a reference knowledge vault.

- **Nexus** (`~/Documents/Obsidian/Nexus`) — the reference vault: real knowledge, real decisions, real projects
- **This vault** — the testbed: magpie-specific content and workflow validation

Content here is either:
1. **Magpie-specific** — ADRs about magpie's design, investigations into its behavior, code docs for the magpie repo
2. **Synthetic test scenarios** — used to exercise a command workflow, not treated as authoritative knowledge

When a test produces broadly applicable insight, surface it for Nexus rather than keeping it here.

---

## Session Start Protocol

Every session, before responding to the user's first message:

1. Read `context.md` — extract all fields
2. **Plugin discovery:** `ls .claude/plugins/` — for each directory, read `.claude/plugins/<name>/PLUGIN.md`
3. Verify the SessionStart hook ran. If absent, emit the briefing:
   ```
   Last focus: <current-focus>
   Next: <next-actions[0]> [+N more]
   Inbox: <inbox-count> items
   ```
4. Increment `session-count` in `context.md`
5. Wait for user command.

---

## Zones

| Zone | Path | Notes |
|------|------|-------|
| Wiki | `wiki/` | Agent read-write |
| Archive | `archive/` | Read-only after archiving |
| Dev | `dev/` | Plan before write |

---

## Plan-Before-Write Protocol

Before creating, moving, or modifying files in this vault:

```
📋 Plan
  Operation: [what and why]
  Files to create: [list or "none"]
  Files to move: [source → destination, or "none"]
  Files to update: [list or "none"]

Proceed? (y/n)
```

**Exception:** `vault-tools log-append` and `vault-tools update-inbox-count` need no plan.

---

## Available Skills

| Skill | Path |
|-------|------|
| session-context | `.claude/skills/session-context/SKILL.md` |
| source-annotation | `.claude/skills/source-annotation/SKILL.md` |

---

## Available Commands

| Command | Purpose |
|---------|---------|
| `/wrap` | End session — update context, log summary |
| `/learn` | Add session note to a learning track |
| `/review` | Synthesize a learning track into a wiki page |
| `/resume` | Restore session context mid-session |
| `/lint` | Health-check the vault |
| `/capture` | Zero-friction inbox dump |
| `/ingest` | Process inbox items into archive and wiki |
| `/query` | Synthesize answer from vault |
| `/park` | Shelve an idea, return to focus |
| `/focus` | Check for topic drift |
| `/next` | Break a topic into concrete first steps |
| `/project` | Manage project files |
| `/adr` | Create or update Architecture Decision Records |
| `/debrief` | Write a blameless post-mortem |
| `/inbox` | List inbox contents |
| `/investigate` | Open or update an investigation |
| `/code` | Import or refresh code documentation |

---

## Installed Plugins

| Plugin | Zone | Testbed use |
|--------|------|-------------|
| learning | `dev/learning/` | magpie-specific learning tracks |
| projects | `dev/projects/` | (empty — projects moved to Nexus) |
| adr | `dev/adr/` | magpie design decisions + workflow testing |
| code | `dev/code/` | magpie repo documentation (`/code magpie`) |
| debrief | `dev/debriefs/` | magpie incidents + workflow testing |
| investigation | `dev/investigations/` | magpie behavior questions + workflow testing |
