---
title: "Obsidian Plugin"
type: backlog-item
project: magpie
status: todo
priority: low
due: ""
created: 2026-06-03
updated: 2026-06-03
tags: [research, obsidian, plugin]
---

# Obsidian Plugin

Research whether an Obsidian plugin that interfaces with magpie is worth building, and what shape it would take.

## Possibilities to evaluate

- **Command palette integration** — expose magpie operations (capture, inbox, archive) as Obsidian commands/hotkeys, eliminating terminal context switching
- **Live graph views** — surface orphan pages, stale notes, broken wikilinks as interactive views rather than CLI lint output
- **Contextual capture modal** — capture with immediate context: drop directly into a project, tag, link, without leaving the current document
- **Sync status overlay** — show preflight-sync-go state in the status bar; flag conflicted files inline
- **Bidirectional state** — Obsidian knows active file and cursor position; magpie knows vault structure and project state. Possibilities: surface next-action when opening a project note, offer to update the learning track index after editing a wiki page

## Prerequisites

- Magpie needs a stable API surface before a plugin is worth designing against
- Worth revisiting after the core CLI phases are complete
