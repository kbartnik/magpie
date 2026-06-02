---
title: "Tmux"
type: concept
status: active
created: 2026-05-30
updated: 2026-05-30
sources:
  - "archive/clippings/2026-05-30-tmux-ninja.md"
  - "archive/books/2026-06-01-tmux-2.md"
related: []
tags:
  - tmux
  - developer-tools
  - terminal
---

# Tmux

**Terminal Multiplexer.** Manages persistent terminal sessions that survive SSH disconnects, with windows (tabs) and panes (splits) inside each session.

## Mental Model

```
Server (background process)
└── Session (named workspace, persists after detach)
    ├── Window 1 (like a tab) — active window has * in status bar
    │   ├── Pane A
    │   └── Pane B
    └── Window 2
```

The **prefix key** is `Ctrl+b`. Every keybinding is `prefix → key`.

---

## Sessions

| Action | Command |
|--------|---------|
| New named session | `tmux new -s <name>` |
| List sessions | `tmux ls` |
| Attach to session | `tmux attach -t <name>` |
| Detach (keep running) | `prefix → d` |
| List sessions (inside) | `prefix → s` |
| Rename session | `tmux rename-session -t <old> <new>` |
| Kill session | `tmux kill-session -t <name>` |

---

## Windows (Tabs)

| Action | Binding |
|--------|---------|
| New window | `prefix → c` |
| Rename window | `prefix → ,` |
| Next window | `prefix → n` |
| Previous window | `prefix → p` |
| List all windows | `prefix → w` (navigate with arrows) |
| Kill highlighted window | `prefix → w` then `x` |

---

## Panes (Splits)

| Action | Binding |
|--------|---------|
| Split horizontal (top/bottom) | `prefix → "` |
| Split vertical (left/right) | `prefix → %` |
| Switch pane | `prefix → arrow keys` |
| Zoom pane (toggle fullscreen) | `prefix → z` |
| Show pane numbers | `prefix → q` (press number to jump) |
| Resize pane | `prefix → Alt+arrow` |
| Kill pane | `exit` or `prefix → x` |

---

## Copy Mode

Scroll and copy terminal output vim-style.

| Action | Binding |
|--------|---------|
| Enter copy mode | `prefix → [` |
| Exit copy mode | `q` |
| Start selection | `Ctrl+Space` then arrow keys |
| Copy selection | `Alt+w` |
| Paste | `prefix → ]` |

Enable vi keybindings in `~/.tmux.conf`:

```
set-window-option -g mode-keys vi
bind h select-pane -L
bind j select-pane -D
bind k select-pane -U
bind l select-pane -R
```

---

## Other Bindings

| Action | Binding |
|--------|---------|
| Command prompt | `prefix → :` |
| Show all keybindings | `prefix → ?` |
| Kill server (all sessions) | `prefix → :` then `kill-server` |
| Show time | `prefix → t` |

---

## Configuration (`~/.tmux.conf`)

```bash
# Enable mouse (resize panes, click to select)
set -g mouse on

# Vi keybindings in copy mode
set-window-option -g mode-keys vi
bind h select-pane -L
bind j select-pane -D
bind k select-pane -U
bind l select-pane -R
```

Reload config without restarting:

```bash
tmux source ~/.tmux.conf
# or inside a session: prefix → : then "source ~/.tmux.conf"
```

---

## Scripted Environments

Source: [[tmux 2]] (Hogan, 2016). The real productivity multiplier: define the full development environment for a project once, then launch it in one command.

### Shell Script Approach

```bash
#!/usr/bin/env bash
SESSION="myproject"
tmux new-session -d -s $SESSION          # create detached session

tmux rename-window -t $SESSION:0 'editor'
tmux send-keys -t $SESSION:0 'nvim .' Enter

tmux new-window -t $SESSION -n 'server'
tmux send-keys -t $SESSION:1 'npm run dev' Enter

tmux new-window -t $SESSION -n 'tests'
tmux send-keys -t $SESSION:2 'npm test -- --watch' Enter

tmux select-window -t $SESSION:0        # focus editor window
tmux attach-session -t $SESSION
```

Run once; from then on just `tmux attach -t myproject`.

### Tmuxinator

Manages per-project tmux environments via YAML configs in `~/.tmuxinator/`:

```yaml
# ~/.tmuxinator/myproject.yml
name: myproject
root: ~/Source/myproject

windows:
  - editor:
      layout: main-vertical
      panes:
        - nvim .
        - git log --oneline -10
  - server: npm run dev
  - tests: npm test -- --watch
```

```bash
tmuxinator start myproject     # launch environment
tmuxinator stop myproject      # kill it
tmuxinator list                # list all configs
```

The YAML config is the project's development environment as code — commit it to the repo.

## Pair Programming

### Shared Account (Read-Write)

Two users SSH into the same machine with the same account; both attach to the same session:

```bash
# User 1 creates the session
tmux new-session -s pairing

# User 2 attaches (both see same screen, both can type)
tmux attach-session -t pairing
```

**Problem:** only one active window, full screen-share equivalent.

### Grouped Sessions (Independent Windows)

Better approach: both users see the same session but can independently browse different windows:

```bash
# User 1 creates the session
tmux new-session -s main

# User 2 creates a grouped session (shares windows, independent current window)
tmux new-session -t main -s pair
```

Both see the same windows list but each has their own active window pointer. They collaborate on a pane only when both deliberately switch to the same window. This is better than full screen-share for async pairing.

### tmate (Quick Remote Pairing)

```bash
brew install tmate
tmate              # starts a session; outputs a URL and SSH command for your pair
```

tmate proxies through tmate.io — no shared server or SSH config needed. The pair gets a read-write link; a read-only link is also available.

## Plugin Manager (TPM)

```bash
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
```

Add to **bottom** of `~/.tmux.conf`:

```bash
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'
run '~/.tmux/plugins/tpm/tpm'
```

Install plugins: `prefix → I` (capital I) inside a session.

### Recommended Plugins

| Plugin | Purpose |
|--------|---------|
| `tmux-plugins/tmux-sensible` | Sane defaults (UTF-8, larger history, etc.) |
| `catppuccin/tmux` | Clean status bar theme |

Minimal config with TPM + Catppuccin:

```bash
set -g mouse on

set -g @plugin 'catppuccin/tmux#v2.1.3'
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'

run '~/.tmux/plugins/tpm/tpm'
```

## See Also

- [[LazyVim]] — Neovim editor that pairs with tmux; covers when the built-in Neovim terminal is sufficient vs. when tmux is needed for session detachment
- [[Progressive Disclosure Architecture]] — tmuxinator YAML configs apply the same principle: define the environment declaratively, load it on demand
