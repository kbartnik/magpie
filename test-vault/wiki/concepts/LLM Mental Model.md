---
title: "LLM Mental Model"
type: concept
status: active
created: 2026-05-31
updated: 2026-06-01
sources:
  - "archive/clippings/2026-05-31-ai-without-illusions-part-2.md"
  - "archive/clippings/2026-05-31-ai-without-illusions-part-4.md"
  - "archive/clippings/2026-06-01-llm-slm-fm-model-selection.md"
  - "archive/clippings/2026-06-01-what-is-multimodal-ai.md"
  - "archive/clippings/2026-06-01-karpathy-vibe-coding-to-agentic-engineering.md"
related:
  - "Harness Engineering"
  - "Vibe-Coding Anti-Pattern"
  - "Context Rot"
  - "AI Without Illusions (Series)"
tags:
  - llm-fundamentals
  - ai-discipline
  - cca-f
---

# LLM Mental Model

The mental model that avoids the two common failure modes: treating LLMs as magic minds (over-trust) or dismissing them as "just autocomplete" (under-use). The useful framing is:

> **LLMs are high-capacity pattern learners that generate tokens under uncertainty.**

They are useful when the patterns correlate with what you want, and dangerous when fluent pattern completion is mistaken for truth.

## The Two Wrong Mental Models

**"Magic mind"**: assumes stable beliefs, coherent goals, internal sense of truth. Leads to reading intent into outputs and interpreting confident tone as confident substance. Vulnerable to hallucinations that sound authoritative.

**"Just autocomplete"**: misses how much structure is captured by training at scale. Leads to underestimating both the genuine utility and the genuine risks. A simple system fails simply; this one doesn't.

The key: *a simple outer loop + complex learned internal representation*. The mechanism is simple; what scale does to that mechanism is not.

## Software 1.0 / 2.0 / 3.0 (Karpathy)

A framing for LLMs as a new computing paradigm, not just a faster tool:

| Era | How you program | What you work with |
|-----|----------------|-------------------|
| **1.0** | Write explicit rules | Structured data + deterministic code |
| **2.0** | Arrange datasets, train neural networks | Learned weights (the program *is* the weights) |
| **3.0** | Prompt; manage context window | LLM as interpreter; context window is your lever |

Software 3.0 enables genuinely new things that couldn't exist before — not just speedups of existing tasks. Example: compiling a wiki from raw documents is not "code over structured data," it's information reprocessing that has no Software 1.0 equivalent.

**The spurious app problem.** When a new paradigm arrives, apps built in the previous paradigm become *structurally obsolete* — not just slower, but architecturally unnecessary. A full OCR + image generation pipeline to show menu photos becomes spurious when a multimodal model can do it directly from the same prompt. Many currently-built apps may be in this category without knowing it yet.

## Ghosts, Not Animals

A third mental model alongside the two wrong ones:

> LLMs are "jagged, statistical, summoned entities" — not animal intelligences. — Andrej Karpathy, AI Ascent 2026

Animals have intrinsic motivation, curiosity, and empowerment shaped by evolution. LLMs have none of this. They are statistical simulation circuits: the substrate is pre-training statistics, with RL bolted on top. Yelling at them has no effect. Praising them has no effect.

The framing matters not for its philosophical precision but for its *building implications*: anthropomorphizing (treating as animals) leads to incorrect intuitions about reliability, consistency, and what kinds of interventions work. Treating them as powerful, strange, statistically-grounded ghosts is more predictive.

## Jagged Intelligence and Verifiability

LLMs are *jagged* — they can refactor a 100,000-line codebase and simultaneously tell you to walk to a car wash. The jaggedness isn't noise; it maps to a structural pattern:

**Peaks** (RL-trained, verifiable domains): math, code, formal reasoning. Frontier labs train these models with RL on environments where correctness can be verified — the model improves dramatically in these circuits.

**Rough edges** (unverifiable or low-priority): commonsense spatial reasoning, aesthetics, judgment tasks without clear reward signals. The model wasn't reinforced here.

Two factors compound: (1) verification difficulty — RL requires a reward signal, which requires verifiability; (2) lab prioritization — chess capability in GPT-3.5→4 improved dramatically because someone added chess data, not from general capability growth. You're partly at the mercy of what the labs decided to put in the mix.

**Practical implications:**
- Know which circuits your application uses. In-distribution = the model flies. Out-of-distribution = expect struggle and consider fine-tuning.
- The jaggedness is a reason to keep humans in the loop — not because the model is generally untrustworthy, but because the failures are *unpredictable* and don't correlate with the model's overall capability level.

## Model Scale Taxonomy

LLM is the umbrella term. The sub-labels signal how models are *used*, not what they fundamentally are.

| Label | Scale | Character | Selection signal |
|-------|-------|-----------|-----------------|
| **SLM** (small language model) | <10B parameters | Efficient specialist — optimized for narrow tasks | Need speed, low cost, or on-prem governance |
| **LLM** (large language model) | Tens of billions | Generalist — broad knowledge, handles variability | Need multi-domain synthesis or nuanced reasoning |
| **FM** (frontier model) | Hundreds of billions+ | Cutting edge — defined by agentic capability, not just size | Need multi-step autonomous reasoning with tool use |

**SLMs aren't just worse LLMs.** Well-tuned SLMs can match or beat larger models on focused tasks (classification, summarization, routing) — faster inference, lower cost, and deployable on-premises.

**FMs are defined by agentic capability, not parameter count.** What makes a model "frontier" is the ability to plan multi-step workflows, call APIs, evaluate intermediate results, and adjust approach — not raw size. Plenty of smaller models can use tools; FMs can *orchestrate* tool chains autonomously.

