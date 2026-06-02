---
title: "SSH Config"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - dev/learning/magpie-go/2026-06-02.md
related:
  - Git
tags:
  - ssh
  - tooling
  - security
---

# SSH Config

`~/.ssh/config` controls SSH client behavior per host. Without it, SSH offers all loaded keys to every server — the right key for a specific host must be declared explicitly.

## Per-Host Key Selection

```
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_github
```

Without this block, a key named `id_ed25519_github` is **never offered** to GitHub during handshake — it exists but is invisible to the SSH agent negotiation. A key file in `~/.ssh/` is not automatically used; it must be mapped to a host.

## Wildcard Block for Shared Settings

```
Host *
  UseKeychain yes
  AddKeysToAgent yes
  ServerAliveInterval 60
```

`UseKeychain yes` — stores passphrases in the macOS keychain so you aren't prompted after first use.
`AddKeysToAgent yes` — loads the key into the running SSH agent on first use.

The wildcard block applies to all hosts. Per-host blocks add or override settings for specific hosts.

## Block Evaluation Order

SSH reads the config top-to-bottom and applies the **first match** for each directive. Put specific `Host` blocks before the wildcard `Host *` — the wildcard fills in anything not already set.

## Common Patterns

```
# Multiple GitHub accounts
Host github-personal
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_personal

Host github-work
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_work
```

Clone with `git clone git@github-personal:user/repo` to select the right identity.

## Testing Auth

```bash
ssh -T git@github.com
# Hi username! You've successfully authenticated...
# Exit code 1 is expected — GitHub denies shell access.
```

Exit code 1 with the greeting message means auth succeeded. Exit code 255 means key negotiation failed.
