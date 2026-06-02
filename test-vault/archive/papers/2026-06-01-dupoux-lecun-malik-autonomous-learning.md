---
title: "Why AI systems don't learn and what to do about it"
type: paper
captured-date: 2026-06-01
source-url: "https://arxiv.org/abs/2603.15381"
authors: "Emmanuel Dupoux, Yann LeCun, Jitendra Malik"
year: 2026
venue: "arXiv (FAIR/META, NYU, UC Berkeley)"
---

Position paper from LeCun, Malik (FAIR/META, NYU, UC Berkeley) arguing that autonomous learning — not larger models — is the next frontier in AI. Proposes a three-system architecture (System A/B/M) inspired by human and animal cognition. Explicitly positions against the hyperscaling paradigm.

## Deep Read

**Key Insight:** Current AI doesn't learn post-deployment — learning is fully "outsourced" to human data scientists and ML engineers. This isn't a missing feature; it's a structural property of how current training pipelines are designed. The paper argues this is the root cause of domain mismatch failures and the inability to adapt to non-stationary real-world data.

**What Surprised Me:** The authors use the *Drosophila* / fruit fly result (Eon Systems fully reproduced the 140,000-neuron nervous system and got complex survival behaviors without training) as evidence that intelligence may be *in the wiring*, not the weights — a direct challenge to the parameterization view that LLMs embody. They cite this in the context of arguing that architectural fidelity matters more than scale on fixed data.

**Open Questions:**
- System M (meta-control) is the novel proposed piece, but the paper is a roadmap, not an implementation. What does a concrete System M look like at the scale of a useful deployed agent — not a toy cognitive architecture?
- The bilevel evolutionary optimization approach (evolve System M over System A/B) is suggested for building these architectures. How does training cost compare to pretraining a frontier LLM? The paper doesn't address economics.
- The paper argues current SSL/LLM models *are* System A, and RL agents *are* System B — so these exist. The gap is System M. Is there any current work that comes close, or is this purely aspirational?

**Wikilink Candidates:**
- [[Autonomous Learning Architecture]] — the System A/B/M framework; page created
- [[Neuromorphic and Bio-Inspired AI]] — software-side counterpart to this paper's architecture
- [[LLM Mental Model]] — the paper adds "frozen after training" as a structural limit to add to this page

**Connections:**
- [[Autonomous Learning Architecture]] — primary conceptual home for the System A/B/M framework
- [[Neuromorphic and Bio-Inspired AI]] — LeCun/Malik's proposed System A/B/M is the software-side design; Onto AI's 4-layer architecture is a prior implementation attempt solving an overlapping problem
- [[Agent Memory Architectures]] — System M is a principled architecture for what Hebbian Knowledge Graph (Architecture 4) approximates; both are trying to get learning out of the model weights and into a persistent adaptive structure
- [[Bio-Inspired Computing]] — bilevel evolutionary optimization for System M is directly in-scope for this page

**Image Candidates:** Figure 1 (standard ML pipeline vs. autonomous ML pipeline) — informational diagram showing the assembly line of human experts vs. autonomous learning loop. Not embedded as Obsidian media (native PDF).