**Governance as a selection driver.** In regulated industries (finance, healthcare), on-prem SLM deployment is often non-negotiable. Data sovereignty requirements set a compliance ceiling that sits below the capability ceiling — the question isn't "which model is best?" but "which models are permissible?"

**The co-pilot reality.** Most production deployments today run FMs as co-pilots with guardrails and human sign-off — not fully autonomous agents. The agentic capability exists in the model; the decision to deploy it autonomously is an organizational and liability question.

## Training vs Inference

**Training** (changes the model):
- Pretraining: next-token prediction across enormous corpora; the model learns to compress regularities in language
- Instruction tuning: supervised fine-tuning on examples of desired behavior; teaches the model to act like an assistant
- Preference optimization (RLHF/DPO/etc): trains the model to produce responses rated higher by humans/models

**Inference** (uses the model):
1. Take current token sequence (prompt + everything generated so far)
2. Compute probability distribution over next token
3. Sample using a decoding strategy
4. Append token, repeat

There is no hidden scratchpad at inference. There is no internal goal unless given via context. Every generation is a fresh prediction — "searching a space of continuations under probabilistic constraints shaped by training."

Training changes the model; inference uses it. Individual prompts don't update model weights (unless the provider is collecting data for future training runs — a governance question, not a capability one).

## Tokens

The model's basic unit. Not words — chunks. Common chunks become single tokens; rare strings get split. Implications:
- **Cost**: APIs charge per input + output token
- **Context window**: tokens compete for a finite budget
- **Latency**: more tokens = more compute
- **Behavior**: token boundaries affect how the model "sees" identifiers, code symbols, unusual strings

Treat context windows like memory: constrained resources to budget deliberately, not to fill indiscriminately.

**Tokens aren't just text.** In multimodal models, images are chopped into pixel patches and audio into segments — each embedded as a vector in the same high-dimensional space as text tokens. A cat image and the word "cat" end up near each other in this shared space. The token prediction mechanism is identical; only the tokenization step differs by modality. See [[Multimodal AI]].

## Alignment ≠ Truthfulness

Instruction tuning and safety training make models *helpful and polite*. These are not the same as *truthful*:

- A model can be safe and wrong
- A model can be polite and wrong
- A model can be aligned and still hallucinate
- "Being helpful" rewards producing plausible answers, not calibrated uncertainty

The practical consequence: a model optimized to be helpful will often answer rather than say "I don't know." This is great for brainstorming; dangerous for factual precision unless grounding and verification are in the workflow.

**Grounding** (retrieval, tool use, citations) is the mechanism for truthfulness. Alignment is not.

## Fluency Is Cheap; Truth Is Not

LLMs produce text that sounds like the kind of text that would appear in their training data — including confident explanations, authoritative tone, step-by-step reasoning. This fluency is an optimization target, not a signal of correctness.

The **fluency heuristic**: humans equate ease of processing with truth. LLMs exploit this accidentally — they produce fluent, well-structured text because that's what training rewarded.

Counterintuitively, experienced engineers are *more* susceptible, not less. AI-generated code resembles their own style, triggering the same pattern-match shortcuts that let them read code fast.

Fluency is necessary but nowhere near sufficient. The three levels of trust:

| Level | What it means | Sufficient for? |
|-------|--------------|-----------------|
| **Plausible** | Sounds reasonable, uses right terminology, no obvious contradictions | Human brainstorming |
| **Evidence-backed** | Cites sources, references verifiable data | Starting point for decisions |
| **Verified** | Checked against independent sources or real systems | Production deployment |

Most AI output lives at plausible. Only verified carries operational trust.

## Practical Mitigations

Stack the deck at the system level:
- **Ground with retrieval** when facts matter — bind the model to external sources rather than training memory
- **Constrain with schemas** when structure matters — make claims checkable by making them typed
- **Run tests** when code matters — inference is not code execution, even if the model can describe execution traces
- **Review changes** when production matters — the model is not accountable; you are
- **Monitor and evaluate** when reliability matters — treat the model as a probabilistic component, not a deterministic function

## The Frozen-After-Training Property

LLMs are System A systems (in the Dupoux/LeCun/Malik [[Autonomous Learning Architecture]] framework) — they learn through passive observation of training data, then stop. This is not a missing feature but an architectural property of how training pipelines are designed. The practical consequence: domain mismatch is structural, not fixable by scaling alone. Real-world data is heavy-tailed and non-stationary; a frozen model cannot adapt to cases outside its training distribution.

See [[Autonomous Learning Architecture]] for the proposed path forward.

## See Also (Mechanics)

For the underlying mathematical machinery — neurons, activations, weights, biases, forward pass, gradient descent, backpropagation — see [[Neural Network Mechanics]]. The LLM Mental Model page covers *what to expect from* LLMs; the mechanics page covers *how the parameters are structured and learned*.

## See Also

- [[AI Scraping Resistance]] — exploits the frozen-after-training property offensively: poisoned training data degrades model behavior permanently, since there's no runtime correction
- [[Harness Engineering]] — system-level design that accounts for these properties
- [[Context Rot]] — what happens to model output when context degrades
- [[Vibe-Coding Anti-Pattern]] — what goes wrong when practitioners ignore these properties
- [[Prompting as Specification]] — how to work with the probabilistic nature deliberately
- [[AI Without Illusions (Series)]] — source course
- [[Probability and Statistics Foundations]] — the mathematical vocabulary underlying "probabilistic pattern machine": distributions, inference, the theoretical/empirical distinction
- [[Data Visualization Principles]] — visualizing model behavior and outputs requires the same encoding discipline as any data visualization
