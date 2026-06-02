---
title: "Adversarial Examples"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources: []
related:
  - "AI Scraping Resistance"
  - "LLM Mental Model"
  - "Neural Network Mechanics"
tags:
  - machine-learning
  - adversarial-ml
  - security
  - ai-resistance
---

# Adversarial Examples

Inputs specially crafted to mislead a machine learning model while appearing normal — or being indistinguishable — to a human observer. First formally described by Szegedy et al. (2014), who found that neural network image classifiers could be reliably fooled by adding imperceptible pixel noise.

## The Core Phenomenon

Neural networks process inputs as high-dimensional vectors and learn decision boundaries shaped by the training distribution. Those boundaries are smooth near training examples but can be sharp and arbitrary in directions that don't vary much in natural data. Adversarial examples exploit this: a small step in a carefully chosen direction crosses a decision boundary without visibly changing the input.

Human perception is robust in these directions — we recognize a panda whether or not it has imperceptible noise added. Models are not. The attack works because human and machine "similarity" operate on different geometries.

## Attack Types

**White-box attacks** have full gradient access to the target model:
- **FGSM** (Fast Gradient Sign Method) — one step in the sign of the gradient; cheap and surprisingly effective
- **PGD** (Projected Gradient Descent) — iterated FGSM with projection back to an ε-ball; stronger and more robust

**Black-box attacks** work without model access, exploiting *transferability* — adversarial examples crafted for one model often fool other models trained on the same data distribution. This is the property that makes audio/image poisoning tools viable against proprietary commercial scrapers.

## Transferability

The transferability property is what makes adversarial examples practically threatening (and practically useful for resistance):
- A perturbation that fools one image classifier tends to fool others
- This holds across architectures, because adversarial directions roughly align with the class of models trained on the same distribution
- Nightshade and Harmony Cloak rely on this: they're crafted against known architectures but transfer to commercial systems the attacker never touches

## AI Scraping Resistance Tools

[[AI Scraping Resistance]] weaponizes adversarial examples for creator protection rather than system compromise:

| Tool | Target | Attack type |
|------|--------|-------------|
| **Nightshade** | Image style extractors | Training-time; corrupts concept associations |
| **Glaze** | Artist style recognition | Inference-time; causes misidentification of style |
| **Harmony Cloak** | Audio melody/rhythm extractors | Training-time; spectrogram unreadable |
| **Poisonify** | Instrument classifiers | Training-time; triggers misclassification snowball |

All exploit the perceptual gap: the perturbation is inaudible/invisible to humans but mathematically decisive for the model.

## Relationship to ML Safety

Adversarial robustness is a core open problem in ML safety. Standard defenses (adversarial training, certified defenses, randomized smoothing) impose significant accuracy/cost trade-offs and have been repeatedly broken by stronger attacks. This is an active research area with no settled solutions as of 2026.

The cat-and-mouse structure: each defense generates a stronger attack surface; each attack generates pressure for better defenses. See [[AI Scraping Resistance]] for the practical implication of this for creator tools (Jordan's obfuscation strategy directly addresses it).
