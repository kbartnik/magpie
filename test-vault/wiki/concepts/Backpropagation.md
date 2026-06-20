---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["backprop", "chain rule", "automatic differentiation", "autograd", "gradient computation"]
related: ["Stochastic Gradient Descent", "Vanishing Gradient Problem", "Transformer Architecture", "RLHF", "Loss Functions"]
sources:
  - "[[archive/videos/2026-06-04-3b1b-backpropagation]]"
---

# Backpropagation

The algorithm that computes the gradient of the loss function with respect to every weight in a neural network. Efficient: a single backward pass costs roughly the same as the forward pass.

## The Key Insight

A neural network is a composition of differentiable functions. By the chain rule:

```
L = f(g(h(x)))
dL/dx = (dL/df)(df/dg)(dg/dh)(dh/dx)
```

Each term is computed once during the backward pass and reused — avoiding the naïve approach of evaluating the loss N times (once per parameter).

## How It Works

1. **Forward pass:** Compute activations at each layer; store intermediate values
2. **Compute loss:** Apply loss function to final output vs. target
3. **Backward pass:** Traverse the computation graph in reverse; at each layer, compute:
   - Gradient of loss w.r.t. inputs (pass backward)
   - Gradient of loss w.r.t. parameters (update weights)

## Automatic Differentiation

Modern frameworks (PyTorch, JAX) record the forward computation as a graph and traverse it in reverse automatically. This enables backpropagation through any differentiable computation.

## Vanishing Gradients

Products of many derivatives shrink exponentially in deep networks if each derivative < 1 (sigmoid: max derivative 0.25). ReLU (derivative = 1 for positive inputs) fixes this.

## Connections

- [[Stochastic Gradient Descent]] — backpropagation computes the gradients that SGD uses to update weights
- [[Vanishing Gradient Problem]] — deep networks with sigmoid activations have shrinking gradients; ReLU is the fix
- [[Transformer Architecture]] — transformers are trained via backpropagation through their QKV matrices and MLP layers
- [[Loss Functions]] — the loss function is the terminal node of the computational graph; its derivative starts the backward pass
- [[RLHF]] — SFT stage uses standard backpropagation; PPO uses policy gradient methods (a generalization)
