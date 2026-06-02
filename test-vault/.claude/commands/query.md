---
description: Synthesize an answer from vault knowledge
argument-hint: "your question"
allowed-tools: ["Bash(grep:*)", "Bash(find:*)", "Bash(cat:*)", "Bash(ls:*)"]
---

1. Read wiki/index.md to identify candidate pages.
2. Grep wiki/ for terms from $ARGUMENTS to find additional candidates.
3. Read up to 10 candidate files — prioritize wiki/ over archive/.
4. Answer: direct prose first (1-2 paragraphs), then evidence with [[wikilink]] citations, then gaps ("wiki is thin on X — worth ingesting?"), then connections.
5. Surface unresolved `Open:` items from recent log entries relevant to this question.
6. Offer: "Want me to file this as `wiki/syntheses/<slug>.md`?"
7. If yes: present draft, wait for approval, write it, update wiki/index.md, call `"${VAULT_PATH:-$(pwd)}/.claude/tools/vault-tools" log-append query "<title>" "<files>" "<summary>" "<open>"`.
