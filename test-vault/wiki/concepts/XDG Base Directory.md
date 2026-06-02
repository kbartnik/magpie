---
title: "XDG Base Directory"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - dev/learning/magpie-go/2026-06-02.md
related:
  - Go Modules and Packages
tags:
  - linux
  - tooling
  - configuration
  - go
---

# XDG Base Directory

The XDG Base Directory Specification defines standard locations for config, data, and cache files so applications don't scatter files across `$HOME`.

## Standard Directories

| Variable | Default | Purpose |
|----------|---------|---------|
| `XDG_CONFIG_HOME` | `~/.config` | User-specific config files |
| `XDG_DATA_HOME` | `~/.local/share` | User-specific data files |
| `XDG_CACHE_HOME` | `~/.cache` | Non-essential cached data |
| `XDG_STATE_HOME` | `~/.local/state` | Persistent state (logs, history) |

An app named `magpie` stores its config at `$XDG_CONFIG_HOME/magpie/config.yaml` — or `~/.config/magpie/config.yaml` when the env var is unset.

## Cross-Platform in Go: os.UserConfigDir()

XDG is Linux-native. `os.UserConfigDir()` provides the platform-correct equivalent:

- **Linux**: `$XDG_CONFIG_HOME` or `~/.config`
- **macOS**: `~/Library/Application Support`
- **Windows**: `%AppData%`

**Gotcha:** `os.UserConfigDir()` on macOS ignores `XDG_CONFIG_HOME` entirely. To support power users on any platform who've set it, check it manually first:

```go
func configPath() string {
    if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
        return filepath.Join(xdg, "magpie", "config.yaml")
    }
    base, err := os.UserConfigDir()
    if err != nil {
        base = filepath.Join(os.Getenv("HOME"), ".config")
    }
    return filepath.Join(base, "magpie", "config.yaml")
}
```

This gives XDG honoring on all platforms while falling back to the OS-native location when unset.

## Testing configPath()

Because the result is platform-dependent, tests should not hardcode `~/.config`. Instead, compute the expected value the same way the implementation does:

```go
func TestConfigPathUnset(t *testing.T) {
    t.Setenv("XDG_CONFIG_HOME", "")
    base, _ := os.UserConfigDir()
    want := filepath.Join(base, "magpie", "config.yaml")
    if got := configPath(); got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

## Why It Matters

Without XDG, apps create `~/myapp/`, `~/.myapp/`, `~/.myapprc` etc. — there's no standard. XDG lets users relocate all config with a single env var and keeps `$HOME` clean.
