---
title: "Measuring the Impact of Early-2025 AI on Experienced Open-Source Developer Productivity"
type: paper
captured-date: 2026-06-01
source-url: "https://arxiv.org/abs/2507.09089"
authors: "Joel Becker, Nate Rush, Beth Barnes, David Rein"
year: 2025
venue: "arXiv (METR)"
---

Randomized controlled trial measuring how AI tools at the February–June 2025 frontier affect the productivity of experienced open-source developers. 16 developers, 246 tasks, mature repos (avg. 23k stars, 1.1M LOC), developers averaging 5 years on target repo.

## Deep Read

**Key Insight:** Allowing AI tools increased task completion time by 19% despite developers predicting 20–24% speedup and ML/economics experts predicting 38–39% speedup. The expectation gap is itself the finding: developers and experts systematically overestimate AI speedup, even after months of active use.

**What Surprised Me:** After completing the study, developers *still* estimated they had been sped up by 20% — despite the measured slowdown of 19%. They were wrong about their own performance in the wrong direction, by nearly 40 percentage points, after having direct experience. This is a stronger form of the fluency heuristic problem than previously documented.

**Open Questions:**
- Developers accepted <44% of AI generations and spent 9% of time reviewing/cleaning output. How does acceptance rate change as model capability improves — and does the review overhead scale proportionally or does it decrease faster?
- The slowdown was specific to experienced devs in mature, high-quality-bar repos. At what point on the experience/codebase-maturity spectrum does AI flip from net cost to net benefit?
- The paper notes "further improvements to current AI systems (e.g. better prompting/agent scaffolding, or domain-specific finetuning) could yield positive speedup in this setting" — what harness changes would most directly address the identified overhead sources?

**Wikilink Candidates:**
- [[AI Productivity Research]] — this paper is the anchor study for the "experienced devs, real tasks" cell; page created
- [[Vibe-Coding Anti-Pattern]] — empirical confirmation of the experienced-engineer trap and surface trust failure mode
- [[IDSD]] — the METR finding was already cited here from a secondary source; this is the primary

**Connections:**
- [[AI Productivity Research]] — primary anchor study; sets the "real tasks, experienced devs, fixed outcome measure" methodological standard
- [[Vibe-Coding Anti-Pattern]] — the 19% slowdown is the quantified cost of the anti-pattern; developers' 20% post-hoc overestimate confirms the "experienced engineer trap" section
- [[IDSD]] — cited as "METR controlled trial finding" in the IDSD page; upgrades that citation to primary source
- [[Harness Engineering]] — the 9% review overhead is a harness cost; good scaffolding would address it

**Image Candidates:** none (PDF — figures present in source but not embedded as Obsidian media)
