---
title: "Autonomous Learning Architecture"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "[[archive/papers/2026-06-01-dupoux-lecun-malik-autonomous-learning]]"
related:
  - "[[Neuromorphic and Bio-Inspired AI]]"
  - "[[LLM Mental Model]]"
  - "[[Agent Memory Architectures]]"
  - "[[Bio-Inspired Computing]]"
tags:
  - ai
  - architecture
  - learning
  - cognitive-science
---

# Autonomous Learning Architecture

A proposed framework by Dupoux, LeCun, and Malik (2026) for AI systems that can learn post-deployment — adapting continuously to new data without human-curated retraining pipelines. The central claim: autonomous learning is not a feature to bolt on, but an architectural property that must be built in from the start.

## The Problem: Externalized Learning

Current AI systems have learning fully outsourced to human experts:
- Data scientists collect and curate training data
- ML engineers engineer loss functions and training recipes
- The deployed model is then frozen — it learns nothing from operational experience

This creates an unavoidable **domain mismatch** problem: real-world data is heavy-tailed and non-stationary. Models trained on fixed datasets fail unpredictably when confronted with data outside their training distribution. Larger pretraining datasets address the distribution partially but can never cover the long tail of novel real-world cases.

Children continuously adapt; deployed LLMs do not. *You can't retrofit post-deployment learning onto a frozen weight matrix.*

## The Three-System Architecture

### System A — Learning from Observation
The passive learning system. The agent observes the world and builds a statistical or predictive model of its input data without taking actions. In current AI: Self-Supervised Learning (SSL), language modeling, vision transformers.

From cognitive science: infants' perceptual specialization (face discrimination, phoneme learning) is System A — driven by passive observation of environmental statistics.

### System B — Learning from Action  
The active learning system. The agent takes actions, observes feedback from the environment, and adjusts behavior. In current AI: Reinforcement Learning (RL), reward-based fine-tuning.

From cognitive science: a toddler exploring a toy randomly to discover its properties; instrumental conditioning.

### System M — Meta-Control
The novel proposed piece. A controller that:
- Monitors the internal state of System A and System B
- Autonomously decides which learning mode to activate and when
- Generates the "data curation" and "loss engineering" decisions that currently require human experts
- Enables higher-order learning modes: learning through communication (following verbal instruction), learning through imagination (mental simulation)

System M is what allows the agent to switch between observation mode and action mode as a function of what the environment demands — reproducing the flexible learning strategy of biological organisms. Without System M, System A and System B remain siloed subfields, requiring human orchestration to combine.

## Three Roadblocks to Building This

1. **Conceptual fragmentation** — System A and System B have developed as separate subfields with different terminology, methods, and communities. Integration requires a unified framework that treats them as complementary modes of a single learning system.

2. **Externalization of learning** — Current training pipelines are optimized for the assumption that humans will orchestrate learning. Autonomous learning requires making that orchestration internal to the system — System M must perform what ML engineers currently do by hand.

3. **No effective methods at scale** — Building System M requires solving how to jointly learn the meta-controller and the initial states of System A and B. The paper proposes a **bilevel evolutionary optimization** approach: outer loop evolves System M; inner loop runs System A/B. This is computationally expensive and largely unvalidated at frontier scale.

## Relationship to Current AI

Existing paradigms partially implement the framework:
- LLMs = System A (large-scale passive observation of text)
- RL fine-tuning (RLHF, DPO) = System B applied to language models
- But the combination is always orchestrated by humans — there is no System M in any deployed system

The hyperscaling bet (bigger models, more data, more compute) scales System A. The paper argues this has diminishing returns and doesn't address domain mismatch, non-stationarity, or the learning modes (communication, imagination) that depend on System M.

## Connection to Bio-Inspired Hardware

The [[Neuromorphic and Bio-Inspired AI]] page documents hardware companies (Cortical Labs, FinalSpark, Eon Systems) building actual biological learning systems. The Eon Systems *Drosophila* result — complex behaviors emerging from architectural fidelity alone, no training required — is cited by Dupoux et al. as evidence that learning capacity may be structural, not parametric.

The System A/B/M framework is the *software design* counterpart to these hardware efforts: both are arguing that the architecture, not just the scale, is what enables biological-quality learning.

## Connections

- [[Neuromorphic and Bio-Inspired AI]] — the software-side counterpart; Onto AI's 4-layer brain (Sensorium/Hippocampus/Cortex/Meta-Cognition) is an independent implementation attempt targeting the same architectural gap
- [[Agent Memory Architectures]] — System M is a principled architecture for what the Hebbian Knowledge Graph (Architecture 4) approximates; both try to move learning out of frozen model weights into a persistent adaptive structure
- [[LLM Mental Model]] — the "frozen after training" property is structural, not incidental; this paper explains why
- [[Bio-Inspired Computing]] — bilevel evolutionary optimization (proposed for System M) is directly in this domain
- [[System M Achievability]] — open question: whether System M is architecturally realizable at inference time
