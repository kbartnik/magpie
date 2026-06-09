## [2026-06-04 11:41] ingest | The Eleven Patterns Behind Every Production Agentic System

**Files:** archive/clippings/2026-06-04-eleven-agentic-patterns.md
**Summary:** 11 agentic workflow patterns each with a trigger phrase; JSON schema discipline reduces to one diagnostic question (does code branch on this output?); dynamic orchestrators and schemas are compatible.
**Open:** Do 3-4 patterns dominate in practice? What is the debugging story when hub-and-spoke context isolation fails silently?

## [2026-06-04 11:42] ingest | Learning Go, 2nd Edition

**Files:** archive/books/2026-06-04-learning-go-2e.md
**Summary:** All Go design choices stem from making dependencies, failures, and data flow visible in the type system. Interfaces are implicitly satisfied; accept interfaces, return structs. MVS keeps builds reproducible.
**Open:** How does the Go scheduler interact with CGo calls? Is a generics-based context.WithValue replacement being considered?

## [2026-06-04 11:43] ingest | Transformers, the tech behind LLMs | Deep Learning Chapter 5

**Files:** archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms.md
**Summary:** Complete transformer data flow: tokenize→embed→[attn+MLP]×N→unembed→softmax. Embedding geometry is emergent. W_E and W_U are approximate transposes — some models tie the weights entirely.
**Open:** Where are the other ~174B parameters beyond W_E and W_U? What actually breaks with larger context windows — compute, positional encoding, or something else?

## [2026-06-04 11:43] ingest | Notes: ADHD Neurobiology Primer (Rege, PsychSceneHub)

**Files:** archive/papers/2026-06-04-adhd-neurobiology-primer-notes.md
**Summary:** Frontostriatal circuit (PFC↔basal ganglia) is most consistently implicated in ADHD. Structural findings: reduced caudate, PFC white matter, corpus callosum, cerebellar vermis. Secondary source.
**Open:** Does cortical thinning resolution in adolescence predict long-term outcomes? Does ventral striatum size track symptom severity in adults?

## [2026-06-04 11:45] ingest | Go Concurrency Patterns (Rob Pike, Google I/O 2012)

**Files:** archive/videos/2026-06-04-go-concurrency-patterns.md
**Summary:** Channels are values — patterns compose fractally. Generator, fan-in, done-channel, worker-pool are the canonical four. Buffer capacity = concurrency bound. Done-channel is ancestor of context.WithCancel.
**Open:** Does context.Context still use a channel internally? When does a simple channel pool require errgroup?

## [2026-06-04 11:45] ingest | Notes: Cognitive Load Theory and Working Memory (Sweller 1988 / Baddeley 2000)

**Files:** archive/papers/2026-06-04-cognitive-load-working-memory.md
**Summary:** Learning fails from extraneous load exhausting working memory. Central executive (~4 chunks) is the bottleneck. Reducing friction ≠ reducing learning. LLM context window is a functional analogue to WM.
**Open:** Does CLT transfer to expert-level acquisition? Is ADHD specifically a central executive deficit?

## [2026-06-04 11:47] ingest | Implementing LLM Tool Use in Go: A Practical Guide

**Files:** archive/clippings/2026-06-04-llm-tool-use-in-go.md
**Summary:** Tool calling in Go: JSON schema contracts, HTTP client patterns for LLM APIs, function dispatch via reflection or type switches. B↔A bridge: Go HTTP clients implement the execution layer for LLM tool calls.
**Open:** How does Go's type system compare to dynamic languages for tool dispatch? What's the idiomatic retry pattern for rate-limited LLM APIs?

## [2026-06-04 11:47] ingest | Attention in Transformers, Visually Explained | Deep Learning Chapter 6

**Files:** archive/videos/2026-06-04-3b1b-attention-transformers.md
**Summary:** QKV formulation: queries ask questions, keys offer answers, values carry content. Scaled dot-product attention. Multi-head attention runs H parallel heads, each specializing in different relationship types. Causal masking for autoregressive generation.
**Open:** Are early-layer heads consistently syntactic and late-layer heads semantic? What property makes RoPE generalize better to longer contexts?

## [2026-06-04 11:47] ingest | Notes: ADHD and the Nature of Self-Control — Barkley (1997)

**Files:** archive/papers/2026-06-04-barkley-adhd-executive-function.md
**Summary:** ADHD is a disorder of behavioral inhibition, not attention. Inattention and WM deficits are downstream of impaired inhibition. Prospective nonverbal WM specifically impaired — the ability to hold a future state and work toward it.
**Open:** Does 25 years of imaging support the inhibition→WM directionality? Is the hierarchical EF model real or a modeling assumption?

## [2026-06-04 11:47] ingest | What Is Executive Function? The Science of Self-Management

**Files:** archive/clippings/2026-06-04-executive-function-science.md
**Summary:** EF is the bridge between knowing and doing. Three components: inhibition, WM updating, cognitive flexibility. Hot vs cool EF dissociation explains ADHD lab vs real-world divergence. EF outpredicts IQ for life outcomes. AI agents as external EF scaffolding.
**Open:** Is ADHD in high-adversity environments a disorder or adaptive response? Does AI EF scaffolding improve functional outcomes or substitute without building EF?

