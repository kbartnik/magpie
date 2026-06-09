---
type: index
clusters: [agents, go, transformers, adhd, ml-fundamentals, learning-science]
---

# Wiki

## Agents

- [[Agentic Workflow Patterns]] — 11 named patterns with triggers; the skill is matching trigger to pattern
- [[Context Engineering]] — most production AI failures are context failures, not prompt failures; context is the instruction set
- [[JSON Schema Discipline]] — single diagnostic question: does code branch on this output?
- [[LLM Tool Calling]] — tool use is structured output with a side effect; the model never executes
- [[Prompt Engineering]] — few-shot beats zero-shot for format; chain-of-thought beats direct for multi-step reasoning
- [[Retrieval-Augmented Generation]] — RAG trades hallucination risk for retrieval noise; chunk size is the key parameter
- [[Go HTTP Client Patterns]] — client-per-dependency, shared transport pool, context-driven cancellation

## Go

- [[Go Channel Concurrency Patterns]] — generator, fan-in, done channel; channels are values, patterns compose
- [[Go Context Patterns]] — context.Context is a cancellation tree, not a DI container; values only for request-scoped data
- [[Go Error Handling]] — errors are values; %w for wrapping; errors.Is/As for unwrapping
- [[Go Escape Analysis]] — compiler decides heap vs stack; interface conversions and closures are the main escape triggers
- [[Go Functional Options Pattern]] — variadic options avoid boolean telescoping; options are functions, not config structs
- [[Go Generics Type Constraints]] — `~T` underlying-type operator; interfaces as constraint sets; instantiation at compile time
- [[Go Interfaces]] — implicit satisfaction, nil gotcha, accept interfaces/return structs
- [[Go Memory Model]] — data races are undefined behavior per 2022 revision; happens-before via sync primitives
- [[Go Modules and Packages]] — MVS, workspaces, versioning, go.mod
- [[Go Package Design]] — every exported name is a cost paid forever; small surface, large interior
- [[Go Select Statement]] — wait on multiple channels; random selection; timeout and cancellation primitives
- [[Go Structured Logging]] — log/slog native; attrs for machine-readable key-value pairs
- [[Go Testing Patterns]] — table-driven tests, subtests, testdata; use -race in CI
- [[Go Tooling]] — staticcheck, golangci-lint, govulncheck, build tags, go generate
- [[Go Worker Pool Pattern]] — N goroutines on a shared channel; buffer capacity is the concurrency bound

## Transformers

- [[Attention Mechanism]] — query-key-value: which context tokens to weight for each output position
- [[Embeddings]] — similarity ≠ semantic similarity; it's training distribution similarity; dimensions are not interpretable
- [[Multi-Head Attention]] — parallel attention heads; each learns different relationship types; concatenated and projected
- [[Softmax]] — logits to probability distribution; temperature scaling; used in attention too
- [[Tokenization]] — text to token IDs; BPE; why context limits are in tokens not chars
- [[Transformer Architecture]] — end-to-end data flow, embedding geometry, weights vs. data
- [[Vector Similarity]] — cosine similarity is angle, not magnitude; dot product is the fast approximation

## ADHD

- [[ADHD]] — structural findings, frontostriatal circuit, source quality notes
- [[Barkley Inhibition Model]] — behavioral inhibition as the foundational deficit; EF downstream of inhibition
- [[Cognitive Load Theory]] — intrinsic/extraneous/germane load; reducing friction ≠ reducing learning
- [[Executive Function]] — planning, inhibition, working memory, cognitive flexibility; PFC-based; trainable
- [[Frontostriatal Circuit]] — PFC ↔ basal ganglia; most consistently implicated circuit in ADHD
- [[Working Memory]] — Baddeley's four-component model; central executive is the bottleneck; LLM context window analogue

## ML Fundamentals

- [[Backpropagation]] — chain rule on computational graphs; one backward pass costs same as one forward pass
- [[Batch Normalization]] — normalize layer activations to zero mean/unit variance; enables deeper networks
- [[Benchmark Contamination]] — benchmarks decay in validity as training data expands; contamination is structural
- [[Evaluation Metrics]] — no single benchmark number is trustworthy; multi-dimensional evaluation is more robust
- [[Fine-Tuning vs Prompting]] — LoRA fine-tunes 0.06% of parameters; prompting externalizes knowledge at inference cost
- [[Loss Functions]] — the choice of loss encodes the optimization target; cross-entropy for classification
- [[Reward Modeling]] — KL penalty against reference policy prevents reward hacking; essential in RLHF
- [[RLHF]] — base models are document completers; RLHF models are assistants; fundamentally different objects
- [[Stochastic Gradient Descent]] — learning rate schedule as important as architecture; minibatch gradient is noisy
- [[Transfer Learning]] — pretrained representations transfer; LoRA updates are intrinsically low-rank
- [[Vanishing Gradient Problem]] — gradients shrink through backprop in deep networks; residual connections are the fix

## Learning Science

- [[Autonomy]] — self-authored behavior; autonomy ≠ independence; controlling language undermines motivation
- [[Deliberate Practice]] — practice at limits of competence with immediate feedback; not accumulated hours
- [[Desirable Difficulties]] — strategies that feel hardest produce the most durable learning
- [[Expertise Development]] — domain-specific chunking accumulated through deliberate practice; does not transfer across domains
- [[Interleaving]] — ABCABC beats AABBCC at test time; the discrimination demand is the learning mechanism
- [[Intrinsic Motivation]] — produced by competence + autonomy + relatedness; rewards undermine it (over-justification effect)
- [[Metacognition]] — cognition about cognition; calibration; Dunning-Kruger as schema failure, not arrogance
- [[Retrieval Practice]] — testing yourself produces more durable retention than re-reading; effort of retrieval is the mechanism
- [[Scaffolding]] — support calibrated to ZPD that fades as competence grows; expertise reversal effect
- [[Self-Regulated Learning]] — plan-monitor-evaluate-adjust cycle; distinguishes effective from time-on-task learners
- [[Spaced Repetition]] — review when about to forget; partial forgetting forces harder retrieval, strengthening the memory
- [[Zone of Proximal Development]] — learning happens only within ZPD; below is fluency, above is frustration

## Entities

## Questions

## Syntheses
