---
title: "Backpropagation, Visually Explained | Deep Learning Chapters 3 & 4"
type: video
captured-date: 2026-06-05
source-url: "https://www.youtube.com/watch?v=Ilg3gGewQ5U"
---

# Backpropagation, Visually Explained | Deep Learning Chapters 3 & 4

**Channel:** 3Blue1Brown
**Duration:** ~26 min (Ch.3) + ~10 min (Ch.4)
**Series:** Deep Learning / Neural Networks — Chapters 3 and 4

## Overview

Chapter 3 explains what backpropagation computes and why, using visual intuitions. Chapter 4 works through the calculus explicitly. Together they make the "magic" of neural network training concrete: backpropagation is the chain rule applied to a computational graph.

## What Backpropagation Is

Backpropagation computes the gradient of the loss function with respect to every weight in the network. The gradient tells you how much changing each weight would increase or decrease the loss — it's the information needed to update weights in the direction that improves performance.

The key insight: gradients can be computed in a single backward pass through the network, and this pass costs roughly the same as a single forward pass. This is not obvious — a network with millions of parameters seems to require millions of separate computations, but the chain rule eliminates the redundancy.

## The Chain Rule Connection

If `L = f(g(x))`, then `dL/dx = (dL/df) × (df/dg) × (dg/dx)`.

In a neural network, the loss is a composition of layer operations. Backpropagation computes these derivatives layer by layer, starting from the output and moving backward. Each layer needs to compute two things:
1. The gradient of its output with respect to its inputs (for passing backward)
2. The gradient of the loss with respect to its parameters (for updating weights)

## Gradient Descent Update

Once gradients are computed, weights are updated:
```
w = w - learning_rate × (dL/dw)
```

This moves each weight slightly in the direction that reduces the loss. Repeat over many examples (stochastic gradient descent) and the network learns.

## The Vanishing Gradient Problem

Gradients are products of many derivatives chained together. If each derivative is < 1 (which sigmoid activations guarantee — derivative ≤ 0.25), the product shrinks exponentially with depth. Deep networks trained with sigmoid activations fail to learn because gradients in early layers become negligible.

**ReLU** (max(0, x)) fixes this: its derivative is 1 for positive inputs (no shrinkage). Modern deep networks use ReLU or variants (Leaky ReLU, GELU) in hidden layers.

## Computational Graph Perspective

Modern deep learning frameworks (PyTorch, JAX) implement backpropagation via **automatic differentiation**: they record the forward computation as a graph and compute gradients by traversing this graph in reverse. This allows backpropagation through any differentiable computation, not just the layer types defined in advance.

## Deep Read

**Key Insight:** Backpropagation is efficient because it reuses intermediate computations. The naive approach to computing all gradients would require evaluating the loss function once per parameter — millions of times. The chain rule restructures this into a single backward pass that reuses the activations computed during the forward pass. The whole field of differentiable programming rests on this insight.

**What Surprised Me:** The vanishing gradient problem was the primary blocker for deep learning for decades (roughly the 1990s through the 2000s). The solution — ReLU activations — is trivially simple: just clip negative values to zero. The gap between "we have the theoretical framework" (backpropagation was known since the 1970s) and "we can train deep networks" (2010s) was largely due to this activation function choice.

**Open Questions:**
- ReLU fixes vanishing gradients but introduces "dying ReLU" — neurons that always output zero because their inputs are always negative, producing zero gradient and no learning. Leaky ReLU and GELU mitigate this. Is there a principled way to choose activation functions, or is it still empirical?
- Automatic differentiation (autograd) records the forward computation for replay during backward. For very large models, this graph can be enormous. How do frameworks handle memory for the backward pass — is the full computation graph stored, or is it recomputed?
- Gradient checkpointing trades compute for memory by recomputing intermediate activations during the backward pass instead of storing them. At what network depth does this tradeoff become worthwhile?

**Wikilink Candidates:**
- [[Backpropagation]] — chain rule applied to computational graphs; single backward pass cost; automatic differentiation; vanishing gradient problem; not yet a wiki page
- [[Vanishing Gradient Problem]] — sigmoid derivative ≤ 0.25; exponential shrinkage in deep nets; ReLU as the fix; historical blockage for deep learning; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — transformers are trained via backpropagation; the attention mechanism's QKV matrices are learned by gradient descent through millions of training examples
- [[Softmax]] — softmax is differentiable, which is why it's used at the output layer; the gradient of softmax is well-defined and flows cleanly through backpropagation
- [[RLHF]] — the reward model and SFT model are both trained via backpropagation; the RL stage (PPO) uses policy gradients, which are a different flavor of gradient computation
