---
title: "Understanding Deep Learning — Chapter 5: Training"
type: book
captured-date: 2026-06-05
source-url: ""
author: "Simon J.D. Prince"
publisher: "MIT Press"
year: 2023
isbn: "978-0-262-04869-8"
---

# Understanding Deep Learning — Prince (2023), Chapter 5: Training

**Source:** `understanding-deep-learning-prince.pdf`, Chapter 5: Training
**Coverage:** Loss functions, stochastic gradient descent, mini-batches, learning rate schedules, batch normalization, regularization, early stopping, hyperparameter tuning

## Loss Functions

The loss function measures how wrong the network is. The choice of loss function depends on the task:

**Regression:** Mean squared error (MSE): `L = (1/N) Σ (y_pred - y_true)²`. Penalizes large errors quadratically.

**Binary classification:** Binary cross-entropy: `L = -[y log(p) + (1-y) log(1-p)]`. Penalizes confident wrong predictions more severely than uncertain ones.

**Multi-class classification:** Categorical cross-entropy: `L = -Σ y_c log(p_c)`. Works with softmax output; maximizes log-likelihood of correct class.

The loss function must be differentiable — backpropagation requires gradients, which requires differentiability.

## Stochastic Gradient Descent

Computing the gradient over the entire training dataset (batch gradient descent) is too expensive for large datasets. **Stochastic gradient descent (SGD)** approximates the true gradient using a single example per update. **Mini-batch SGD** — the standard in practice — uses a batch of 32–512 examples:

```
for each mini-batch:
    compute gradient over mini-batch
    update weights: w = w - lr × gradient
```

The noise introduced by mini-batch approximation is not purely a downside — it acts as a regularizer, helping the network escape sharp local minima and find flatter minima that generalize better.

## Learning Rate and Schedules

The learning rate (lr) controls step size. Too high: loss diverges or oscillates. Too low: training is slow and may get stuck.

**Learning rate schedules** vary lr during training:
- **Warmup:** Start with a very small lr, increase linearly for the first N steps. Prevents instability at initialization.
- **Cosine decay:** Decrease lr following a cosine curve from max to near-zero. Improves final model quality.
- **Cyclical learning rates:** Oscillate lr between bounds. Can help escape local minima.

Transformer models (including LLMs) universally use warmup + cosine decay.

## Batch Normalization

Normalizes activations within each mini-batch to zero mean and unit variance, then applies learned scale and shift parameters. Effects:
- Reduces internal covariate shift (activations changing distribution as earlier layers update)
- Allows higher learning rates without instability
- Acts as a regularizer

Batch norm is less common in transformers (where layer normalization is preferred) but foundational in convolutional networks.

## Regularization

**L2 regularization (weight decay):** Adds `λ ||w||²` to the loss, penalizing large weights. Equivalent to placing a Gaussian prior on weights.

**Dropout:** During training, randomly zero out a fraction of neurons' activations. Forces the network to learn redundant representations. Disabled at inference time.

**Early stopping:** Monitor validation loss; stop training when it starts increasing. Prevents overfitting to the training set.

## Deep Read

**Key Insight:** The learning rate schedule is as important as the architecture. Modern large model training uses warmup + cosine decay almost universally — not because it's theoretically optimal but because it's empirically robust across architectures and datasets. The warmup period prevents the unstable early updates that occur when gradient magnitudes are large and the model hasn't yet found a productive region of parameter space.

**What Surprised Me:** Mini-batch noise is a feature, not a bug. Sharp minima (where the loss surface has steep walls) tend to generalize poorly — small changes in the test distribution move you off the minimum. Flat minima generalize well. The noise from mini-batch sampling helps the optimizer find flat minima by "shaking" it out of sharp ones. This is a non-obvious argument for why SGD outperforms full-batch gradient descent in practice.

**Open Questions:**
- The "flat minima generalize better" argument is intuitive but contested. Hochreiter & Schmidhuber (1997) formalized it, but Dinh et al. (2017) showed you can reparametrize a sharp minimum into a flat one without changing generalization. Is flatness of the loss landscape a real cause of generalization or a correlation?
- Batch normalization has hyperparameters (momentum for running statistics, epsilon for numerical stability) and interacts badly with small batch sizes. As hardware pushes toward larger batch sizes, does batch norm become less necessary?
- The learning rate schedule interacts with the optimizer (Adam vs SGD vs AdamW). Is there a principled way to set the schedule given the optimizer, or is it always empirical?

**Wikilink Candidates:**
- [[Stochastic Gradient Descent]] — mini-batch SGD; noise as regularizer; flat vs sharp minima; not yet a wiki page
- [[Loss Functions]] — MSE, cross-entropy, categorical cross-entropy; differentiability requirement; not yet a wiki page
- [[Batch Normalization]] — zero-mean unit-variance normalization; internal covariate shift; less used in transformers; not yet a wiki page

**Connections:**
- [[Backpropagation]] — gradient descent requires gradients; backpropagation computes them; the loss function is the terminal node of the computational graph
- [[Transformer Architecture]] — transformers use AdamW optimizer + warmup/cosine schedule universally; layer norm instead of batch norm; this chapter provides the training-loop context for the architecture description
- [[Vanishing Gradient Problem]] — batch normalization was partially motivated by the vanishing gradient problem; normalizing activations keeps them in the useful gradient range
