---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["vanishing gradients", "exploding gradients", "gradient shrinkage", "sigmoid saturation", "deep learning blockage"]
related: ["Backpropagation", "Stochastic Gradient Descent", "Transformer Architecture"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-backpropagation]]"
---

# Vanishing Gradient Problem

In deep networks, gradients shrink exponentially as they propagate backward through layers. Early layers receive negligible gradient updates and don't learn.

## The Mechanism

Gradients are products of derivatives chained through layers. Sigmoid's maximum derivative is 0.25. In a 10-layer network: `0.25^10 ≈ 0.000001`. The gradient vanishes before it reaches early layers.

## Why It Blocked Deep Learning

The theoretical framework (backpropagation, gradient descent) was established in the 1970s–80s. Deep networks couldn't be trained in practice until the 2000s–2010s. The primary blocker was vanishing gradients — and the fix was trivially simple.

## The Fix: ReLU

`ReLU(x) = max(0, x)`

Derivative: 1 for positive inputs, 0 for negative. Gradients for active neurons don't shrink through ReLU layers.

**Dying ReLU:** Neurons with always-negative inputs always output 0 → zero gradient → no learning. Leaky ReLU and GELU mitigate this by allowing small negative outputs.

## Modern Mitigations

- **ReLU and variants** — the primary fix
- **Residual connections (ResNets)** — gradient highway that bypasses layers
- **Batch/Layer normalization** — keeps activations in the productive gradient range
- **Careful initialization** — He/Xavier initialization sets initial weight scales to avoid saturation

## Connections

- [[Backpropagation]] — vanishing gradients are a failure mode of backpropagation in deep networks
- [[Stochastic Gradient Descent]] — without gradient signal, SGD cannot update early layers
- [[Transformer Architecture]] — transformers use layer normalization and residual connections to maintain gradient flow
