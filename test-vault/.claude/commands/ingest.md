---
description: Process inbox items into archive and wiki
argument-hint: "[url | file | empty for all inbox]"
allowed-tools: ["Bash(bash:*)", "Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)", "WebFetch"]
---

Load `.claude/skills/source-annotation/SKILL.md` before reading any source.

**Route by argument:**
- URL: fetch with WebFetch, save to `inbox/YYYY-MM-DD-<slug>.md` with `source-url:` in frontmatter, then process as inbox item
- File path: process that file only
- Empty or "inbox": process all non-.keep files in inbox/ sequentially

**For each file:**
1. Read fully. If it's a daily-inbox file (multiple bullet captures): split into logical items and process each separately.
2. Classify: clipping, paper, book, daily, idea, or doc. For unambiguous cases (PDF → paper, dated daily-inbox → daily) proceed without asking. For ambiguous cases: one-line prompt, wait.
3. Surface 3-5 key takeaways, anything surprising, existing wiki pages this connects to, and any informational images identified by the Image Scan (diagrams, charts, flowcharts worth placing in wiki pages).
4. Ask: "Any angle or emphasis before I write?" Wait.
5. Present plan:
   ```
   📋 Ingest plan: <source title>
     Archive: inbox/<file> → archive/<type>/YYYY-MM-DD-<slug>.md
     Wiki pages to create: [list or "none"]
     Wiki pages to update: [list or "none"]
     Scope: ~N pages
   Proceed? (y/n)
   ```
6. On approval: `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" archive-file <source> <type> <slug>`, then create/update wiki pages and wiki/index.md. Place any informational images identified in the Image Scan at the appropriate wiki page locations using `![[resources/media/...]]` embeds — prefer this over writing text descriptions of what the image already shows.
6a. If `archive-file` exits with code 2 (duplicate warning): surface the reported existing file(s) to the user:
    "⚠ This slug is already archived at [path]. Archive again anyway? (y/n)"
    - If yes: re-run with `--force`: `vault-tools archive-file --force <source> <type> <slug>`
    - If no: skip archiving; proceed to wiki annotation steps only if the user wants to update existing knowledge.
7. `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append ingest "<title>" "<files summary>" "<summary>" "<open items>"`
8. Report: `✓ Ingested: <title>. Created N pages, updated M.`
