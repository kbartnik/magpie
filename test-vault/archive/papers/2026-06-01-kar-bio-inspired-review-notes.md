---
title: "Notes: Bio-Inspired Computing — A Review of Algorithms and Scope of Applications (Kar, Expert Systems with Applications)"
type: paper-notes
source: "2026-06-01-kar-bio-inspired-review.pdf"
---

## Deep Read

**Key Insight:** The bio-inspired literature is heavily skewed toward ~4 popular algorithms (neural networks, genetic algorithms, PSO, ACO), causing researchers to "force-fit" familiar tools rather than select the approach best suited to their problem. Kar reviews 12 algorithms with their application domains to address this awareness gap.

**What Surprised Me:** The breadth of algorithms with meaningful application histories extends well beyond the usual suspects — firefly algorithm, cuckoo search, bacterial foraging optimization, artificial bee colony, and artificial plant optimization all have >15 published applications. The perception that "only a few real algorithms exist" understates the field's actual diversity.

**Open Questions:**
- The review explicitly excludes pseudocode and implementation details to stay high-level. How useful is algorithm selection guidance without enough detail to distinguish true behavioral differences between candidates?
- Most reviewed algorithms predate deep learning's dominance (2015+). What is their current competitive position against DL baselines on the application domains Kar covers?
- "Force-fitting" is critiqued as a systemic problem but not illustrated with case studies. What signals should a practitioner use to identify when they are force-fitting rather than legitimately applying a familiar tool?

**Wikilink Candidates:**
- [[Bio-Inspired Computing]] — this review provides the practitioner lens (algorithm selection, application scope) that complements the field-level critique in Somvanshi 2025

**Connections:**
- [[Bio-Inspired Computing]] — the force-fit problem Kar identifies is a human/community factor, not just an algorithmic one
- [[archive/papers/2026-06-01-somvanshi-2025-bio-inspired-critique-notes|Somvanshi 2025]] — convergent finding: both papers diagnose limited awareness as a root cause of poor algorithm selection
