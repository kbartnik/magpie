---
title: "Neural Network Mechanics"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/clippings/2026-06-01-3b1b-what-is-neural-network.md"
  - "archive/clippings/2026-06-01-3b1b-gradient-descent.md"
  - "archive/videos/2026-06-01-3b1b-backpropagation-intuitively.md"
  - "archive/videos/2026-06-01-3b1b-backpropagation-calculus.md"
related:
  - "LLM Mental Model"
  - "Multimodal AI"
  - "Probability and Statistics Foundations"
tags:
  - llm-fundamentals
  - deep-learning
  - mathematics
---

# Neural Network Mechanics

The mathematical machinery underlying neural networks. Built from the 3Blue1Brown deep learning series — updated as each chapter is ingested.

**Series progress:** Ch. 1 ✓ · Ch. 2 ✓ · Ch. 3 ✓ · Ch. 4 ✓ — complete

*Ch. 5+ (Transformers) continued on [[Transformer Architecture]].*

---

## Chapter 1: Structure

*Source: "But what is a neural network?" (3B1B, 2017)*

### The Basic Unit: Neurons and Activations

A **neuron** holds a single number — its **activation** — between 0 and 1. That's all it is. Nothing more.

In the digit-recognition example:
- **Input layer**: 784 neurons (one per pixel of a 28×28 image; grayscale value 0–1)
- **Hidden layers**: two layers of 16 neurons each
- **Output layer**: 10 neurons (one per digit 0–9; the brightest = network's answer)

### The Forward Pass: How Layers Connect

Activations in one layer determine activations in the next. The full transition from layer to layer:

```
a' = σ(W·a + b)
```

| Symbol | Meaning |
|--------|---------|
| `a` | Activation vector of current layer |
| `W` | Weight matrix — rows = next layer neurons, columns = current layer neurons |
| `b` | Bias vector |
| `σ` | Activation function applied elementwise |
| `a'` | Activation vector of next layer |

This is all a neural network layer is: a **linear transformation** (matrix multiply + bias) followed by a **nonlinearity** (activation function). Stacking these is all of deep learning.

### Weights

Each connection between neurons in adjacent layers has a **weight** — a real number. Positive weights increase the downstream neuron's activation; negative weights suppress it. A neuron detecting a specific edge: high positive weights on edge pixels, negative weights on surrounding pixels. The weighted sum is large when the edge is present.

### Biases

Each neuron also has a **bias** — a number added to its weighted sum before the activation function. The bias sets the threshold: how strongly the weighted sum must be before the neuron activates meaningfully.

### Activation Functions

**Sigmoid** (original): `σ(x) = 1 / (1 + e^(-x))`
- Squishes any real number into (0, 1)
- Smooth S-curve; differentiable everywhere
- Biologically motivated (neuron firing rate)
- **Problem**: causes vanishing gradients in deep networks; mostly obsolete

**ReLU** (modern standard): `ReLU(x) = max(0, x)`
- Zero for negative inputs; identity for positive inputs
- Trains dramatically better in deep networks
- Rectified Linear Unit — simpler, but more effective

Without nonlinearity, stacking layers would produce only another linear function — a stack of linear transforms collapses to a single linear transform. The activation function is what gives deep networks their expressive power.

### Parameters and Scale

For the digit network (~13,000 parameters):
- 784 × 16 weights (layer 1→2) + 16 biases
- 16 × 16 weights (layer 2→3) + 16 biases
- 16 × 10 weights (layer 3→4) + 10 biases

**Training** = finding the values of all these parameters that make the network give correct outputs on the training data. The parameters are the "knobs and dials."

### The Network as a Function

The entire network is just a function: 784 inputs → 10 outputs. It's complicated (13,000 parameters, many nonlinearities), but just a function. This framing is clarifying: there's no magic, only a parameterized function whose parameters have been tuned by training.

### Hierarchical Feature Detection (The Hope)

The layered structure is designed to *allow* hierarchical feature learning:
- Layer 2 could detect edges
- Layer 3 could detect patterns (loops, lines made of edges)
- Layer 4 could identify digit components

**Important caveat:** this is the hope, not the guarantee. Whether trained networks actually learn interpretable hierarchical features is an open question — they often don't. The architecture allows it; training finds whatever minimizes the loss, interpretable or not.

---

## Chapter 2: Learning (Gradient Descent)

*Source: "Gradient descent, how neural networks learn" (3B1B, 2017)*

### The Cost Function

**Per-example cost:** Sum of squared differences between actual output activations and desired values.

```
C(example) = Σ (actual_output_i - desired_output_i)²
```

Small when the network is confident and correct. Large when wrong or uncertain.

**Overall cost:** The average across all training examples — a single "lousiness" scalar. It takes all ~13,000 weights and biases as input and outputs one number.

### Gradient Descent

The algorithm for minimizing the cost function over the 13,000-dimensional parameter space:

1. Initialize all weights and biases randomly
2. Compute the gradient of the cost — a vector with one component per parameter
3. Step in the **negative gradient direction** (downhill, toward lower cost)
4. Repeat

Step size is proportional to gradient magnitude: near a minimum the gradient shrinks, so steps naturally get smaller. No guarantee of reaching the global minimum — gradient descent finds local minima.

### What the Gradient Vector Means

Each component of the gradient has two pieces of information:
- **Sign** — nudge this parameter up or down
- **Relative magnitude** — how much impact this parameter has on cost; which changes carry the most "bang for your buck"

The gradient encodes the *relative importance* of every weight and bias for the current training data.

### Backpropagation (Preview)

The algorithm for efficiently computing this gradient is **backpropagation** — applying the chain rule backward through the network. See Ch. 3 below for the full treatment.

### What the Network Actually Learns

When trained, the digit recognition network achieves 96–98% accuracy — but its hidden-layer neurons don't detect the clean edges and patterns the architecture was designed to encourage. They look nearly random. The network found a *different* local minimum: one that works, but isn't interpretable.

**Key implication:** Architecture enables hierarchical features; training doesn't enforce them. Cost minimization finds *any* solution that works, not the one that matches human intuition.

### Confident Nonsense

Feed a random noise image → confident classification as some digit. The network has no concept of "this doesn't look like a digit" because:
1. Training data contained only clean digits
2. The cost function never penalized overconfidence on out-of-distribution inputs
3. The network's "universe" was exclusively valid digit images

This is a training distribution failure, not a reasoning failure.

### Local Minima and Structured Data

The memorization experiment (Leysa Lee): a network can memorize randomly labeled training data given enough parameters, but structured (correctly labeled) data shows a faster, steeper loss curve — easier for gradient descent to exploit.

Local minima found on structured data tend to be of roughly equal quality. "Bad local minima" is less severe a concern in practice than expected.

---

## Chapter 3: Backpropagation (Intuition)

*Source: "Backpropagation, intuitively" (3B1B, 2017)*

### What Backprop Actually Does

Backpropagation computes the gradient by propagating "desired nudges" backward through the network. The key reframe: instead of thinking of the gradient as a direction in 13,000-dimensional space, think of each component's magnitude as measuring how *sensitive* the cost is to that weight. A gradient component of 3.2 vs 0.1 means the cost changes 32× more when you wiggle that first weight.

### The Three Levers

For any neuron you want to make more active, there are exactly three avenues:

| Lever | Effect |
|-------|--------|
| Increase the **bias** | Shifts activation threshold |
| Increase **weights** | Especially those connected to already-bright (active) neurons — those weights multiply larger values |
| Push **upstream activations** | In the direction that increases this neuron — becomes the backprop signal for the previous layer |

The weight lever has a key property: **connections to brighter neurons have more influence**, because the weight multiplies the upstream activation. This means gradient updates are proportional to activation magnitude.

### The Hebbian Echo

This weight-update rule naturally produces Hebbian-like learning: the biggest weight strengthening happens between the most co-active neurons. "Neurons that fire together, wire together."

This is a *mathematical consequence* of minimizing squared error — not something explicitly encoded. 3B1B is careful to note this is a loose analogy; artificial networks may not behave like biological brains. But the structural similarity is striking.

### Propagating Backwards

For a training example showing a "2":
1. Compute desired nudges for the output layer (up on neuron 2, down on all others — proportional to distance from target)
2. For each output neuron, trace back what changes to the second-to-last layer would help
3. Sum all output neurons' desires for the previous layer (weighted by connection strength and how much each output needs to change)
4. This summed signal becomes the "desired nudge" for the second-to-last layer — apply the same process recursively

One example gives one set of desired nudges. Run this for every training example, then **average** the results. That average is (loosely) the negative gradient.

### Stochastic Gradient Descent

Computing the true gradient over all training examples every step is prohibitively slow. The practical solution:

1. Randomly shuffle training data
2. Divide into **mini-batches** (~100 examples each)
3. Compute a gradient step using one mini-batch
4. Move to the next mini-batch; repeat

Each mini-batch gives a good *approximation* of the true gradient with a major computational speedup. The network's path through cost-space resembles "a drunk man stumbling down a hill but taking quick steps" rather than a precise surveyor taking perfect steps. This is **stochastic gradient descent (SGD)**.

**Why SGD can work better than full-batch:** The noise in mini-batch estimates can help escape shallow local minima. Full-batch gradient descent would follow the exact gradient but may settle into worse local optima.

### The Labeled Data Requirement

Backprop requires labeled training data — every example needs a known correct output to compute the nudge direction. MNIST (60,000+ labeled handwritten digits) is the canonical example. Obtaining sufficient labeled data is the central practical challenge in ML, not the algorithm itself.

---

## Chapter 4: Backpropagation (Calculus)

*Source: "Backpropagation calculus" (3B1B, 2017)*

### Notation

Working with one neuron per layer first. Let:

| Symbol | Meaning |
|--------|---------|
| `a^(L)` | Activation of layer L (superscript = layer index, not exponent) |
| `y` | Desired output for this training example |
| `C₀ = (a^(L) - y)²` | Cost for one training example |
| `z^(L) = w^(L)·a^(L-1) + b^(L)` | Weighted sum (pre-activation) |
| `a^(L) = σ(z^(L))` | Activation = nonlinearity applied to z |

Naming the weighted sum `z` is the key notational move — it splits the linear combination from the nonlinearity, giving the chain rule two clean stages to work on.

### The Chain Rule Decomposition

```
∂C₀/∂w^(L) = (∂z^(L)/∂w^(L)) × (∂a^(L)/∂z^(L)) × (∂C₀/∂a^(L))
```

Each factor:

| Factor | Value | Meaning |
|--------|-------|---------|
| `∂C₀/∂a^(L)` | `2(a^(L) - y)` | How wrong is the output? |
| `∂a^(L)/∂z^(L)` | `σ'(z^(L))` | Slope of activation function at z |
| `∂z^(L)/∂w^(L)` | `a^(L-1)` | How active was the previous neuron? |

**The critical line:** `∂z/∂w = a^(L-1)`. The weight gradient *is* the previous activation. This is the mathematical proof of Ch. 3's Hebbian observation — not a loose analogy, a direct consequence of differentiating the weighted sum. Connections to more-active neurons get proportionally larger updates.

### Bias and Backward Signal

**Bias gradient:** `∂z/∂b = 1` always. The bias update collapses to just `σ'(z) × 2(a - y)` — no dependence on connection strength.

**Backward propagation signal:** `∂z^(L)/∂a^(L-1) = w^(L)`. The cost's sensitivity to the previous activation equals the weight. This is what propagates backward: multiply by the weight, then recurse.

### Multiple Neurons: Same Idea, More Indices

With multiple neurons: `k` indexes layer L-1, `j` indexes layer L. Weight `w^(L)_jk` connects neuron k → j (index order matches the weight matrix row/column convention).

Cost sums over all output neurons: `C₀ = Σⱼ (a^(L)_j - y_j)²`

The chain rule for any specific weight looks identical to the single-neuron case. What changes: neuron k in L-1 influences cost through *every* j in L, so its gradient is a sum:

```
∂C₀/∂a^(L-1)_k = Σⱼ [ w^(L)_jk · σ'(z^(L)_j) · 2(a^(L)_j - y_j) ]
```

The apparent complexity in backprop derivations is almost entirely index-chasing, not new concepts.

### What the Code Does

Every line of backprop code implements one of three factors, applied layer-by-layer from output to input:
1. Compute `∂C/∂a` at the output (the error signal)
2. Multiply by `σ'(z)` (activation derivative — this is the backward pass through the nonlinearity)
3. Multiply by `a^(L-1)` for weight gradients; multiply by `w` to propagate the signal to the next layer back

The full gradient vector is the assembly of these partial derivatives across all weights and biases in all layers.

---

## See Also

- [[Transformer Architecture]] — Ch. 5+: how the neural network mechanics here scale into GPT-style models; embedding space, attention, unembedding
- [[LLM Mental Model]] — LLMs use the same weight/parameter/training framework at vastly larger scale; "training changes the model, inference uses it"
- [[Multimodal AI]] — image patch embeddings extend the activation framework to non-text modalities; the shared vector space is a high-dimensional activation space
- [[Probability and Statistics Foundations]] — network outputs are probability distributions; the mathematics of training (cross-entropy loss, optimization) builds on probability theory
- [[Stochastic Gradient Descent]] — the mini-batch approximation algorithm described in Ch. 3; dedicated page
- [[Vanishing Gradient Problem]] — why sigmoid activations are obsolete in deep networks; Ch. 4 context; dedicated page
