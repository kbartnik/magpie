---
title: "Archive Similarity Scan"
type: backlog-item
project: magpie
status: todo
milestone: "post-1.0"
priority: low
due: ""
created: 2026-05-30
updated: 2026-05-30
tags: [post-1.0, archive, dedup, similarity]
---

# Archive Similarity Scan

vault-tools' `find-duplicates` catches exact slug matches (same filename stem, different date prefix). That misses cases where the same content arrives under a different slug — e.g., captured from a mirror domain, reformatted, or re-clipped with different wording in the title.

Magpie should offer a `magpie archive scan --deep` mode that normalizes each file's body (strip frontmatter, lowercase, remove punctuation, tokenize into word sets) and computes pairwise containment scores (`|A∩B| / min(|A|,|B|)`). Containment is preferred over Jaccard because an annotated re-capture is a superset of the original — Jaccard penalizes the added words; containment does not.

## Sketch

- `normalizeWords(content []byte) map[string]bool` — strip frontmatter, lowercase, remove non-alphanumeric, drop tokens ≤ 2 chars
- `contentSimilarity(a, b map[string]bool) float64` — containment: `|A∩B| / min(|A|,|B|)`
- `deepScan(dir string, threshold float64) []SimilarityResult` — O(n²) pairwise, fine for archive sizes
- Report format: `archive/clippings/file-a.md | archive/clippings/file-b.md | 87% match`
- Default threshold: 0.70; configurable via flag

## Notes

- O(n²) is acceptable for archive sizes (typically < 200 files per type)
- The two-tier detection (slug surface + similarity deep) makes the surface scan fast for real-time use in `archive-file`, and the deep scan a periodic maintenance command
- Threshold tuning: the `skills-deep-dive-part-2` pair (original + annotated) is a good calibration case — the annotated version should score ≥ 80%
