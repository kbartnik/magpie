---
title: "Sparse Autoencoder"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-3b1b-how-llms-store-facts.md"
  - "archive/clippings/2026-06-01-welch-labs-how-ai-images-videos-work.md"
related:
  - "[[Superposition Hypothesis]]"
  - "[[Transformer Architecture]]"
  - "[[Multimodal AI]]"
tags:
  - llm-fundamentals
  - interpretability
  - deep-learning
---

# Sparse Autoencoder

A tool from mechanistic interpretability used to extract the true feature basis from neural network activations. The [[Superposition Hypothesis]] predicts that features are encoded as nearly-orthogonal directions superimposed across many neurons — individual neurons don't cleanly represent features. Sparse autoencoders (SAEs) attempt to recover the underlying features.

## The Problem They Solve

A trained network's internal activations are a superposition of many features. When you probe an individual neuron, it responds to unrelated concepts (polysemanticity) because it participates in many superimposed features simultaneously. There is no direct way to read out "the features this layer knows about."

## Architecture

An SAE has a simple encoder-decoder structure applied to a specific layer's activation vector `x`:

```
h = ReLU(W_enc · x + b_enc)    # encode: expand to large sparse space
x̂ = W_dec · h + b_dec          # decode: reconstruct from sparse code
```

The dictionary `W_dec` is overcomplete: far more columns than the activation dimension. For GPT-3's 12,288-dim embedding space, a SAE might have 50,000–1,000,000 dictionary directions.

**Training objective:** minimize reconstruction error `||x - x̂||²` + sparsity penalty `λ||h||₁`. The sparsity term (L1 on `h`) forces the model to explain each activation using as few dictionary directions as possible.

## What It Finds

If the superposition hypothesis is correct, the SAE's dictionary directions correspond to actual model features — concepts, syntactic patterns, factual associations. An SAE trained on a specific layer of Claude or GPT has revealed:

- Features for specific proper nouns ("the Eiffel Tower"), positions ("in the third paragraph"), tones ("sycophantic"), code constructs
- Features that appear to mediate model behavior (activating the feature for "Assistant" affects compliance)
- Features that are interpretable to humans even though no individual neuron represents them

**Caveat from [[Superposition Hypothesis]]:** Do SAE features represent what the model actually uses for computation, or a convenient decomposition that may not match internal circuits? Active research question.

## Anthropic's Interpretability Work

Anthropic's interpretability team has published extensively on SAE-based feature extraction. Their work on Claude's residual stream has extracted millions of features and is one of the primary empirical bases for the superposition hypothesis.

## Connection to Generative Models

SAEs also appear in generative contexts — [[Diffusion Models]] for image generation involve learned latent representations that can be analyzed similarly. The [[Multimodal AI]] context: CLIP embeddings and diffusion latents are high-dimensional superpositions that SAE-like methods help decompose.

## See Also

- [[Superposition Hypothesis]] — the geometric theory SAEs are designed to test and exploit
- [[Transformer Architecture]] § Ch. 7 — MLP blocks as the primary home of factual storage and the target layer for many SAE studies
- [[Multimodal AI]] — SAE methods extend to vision and multimodal latent spaces
