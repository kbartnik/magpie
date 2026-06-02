---
title: "Fsnotify Filesystem Watching"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Bubbletea Async Patterns", "Go Channel Concurrency Patterns"]
tags: [go, fsnotify, filesystem, concurrency, preflight-sync-go]
---

# Fsnotify Filesystem Watching

`github.com/fsnotify/fsnotify` is the standard cross-platform filesystem event library for Go. It wraps OS-specific APIs (inotify on Linux, FSEvents on macOS, ReadDirectoryChangesW on Windows) behind a unified interface.

## Basic watcher setup

```go
watcher, err := fsnotify.NewWatcher()
if err != nil {
    return
}
defer func() { _ = watcher.Close() }()

watcher.Add("/path/to/watch")

for {
    select {
    case event, ok := <-watcher.Events:
        if !ok {
            return
        }
        if event.Has(fsnotify.Create) {
            // new file at event.Name
        }
    case <-watcher.Errors:
        // ignore or log
    case <-ctx.Done():
        return
    }
}
```

Event types: `Create`, `Write`, `Remove`, `Rename`, `Chmod`. Only `Create` is relevant for detecting new files. `event.Name` is the full path to the affected file.

`watcher.Close()` returns an error — discard explicitly with `_ =` to satisfy `errcheck`.

## Connecting to bubbletea: two-message async pattern

When the watcher sends a `Create` event, don't run the stability check directly in `Update` — that would block the UI. Instead, use the two-message async pattern:

1. Watcher goroutine calls `p.Send(NewFileDetectedMsg{})` — non-blocking, fire and forget
2. `Update` handles `NewFileDetectedMsg` → returns a `Cmd` that runs the stability check off the UI goroutine
3. `Cmd` completes → returns `NewFileReadyMsg` → `Update` handles the result, adds file to queue

```go
// First handler — kicks off the blocking work
case NewFileDetectedMsg:
    return m, checkStability(msg.File, pause)

// Cmd factory — runs off UI goroutine
func checkStability(file models.ScannedFile, pause time.Duration) tea.Cmd {
    return func() tea.Msg {
        if scan.CheckStability(file.Path, pause, false, time.Sleep) {
            return NewFileReadyMsg{File: file}
        }
        return nil  // silently drop unstable files
    }
}

// Second handler — acts on the result
case NewFileReadyMsg:
    m.queue = append(m.queue, msg.File)
    // dispatch if worker idle...
```

## Graceful shutdown

The watcher goroutine selects on both `watcher.Events` and a stop signal. With `context.Context`:

```go
select {
case event, ok := <-watcher.Events:
    if !ok { return }
    // handle event
case <-ctx.Done():
    return
}
```

With a done channel (`chan struct{}`):

```go
select {
case event, ok := <-watcher.Events:
    if !ok { return }
case <-stopCh:
    return
}
```

Both patterns let the goroutine exit cleanly when the application quits without leaking goroutines.

*Related: [[Bubbletea Async Patterns]] (two-message pattern, p.Send), [[Go Channel Concurrency Patterns]] (select, channel close)*
