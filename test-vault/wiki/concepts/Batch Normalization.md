---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["batch norm", "layer normalization", "layer norm", "normalization layers"]
related: ["Stochastic Gradient Descent", "Transformer Architecture", "Vanishing Gradient Problem"]
sources:
  - "[[archive/books/2026-06-04-neural-network-training-loop]]"
---

# Batch Normalization

Normalizes layer activations to zero mean and unit variance within each mini-batch, then applies learned scale (γ) and shift (β) parameters.

## Effects

- Reduces **internal covariate shift**: activations don't shift distribution as earlier layers update
- Allows **higher learning rates** without instability
- Acts as a **regularizer** (noise from mini-batch statistics)

## Layer Normalization

Normalizes across the feature dimension (per sample) rather than across the batch dimension. **The standard in transformers** — batch norm is problematic with variable sequence lengths and small batch sizes.

```
# Batch norm: normalize across batch for each feature
# Layer norm: normalize across features for each sample
```

## Where Each Is Used

| Architecture | Normalization |
|---|---|
| CNNs (ResNet, VGG) | Batch Normalization |
| Transformers (BERT, GPT) | Layer Normalization |
| RNNs | Layer Normalization |

## Connections

- [[Stochastic Gradient Descent]] — batch norm was motivated by the instability of training with high learning rates; it stabilizes the optimization landscape
- [[Transformer Architecture]] — transformers use layer norm (not batch norm); applied before each sub-layer (Pre-LN) in modern architectures
- [[Vanishing Gradient Problem]] — normalization helps maintain activations in the range where gradients are useful (not saturated)
