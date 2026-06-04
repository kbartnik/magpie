# CLAUDE.md — magpie

> **Transitional note:** Skills and commands currently use `vault-tools` (Nexus binary, copied to `test-vault/.claude/tools/`). Phase 8 (Migration) replaces all `vault-tools` calls with `magpie` subcommands and removes the copied binary.

---

## Testbed vs. Reference

`test-vault/` is a **testbed** — its purpose is exercising vault command workflows and validating magpie CLI behavior. It is not a reference knowledge vault.

- **Nexus** (`~/Documents/Obsidian/Nexus`) is the reference vault: real knowledge, real decisions, real projects
- **test-vault** is the testbed: synthetic scenarios, magpie-specific content, and workflow validation

Content created in the test vault is either:
1. Magpie-specific (ADRs about magpie's design, investigations into magpie's behavior, code docs for the magpie repo)
2. Synthetic test scenarios (used to exercise a command workflow, not treated as authoritative knowledge)

When a test produces broadly applicable insight, surface it for Nexus rather than keeping it here.

---

## Vault

The test vault lives at `test-vault/` in this repo. It is a fully functional magpie vault with `.magpie/` sentinel, wiki, archive, and dev zones for all plugins.

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
| learning | `test-vault/dev/learning/` | magpie-specific learning tracks |
| projects | `test-vault/dev/projects/` | (empty — projects moved to Nexus) |
| adr | `test-vault/dev/adr/` | magpie design decisions + workflow testing |
| code | `test-vault/dev/code/` | magpie repo documentation (`/code magpie`) |
| debrief | `test-vault/dev/debriefs/` | magpie incidents + workflow testing |
| investigation | `test-vault/dev/investigations/` | magpie behavior questions + workflow testing |
