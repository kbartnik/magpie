---
title: "Superposition Hypothesis"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/videos/2026-06-01-3b1b-how-llms-store-facts.md"
related:
  - "Transformer Architecture"
  - "LLM Mental Model"
tags:
  - llm-fundamentals
  - interpretability
  - deep-learning
---

# Superposition Hypothesis

A hypothesis in AI mechanistic interpretability: neural networks represent far more features than they have dimensions by encoding features as *nearly orthogonal* directions rather than strictly orthogonal ones. This explains why individual neurons appear polysemantic and why model capability scales disproportionately with size.

---

## The Core Geometry

**Strict orthogonality constraint:** If each feature must occupy a direction perpendicular to all others (so features don't interfere), the maximum number of features = number of dimensions. For GPT-3's 12,288-dim embedding space, this caps features at 12,288.

**Relaxed constraint:** Allow features to be *nearly* perpendicular — say, 89–91° apart rather than exactly 90°. By the **Johnson-Lindenstrauss lemma**, the number of nearly-orthogonal vectors that fit in an n-dimensional space grows **exponentially** with n.

Empirical illustration: 10,000 vectors in a 100-dimensional space can be optimized to all fall within a 1° band around 90° of each other. The space stores 100× more "independent" directions than its dimension count would suggest.

**Trade-off:** Each pair of nearly-orthogonal features introduces small cross-feature noise. The model tolerates this noise in exchange for massively expanded representational capacity.

---

## Consequences for LLMs

**Polysemantic neurons:** If features are superimposed across many neurons, no single neuron corresponds to a single feature. Instead, a feature activates a *specific combination* of neurons — a superposition. This is why neurons in trained models respond to apparently unrelated inputs.

**Scaling disproportionality:** Doubling the number of dimensions doesn't double feature capacity — it exponentially expands the number of nearly-orthogonal directions. This may partially explain why larger models don't just perform better proportionally but seem to develop qualitatively new capabilities.

**Fact storage:** MLP blocks in Transformers likely exploit superposition to store vastly more facts and associations than their neuron count would linearly suggest. See [[Transformer Architecture]] § Ch. 7.

---

## Implications for Interpretability

Because features are superimposed rather than neuron-aligned, standard techniques that probe individual neurons (activation patching, neuron ablation) can miss or misattribute features.

**Sparse autoencoders (SAEs):** The primary tool for extracting the true feature basis. An SAE is trained to reconstruct a layer's activations using a sparse combination of a large overcomplete dictionary of learned directions. The hypothesis is that these directions approximate the model's actual features, even though they're not visible in individual neuron activations. Anthropic's interpretability team has published extensively on SAE-based feature extraction.

---

## Open Questions

- Do sparse autoencoders find features the model actually uses, or a convenient decomposition that may not match internal computations?
- Is superposition uniform across layers, or do earlier layers use strict orthogonality while later layers leverage superposition more aggressively?
- How does the degree of superposition interact with training — do models learn to use superposition, or does it emerge passively from gradient descent?

---

## See Also

- [[Transformer Architecture]] — MLP blocks as the likely home of factual knowledge; superposition in the neuron space
- [[LLM Mental Model]] — why LLMs seem to "know" far more than their parameter count would linearly suggest
- [[Neural Network Mechanics]] — the basic neural net structure (matrix × nonlinearity) that MLP blocks instantiate
- [[Sparse Autoencoder]] — the primary tool for empirically investigating and exploiting superposition; dedicated page
