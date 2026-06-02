---
title: "Go Testing Patterns"
type: concept
status: active
created: 2026-05-29
updated: 2026-05-29
sources: []
related: ["Go Error Handling", "Go File IO"]
tags: [go, testing, tdd, testify, preflight-sync-go]
---

# Go Testing Patterns

Go's testing package is built into the standard library. Test functions are discovered by the runner — no registration framework needed.

## Test function signature

```go
func TestLoadConfig_MissingFile(t *testing.T) {
```

Convention: `TestFunctionName_CaseName`. Must start with `Test`, take `*testing.T`, live in `_test.go` files.

## t.Fatalf vs t.Errorf

The distinction is about whether continuing after failure would be meaningful:

- `t.Fatalf` — stops the test immediately. Use when subsequent lines would panic or produce meaningless results (setup steps, loading fixtures, any step where the rest of the test depends on the result).
- `t.Errorf` — records a failure and keeps running. Use when remaining assertions are still independent and useful.

```go
cfg, err := config.Load("testdata/valid.toml")
if err != nil {
    t.Fatalf("unexpected error: %v", err)  // cfg is nil if we continue
}
if len(cfg.MediaTypes) != 2 {
    t.Errorf("expected 2 types, got %d", len(cfg.MediaTypes))
}
```

Never discard errors from the function under test with `_` — a failure produces a panic instead of a clean message.

## testify

`github.com/stretchr/testify` provides two packages:

- `assert` — records a failure and continues (like `t.Errorf`)
- `require` — stops immediately on failure (like `t.Fatalf`)

```go
// stdlib
if got != want {
    t.Errorf("expected %v, got %v", want, got)
}

// testify/assert — continues on failure
assert.Equal(t, want, got)

// testify/require — stops immediately
require.Equal(t, want, got)
```

Rule of thumb: `require` for setup steps where continuing would panic; `assert` for independent value assertions.

## Table-driven tests

The idiomatic Go approach to parametrized tests:

```go
func TestClassifyFile(t *testing.T) {
    cases := []struct {
        name  string
        input string
        want  string
    }{
        {"show", "Show (2024) - S01E01 - Pilot.mkv", "show"},
        {"movie", "Movie (2024).mkv", "movie"},
        {"no match", "readme.txt", ""},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := ClassifyFile(tc.input)
            assert.Equal(t, tc.want, got)
        })
    }
}
```

`t.Run` creates a named subtest. Failure output reads `TestClassifyFile/show` — the failing case is immediately identifiable. Adding a new case is one line in the struct slice.

## t.Cleanup

Registers a function that runs after the test completes, regardless of pass or fail:

```go
t.Cleanup(func() {
    os.WriteFile(path, []byte("content"), 0644)  // restore fixture file
})
```

Preferred over `defer` in tests because it works correctly in helper functions called from the test. Multiple `t.Cleanup` calls run in LIFO order.

## Test fixtures

Each package that needs test fixtures gets a `testdata/` subdirectory alongside its `_test.go` files. Go's test runner sets the working directory to the package directory, so `testdata/foo.toml` resolves relative to the package, not the repo root.

## Dependency injection via function parameters

Go's idiomatic alternative to mocking frameworks — pass a function as a parameter:

```go
func CheckStability(path string, pause time.Duration, skip bool,
    sleep func(time.Duration)) bool {
    // ...
    sleep(pause)  // caller decides what "sleep" means
}

// Production
CheckStability(path, 2*time.Second, false, time.Sleep)

// Test — no-op sleep
CheckStability(path, 0, false, func(time.Duration) {})

// Test — verify sleep was called
sleepCalled := false
CheckStability(path, 0, false, func(time.Duration) { sleepCalled = true })
```

A single function with a known signature needs no interface, no mock type, no framework.

## TDD discipline

**Red before green:** a test that passes before its implementation exists is not a red test — it's a broken test. Always confirm a test fails for the right reason before writing the implementation.

**Choosing what to test first:** ask before writing a test — *can this fail against a stub that returns zero values?*

- `len(files) != 0` against a nil return — passes trivially, not a useful red test
- `len(files) != 2` against a nil return — fails immediately, useful red test

When the simplest test is also the least discriminating, start with a test that has real fixture data and concrete assertions. Come back to degenerate cases (empty input, error paths) once the implementation exists — at that point they guard against regressions.

*Related: [[Go Error Handling]] (never discard errors in tests), [[Go File IO]] (testdata fixtures, t.TempDir)*
