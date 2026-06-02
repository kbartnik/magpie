---
title: "Bubbletea Async Patterns"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Bubbletea Elm Architecture", "Fsnotify Filesystem Watching"]
tags: [go, bubbletea, tui, async, concurrency, preflight-sync-go]
---

# Bubbletea Async Patterns

Patterns for handling background work, long-running operations, and event streams in a bubbletea TUI without blocking the UI goroutine.

## The two-message async pattern

When a bubbletea event triggers a **blocking operation** (stability check, network call, file stat), it can't run directly in `Update` — that would freeze the display. The pattern:

1. First message arrives → `Update` returns a **Cmd** that runs the blocking work off the UI goroutine
2. The Cmd completes → returns a second message → `Update` handles the result

```go
// First handler — kicks off the work
case NewFileDetectedMsg:
    pause := time.Duration(cfg.StabilityPauseSecs * float64(time.Second))
    return m, checkStability(msg.File, pause)

// Cmd factory — runs off the UI goroutine
func checkStability(file models.ScannedFile, pause time.Duration) tea.Cmd {
    return func() tea.Msg {
        if scan.CheckStability(file.Path, pause, false, time.Sleep) {
            return NewFileReadyMsg{File: file}
        }
        return nil  // drop unstable files silently
    }
}

// Second handler — acts on the result
case NewFileReadyMsg:
    m.queue = append(m.queue, msg.File)
    // ...
```

This keeps `Update` non-blocking while preserving a sequential async flow driven by messages.

## Progress streaming with listenForProgress

For operations that produce a stream of updates (chunked file copy), a per-file channel carries progress messages. A `Cmd` listens on the channel and re-subscribes after each message:

```go
// Cmd that blocks waiting for the next progress event
func listenForProgress(ch chan models.ProgressMsg) tea.Cmd {
    return func() tea.Msg {
        msg, ok := <-ch
        if !ok {
            return nil  // channel closed — copy finished
        }
        return msg
    }
}

// In Update — re-subscribe after each progress message
case models.ProgressMsg:
    m.progress = msg.Percent
    m.copyState = msg.State
    return m, listenForProgress(m.currentProgressCh)  // subscribe again
```

Each `Cmd` handles exactly one event, then the next `Update` call subscribes again. This is how a channel-based event stream maps onto bubbletea's message loop.

## Worker + channel dispatch

Background goroutines communicate back to the TUI via `p.Send(msg)` (fire and forget) or by returning messages from Cmds. The worker pattern:

```go
// Cmd that sends a file to the worker channel
func dispatchFile(workCh chan models.ScannedFile, file models.ScannedFile) tea.Cmd {
    return func() tea.Msg {
        workCh <- file
        return nil  // no message back — worker sends its own messages via p.Send
    }
}
```

The worker goroutine calls `p.Send(WorkerReadyMsg{})`, `p.Send(FileDoneMsg{...})` etc. directly. These arrive in `Update` just like any other message.

## WorkerReadyMsg as flow control

A `WorkerReadyMsg` signals that the worker is ready for the next file. The TUI maintains an explicit `workerIdle bool` field set `false` on dispatch and `true` on `WorkerReadyMsg`. This closes a race window where `FileDoneMsg` (which clears `current`) arrives before `WorkerReadyMsg`, making the worker appear idle when it isn't.

```go
case WorkerReadyMsg:
    m.workerIdle = true
    if len(m.queue) > 0 && !m.paused {
        return m.tryDispatchNext()
    }
    // ...
```

## Async scan with streaming messages

`ScanFiles` runs in a goroutine launched from `Init`, sending one `NewFileReadyMsg` per discovered file. The TUI opens immediately with an empty queue and populates as files arrive — no blocking upfront scan.

```go
func (m AppModel) Init() tea.Cmd {
    return tea.Batch(
        tea.EnterAltScreen,
        func() tea.Msg {
            go func() {
                files, _ := scan.ScanFiles(cfg)
                for _, f := range files {
                    p.Send(NewFileReadyMsg{File: f})
                }
            }()
            return nil
        },
    )
}
```

*Related: [[Bubbletea Elm Architecture]] (the Cmd mechanism), [[Fsnotify Filesystem Watching]] (watcher feeds the same NewFileReadyMsg path)*
