---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["loss function", "cost function", "cross-entropy loss", "MSE", "training objective"]
related: ["Backpropagation", "Stochastic Gradient Descent", "Softmax", "Transformer Architecture"]
sources:
  - "[[archive/books/2026-06-04-neural-network-training-loop]]"
---

# Loss Functions

Measures how wrong the network is. The choice of loss function depends on the task and must be differentiable (backpropagation requires gradients).

## Common Loss Functions

**Mean Squared Error (MSE):** `L = (1/N) Σ (y_pred - y_true)²`
- Regression tasks
- Penalizes large errors quadratically

**Binary Cross-Entropy:** `L = -[y log(p) + (1-y) log(1-p)]`
- Binary classification
- Penalizes confident wrong predictions severely

**Categorical Cross-Entropy:** `L = -Σ y_c log(p_c)`
- Multi-class classification; used with softmax output
- Maximizes log-likelihood of the correct class

**Next-Token Prediction (Language Models):** Categorical cross-entropy over the vocabulary. The loss for a sequence is the average per-token loss. Minimizing it maximizes the probability the model assigns to the correct next token.

## Differentiability Requirement

The loss function is the terminal node of the computational graph. Backpropagation computes `dL/dw` for every weight — this requires the loss to be differentiable.

## Connections

- [[Backpropagation]] — the loss function is the root of the computational graph; its derivative starts the backward pass
- [[Stochastic Gradient Descent]] — SGD minimizes the loss; the landscape (convex vs non-convex) determines convergence behavior
- [[Softmax]] — cross-entropy loss is computed from softmax output probabilities; the combination is numerically stable (log-softmax)
- [[Transformer Architecture]] — language models are trained with next-token cross-entropy loss across the full vocabulary
