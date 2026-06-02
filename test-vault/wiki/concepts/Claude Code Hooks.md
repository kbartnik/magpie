---
title: "Claude Code Hooks"
type: concept
status: active
created: 2026-05-20
updated: 2026-05-31
sources:
  - "archive/clippings/2026-05-20-hooks-explained.md"
  - "archive/clippings/2026-05-20-hooks-running-silently.md"
  - "archive/clippings/2026-05-29-hooks-5-automations.md"
  - "archive/clippings/2026-05-30-2026-04-26-claude-code-hooks-lifecycle.md"
  - "archive/clippings/2026-05-31-hooks-production-patterns.md"
related:
  - "Harness Engineering"
  - "Agentic Workflow Patterns"
  - "Nexus Vault Template"
tags:
  - claude-code
  - hooks
  - harness-engineering
  - cca-f
---

# Claude Code Hooks

Hooks are lifecycle control points that introduce **deterministic behavior** at specific moments in Claude Code's execution. They exist because some behaviors should not depend on model judgment at all.

> "You are relying on a probabilistic system to behave deterministically. That's the gap hooks fill."

This maps directly to the [[Agentic Workflow Patterns]] **Programmatic Prerequisites** pattern — hooks are that pattern implemented in Claude Code's runtime.

## The Core Pipeline

```
event → matcher → structured input → handler → outcome
```

Every hook is defined by these four elements:

**Event** — when in the lifecycle this fires (see table below)
**Matcher** — regex filter on `tool_name`; narrows when the hook runs. `Edit|Write`, `Bash`, `mcp__.*`. No matcher = runs every time for matcherless events.
**Structured input** — JSON receipt Claude sends to your handler: event type, tool name, tool inputs, file paths. Your handler reacts to a precise description, not a guess.
**Handler** — what actually runs.

## Five Events Worth Knowing First

| Event | When it fires | Primary use |
|-------|--------------|-------------|
| **SessionStart** | On session open | Initialise context, emit briefing |
| **PreToolUse** | Before a tool call executes | Intercept, validate, block |
| **PostToolUse** | Immediately after a tool call | React, update state, validate output |
| **Stop** | Before Claude returns its response | Final quality gate |
| **TaskCompleted** | When Claude considers the task done | Enforce definition of done |

The key question isn't "which event?" — it's "**at what point do I want to intervene?**"

## Four Handler Types

| Type | What it does | Best for |
|------|-------------|---------|
| `command` | Runs a local shell script | Mechanical checks (did file change? did tests run?) |
| `http` | POSTs event to external service | Slack notifications, webhooks, logging |
| `prompt` | Runs an LLM check inline | Semantic validation ("is the README actually accurate?") |
| `agent` | Delegates to a subagent | Complex review, multi-step validation |

### Mechanical vs. Semantic Checks

**Command hooks** handle mechanical questions: "Did README.md change?" "Did tests run?" "Was this file edited?"

**Prompt/agent hooks** handle semantic questions: "Is the README *actually up to date* with the changes?" "Does this implementation match the original request?" These pause execution and introduce an extra layer of reasoning — hooks become checkpoints for thinking, not just automation.

## Exit Codes: The Critical Distinction

**Exit code 2 = hard block.** The action does not run. Claude cannot override this.  
**Exit code 1 = logged warning.** Claude proceeds anyway. You have *logging*, not enforcement.

This is the most commonly misunderstood detail when building security hooks. A shell-script security gate that checks for `rm -rf` patterns but exits 1 on a match provides zero protection — it just logs the attempt. The block requires exit 2.

```bash
if echo "$COMMAND" | grep -qiE "$BLOCKED_PATTERNS"; then
  echo "Blocked: destructive command detected" >&2
  exit 2  # EXIT 2 = HARD BLOCK — the only real enforcement
fi
exit 0    # EXIT 0 = proceed
```

## Governance Ladder: Command → Prompt → Agent

Hooks form a 3-tier governance system. **Start at the lowest tier that works:**

| Tier | Handler type | Use when |
|------|-------------|----------|
| **Automated compliance** | `command` | The check is mechanical — does a pattern match? Did a file change? |
| **Judgment** | `prompt` | The check requires reasoning — "does this code match the intent?" |
| **Investigation** | `agent` | The check requires multi-step file access and analysis |

Agent hooks cost tokens. Graduate to prompt or agent only when a shell script genuinely can't do the check.

## "AI Checking AI" — The Governance Endgame

Agent Stop hooks implement a second AI that verifies the first AI's work with full codebase access:

```json
"Stop": [{
  "hooks": [{
    "type": "agent",
    "prompt": "You are a code reviewer. Check: (1) Do changes match the original intent? (2) Unintended side effects? (3) Tests passing? Return ok:false with reason if ANY check fails.",
    "model": "claude-3-5-haiku-20241022",
    "timeout": 120
  }]
}]
```

This isn't theoretical governance — it's practical, affordable (Haiku), and catches real bugs. The agent has read access to the entire codebase. Every team using AI assistants should adopt this pattern.

## Permission Audit Paradox

More permission prompts ≠ more safety. Excessive prompts exhaust attention and train autopilot approval — which is exactly when dangerous actions slip through. The smart permission filter pattern (auto-approve read-only tools, require manual review only for writes/deletes/network) reduces noise so humans focus on genuinely risky decisions.

## Four Safety Levels (hooks-running-silently)

When building hooks that modify or block actions, layer safety from lowest to highest:

