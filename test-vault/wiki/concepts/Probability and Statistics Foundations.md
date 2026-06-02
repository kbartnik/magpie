---
title: "Probability and Statistics Foundations"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/books/2026-06-01-seeing-theory.md"
related:
  - "LLM Mental Model"
  - "AI Productivity Research"
  - "Digital Biomarkers for ADHD"
tags: [probability, statistics, inference, foundations]
---

# Probability and Statistics Foundations

Core vocabulary for reasoning under uncertainty. Source: [[Seeing Theory]] (Kunin et al., Brown University 2018); companion visualizations at seeing-theory.brown.edu.

## The Core Distinction: Theoretical vs. Empirical

**Theoretical probability**: the true, underlying probability of an event — a property of the process, not of any particular outcome.

**Empirical frequency**: the observed frequency in a finite sample — always noisy, always different from the theoretical probability.

Every inferential error traces back to conflating these two. A p-value of 0.03 doesn't mean "there's a 97% chance the hypothesis is true" — it's a statement about the frequency of test statistics at least as extreme as observed *under the assumption the null is true*, in a hypothetical infinite repetition of the experiment.

## Probability Basics

**Sample space (Ω):** the set of all possible outcomes.  
**Event:** a subset of Ω.  
**Probability P(A):** a number in [0,1] satisfying:
- P(Ω) = 1
- P(A ∪ B) = P(A) + P(B) for mutually exclusive A, B

**Conditional probability:** P(A|B) = P(A ∩ B) / P(B) — the probability of A given B has occurred.

**Independence:** A and B are independent if P(A ∩ B) = P(A)·P(B), equivalently P(A|B) = P(A). Knowing B gives no information about A.

**Bayes' theorem:**
```
P(A|B) = P(B|A) · P(A) / P(B)
```
The probability of A given B equals the probability of B given A, scaled by the ratio of base rates. This is the mathematical foundation for updating beliefs given evidence.

## Probability Distributions

A **random variable** X maps outcomes to numbers. Its **distribution** describes the probability of each value.

| Distribution | Shape | Use |
|-------------|-------|-----|
| **Uniform** | Flat | Equal probability for all outcomes |
| **Binomial** | Bell-shaped for large n | Count of successes in n independent trials |
| **Normal (Gaussian)** | Symmetric bell | Sums of many independent variables (CLT) |
| **Poisson** | Right-skewed | Count of rare events in fixed interval |
| **Exponential** | Right-skewed | Time between independent events |

**Central Limit Theorem:** The mean of n independent, identically distributed random variables approaches a normal distribution as n grows, regardless of the original distribution. This is why the normal distribution appears everywhere in statistics — not because data is normal, but because *averages* are.

## Frequentist Inference

The classical framework. Parameters are fixed but unknown; data is random.

**Hypothesis testing:**
1. State H₀ (null hypothesis) and H₁ (alternative)
2. Choose a test statistic that measures evidence against H₀
3. Compute the **p-value**: probability of observing a test statistic at least as extreme, *if H₀ were true*
4. Reject H₀ if p < α (typically 0.05)

**What p-values are NOT:**
- Not the probability the null is true
- Not the probability the result is due to chance
- Not the probability that replication will succeed

**Confidence intervals:** A 95% CI is a procedure that, if repeated many times on different samples, would contain the true parameter 95% of the time. It does NOT mean "95% probability the parameter is in this interval" (the parameter is fixed; it either is or isn't).

## Bayesian Inference

The alternative framework. Parameters have probability distributions; data updates those distributions.

**Prior P(θ):** belief about parameter θ before seeing data.  
**Likelihood P(data|θ):** probability of observed data given θ.  
**Posterior P(θ|data):** updated belief after seeing data, via Bayes' theorem:
```
P(θ|data) ∝ P(data|θ) · P(θ)
```

**Credible interval:** a 95% Bayesian credible interval *does* mean "95% probability the parameter is in this interval" — but relative to the prior, not in a frequency sense.

**Prior sensitivity:** Bayesian conclusions depend on prior choice. With large samples, the data overwhelms the prior. With small samples (common in ADHD research), the prior matters significantly — a diffuse prior and an informative prior can produce different posteriors from the same data.

## Regression Analysis

A model for the relationship between variables.

**Simple linear regression:** Y = β₀ + β₁X + ε, where ε is random noise.
- β₀: expected value of Y when X = 0
- β₁: expected change in Y per unit change in X
- ε: residual — what the model doesn't explain

**Key distinction:** regression describes *association*, not *causation*. β₁ ≠ 0 means X and Y co-vary; it doesn't mean X causes Y. Causal claims require experimental design (randomization) or causal modeling (DAGs).

**R²:** fraction of variance in Y explained by the model. R² = 0.3 means the model explains 30% of the variation; 70% remains unexplained.

## See Also

- [[LLM Mental Model]] — language models are probability distributions over token sequences; the theoretical/empirical distinction applies directly
- [[AI Productivity Research]] — the METR RCT uses frequentist inference; reading it critically requires understanding confidence intervals and p-values correctly
- [[Digital Biomarkers for ADHD]] — passive sensing studies use regression and hypothesis testing throughout; the prior sensitivity issue matters for small-n ADHD studies
