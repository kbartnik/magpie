---
title: "ADHD"
type: hub
status: active
created: 2026-05-31
updated: 2026-06-01
sources: []
related:
  - "Digital Biomarkers for ADHD"
  - "Digital Therapeutics for ADHD"
  - "ADHD in Software Engineering"
  - "Digital Phenotyping"
tags:
  - adhd
  - neurodiversity
---

# ADHD

Attention-Deficit/Hyperactivity Disorder — a neurodevelopmental condition characterized by persistent inattention, hyperactivity, and/or impulsivity that interferes with functioning. Estimated prevalence: 5.0–7.1% of the global population. Symptoms often persist into adulthood, though adult ADHD remains underdiagnosed.

**Core dimensions:**
- **Inattention** — difficulty sustaining focus, following through on tasks, organizing work, resisting distraction
- **Hyperactivity/Impulsivity** — excessive movement, interrupting, difficulty waiting; in adults often internalized as restlessness and racing thoughts
- **Executive function deficits** — time estimation, working memory, task initiation, cognitive flexibility

ADHD is heterogeneous: three subtypes (predominantly inattentive, predominantly hyperactive/impulsive, combined) and significant individual variation in symptom expression, severity, and coping.

## This Vault's Coverage

This is a research domain topic covering the intersection of ADHD with digital health, passive sensing, and workplace technology. Four concept pages:

| Page | Coverage |
|---|---|
| [[Digital Biomarkers for ADHD]] | Keystroke, mouse movement, eye-tracking, EEG — passive sensing for ADHD assessment |
| [[Digital Therapeutics for ADHD]] | FDA-cleared digital treatments, DHI taxonomy, evidence base |
| [[ADHD in Software Engineering]] | Challenges, strengths, strategies, and tools for developers with ADHD |
| [[Digital Phenotyping]] | Infrastructure layer: passive sensing platforms, governance, privacy |

## Neurobiology

The most consistent structural finding is an overall reduction in total brain size with specific changes in the **caudate nucleus**, prefrontal cortex white matter, corpus callosum, and cerebellar vermis. The **frontostriatal circuit** (prefrontal cortex ↔ basal ganglia) is the most convincingly implicated system.

**Dopamine transfer deficit model (Tripp & Wickens 2009):** Dopamine cells normally fire at rewarding events and then "transfer" their firing to earlier predictive cues as behavior is learned. ADHD may involve a failure of this transfer — the system remains hypersensitive to immediate reward and fails to represent delayed ones. This explains delay intolerance, impulsivity, and why immediate feedback (games, stimulants) dramatically changes behavior.

**Genetics:** Heritability ~0.8 (comparable to height). Massively polygenic — no single gene accounts for more than ~3.6% of symptom variance. Most studied variants: DRD4, DAT1, DRD5, serotonin transporter. Gene-environment interactions matter (e.g., prenatal smoking × DAT1 genotype).

Sources: [[archive/papers/2026-06-01-tripp-wickens-adhd-neurobiology-notes|Tripp & Wickens 2009]], [[archive/papers/2026-06-01-adhd-neurobiology-primer-notes|Rege primer]]

## Key Facts for This Research Domain

- Current diagnosis relies entirely on clinical interview and behavioral questionnaires — no objective biomarkers in clinical use
- Medication (stimulants, non-stimulants) is effective but has side effects and access barriers
- Digital therapeutics represent a new regulatory category: FDA-cleared software as prescription treatment
- ADHD in adults is often undiagnosed (estimated ~75% undetected in adulthood)
- In software engineering specifically, **10.6% of programmers** report a concentration/memory disorder (Stack Overflow 2022) — higher than the general population estimate

## See Also

- [[Digital Biomarkers for ADHD]] — the sensing modality research
- [[Digital Therapeutics for ADHD]] — the intervention research
- [[ADHD in Software Engineering]] — the workplace/professional context
- [[Digital Phenotyping]] — the infrastructure and governance layer
- [[Employer Monitoring as ADHD Sensing]], [[Neurodiverse SE Research Post-LLM Validity]], [[ADHD Disclosure Risk and Policy]], [[Dopamine Transfer Deficit Subtypes]] — open questions on ADHD research
