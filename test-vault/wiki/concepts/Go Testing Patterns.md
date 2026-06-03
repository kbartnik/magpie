---
title: "Go Testing Patterns"
type: concept
status: active
created: 2026-05-29
updated: 2026-06-02
sources:
  - dev/learning/magpie-go/2026-06-02.md
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

`github.com/stretchr/testify` provides two packages with identical APIs:

- `assert` — records failure and continues (like `t.Errorf`)
- `require` — stops immediately on failure (like `t.Fatalf`)

```go
import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)
```

Rule of thumb: `require` for setup/preconditions where continuing would panic or produce meaningless output; `assert` for independent value checks.

### Common assertions

```go
assert.Equal(t, want, got)          // deep equality
assert.NotEqual(t, a, b)

assert.Nil(t, err)
assert.NotNil(t, cfg)

assert.True(t, ok)
assert.False(t, cfg.Verbose)

assert.Len(t, results, 3)           // checks len(results) == 3
assert.Empty(t, list)               // len == 0, nil, "", etc.
assert.NotEmpty(t, list)

assert.Contains(t, str, "substr")   // works on strings, slices, maps
assert.ElementsMatch(t, want, got)  // same elements, any order
```

### Error assertions

```go
require.NoError(t, err)                          // fails + stops if err != nil
assert.Error(t, err)                             // fails if err == nil

assert.ErrorIs(t, err, fs.ErrNotExist)           // unwraps error chain
assert.ErrorAs(t, err, &pathErr)                 // like errors.As

assert.EqualError(t, err, "config: missing url") // exact message
```

`require.NoError` is the most common setup guard — after loading a fixture or calling an initialiser, if it errors the rest of the test is meaningless.

### Suite — shared setup and teardown

`testify/suite` provides xUnit-style setup/teardown via a struct. The struct fields hold shared state; the lifecycle methods initialise and clean it up.

```go
import "github.com/stretchr/testify/suite"

type LoadSuite struct {
    suite.Suite
    cfgPath string
}

// Runs once before the entire suite
func (s *LoadSuite) SetupSuite() {
    s.cfgPath = filepath.Join(s.T().TempDir(), "config.yaml")
}

// Runs before each test method
func (s *LoadSuite) SetupTest() { ... }

// Runs after each test method — always, even on failure
func (s *LoadSuite) TearDownTest() { ... }

func (s *LoadSuite) TestValidFile() {
    cfg, err := config.Load(s.cfgPath)
    s.Require().NoError(err)
    s.Equal("https://example.com", cfg.URL)
}

// Entry point — one standard Test* function per suite
func TestLoad(t *testing.T) {
    suite.Run(t, new(LoadSuite))
}
```

Inside suite methods use `s.Require()` / `s.Assert()`, or call assertions directly on `s` (`s.Equal`, `s.NoError`) — the suite embeds them. The `*testing.T` receiver is `s.T()`.

| Hook | When it runs |
|---|---|
| `SetupSuite` / `TearDownSuite` | once for the whole suite |
| `SetupTest` / `TearDownTest` | before/after each test method |

### Suite vs table-driven — when to use which

They solve different problems and compose naturally.

**Table-driven:** same function, many inputs. The question is "does this function return the right output for all these cases?" All rows share the same assertion logic; only the data varies.

**Suite:** shared infrastructure, many functions. The question is "how do I avoid repeating setup across multiple test functions that all need the same fixture?" The struct holds the state; each method tests something different.

A suite method can contain a table-driven loop:

```go
func (s *LoadSuite) TestMalformedURLs() {
    cases := []struct{ input, wantErr string }{
        {"not-a-url", "invalid URL"},
        {"ftp://bad", "unsupported scheme"},
    }
    for _, tc := range cases {
        s.Run(tc.input, func() {  // s.Run inside a suite, not t.Run
            _, err := config.Load(tc.input)
            s.ErrorContains(err, tc.wantErr)
        })
    }
}
```

Reach for a suite when you notice yourself copy-pasting the same setup block across test functions. For one or two tests with a one-line setup, `t.Cleanup` is lighter.

### Adding context to failures

Every assertion accepts an optional message as final args, `printf`-style (add `f` suffix):

```go
assert.Equalf(t, want, got, "case %q: unexpected result", tc.name)
require.NoErrorf(t, err, "loading fixture %s", path)
```

Without a message, testify prints a diff automatically. Add messages when the diff alone wouldn't identify *which* case or input caused the failure.

## Grouping subtests with t.Run

When cases aren't parametrized but you still want named, independent subtests under one `Test*` function:

```go
func TestLoad(t *testing.T) {
    t.Run("valid YAML file returns correct fields", func(t *testing.T) {
        // ...
    })

    t.Run("missing file returns empty Config, no error", func(t *testing.T) {
        // ...
    })

    t.Run("malformed YAML returns an error", func(t *testing.T) {
        // ...
    })
}
```

Each subtest gets its own pass/fail status and can be targeted individually: `go test -run TestLoad/missing`. Unlike table-driven tests, there's no shared struct — each case has its own setup inline.

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

### Locating testdata from a shared helper

If a `testdata()` helper lives in a shared package (e.g. `internal/testhelpers/`), use `runtime.Caller(1)` to anchor to the *calling* test file, not the helper's own file:

```go
// skip=1 resolves relative to the caller of this function
func Fixture(name string) string {
    _, file, _, _ := runtime.Caller(1)
    return filepath.Join(filepath.Dir(file), "testdata", name)
}
```

With `skip=0`, the path would resolve relative to the helper file — pointing at `testhelpers/testdata/`, not the calling package's `testdata/`. Use `skip=0` only when the helper and its testdata live in the same package.

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

## t.Skip — skipping and stubbing

```go
t.Skip("not yet implemented")   // marks test as stub; suite stays green
t.Skip("flaky on CI, see #42")  // skip with reason

if runtime.GOOS == "windows" {
    t.Skip("unix-only")         // conditional skip
}
```

`t.Skip` stops the current test and marks it as skipped (not failed). In TDD, writing `t.Skip("not yet implemented")` inside each `t.Run` block is the standard way to scaffold a test suite before the implementation exists — the structure is visible in the file and in test output, but the suite stays green.

`t.SkipNow()` is the same but takes no message; `testing.Short()` pairs with `-test.short` to skip slow tests in fast-feedback loops.

## TDD discipline

**Red before green:** a test that passes before its implementation exists is not a red test — it's a broken test. Always confirm a test fails for the right reason before writing the implementation.

**Choosing what to test first:** ask before writing a test — *can this fail against a stub that returns zero values?*

- `len(files) != 0` against a nil return — passes trivially, not a useful red test
- `len(files) != 2` against a nil return — fails immediately, useful red test

When the simplest test is also the least discriminating, start with a test that has real fixture data and concrete assertions. Come back to degenerate cases (empty input, error paths) once the implementation exists — at that point they guard against regressions.

*Related: [[Go Error Handling]] (never discard errors in tests), [[Go File IO]] (testdata fixtures, t.TempDir)*
