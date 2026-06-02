---
title: "Go File IO"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Go Error Handling", "Go Testing Patterns"]
tags: [go, io, filesystem, preflight-sync-go]
---

# Go File IO

Core patterns for reading, writing, and walking the filesystem in Go.

## filepath.WalkDir

Recursively walks a directory tree, calling a function for each entry:

```go
filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err       // propagate walk errors
    }
    if d.IsDir() {
        return nil       // skip directories, continue walking
    }
    // process file at path
    return nil
})
```

Callback return value controls walk behavior:
- `nil` — no error, continue
- `error` — stop walking, return error from `WalkDir`
- `fs.SkipDir` — skip current directory, continue others
- `fs.SkipAll` — stop walking entirely, return nil (Go 1.20+)

`d.Info()` returns `(fs.FileInfo, error)` — efficient because `WalkDir` often has it cached from the directory read.

## os.Stat

Returns metadata about a file without opening it:

```go
info, err := os.Stat(path)
if err != nil {
    // file doesn't exist, permission denied, etc.
}
info.Size()  // file size in bytes
info.IsDir() // true if directory
```

Two `os.Stat` calls with a sleep between them is the standard stability check pattern — if size is unchanged, the file is probably done being written.

## os.MkdirAll

Creates a directory and all necessary parents — equivalent to `mkdir -p`:

```go
if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
    return err
}
```

Always call before `os.Create` when the destination path may not exist. `0o755` is the standard directory permission (owner read/write/execute, group and others read/execute).

## defer for cleanup

`defer` registers a call to run when the surrounding function returns, regardless of how:

```go
f, err := os.Open(path)
if err != nil {
    return err
}
defer f.Close()  // runs when the function exits, whatever happens
```

For read-only files, discarding the close error is acceptable:

```go
defer func() { _ = f.Close() }()  // explicit discard — linter accepts this
```

Multiple `defer` calls run in LIFO order.

## Named returns and defer for close errors

For write operations, `Close()` errors are meaningful — they can indicate data wasn't flushed (especially over NFS). Use a named return to capture close errors:

```go
func writeFile(path string) (err error) {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer func() {
        if cerr := f.Close(); cerr != nil && err == nil {
            err = cerr  // only set if no other error already occurred
        }
    }()
    // ... write ...
    return nil
}
```

The named return `(err error)` makes the return variable accessible inside the deferred function.

## Closures for accumulation

Anonymous functions capture variables from their surrounding scope. `filepath.WalkDir` uses this pattern to collect results:

```go
results := make([]ScannedFile, 0)

filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
    // results is captured from the outer scope
    results = append(results, file)
    return nil
})

// results contains everything the closure accumulated
```

The closure doesn't receive `results` as a parameter — it accesses it directly from the enclosing scope. Changes inside the closure are visible outside.

## NFS flushing: dst.Sync() before close

On NFS, writes go into the kernel page cache first. If you rely on `Close()` to flush, the flush is silent — the copy appears to finish at 100% and then the program hangs with no feedback.

Calling `dst.Sync()` before `Close()` makes the flush an explicit step you can report:

```go
progressCh <- models.ProgressMsg{Percent: 100, State: models.CopyStateSyncing}
if err := destFile.Sync(); err != nil {
    return &models.TransferError{Filename: dest, Err: err}
}
```

The user sees "syncing..." instead of a frozen progress bar at 100%.

*Related: [[Go Error Handling]] (named returns, error wrapping), [[Go Testing Patterns]] (testdata fixtures, t.TempDir for temp files)*
