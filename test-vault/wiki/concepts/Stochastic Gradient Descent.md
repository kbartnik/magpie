---
tags: [concept, ml-fundamentals]
cluster: ml-fundamentals
aliases: ["SGD", "mini-batch gradient descent", "gradient descent", "optimizer", "learning rate"]
related: ["Backpropagation", "Loss Functions", "Batch Normalization", "Transformer Architecture", "Vanishing Gradient Problem"]
sources:
  - "[[archive/books/2026-06-04-neural-network-training-loop]]"
---

# Stochastic Gradient Descent

The optimization algorithm that trains neural networks by repeatedly computing gradients on mini-batches and updating weights in the direction that reduces loss.

## Algorithm

```
for each epoch:
    shuffle training data
    for each mini-batch of size B:
        compute gradient of loss over mini-batch
        w = w - lr × gradient
```

## Mini-Batch Noise as Regularizer

Full-batch gradient descent computes the exact gradient but is too expensive for large datasets. Mini-batch SGD approximates it with noise. This noise is not purely a downside: it helps the optimizer escape sharp local minima and find flatter minima that generalize better.

## Learning Rate

The single most important hyperparameter. Standard schedule for large models:
1. **Warmup:** Linear increase from near-zero to peak LR for the first N steps
2. **Cosine decay:** Smooth decrease from peak to near-zero

Warmup prevents unstable early updates; cosine decay improves final quality.

## Adam and AdamW

Adaptive optimizers track per-parameter gradient history to set effective learning rates:
- **Adam:** Momentum (first moment) + adaptive scale (second moment)
- **AdamW:** Adam + weight decay decoupled from the gradient update

AdamW is the standard optimizer for large language models.

## Connections

- [[Backpropagation]] — backpropagation computes the gradients that SGD uses; they're the input-output pair of the optimization loop
- [[Loss Functions]] — SGD minimizes the loss; the loss function shape determines the optimization landscape
- [[Batch Normalization]] — batch norm stabilizes training by reducing internal covariate shift; allows higher learning rates
- [[Transformer Architecture]] — transformers are trained with AdamW + warmup/cosine decay universally