1. **Dry-run mode** — show what *would* change without applying it
2. **Diff preview** — display the exact diff before applying
3. **Explicit approval gate** — require a `y/n` before any action
4. **Hard block** — return `decision: block` unconditionally; nothing runs

Start at level 1 when iterating; move to level 3-4 for production workflows that touch shared state.

## Minimal Hook JSON

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Edit|Write",
        "hooks": [
          {
            "type": "command",
            "command": "/path/to/check.sh"
          }
        ]
      }
    ]
  }
}
```

Hook config lives in `settings.json` — project-level (`.claude/settings.json`), personal (`.claude/settings.local.json`), or global (`~/.claude/settings.json`). Team-level non-negotiables go in project settings; personal preferences go in local. Hooks in this vault: `SessionStart` (briefing), `PostToolUse` (inbox count), `Stop` (session-end timestamp). See [[Nexus Vault Template]].

**Matcher syntax:** `Write|Edit` works; `Write | Edit` (with spaces) does not. Test hooks with `2>&1 | tee ~/.claude/hook-debug.log` if debugging silent failures. Hooks have a 60-second timeout.

## Five Practical Patterns

These move specific decisions from developer attention into configuration ("you decide once, Claude executes forever"):

**1. Auto-format on every write (PostToolUse)**
```json
{
  "hooks": {
    "PostToolUse": [{
      "matcher": "Write|Edit",
      "hooks": [
        {"type": "command", "command": "npx prettier --write \"$CLAUDE_TOOL_INPUT_FILE_PATH\""},
        {"type": "command", "command": "npx eslint --fix \"$CLAUDE_TOOL_INPUT_FILE_PATH\""}
      ]
    }]
  }
}
```
Multiple hooks in the same event run in parallel.

**2. Session context injection (SessionStart)**
```json
{
  "hooks": {
    "SessionStart": [{
      "hooks": [{"type": "command", "command": "echo '## Current State' && git status --short && echo '## Active TODOs' && grep -r 'TODO:' src/ --include='*.ts' | head -5"}]
    }]
  }
}
```
Claude sees uncommitted changes and active TODOs before the first message.

**3. Auto-approve trusted commands (PermissionRequest)**
```json
{
  "hooks": {
    "PermissionRequest": [{
      "matcher": "Bash(npm test*)",
      "hooks": [{"type": "command", "command": "echo '{\"decision\": \"approve\"}'"}]
    }]
  }
}
```
Returns structured JSON to bypass the permission dialog for pre-trusted commands.

**4. Run tests when Claude finishes (Stop)**
```json
{
  "hooks": {
    "Stop": [{
      "hooks": [{"type": "command", "command": "npm test --passWithNoTests 2>&1 | tail -20"}]
    }]
  }
}
```
Tests run automatically; failures appear immediately without needing to ask.

**5. Skill activation without prompting (UserPromptSubmit)**
```json
{
  "hooks": {
    "UserPromptSubmit": [{
      "hooks": [{"type": "command", "command": "node .claude/hooks/SkillActivationHook/skill-activation-prompt.mjs"}]
    }]
  }
}
```
Appends skill-activation guidance to prompts before Claude sees them — consistent behavior without relying on model memory.

## DS/ML Use Cases

Three patterns from data science workflows (via Jose Parreño):

**Experiment documentation (PostToolUse or Stop)**
After a training run, trigger a hook to extract key parameters, capture metrics/outputs, and update a README or experiment log. With Databricks MCP, this can verify that runs are properly logged and tagged. Stops documentation from being a manual afterthought.

**ML training sanity check (Stop, agent handler)**
After training scripts execute, run an agent hook that reviews the results for overfitting signals, metric reasonableness, and alignment with the original objective:
```json
{
  "hooks": {
    "Stop": [{
      "hooks": [{
        "type": "agent",
        "prompt": "Review the latest model results. Check whether the evaluation is sound, whether there are signs of overfitting, and whether the conclusions are justified."
      }]
    }]
  }
}
```
Combine with a matcher on training script file edits to avoid firing on every stop.

**Async workflow notifications (Stop or Notification)**
For long-running jobs (training, large data transformations), fire a Slack webhook or HTTP hook on completion. The system tells you it's done instead of you polling.

The underlying principle in all three: decisions that used to require developer attention — "did I document this run?", "did I check for overfitting?", "is it done yet?" — become enforced side effects of doing the work.

## CCA-F Connection

Hooks implement two patterns from [[Agentic Workflow Patterns]]:
- **Programmatic Prerequisites** — `PreToolUse` hooks that gate tool calls until conditions are met
- **Evaluator-Optimizer** (prompt/agent hooks) — semantic validation that loops back to the generator

The key exam distinction: hooks are *deterministic*; model instructions are *probabilistic*. If a behavior must always happen, hooks — not CLAUDE.md rules.

## See Also

- [[Harness Engineering]] — hooks are the harness's "architectural constraints" pillar
- [[Agentic Workflow Patterns]] — Programmatic Prerequisites pattern
- [[Claude Code Skills]] — skills and hooks compose; skills define what to do, hooks enforce when
- [[Nexus Vault Template]] — this vault's three hooks (SessionStart, PostToolUse, Stop)
- [[rtk]] — production-scale PreToolUse hook that rewrites Bash commands to token-efficient versions; real-world example of the pattern at scale
- [[Hard Deny Hooks Interaction Order]], [[CLAUDE.md Degradation Mechanism]], [[AI-Checking-AI Stop Hooks]], [[Vault Security Gate Exit Code]] — open questions on Claude Code behavior
