---
title: "AI Scraping Resistance"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/videos/2026-06-01-benn-jordan-poison-pilling-music.md"
  - "archive/videos/2026-06-01-addie-lamarr-data-poisoning-surveillance.md"
related:
  - "Digital Surveillance Resistance"
  - "LLM Mental Model"
  - "Community Self-Defense"
  - "Threat Modeling"
tags:
  - ai-resistance
  - privacy
  - creator-rights
  - surveillance
---

# AI Scraping Resistance

Technical resistance against unauthorized AI data collection and algorithmic profiling — not opting out, but actively poisoning the pipeline. Covers two distinct layers of attack:

- **Training-time poisoning** — corrupt the data a model trains on; affects model weights permanently
- **Inference-time poisoning** — corrupt the behavioral signals a running model uses to profile you; disrupts predictions without touching model weights

Distinct from traditional privacy protection (defensive data hygiene) and from legal/policy approaches (robots.txt, ToS, litigation): this is adversarial resistance at the data layer.

## The Framing Shift: Opt-Out vs. Active Poisoning

Most creator protection assumes a cooperative adversary — robots.txt assumes scrapers obey it; Terms of Service assumes compliance. When the adversary doesn't comply, passive opt-out offers no actual protection.

Active data poisoning inverts the logic: instead of asking not to be scraped, deliberately feed the system corrupt data. The goal is not to stop collection but to degrade what's collected.

## Audio Poisoning: Harmony Cloak + Poisonify (Benn Jordan, 2025)

The primary technical source is Benn Jordan's April 2025 video (summarized by TUNED INTO TECH in 2026). Jordan — independent musician, co-founder of consensual-AI platform Voice Swap AI — developed two complementary tools after AI companies scraped his catalog without consent.

**How AI processes music:** neural networks read audio as spectrograms (images of frequency over time). They learn two classes: (1) melody, rhythm, chord progressions — the *musical* patterns; (2) instrument sounds — stem separation, timbre classification.

**Harmony Cloak** (University of Tennessee Knoxville — Sedon Ali Mirza + Gian Lu): adversarial noise targeting class (1). Encodes the file so the spectrogram is unreadable for melody/rhythm extraction. Tensor board validation shows training on Harmony Cloak files plateaus almost immediately — the model stops improving from the start. 30-participant study: perturbed files rated nearly identical quality to clean files by listeners.

**Poisonify** (Jordan): adversarial noise targeting class (2) instrument classifiers. Makes a cymbal read as harmonica; a synthesizer read as string quartet. **Snowball effect**: once the model learns the misclassification, it applies it to all similar sounds in future training — the false positive reinforces and spreads.

**Combined effect**: Harmony Cloak + Poisonify produces music that is not only untrainable for that file but actively degrades the model's overall quality. Live tests:

| Service | Unencoded | Encoded |
|---------|-----------|---------|
| Suno | Coherent extension | "Airport spa music from Napster 1999" |
| Minimax Audio | Coherent fine-tune | Nightmare-fuel output |
| Meta MusicGen | Coherent audio | Hangs and crashes |

**Obfuscation strategy**: Jordan encodes different tracks with different combinations — Harmony Cloak, Poisonify, both, or inaudible no-op noise that does nothing — without disclosing which is which. This denies AI companies the ability to identify the attack vector and train around it. OPSEC applied to data poisoning.

**Live performance protection**: inaudible adversarial signals can disable smart home microphones (Alexa, Echo) from recording during a performance. Jordan demonstrates device-specific audio files from his phone that prevent logging.

## The Pareto Plateau: Existential Threat to AI Music

The 80/20 rule: AI generators quickly improved 80% with 20% of the effort. The remaining 20% improvement is expensive, complicated grinding — and Jordan argues this may be structurally impossible if the best data is poisoned.

The *best* 20% of training data — the most skilled, most distinctive artists — produces 80% of model quality. If that 20% adopts poisoning, the model hits its quality plateau much earlier and much lower. This is not a copyright problem or a PR problem: it's a data quality ceiling that no amount of compute can overcome, because the signal was never there to train on.

High-signal artists have the most leverage and the lowest adoption threshold needed to matter.

## Historical Parallels

- **Operation Mincemeat (WWII)** — feeding false intelligence to shape strategic outcomes
- **Record labels vs. P2P (early 2000s)** — flooding Napster-era networks with degraded files to frustrate users into buying albums

The same tactic now runs in reverse: creators using the music industry's own weapons against the new extractive actor.

## Image Poisoning (Same Family)

