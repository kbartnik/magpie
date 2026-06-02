---
title: "Diffusion Models"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/videos/2026-06-01-welch-labs-how-ai-images-videos-work.md"
related:
  - "Multimodal AI"
  - "Transformer Architecture"
  - "Superposition Hypothesis"
tags:
  - deep-learning
  - generative-ai
  - image-generation
---

# Diffusion Models

The generative mechanism behind text-to-image and text-to-video models (Stable Diffusion, DALL-E 2, WAN 2.1, Sora). The core idea: corrupt data with noise, then train a network to reverse the corruption. The connection to physics (Brownian motion, Fokker-Planck) is exact, not metaphorical — real algorithms fall out of the physics.

---

## The Forward Process: Adding Noise

An image can be treated as a point in high-dimensional space — one dimension per pixel. Real images occupy a structured manifold within that space.

The **forward diffusion process** repeatedly adds small Gaussian noise to an image across T steps. After enough steps, the image is pure noise — indistinguishable from a sample from a Gaussian distribution. In pixel-space terms, this is a **random walk** away from the image manifold, equivalent to the Brownian motion of a particle diffusing through a medium.

```
x₀ (clean image) → x₁ → x₂ → ... → xT (pure noise)
```

---

## The Reverse Process: Denoising

**Goal:** Learn to run the forward process backwards — start from pure noise and recover a realistic image.

A neural network is trained to predict, at each step t, the direction from a noisy image xₜ back toward less-noisy data. At inference time, start from pure noise xT and iteratively apply the network.

### DDPM Training Objective

*Berkeley team (Ho et al., 2020)*

**Naive objective:** Given xₜ, predict xₜ₋₁. Problem: the direction of the last noise step is random → high variance training signal.

**DDPM insight:** Instead, train the model to predict the **total noise added** across the entire walk — the vector from xₜ back to the original x₀. This is mathematically equivalent (provable via ELBO) but has much lower variance.

The network learns a **time-varying score function** (vector field): at each point in noisy-image space, at each time t, it points toward more probable (less noisy) data.

| Time t | What the vector field encodes |
|--------|------------------------------|
| Large (t ≈ T) | Coarse direction toward the data mean |
| Small (t ≈ 0) | Fine structure of the image manifold |

The model is conditioned on t to learn both scales simultaneously.

---

## Why Random Noise During Generation Is Required

The model learns to predict the **mean** of a conditional Gaussian distribution (proven via the Fokker-Planck equation from statistical mechanics). To actually *sample* from that distribution, Gaussian noise must be added at each step.

Without noise: every generated image converges to the conditional mean — a blurry average of training images. With noise: generated points explore the full distribution.

This is not a heuristic. The stochasticity is load-bearing.

---

## DDIM: Deterministic Sampling

*Song et al. (Stanford) + Google Brain, ~2020*

**Problem with DDPM:** Hundreds of network passes required per image — slow at scale.

**The Fokker-Planck bridge:** The SDE (stochastic differential equation) governing DDPM has a mathematically equivalent **ODE** (ordinary differential equation) that produces the *same marginal distribution* of outputs without any random component. Same training, different sampler.

**DDIM result:** Deterministic generation in significantly fewer steps (~20–50 vs. hundreds), no change to training. The deterministic trajectory follows the contour lines of the learned vector field rather than making noisy jumps.

**Flow matching** (WAN 2.1, Stable Diffusion 3): A generalization of DDIM with a simpler, straighter trajectory formulation. The current standard in production models.

---

## Text Conditioning

Images are generated to match a text description by conditioning the denoising network on a text embedding.

**CLIP** (OpenAI): Joint image-text encoder trained on image-caption pairs. Produces a shared embedding space where similar text and images land near each other. CLIP text embeddings are a natural steering signal for diffusion.

**DALL-E 2 (unCLIP):** Train a diffusion model to invert the CLIP image encoder. Use CLIP text embeddings as the generation target. Result: strong prompt adherence from day one.

**Conditioning alone is insufficient:** The network simultaneously learns "produce realistic images" and "produce images matching the text." The former tends to overpower the latter at inference.

---

## Classifier-Free Guidance (CFG)

The technique that made text-to-image models actually useful.

**Training:** For a random subset of training examples, drop the text conditioning. The model learns both a conditioned mode (text-guided) and an unconditioned mode (realism-guided).

**Inference formula:**

```
output = conditioned + α × (conditioned − unconditioned)
```

Amplifies the component of the conditioned output that is *specifically due to the text*, having subtracted the "be realistic in general" direction. The scaling factor α controls the trade-off:

| α | Effect |
|---|--------|
| 0 | Unconditioned (ignore prompt) |
| 1 | Standard conditioning |
| 2–10 | Strong prompt adherence; may sacrifice diversity |

**Why it works geometrically:** At large t, conditioned and unconditioned vectors point roughly the same direction (toward the data mean). As t → 0, they diverge — conditioned points toward the requested subject. Amplifying the difference steers the trajectory precisely to the right region of the manifold.

**Negative prompts (WAN 2.1):** Instead of subtracting the unconditioned output, subtract the output conditioned on *explicitly described undesired features* ("extra fingers, walking backwards, cartoonish proportions..."). Finer control over what the model avoids.

---

## The Full Pipeline

```
Text prompt
  → CLIP/text encoder → text embedding
  → [Start from pure noise xT]
  → for each step t = T ... 0:
      conditioned_score = model(xₜ, t, text_embedding)
      unconditioned_score = model(xₜ, t, ∅)
      guided_score = conditioned + α × (conditioned − unconditioned)
      xₜ₋₁ = DDIM_step(xₜ, guided_score)
  → x₀: generated image/video
```

WAN 2.1 uses a Transformer backbone for the denoising network (same architecture as [[Transformer Architecture]]). The "denoising network" in modern models is not a simple CNN — it's a full-scale Transformer operating on video patch tokens.

---

## Open Questions

- Flow matching uses straight-line trajectories in latent space. Does the straightness actually correspond to anything semantically meaningful, or is it purely an optimization convenience?
- CFG amplification can push generations off the realistic manifold at high α. What determines where "too far" is — is there a principled bound?
- CLIP was trained for retrieval, not generation. Its embeddings steer diffusion remarkably well despite this. Is this because good retrieval representations and good generative representations are the same thing, or is it a coincidence of architecture?

---

## See Also

- [[Multimodal AI]] — shared embedding space architecture; CLIP as the bridge between text and image modalities
- [[Transformer Architecture]] — the denoising network in modern models is a Transformer; same Q/K/V attention, different input modality
- [[Superposition Hypothesis]] — high-dimensional geometry underlies both why diffusion works at scale and why LLMs can store exponentially many features
