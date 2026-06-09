---
tags: [concept, learning-science]
cluster: learning-science
aliases: ["spacing effect", "spaced practice", "Anki", "SuperMemo", "SM-2", "forgetting curve"]
related: ["Retrieval Practice", "Desirable Difficulties", "Deliberate Practice", "Working Memory"]
sources:
  - "[[archive/books/2026-06-04-make-it-stick]]"
---

# Spaced Repetition

Reviewing material at increasing intervals produces far more durable retention than massed review. Discovered by Ebbinghaus (1885); among the most replicated findings in cognitive psychology.

## The Mechanism

Forgetting is not the enemy of learning — some forgetting is necessary for spacing to work. The partial forgetting between review sessions forces harder retrieval, which strengthens the memory more than easy retrieval would. Review before complete forgetting ensures the memory survives; the gap ensures the retrieval is effortful.

Optimal schedule: review when you're about to forget — typically 1 day, 3 days, 1 week, 2 weeks, 1 month, escalating.

## Software Implementation

**SM-2 algorithm** (SuperMemo): rates each card after each review (0–5 difficulty); adjusts the next review interval based on difficulty. Anki uses a modified SM-2. The algorithm approximates the optimal review point across thousands of cards simultaneously.

## Massed Practice vs Spaced Practice

| Practice Type | Feels | Produces |
|---|---|---|
| Massed (cramming) | Productive | Short-term fluency |
| Spaced | Harder, slower | Long-term retention |

The feeling of productivity during cramming is deceptive. Performance on a test the next day may be comparable; performance a week later is not.

## Connections

- [[Retrieval Practice]] — spaced repetition schedules retrieval practice at optimal intervals; the two techniques compound
- [[Desirable Difficulties]] — the spacing effect is a desirable difficulty: the forgetting that makes spaced practice feel harder is the mechanism of learning
- [[Deliberate Practice]] — both deliberate practice and spaced repetition work by practicing at the edge of current ability
