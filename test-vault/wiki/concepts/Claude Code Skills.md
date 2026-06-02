---
title: "Claude Code Skills"
type: concept
status: active
created: 2026-05-20
updated: 2026-05-30
sources:
  - "archive/docs/2026-06-01-complete-guide-building-skills-claude.md"
  - "archive/clippings/2026-05-20-agent-skills-2-0.md"
  - "archive/clippings/2026-05-20-mastering-agentic-skills.md"
  - "archive/clippings/2026-05-20-skills-deep-dive-part-2.md"
  - "archive/clippings/2026-05-20-skills-hidden-gems.md"
  - "archive/clippings/2026-05-29-skills-conceptual-deep-dive.md"
  - "archive/clippings/2026-05-29-skills-deep-dive-part-1.md"
  - "archive/clippings/2026-05-29-skills-deep-dive-part-2.md"
  - "archive/clippings/2026-05-29-project-memory-skill.md"
  - "archive/clippings/2026-05-30-2026-04-29-matt-pocock-skills-collection.md"
  - "archive/clippings/2026-05-30-2026-05-17-claude-code-skills-anatomy.md"
related:
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
  - "Claude Code Hooks"
  - "Nexus Vault Template"
  - "Progressive Disclosure Architecture"
tags:
  - claude-code
  - skills
  - harness-engineering
  - cca-f
---

# Claude Code Skills

Skills are packaged, reusable agent capabilities — markdown files (or directories) that teach Claude a workflow, enforce a process, or inject domain knowledge. As of Skills 2.0, they are programs, not instructions.

> "Skills should be **orchestrators** that load what they need, not encyclopedias that carry everything."

## Official Anthropic Spec (2025)

Source: [[Complete Guide to Building Skills for Claude]].

### File Structure

```
your-skill-name/
├── SKILL.md          # required; must be exactly this name (case-sensitive)
├── scripts/          # optional; Python, Bash, etc.
├── references/       # optional; documentation loaded on demand
└── assets/           # optional; templates, fonts, icons
```

**Critical naming rules:** `SKILL.md` only — `skill.md` and `SKILL.MD` are not accepted. Folder names: kebab-case only, no spaces or capitals.

### YAML Frontmatter Spec

```yaml
---
name: your-skill-name          # required; kebab-case; matches folder name
description: |                 # required; under 1024 chars; no XML tags
  What it does. Use when user asks to [specific trigger phrases].
license: MIT                   # optional
compatibility: Claude Code     # optional; environment requirements
metadata:                      # optional; custom key-value pairs
  author: you
  version: 1.0.0
---
```

**The description field is the most important part.** It must include both *what the skill does* and *when to use it* (specific trigger phrases). Structure: `[What it does] + [When to use it] + [Key capabilities]`.

Good description:
```
Analyzes Figma design files and generates developer handoff documentation. 
Use when user uploads .fig files, asks for "design specs", "component 
documentation", or "design-to-code handoff".
```

Bad descriptions: "Helps with projects." (too vague), "Creates sophisticated documentation systems." (no triggers).

**Security:** No XML angle brackets in frontmatter (it appears in the system prompt). Names containing "claude" or "anthropic" are reserved.

### Use Case Categories

| Category | Use for | Key pattern |
|----------|---------|-------------|
| **Document & Asset Creation** | Consistent, high-quality output (docs, code, designs) | Embedded style guides, quality checklists, templates |
| **Workflow Automation** | Multi-step processes, cross-MCP coordination | Step-by-step gates, iterative refinement loops |
| **MCP Enhancement** | Workflow guidance on top of MCP tool access | Coordinates MCP calls, embeds domain expertise |

### Quality Metrics

**90% trigger rate target:** run 10–20 test queries that should activate the skill; if it triggers on fewer than 90%, refine the frontmatter description.

**Qualitative bar:** users shouldn't need to redirect Claude mid-workflow; running the same request 3–5 times should yield structurally consistent output.

Note: Anthropic acknowledges evaluation is partly "vibes-based" — measurement tooling is still in development.

### Skills + MCP Architecture

MCP provides the professional kitchen: access to tools, data, and external services.
Skills provide the recipes: workflows and best practices for using that access.

Without skills, MCP users must figure out workflows themselves and prompt inconsistently. Skills embed the domain expertise so every invocation follows best practices automatically.

## Skills 2.0: What Changed

Commands and skills are now unified. Files in `.claude/commands/` still work; `.claude/skills/` is the recommended path for richer capabilities.

| Capability | Description |
|-----------|-------------|
| **Subagent execution** (`context: fork`) | Skill runs in an isolated 200K-token context window; main conversation stays clean |
| **Dynamic context injection** (`!` backtick) | Shell commands run at load time; Claude receives live output, not stale text |
| **Granular permissions** | Control who invokes (human/Claude/both), which tools are accessible, which model runs |
| **Lifecycle hooks** | Skills can attach hook behaviors |

## Progressive Disclosure Architecture (PDA)

See [[Progressive Disclosure Architecture]] for the full model. The summary:

```
Tier 1: Metadata scan (~100 tokens/skill)   — frontmatter at session start
Tier 2: SKILL.md (1–5K tokens)              — loaded when skill is selected
Tier 3: References + scripts (0 tokens)     — loaded on-demand by Tier 2 instructions
```

