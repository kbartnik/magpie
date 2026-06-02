---
title: "Vanishing Gradient Problem"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-3b1b-backpropagation-calculus.md"
related:
  - "[[Neural Network Mechanics]]"
  - "[[Stochastic Gradient Descent]]"
tags:
  - llm-fundamentals
  - deep-learning
  - optimization
---

# Vanishing Gradient Problem

In deep neural networks, gradients computed by backpropagation must be multiplied backward through every layer. If each multiplication shrinks the gradient, it can become so small by the time it reaches the early layers that those layers effectively stop learning.

## The Sigmoid Root Cause

The sigmoid activation function `σ(x) = 1 / (1 + e^(-x))` squishes inputs into (0, 1). Its derivative `σ'(x) = σ(x)(1 - σ(x))` has a maximum value of **0.25** — at the midpoint x = 0.

Backpropagation multiplies `σ'(z)` at each layer (the chain rule factor `∂a/∂z`). In a 10-layer network:

```
gradient shrinkage ≤ 0.25¹⁰ ≈ 0.000001
```

Early layers receive a gradient one-millionth the size of the output layer's gradient. Their weights update negligibly or not at all — the network fails to learn useful early features.

## Why ReLU Solves It

**ReLU (Rectified Linear Unit)**: `max(0, x)`

Its derivative is 0 for negative inputs and **1** for positive inputs. Multiplying by 1 doesn't shrink the gradient. A deep network with ReLU activations can propagate gradients back through many layers without exponential decay.

The trade-off: the "dying ReLU" problem — neurons whose input is always negative have derivative 0 and never update. Addressed by variants like Leaky ReLU and ELU.

## Relationship to Architecture Choices

The vanishing gradient problem is one of the core reasons deep learning stalled on sigmoid networks before 2010 and revived with ReLU. It also motivates:

- **Residual connections** (ResNets): add skip connections so gradients have a direct path backward, bypassing the multiplication chain
- **Layer normalization**: keeps activations in a range where gradients stay healthy
- **Careful initialization**: starting weights too large or too small exacerbates the problem

## See Also

- [[Neural Network Mechanics]] § Activation Functions — sigmoid vs. ReLU comparison; why sigmoid is mostly obsolete
- [[Stochastic Gradient Descent]] — the optimization algorithm whose effectiveness depends on gradient quality
- [[Transformer Architecture]] — uses layer norm and residual connections to manage gradient flow across 96 layers