## [2026-06-04 11:47] ingest | The Go Programming Language — Donovan & Kernighan

**Files:** archive/books/2026-06-04-go-programming-language.md
**Summary:** Go is boring by design — every decision answers: what makes large codebases maintainable by teams that weren't there when code was written. Deferred functions can modify named return values. Interface-based patterns predate generics; both now coexist.
**Open:** Is the interface approach deprecated by generics, or do both coexist as idiomatic? Is reflection use declining with generics?

## [2026-06-04 11:47] ingest | Prompt Engineering for LLMs — Berryman & Ziegler

**Files:** archive/books/2026-06-04-prompt-engineering-for-llms.md
**Summary:** Prompt engineering is applied transformer psychology. System prompts are load-bearing, not cosmetic — instructions placed there are more reliably followed. RAG chapter: semantic chunking over fixed-size chunking. Each technique is a lever on the training distribution.
**Open:** Is CoT genuinely explanatory or confabulated? Does context management advice still apply with 100K+ token windows?

## [2026-06-04 11:49] ingest | Practical Go: Real World Advice for Writing Maintainable Go Programs

**Files:** archive/videos/2026-06-04-dave-cheney-practical-go.md
**Summary:** Every exported name is a cost paid forever. Functional options pattern: backward-compatible config, zero-value defaults, self-documenting call sites. Package names should be short and non-generic.
**Open:** Is closure allocation overhead for functional options measurable in hot paths? How does module versioning affect package naming advice?

## [2026-06-04 11:49] ingest | The Go Memory Model

**Files:** archive/papers/2026-06-04-go-memory-model.md
**Summary:** Happens-before ordering defines cross-goroutine observability guarantees. Data races have undefined behavior — not just stale values. 2022 revision makes this explicit: compiler can legally delete racy code.
**Open:** Are memory model rules intuitive enough to use correctly by default? Can escape analysis changes affect observable behavior?

## [2026-06-04 11:49] ingest | Typing [Generic] Code in Go — GopherCon 2022

**Files:** archive/videos/2026-06-04-go-generics-gophercon.md
**Summary:** Type parameters with constraints solve homogeneous-collection/uniform-algorithm cases. The ~ tilde operator for underlying types preserves Go's defined-type idiom. Generics ≠ interfaces; decision point is heterogeneous vs homogeneous.
**Open:** Is type inference limited on method receivers fundamentally? When does GC shape stenciling produce different performance than full monomorphization?

## [2026-06-04 11:49] ingest | Contexts and Structs: How to Use context.Context in Go

**Files:** archive/clippings/2026-06-04-go-context-patterns.md
**Summary:** context.Context is a cancellation tree, not a dependency container. Four rules: first param ctx, never in structs, Background at top of tree, TODO as placeholder. Logger-in-context is the narrow acceptable exception to no-DI-via-context.
**Open:** Is logger-in-context truly the accepted pragmatic exception? Is there a proposal to type-parameterize context values?

## [2026-06-04 11:53] ingest | State of GPT | Microsoft Build 2023

**Files:** archive/videos/2026-06-04-karpathy-state-of-gpt.md
**Summary:** Base models are document completers; RLHF models are assistants — fundamentally different objects. SFT requires only ~10K examples because capability is already latent in the base model. KL penalty prevents reward hacking in PPO stage.
**Open:** How well does reward model generalize to novel capability areas? How does AI feedback compare to human feedback quality?

## [2026-06-04 11:53] ingest | Embeddings: What They Are and Why They Matter

**Files:** archive/clippings/2026-06-04-embeddings-and-semantic-search.md
**Summary:** Embedding similarity ≠ semantic similarity — it's training distribution similarity. Contrastive training produces geometric proximity for co-occurring concepts. Bias is directly encoded from training data.
**Open:** Does cosine similarity in 1536 dimensions correspond to human intuitions? How sensitive is RAG to chunking strategy?

## [2026-06-04 11:53] ingest | Context Engineering: The Hidden Skill Behind Effective AI Systems

**Files:** archive/clippings/2026-06-04-context-engineering-guide.md
**Summary:** Context engineering is what you put in the context window, not how you phrase it. Most production AI reliability failures are context failures, not prompt failures. Larger context windows shift failure from truncation to dilution.
**Open:** How do you detect noisy context (wrong attention targets)? Does persistent AI scaffold availability prevent EF development?

## [2026-06-04 11:53] ingest | Backpropagation, Visually Explained | Deep Learning Chapters 3 & 4

**Files:** archive/videos/2026-06-04-3b1b-backpropagation.md
**Summary:** Backpropagation is the chain rule applied to computational graphs. Efficient: single backward pass costs same as forward pass by reusing intermediate activations. Vanishing gradients blocked deep learning for decades; ReLU is the trivially simple fix.
**Open:** Is dying ReLU principally solved? When does gradient checkpointing become worthwhile?

