---
title: "Digital Therapeutics for ADHD"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/papers/2026-05-31-stamatis-2024-akl-t01-digital-therapeutic.pdf"
  - "archive/papers/2026-05-31-lakes-2022-digital-health-interventions-adhd.pdf"
related:
  - "ADHD"
  - "Digital Biomarkers for ADHD"
  - "ADHD in Software Engineering"
tags:
  - adhd
  - digital-therapeutics
  - intervention
---

# Digital Therapeutics for ADHD

Digital health interventions (DHIs) for ADHD span a spectrum from prescription-grade FDA-cleared treatments to evidence-backed apps to telehealth delivery of existing behavioral therapies. The field is rapidly expanding — a 2022 mapping review found 51 studies across 16 countries and 5 delivery modalities.

## Scoping Review: State of the Field (Schofield et al. 2026)

The most comprehensive map of the DHT landscape for *adult* ADHD — 133 studies across three functional categories:

| Category | N studies | Primary modalities |
|----------|-----------|-------------------|
| Treat a Specific Condition | 63 | Web/app cognitive therapy (26), cognitive training (13), transcranial stimulation (12), neurofeedback (9) |
| Drive Clinical Management | 36 | ML on CPT data (11), neuroimaging (11), VR (5) |
| Diagnose a Specific Condition | 19 | Similar ML classification, without clinical pathway integration |

**Key finding:** Evidence base is large but high-risk-of-bias. Most studies are preliminary efficacy tests without standardized outcomes. A recurring gap: DHTs developed for diagnosis don't specify how they integrate into clinical pathways — a clinician has no decision framework for acting on ML output.

**Context that makes this urgent:** UK wait times for adult ADHD diagnosis now commonly exceed 2 years. Administrative prevalence in GP records is ~0.3% (men) and ~0.07% (women) — far below the expected epidemiological prevalence of 2.58%. The diagnostic gap, not the treatment gap, is the dominant problem.

Source: [[archive/papers/2026-06-01-dht-adhd-scoping-review-2026-notes|Schofield et al. 2026]]

## The DHI Taxonomy (Lakes et al. 2022)

**Delivery modalities** (not mutually exclusive):
- **Serious games / e-learning** — cognitive training via game mechanics; largest evidence base
- **mHealth** — smartphone apps for symptom tracking, reminders, behavioral monitoring
- **Web** — browser-based tools; includes caregiver-facing portals
- **Telehealth** — video/phone delivery of established behavioral therapies
- **VR/AR** — immersive environments for attention training; small but growing literature

**Target domains** (what the DHI is trying to change):
- Cognition (attention, executive function, working memory)
- Social-emotional skills
- Behavior management
- Academic / organizational skills
- Medication adherence
- Vocational skills

The evidence is concentrated in **cognition** (cognitive training + neurofeedback have prior meta-analyses and review cycles) and **behavior management via telehealth** (especially parent training). The other domains are underresearched.

## AKL-T01 (EndeavorRx) — The FDA-Cleared Case

AKL-T01 is a video game-based digital therapeutic developed by Akili Interactive. It received FDA De Novo authorization (2020) for treatment of ADHD in children ages 8–12 — the first prescription video game.

**Mechanism:** Targeted cognitive training for attentional control, using multi-sensory challenge environments requiring selective attention and task-management. The hypothesis is that repeated cognitive engagement transfers to real-world attentional functioning via neuroplasticity.

**Evidence across age groups (Stamatis et al. 2024):**

| Population | Trial | Duration | TOVA-ACS Change | p-value |
|---|---|---|---|---|
| Children 8–12 | STARS-ADHD-Child (RCT) | 4 weeks | +1.0 | <0.001 |
| Adolescents 13–17 | STARS-ADHD-Adolescent (single-arm) | 4 weeks | +2.6 | <0.0001 |
| Adults 18+ | STARS-ADHD-Adult (single-arm) | 6 weeks | +6.5 | <0.0001 |

The adult effect (~7x the pediatric effect) is attributed to self-selection into treatment and the predominance of inattentive subtype in adults — which is precisely the symptom domain AKL-T01 targets.

**Regulatory significance:** The FDA De Novo pathway was used because no legally marketed predicate device existed. This created a new regulatory category for prescription digital therapeutics (PDTs). The precedent matters for any software aiming for "software as medicine" status.

**Key limitation of the adult/adolescent trials:** Both are single-arm (no control group). Without active control conditions, placebo effects can't be ruled out, though the pediatric RCT showed no TOVA placebo effect.

## Telehealth and Parent Training

The Lakes 2022 review highlights that telehealth delivery of *behavioral parent training* — not the child-focused cognitive interventions — has a surprisingly strong evidence base for pediatric ADHD. This shifts the intervention target from child to environment:
- Training parents to structure environments and rewards
- Reducing the burden on the child to self-regulate
- Telehealth removes geographic and access barriers to specialist behavioral therapists

This finding challenges the assumption that DHIs for ADHD primarily mean apps for the person with ADHD.

## Current Evidence Gaps

- **Adult populations:** most DHI research is in children; the adult evidence base is thin (AKL-T01 adult trial is single-arm)
- **Social and identity dimensions:** few DHIs target the social, emotional, and identity aspects of ADHD in adolescents/adults
- **Long-term outcomes:** most trials measure outcomes at 4–6 weeks; long-term maintenance effects are largely unknown
- **Real-world effectiveness:** research samples have high engagement; real-world use patterns with ADHD users likely show higher dropout

## Relationship to Biomarkers

Digital biomarkers and digital therapeutics are increasingly coupled:
- **Adaptive therapeutics:** if [[Digital Biomarkers for ADHD]] can track current attention state (e.g., Shahaf's state-specific BEI'), therapeutics could adapt difficulty in real time
- **Outcome measurement:** the same behavioral sensing used for diagnosis can measure treatment response — closing the loop between biomarker and intervention

## See Also

- [[ADHD]] — clinical overview
- [[Digital Biomarkers for ADHD]] — the sensing/measurement counterpart
- [[ADHD in Software Engineering]] — workplace tools that borrow therapeutic design principles (Tether, gamification)
