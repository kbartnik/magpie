---
title: "Alignment as Iterative Design"
question: "Does the emergent control problem imply that AI alignment cannot be solved by upfront design and must instead be discovered iteratively by running systems?"
type: question
status: open
domain: ai-theory
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-emergent-garden-emergent-complexity.md"
related:
  - "[[Emergence and Complexity]]"
  - "[[Autonomous Learning Architecture]]"
  - "[[LLM Mental Model]]"
tags:
  - ai-theory
  - alignment
---

# Alignment as Iterative Design

*Does the emergent control problem imply that AI alignment cannot be solved by upfront design and must instead be discovered iteratively by running systems?*

The "control problem" framing from emergence theory: emergent systems can only be controlled by modifying the rules that produce the behavior, not by directly patching the behavior. Applied to AI: RLHF, Constitutional AI, and fine-tuning are consistent with this — they modify training objectives (rules). Post-hoc output filters are not — they patch symptoms.

If this framing is correct, there is a strong implication: alignment is not a safety layer added after training, it is constitutive of the training process itself. The question is whether this is actually true or whether sufficiently careful behavior-level patches can reliably prevent misaligned outputs at deployment. The empirical record is mixed: some post-hoc mitigations hold; others are bypassed with modest adversarial effort. Whether this gap is a fundamental property of the approach or a maturity problem is unresolved.

## See Also

- [[Emergence and Complexity]] — the control problem; building-block rules as the intervention point
- [[Autonomous Learning Architecture]] — System M as the post-deployment learning architecture this problem motivates