## [2026-06-04 11:53] ingest | Understanding Deep Learning — Chapter 5: Training

**Files:** archive/books/2026-06-04-neural-network-training-loop.md
**Summary:** Learning rate schedule is as important as architecture. Mini-batch noise is a feature (finds flat minima that generalize better). Warmup prevents unstable early updates. Flat minima argument is contested.
**Open:** Is flatness of loss landscape a real cause of generalization or a correlation? Does batch norm become less necessary with larger batch sizes?

## [2026-06-04 11:53] ingest | Training language models to follow instructions with human feedback (InstructGPT)

**Files:** archive/papers/2026-06-04-rlhf-instructgpt.md
**Summary:** RLHF doesn't add knowledge — it steers which knowledge surfaces. KL penalty is essential; without it PPO immediately finds degenerate reward-hacking outputs. Alignment tax: slight benchmark degradation for large helpfulness/safety gains.
**Open:** Whose values does the reward model reflect? Does alignment tax shrink or grow with scale?

## [2026-06-04 11:53] ingest | Transfer Learning and Fine-Tuning: The Practical Guide

**Files:** archive/clippings/2026-06-04-transfer-learning-fine-tuning.md
**Summary:** Representational hierarchy is what makes transfer work: early layers are universal, late layers are task-specific. LoRA: 7B model fine-tuned with ~4M parameters (0.06%); weight updates are intrinsically low-rank. The paradigm inverted with BERT: one pre-trained model beats task-specific architectures trained from scratch.
**Open:** Why are task adaptation updates intrinsically low-rank? Does negative transfer worsen for multimodal models?

## [2026-06-04 11:53] ingest | Benchmarking Large Language Models: What We Measure and What It Means

**Files:** archive/clippings/2026-06-04-llm-evaluation-benchmarks.md
**Summary:** Benchmarks decay in validity as models improve and training data expands. HellaSwag/WinoGrande are saturated above human performance without demonstrable common-sense reasoning. Multi-dimensional evaluation (accuracy, calibration, robustness, fairness, efficiency) vs single-number rankings.
**Open:** Is time-locked evaluation contamination-resistant? Can AI judges approximate human evaluation reliably?

## [2026-06-04 11:58] ingest | Make It Stick

**Files:** archive/books/2026-06-04-make-it-stick.md; wiki/concepts/Retrieval Practice.md; wiki/concepts/Spaced Repetition.md; wiki/concepts/Interleaving.md; wiki/concepts/Desirable Difficulties.md
**Summary:** Brown, Roediger, McDaniel on evidence-based learning strategies. Core claim: strategies that feel most effective (re-reading, massed practice) are least effective for durable learning.
**Open:** Cross-cluster bridges: Retrieval Practice ↔ ADHD/Working Memory; Spaced Repetition ↔ Deliberate Practice

## [2026-06-04 11:58] ingest | Ericsson Deliberate Practice

**Files:** archive/papers/2026-06-04-ericsson-deliberate-practice.md; wiki/concepts/Deliberate Practice.md; wiki/concepts/Expertise Development.md
**Summary:** Ericsson et al. 1993 landmark paper. Expertise requires practice at limits of competence with immediate feedback — not accumulated hours. 10,000-hour myth debunked.
**Open:** Software development domain structure problem: unclear performance criteria make deliberate practice hard to structure

## [2026-06-04 11:59] ingest | Metacognition and Self-Regulated Learning

**Files:** archive/papers/2026-06-04-metacognition-self-regulation.md; wiki/concepts/Metacognition.md; wiki/concepts/Self-Regulated Learning.md
**Summary:** Schraw & McCrudden review. Metacognition is the executive control layer of learning. Dunning-Kruger as calibration failure; self-regulation cycle (plan-monitor-evaluate-adjust) is trainable.
**Open:** ADHD connection strong: metacognitive regulation maps directly onto EF deficits

## [2026-06-04 11:59] ingest | ZPD and Scaffolding

**Files:** archive/books/2026-06-04-zpd-and-scaffolding.md; wiki/concepts/Zone of Proximal Development.md; wiki/concepts/Scaffolding.md
**Summary:** Hattie & Donoghue synthesis of Vygotsky's ZPD. Learning happens only within ZPD; scaffolding fades as competence grows. Expertise reversal effect: scaffolding that helps novices burdens experts.
**Open:** AI tutoring connection: persistent AI assistance may prevent scaffold-fading that produces independence

## [2026-06-04 11:59] ingest | Self-Determination Theory

**Files:** archive/papers/2026-06-04-self-determination-theory.md; wiki/concepts/Intrinsic Motivation.md; wiki/concepts/Autonomy.md
**Summary:** Deci & Ryan 2000. Motivation produced by environment (competence + autonomy + relatedness), not a fixed trait. Over-justification effect: rewards undermine intrinsic motivation for already-interesting activities.
**Open:** ADHD dopamine dysregulation may require different motivation support than standard autonomy-support framework

