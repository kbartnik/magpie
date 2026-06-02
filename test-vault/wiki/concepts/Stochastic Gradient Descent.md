---
title: "Stochastic Gradient Descent"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-3b1b-backpropagation-intuitively.md"
related:
  - "[[Neural Network Mechanics]]"
  - "[[Transformer Architecture]]"
tags:
  - llm-fundamentals
  - deep-learning
  - optimization
---

# Stochastic Gradient Descent

The practical algorithm for training neural networks. Full gradient descent — computing the exact gradient over all training examples before each parameter update — is too slow for datasets of millions or billions of examples. SGD approximates it using mini-batches.

## The Algorithm

1. Randomly shuffle the training data
2. Divide into **mini-batches** (~32–512 examples each)
3. Compute backpropagation on one mini-batch → approximate gradient
4. Step in the negative gradient direction (update all weights and biases)
5. Move to the next mini-batch; repeat until convergence

Each mini-batch gives a noisy approximation of the true gradient over the full dataset — cheap to compute, close enough to be useful.

## Why "Stochastic"

Each mini-batch is a random sample. The gradient estimate fluctuates from batch to batch — hence "stochastic" (random). A full-batch gradient descent step would be exact but requires the entire dataset. A single-example step (pure SGD) is maximally cheap but maximally noisy. Mini-batch SGD is the practical middle ground.

## Why Noise Can Help

Counterintuitively, the noise in mini-batch gradient estimates can improve final model quality. Full-batch gradient descent follows the exact gradient and reliably converges to the nearest local minimum — which may be shallow or mediocre. The noise in SGD can push the parameter path over small bumps and into deeper, better minima.

3B1B's framing: "a drunk man stumbling down a hill but taking quick steps rather than a precise surveyor taking perfect steps."

## SGD vs. Adam

Modern deep learning rarely uses vanilla SGD. **Adam** (Adaptive Moment Estimation) adds per-parameter adaptive learning rates and momentum — effectively a smarter variant of SGD that converges faster and is more robust to learning rate choices. The underlying principle is the same: mini-batch gradient estimation + parameter update.

## See Also

- [[Neural Network Mechanics]] § Ch. 3 — full derivation of backpropagation and SGD mechanics
- [[Transformer Architecture]] — SGD (or Adam) is how all Transformer parameters are learned during training
- [[Vanishing Gradient Problem]] — why the gradient signal can weaken or disappear in deep networks
