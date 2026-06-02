---
title: "Digital Biomarkers for ADHD"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/papers/2026-05-31-alfalahi-2022-keystroke-biomarkers-meta-analysis.pdf"
  - "archive/papers/2026-05-31-leontyev-2019-mouse-movement-adhd-ssrt.pdf"
  - "archive/papers/2026-05-31-liu-2024-eye-tracking-adhd-diagnosis.pdf"
  - "archive/papers/2026-05-31-shahaf-2018-eeg-brain-engagement-adhd.pdf"
related:
  - "ADHD"
  - "Digital Phenotyping"
  - "Digital Therapeutics for ADHD"
tags:
  - adhd
  - digital-biomarkers
  - passive-sensing
---

# Digital Biomarkers for ADHD

Digital biomarkers are objectively measurable physiological or behavioral signals collected from digital devices that indicate health status. For ADHD, the appeal is significant: current diagnosis relies entirely on clinical interview and self-report questionnaires, with no objective measure in clinical use. Digital biomarkers offer the possibility of passive, continuous, ecologically valid assessment.

## The Sensing Modality Landscape

Four modalities are represented in current literature, each capturing a different aspect of ADHD-related function:

### Keystroke Dynamics

Keyboard typing patterns — inter-key intervals, dwell times, flight times, error rates — capture fine motor and cognitive control processes that differ between ADHD and neurotypical populations.

**Evidence (Alfalahi et al. 2022 — systematic review, 41 studies):**
- Pooled sensitivity: 0.86 (95% CI 0.82–0.90)
- Pooled specificity: 0.83 (95% CI 0.79–0.87) for neuropsychiatric disorders broadly
- Psychiatric disorder subgroup (includes ADHD): sensitivity 0.85, specificity 0.82
- Naturalistic typing from everyday use performs comparably to lab-controlled tasks
- Deep learning + multimodal analysis achieves best performance

**Key gap:** All studies are cross-sectional; longitudinal monitoring validity not yet established.

### Mouse Movement

Cursor kinematics during tasks capture impulsivity and response inhibition — the behavioral disinhibition dimension of ADHD.

**Evidence (Leontyev & Yamauchi 2019):**
- Maximum acceleration and maximum velocity are the strongest individual predictors of impulsivity
- Mouse movement measures significantly outperform SSRT (traditional stop-signal reaction time) for impulsivity prediction
- ML models trained on mouse data generalize accurately to unseen participants
- Preset (fixed) stop-signal delays > adaptive delays for detecting impulsivity

**Key insight:** SSRT, the established clinical impulsivity measure, has weak association with questionnaire-rated impulsivity. Mouse kinematics fill a measurement gap that the standard clinical measure was missing.

### Eye Tracking

Saccadic and fixation patterns during cognitive paradigms capture attentional control, working memory load, and processing speed.

**Evidence (Liu et al. 2024 — case-control study, children):**
- Assessment paradigm designed specifically around ADHD's core cognitive deficits (not generic eye-tracking)
- Saccadic movement metrics appear particularly discriminative — plausible given dopaminergic involvement in saccadic control
- Framed as "auxiliary diagnosis" — supports, not replaces, clinical assessment
- Task design matters as much as sensing modality

**Key gap:** Study uses dedicated clinical-grade eye-tracking hardware. Consumer webcam gaze estimation accuracy unknown.

### EEG — Brain Engagement Index (BEI')

Single-channel EEG during a 1-minute auditory paradigm captures neural correlates of attention.

**Evidence (Shahaf et al. 2018):**
- CPT BEI' is *trait-specific*: distinguishes ADHD patients from controls (stable diagnostic marker)
- Oddball BEI' is *state-specific*: tracks current attention level within individuals and monitors medication response
- Single channel, 1-minute protocol — far more practical than clinical EEG
- N=20 ADHD + N=10 controls (small; replication needed)

**Key insight:** Trait vs. state distinction is fundamental. Diagnosis requires trait stability; treatment monitoring requires state sensitivity. BEI' provides both, via different paradigms.

## The Assessment Paradigm Problem

A consistent finding across all modalities: **task design matters more than raw sensing**. Eye-tracking built around generic visual tasks performs worse than eye-tracking built around ADHD's specific cognitive deficits. Mouse movement in the context of a stop-signal task captures impulsivity; mouse movement in free use may not. 

The sensor is only as good as the cognitive challenge it's embedded in.

## From Assessment to Monitoring

A distinctive potential of digital biomarkers: continuous monitoring in naturalistic settings (at the desk, not in a clinic). This shifts the use case from diagnosis to:
- **Treatment monitoring** — does medication produce measurable behavioral changes?
- **State tracking** — when is attention at peak vs. trough, in real time?
- **Early warning** — detecting attention failures before they produce errors

The Shahaf BEI' trait/state distinction is the clearest current example of this duality. Leontyev's mouse measures suggest similar potential for real-world impulsivity monitoring.

## Governance Considerations

When digital biomarkers are derived from behavioral data collected for other purposes (productivity monitoring, keyboard analytics, attention tracking), the governance implications are significant. See [[Digital Phenotyping]] for the full treatment.

The core tension: biomarker data collected passively is health data, even if it wasn't collected as health data.

## See Also

- [[ADHD]] — clinical overview and vault coverage map
- [[Digital Phenotyping]] — infrastructure and governance layer; sensing modality foundations
- [[Digital Therapeutics for ADHD]] — the intervention counterpart; biomarkers and therapeutics increasingly coupled in adaptive systems
- [[ADHD in Software Engineering]] — workplace context where behavioral sensing data is already collected by productivity tools
- [[Probability and Statistics Foundations]] — the statistical methods underlying biomarker studies (regression, frequentist inference, confidence intervals); prior sensitivity is particularly relevant for small-n ADHD studies