Skills don't dump everything into context at once. Tier 3 resources load only when the skill explicitly reads them. A 50KB reference library costs 0 tokens until the skill decides it's needed.

**The orchestrator rule:** a skill should *point to* reference material, not *contain* it. Skills that violate this become encyclopedias that bloat every invocation.

**Token savings:** PDA reduces context consumption 78–93% vs. loading all documentation upfront. A 50KB monolithic skill handling 5 diagram types costs 250KB across 5 requests; the PDA equivalent costs 44KB.

## Dynamic Context Injection

The `!` backtick syntax runs shell commands before the prompt reaches Claude:

```markdown
Current branch: `! git branch --show-current`
Open files: `! git status --short`
```

Claude receives the live output. This is the skill equivalent of [[Claude Code Hooks]]' structured input — live data rather than stale instructions.

## Subagent Execution

```yaml
---
context: fork
---
```

With `context: fork`, the skill's content becomes the task prompt for a fresh subagent with its own context window. The main conversation receives only the result. This is the [[Agentic Workflow Patterns]] **hub-and-spoke** pattern applied to skills: isolated context, typed output back to coordinator.

## Bundled Skills (ship with Claude Code)

| Skill | What it does |
|-------|-------------|
| `/simplify` | Spawns 3 parallel review agents: code reuse, quality, efficiency |
| `/batch` | Decomposes large changes into 5–30 units, presents plan, spawns one background agent per unit in an isolated git worktree |
| `/debug` | Troubleshoots misconfigured MCP, failing tool calls |
| `/claude-api` | Loads Claude API reference for the project's language |

`/batch` is the [[Agentic Workflow Patterns]] **parallelization + orchestrator-workers** pattern as a production-ready bundled skill.

## Cross-Platform Standard

Agentic skills are now a cross-platform open standard. Claude Code, Codex, GitHub Copilot CLI, OpenCode, Cursor, Gemini CLI, and 14+ other agents support the same `SKILL.md` format. Skills written for Claude Code run on other agents with no changes. The Agent Skill Standard at agentskills.io governs the cross-platform spec.

## mattpocock/skills: Community Reference Collection

[mattpocock/skills](https://github.com/mattpocock/skills) is Matt Pocock's personal `.claude` directory made public — 21 skills organized across planning, development, tooling, and knowledge management. 20K+ GitHub stars. Key skills:

| Skill | What it enforces |
|-------|-----------------|
| `tdd` | Red-Green-Refactor cycle; blocks implementation before failing test |
| `git-guardrails-claude-code` | Requires human approval for destructive git operations |
| `write-a-prd` | Interactive PRD interview → GitHub issue, not document-filling |
| `grill-me` | Design review via structured interrogation before implementation |
| `triage-issue` | Root-cause investigation before any fix attempt |

**Install individual skills:**
```bash
npx skills@latest add mattpocock/skills/tdd
npx skills@latest add mattpocock/skills/git-guardrails-claude-code
```

The collection distinguishes itself from capability-extension skill packs (browser control, API access) by focusing entirely on **workflow enforcement** — making Claude Code a participant in your team's existing process rather than a one-shot command executor.

## Project Memory: A First-Skill Pattern

The simplest useful skill: structured markdown files in `docs/project_notes/` that the agent checks before making decisions:

- `bugs.md` — bug log with root cause, solution, prevention (guards against repeat failures)
- `decisions.md` — ADRs with context, decision, alternatives, consequences
- `key_facts.md` — ports, URLs, identifiers (never secrets)
- `issues.md` — work log with ticket references

Key insight: name it `docs/project_notes/` not `ai-memory/` — it looks like engineering documentation, so humans maintain it. This is a **basic skill** (under 5KB, no PDA needed), useful precisely because of its simplicity and longevity.

## Writing Reliable Skills (Production Principles)

From "What the docs don't tell you":

- **Skills are a design pattern, not a feature.** Treat them like a tested interface, not a one-off prompt.
- **Fail explicitly.** A skill that silently does the wrong thing is worse than one that errors.
- **One skill, one job.** Composability beats monoliths. Small skills that call each other are more testable.
- **Feedback loops.** For quality-critical tasks, build evaluation into the skill — not just instructions.

## This Vault's Skill Structure

This vault uses `.claude/skills/` (session-context, source-annotation) and `.claude/commands/` (ingest, query, lint, etc.). The plugin system extends this — each plugin ships a `skill.md` that loads on plugin discovery. See [[Nexus Vault Template]].

## CCA-F Connection

- PDA maps to **context engineering** — managing what the model sees per call
- `context: fork` maps to **hub-and-spoke** isolation — subagents with blank context windows
- Bundled `/batch` maps to **orchestrator-workers + parallelization** — the full pattern in one skill

## See Also

- [[Progressive Disclosure Architecture]] — the full PDA model with token math, reference organization strategies, and the AI resilience layer
- [[Claude Code Hooks]] — hooks and skills compose; hooks enforce *when*, skills define *what*
- [[Harness Engineering]] — skills are the "capability acquisition" layer of the harness
- [[Agentic Workflow Patterns]] — context:fork = hub-and-spoke; /batch = orchestrator-workers
- [[Nexus Vault Template]] — this vault's skill and command architecture
- [[CLAUDE.md Configuration Patterns]] — complements skills by configuring LLM behavior (skills = what; CLAUDE.md = how the LLM approaches it)