- **Nightshade** — imperceptible pixel perturbations that corrupt AI image model training
- **Glaze** — perturbations that cause style-transfer AI to misidentify an artist's style

Audio watermarking, Nightshade, and Glaze all exploit the same structural gap: the difference between human perceptual limits and what models process mathematically. See [[Adversarial Examples]] for the broader class.

## The Arms Race Structure

Each defense creates pressure for the adversary to detect and strip it. Jordan's obfuscation strategy directly addresses this: if the technique can't be identified, it can't be patched. As tools become APIs via distributors like Symphonic Distribution, the scale of adoption changes the calculus — at community scale, verification costs become structurally prohibitive.

## Behavioral Poisoning Against Surveillance (Inference-Time)

Profiling and surveillance algorithms assume behavioral engagement reflects genuine interest — that's the attack surface. Deliberately feeding false signals disrupts the model's predictions about you without affecting its weights.

**Techniques:**
- Watch out-of-context videos; generate misleading engagement signals
- Non-authentic clicks and browsing patterns
- Cultivate contradictory personas across platforms
- Introduce conflicting signals that erode profiling accuracy over time

**The 0.01% attack surface**: research on a 400M-image dataset showed poisoning just 0.01% of data could alter model behavior for ~$60. Counterintuitively, *larger* surveillance datasets are *more* vulnerable — the poisoned minority shrinks as a fraction of total noise and is statistically harder to detect. Scale creates a comforting story of robustness while creating a gigantic shadow map of places to hide poison.

**Anthropic backdoor research**: separately, dropping ~250 carefully crafted documents into a training corpus of *billions* was enough to reliably backdoor a language model regardless of model size. The backdoor triggers on a specific pattern while the model appears normal everywhere else.

**Taleb's minority rule**: Nassim Taleb's concept from *Skin in the Game* — a small, non-scalable, intransigent minority can disproportionately dictate the behavior of an entire ecosystem. Both the 0.01% stat and behavioral poisoning are instances of this: a committed minority introducing false signals can shift what the whole system "believes."

**The feedback loop problem**: "AI learns from us and we adapt to AI." Algorithms shape behavior, which retrains models, which shapes behavior further. This creates a systemic instability — not just individual self-defense but a structural intervention if adopted at scale.

**Model collapse**: modern AIs are continuously refreshed on the open web. But the web is now filling with AI-generated content. Models trained on their own recycled output become blander, more generic, and less connected to reality — slow-motion self-poisoning through the feedback loop.

**High-stakes consequence**: parole eligibility, immigration status, insurance pricing, credit scores — and specifically systems like Palantir's that aggregate billions of data points into risk scores — are determined by profiling systems with the same fragility. When an algorithm can infer a miscarriage from shopping patterns, the stakes are life and death.

**Limitation**: behavioral techniques require no tooling but their effectiveness at individual scale is unclear; collateral effects (degraded recommendations, fraud detection triggers) are real.

## Collective Action Problem

Both training-time and inference-time poisoning only create meaningful friction at community scale. Individual implementation has limited impact:
- Training data: needs a non-trivial fraction of a billions-token corpus
- Behavioral profiling: needs enough participants to make the poisoned signal statistically significant

This mirrors [[Community Self-Defense]]'s core insight: individual protective action is limited; collective, coordinated action changes the calculus for the adversary.

## Legal Ambiguity

The legal status of deliberately poisoning a system that is itself acting unlawfully (violating ToS, copyright, or conducting unauthorized surveillance) is unclear. Does actively corrupting data constitute computer fraud even against an unauthorized accessor? This remains an open question as of 2026.

## Relationship to Other Resistance Frameworks

**[[Digital Surveillance Resistance]]** — defensive posture (communication security, device hygiene); AI scraping resistance adds an offensive layer. The adversaries differ — corporate AI scrapers vs. state surveillance — but the posture is shared: technical resistance when policy fails.

**[[Threat Modeling]]** — both threat models map cleanly onto adversary/asset/countermeasure structure. Training-time: adversary = scraper, asset = creative work/training corpus, countermeasure = subtitle/image poisoning. Inference-time: adversary = profiling system, asset = accurate behavioral profile, countermeasure = signal noise injection.

**[[LLM Mental Model]]** — the frozen-after-training property cuts both ways: training data poisoning is permanent (no runtime correction); behavioral poisoning exploits the same property in real-time profiling systems trained on behavioral corpora.

**[[Training Poisoning Degradation at Scale]]**, **[[Data Poisoning Legal Status]]**, **[[Behavioral Poisoning Adoption Threshold]]**, **[[Poisoning Counter-Adaptation Timeline]]** — open questions on data poisoning
