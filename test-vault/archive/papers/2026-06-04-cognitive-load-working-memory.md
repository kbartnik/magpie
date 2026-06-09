---
title: "Notes: Cognitive Load Theory and Working Memory (Sweller 1988 / Baddeley 2000)"
type: paper-notes
source: "Sweller, J. (1988). Cognitive load during problem solving. Cognitive Science, 12(2), 257–285. + Baddeley, A. (2000). The episodic buffer: a new component of working memory? Trends in Cognitive Sciences, 4(11), 417–423."
---

## Deep Read

**Key Insight:** Working memory is the bottleneck of conscious thought. Sweller's central finding: learning fails not because content is too hard but because the *presentation* imposes unnecessary cognitive load that exhausts working memory before the material can reach long-term memory. Baddeley's model makes this concrete — working memory has distinct subsystems with limited capacity, and the bottleneck is the central executive (prefrontal cortex-mediated control), not raw storage.

**What Surprised Me:** Sweller distinguishes three types of load:
- **Intrinsic load** — the inherent complexity of the material (can't be reduced without changing the content)
- **Extraneous load** — complexity introduced by poor presentation (can and should be minimized)
- **Germane load** — cognitive effort spent constructing schemas (desirable — this is learning)

The implication is counterintuitive: making material *easier to process* (lower extraneous load) frees capacity for schema formation (higher germane load). Reducing friction is not the same as reducing learning. Well-designed worked examples, split-attention effects, and redundancy effects all follow from this distinction.

Baddeley's episodic buffer adds a fourth subsystem to his earlier phonological loop + visuospatial sketchpad model. The episodic buffer integrates information across modalities and from long-term memory into a unified conscious representation. It's capacity-limited (~4 chunks), temporarily bound, and centrally controlled. This is where complex reasoning happens.

**Open Questions:**
- Sweller's original work focused on math problem-solving with novices. How well does CLT transfer to expert-level skill acquisition, where schemas are already rich and intrinsic load is lower?
- The "4 chunks" capacity estimate comes from Cowan (2001), updating Miller's famous "7±2" figure. How does chunk size vary with expertise — can an expert's chunk hold exponentially more information than a novice's?
- Baddeley's central executive maps roughly onto prefrontal cortex function. Given that ADHD involves reduced prefrontal regulation of the frontostriatal circuit, is ADHD best understood as a *central executive* deficit specifically, rather than a working memory storage deficit?

**Wikilink Candidates:**
- [[Working Memory]] — Baddeley's four-component model, capacity limits, central executive; not yet a wiki page
- [[Cognitive Load Theory]] — intrinsic/extraneous/germane load, implications for instructional design; not yet a wiki page

**Connections:**
- [[ADHD]] — ADHD's core impairments (inattention, poor working memory, executive dysfunction) map directly onto central executive deficits; CLT predicts that high extraneous load is disproportionately costly for ADHD individuals because their central executive has less reserve capacity
- [[Frontostriatal Circuit]] — the central executive is prefrontal-cortex-mediated; the frontostriatal circuit's role in regulating attention and inhibition is the neurological substrate of Baddeley's central executive
- [[Transformer Architecture]] — the LLM context window is a functional analogue to working memory: a limited-capacity buffer holding the "active" information for current reasoning. Forgetting beyond the context window mirrors the rapid decay of working memory without rehearsal. Retrieval-augmented generation is roughly analogous to Baddeley's episodic buffer pulling from long-term memory on demand.
