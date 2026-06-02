# CLAUDE.md — magpie

> **Transitional note:** Skills and commands currently use `vault-tools` (Nexus binary, copied to `test-vault/.claude/tools/`). Phase 8 (Migration) replaces all `vault-tools` calls with `magpie` subcommands and removes the copied binary.

---

## Vault

The test vault lives at `test-vault/` in this repo. It is a fully functional magpie vault with `.magpie/` sentinel, wiki, project files, learning track, and archive.

**VAULT_PATH:** `test-vault` (relative to repo root, resolved to absolute at runtime)

All vault operations run against `test-vault/`. Skills use `${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools` — with `VAULT_PATH=/Users/kurt/Source/magpie/test-vault` this resolves correctly.

---

## Session Start Protocol

Every session, before responding to the user's first message:

1. Read `test-vault/context.md` — extract all fields
2. **Plugin discovery:** `ls test-vault/.claude/plugins/` — for each directory, read `test-vault/.claude/plugins/<name>/PLUGIN.md`
3. Verify the SessionStart hook ran. If absent, emit the briefing:
   ```
   Last focus: <current-focus>
   Next: <next-actions[0]> [+N more]
   Inbox: <inbox-count> items
   ```
4. Increment `session-count` in `test-vault/context.md`
5. Wait for user command.

---

## Zones

| Zone | Path | Notes |
|------|------|-------|
| Wiki | `test-vault/wiki/` | Agent read-write |
| Archive | `test-vault/archive/` | Read-only after archiving |
| Dev | `test-vault/dev/` | Plan before write |

---

## Plan-Before-Write Protocol

Before creating, moving, or modifying files in `test-vault/`:

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
| session-context | `test-vault/.claude/skills/session-context/SKILL.md` |
| source-annotation | `test-vault/.claude/skills/source-annotation/SKILL.md` |

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
| `/query` | Synthesize answer from vault |
| `/park` | Shelve an idea, return to focus |
| `/focus` | Check for topic drift |
| `/next` | Break a topic into concrete first steps |
| `/project` | Manage project files |

---

## Installed Plugins

| Plugin | Zone | Commands |
|--------|------|----------|
| learning | `test-vault/dev/learning/` | `/learn`, `/review` |
| projects | `test-vault/dev/projects/` | `/project` |
