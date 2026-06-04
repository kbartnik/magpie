# CLAUDE.md — magpie

> **Transitional note:** Skills and commands currently use `vault-tools` (Nexus binary, copied to `test-vault/.claude/tools/`). Phase 8 (Migration) replaces all `vault-tools` calls with `magpie` subcommands and removes the copied binary.

---

## Test Vault

The test vault lives at `test-vault/` in this repo. See `test-vault/CLAUDE.md` for vault operation instructions.

**VAULT_PATH** must be set to the absolute path when running vault commands from the repo root:
```
VAULT_PATH=/Users/kurt/Source/magpie/test-vault
```

When Claude Code is opened directly from `test-vault/`, `${VAULT_PATH:-$(pwd)}` resolves correctly without this variable.
