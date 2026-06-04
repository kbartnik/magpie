---
description: Seed the inbox with fixture documents for testing
allowed-tools: ["Bash(ls:*)", "Bash(cp:*)"]
---

Populate `inbox/` with the fixture documents from `.fixtures/` so the vault is ready for a test run.

**Step 1 — Check for existing inbox files:**

Run: `ls "${VAULT_PATH:-$(pwd)}/inbox/"` (excluding `.keep` files).

If any non-hidden files are present, print:
```
⚠ inbox/ already contains files:
  <list files>

Overwrite? (y/n)
```

Wait for the user's response. If anything other than `y`, print "Seed cancelled." and stop.

**Step 2 — Copy fixtures:**

Copy all files from `${VAULT_PATH:-$(pwd)}/.fixtures/` into `${VAULT_PATH:-$(pwd)}/inbox/`.

**Step 3 — Update context.md:**

Count the number of files copied. Update the `inbox-count` field in `${VAULT_PATH:-$(pwd)}/context.md` to that count.

**Step 4 — Report:**

Print a summary listing the seeded files, e.g.:
```
✓ Seeded N items into inbox/:
  - 2026-05-20-eleven-agentic-patterns.md
  - ...
```
