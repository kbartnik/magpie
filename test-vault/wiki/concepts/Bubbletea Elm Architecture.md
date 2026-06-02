---
title: "Bubbletea Elm Architecture"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Bubbletea Async Patterns", "Go Defined Types"]
tags: [go, bubbletea, tui, elm, preflight-sync-go]
---

# Bubbletea Elm Architecture

[Bubbletea](https://github.com/charmbracelet/bubbletea) is a Go TUI framework built on the **Elm Architecture** — a functional UI pattern with three parts and no shared mutable state.

## The three parts

### Model
A plain struct containing all application state. Never mutated in place — `Update` receives a copy and returns a new one.

```go
type AppModel struct {
    queue   []models.ScannedFile
    paused  bool
    mode    AppMode
    current *models.ScannedFile
}
```

### Update
A pure function. Receives a copy of the model and a message, returns the new model and an optional command. The bubbletea runtime replaces the old model with the returned one after every call.

```go
func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "p":
            m.paused = !m.paused  // modifying the copy
            return m, nil
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}
```

The value receiver (`m AppModel`, not `*AppModel`) guarantees `Update` never sees the original — every state change is explicit and traceable.

### View
A pure render function. Same model → same output. Called after every `Update`.

```go
func (m AppModel) View() string {
    if m.quitting {
        return "goodbye.\n"
    }
    return renderMainView(m)
}
```

## Messages

Everything that can happen is a message: key press, file done, progress tick, window resize. `Update` dispatches on message type using a **type switch**:

```go
switch msg := msg.(type) {
case tea.KeyMsg:     // msg is now tea.KeyMsg
case FileDoneMsg:    // msg is now FileDoneMsg — access msg.File, msg.Status
case ProgressMsg:    // msg is now ProgressMsg — access msg.Percent
}
```

`switch msg := msg.(type)` reassigns `msg` to the concrete type in each case. No explicit type assertion or cast needed.

## Commands

A `Cmd` is a `func() tea.Msg` — runs outside the UI goroutine and delivers its result as a message. The mechanism for any I/O that would block the display.

```go
func listenForProgress(ch chan models.ProgressMsg) tea.Cmd {
    return func() tea.Msg {
        msg, ok := <-ch
        if !ok {
            return nil
        }
        return msg
    }
}
```

## Key builtins

- `tea.Quit` — stops the program cleanly; bubbletea renders one final `View` frame before exiting
- `tea.EnterAltScreen` — returned from `Init()` to use the terminal's alternate screen buffer; preserves the user's terminal content and restores it on exit

```go
func (m AppModel) Init() tea.Cmd {
    return tea.EnterAltScreen
}
```

## Why no shared state

`Update` never runs concurrently with itself — bubbletea serializes all messages through a single goroutine. Goroutines that do background work (file copy, scan, watcher) communicate back via `p.Send(msg)` or by returning a `Cmd`, never by writing to model fields directly. This eliminates an entire class of data races at the UI layer.

*Related: [[Bubbletea Async Patterns]] (how Cmds compose for async work), [[Go Defined Types]] (AppMode, FileStatus)*
