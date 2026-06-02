---
type: "article"
author:
  - "TheLLMSeeker"
source: "https://hackernoon.com/were-solving-the-wrong-problem-for-llms-and-ai-overall"
published: 2026-05-22T00:00:00-04:00
clipped: "2026-06-01T08:45:05-04:00"
---
# We're Solving the Wrong Problem for LLMs and AI Overall

> [!info] **Metadata**
> 👤 **Author:** TheLLMSeeker  
> 📅 **Published:** 2026-05-22  
> ⏱️ **Reading Time:**   
> 🔗 **Original Link:** [Read on HackerNoon](https://hackernoon.com/were-solving-the-wrong-problem-for-llms-and-ai-overall)

---

## Summary


## Content

Part two of an accidental series. A few weeks ago, I wrote about LLMs having amnesia:

I wrote about LLMs and memory and the a16z continual learning essay, and I got the response I usually get when I write something like that and post on Linkedin: a bunch of people telling me I'm right for the wrong reasons, a few telling me I'm wrong for the right reasons, and one very long DM from someone I won't name who had clearly read every paper I cited and wanted to argue about test-time training for forty-five minutes. I'm not going to write about that email. What I am going to write about is a paper someone dropped in the comments that I almost didn't read because it was fifty pages long and published by a company called Onto AI, predating the a16z article. I assumed it was going to be the usual vendor document: carefully sourced just enough to seem credible, designed to make you think the author's product is the only coherent response to every problem in the field.

It's not that. I mean, it is a little bit that, it's still a whitepaper, and the company's own performance figures should be held loosely. But the architecture section stopped me cold in a way I didn't expect. Specifically, the diagram.

The four layers of their platform are called: the Sensorium, the Hippocampus, the Cortex, and Meta-Cognition. With the LLM sitting on top as a replaceable interface.

They named their memory layer after the part of the brain that builds memory. They named their reasoning layer after the part of the brain that reasons. And when I saw that, I thought, okay, either this is very good branding or someone here has genuinely been thinking about the problem differently than everyone else I've been reading. And then I spent a weekend going down a rabbit hole that this post is going to be the result of.

## The thing the whitepaper said that reframed everything

The whitepaper makes an argument I've heard before: retrieval isn't learning, compression is, static models can't accumulate institutional knowledge. However, it made it in a way that pushed me somewhere new. There's a line in the continual learning section that reads: "A system with infinite external storage — every fact perfectly indexed, instantly retrievable — has not learned. It has been spared the compression that makes learning powerful."

Fine. I made roughly the same point in my last piece. But the paper follows it with something I hadn't quite articulated: it says the reason human intelligence doesn't have this problem is structural. The hippocampus builds episodic memory. The cortex consolidates it. Synaptic weights strengthen with repeated co-activation and weaken with contradiction. The architecture itself is designed for continuous learning, not as a feature that was bolted on, but as the fundamental mechanism from which everything else emerges.

And that's when I started asking a different question than the one I'd been asking in my last piece. In part one, I was asking: how do you add continuous learning to an LLM? This whitepaper made me ask: why doesn't the brain have this problem? And once you ask it that way, the answer is not "because the brain has a clever finetuning approach." It's because the brain was built to learn continuously from day one. It's not a transformer with Hebbian updates bolted on. It's a Hebbian system, period, that also does perception and language as downstream capabilities.

*We built the fast brain first and now we're trying to retrofit the rest of it. But you can't retrofit a hippocampus. You have to design for one.*

## So I started looking at who's actually trying to design for one

The Onto AI architecture is a software interpretation of this idea. Their Hippocampus layer does persistent relational memory with Hebbian self-learning: edge weights in the knowledge graph strengthen when relationships are repeatedly confirmed, weaken when contradicted. It's a direct implementation of the biological mechanism in graph form. Their Cortex layer does deliberate multi-step reasoning over consolidated knowledge, which is roughly what the prefrontal cortex does when you're working through a hard problem rather than pattern-matching. The Meta-Cognition layer monitors reasoning quality in real time. Again, not a metaphor for something vaguely self-aware, but a concrete observational system that watches for low-confidence synthesis, sparse evidence, and frequent contradictions.

What struck me is that this architecture doesn't require solving catastrophic forgetting, doesn't require opening the weight boundary, and doesn't need to touch a single parameter in the underlying LLM. The learning happens in the graph. And because the graph is transparent (every edge has a provenance trail, every weight change is traceable to the evidence that caused it) you get continuous learning that remains auditable. Which is the property that makes parametric continual learning so hard to deploy safely, and which the graph approach gets for free.

But here's where it got more interesting. Because once I was thinking about the brain as a design document rather than a metaphor, I started noticing that there are companies taking this idea a lot more literally than Onto AI does.

## Some people are copying the brain with extreme prejudice

There's a whole conversation happening in AI hardware and neurotechnology that barely intersects with the LLM discourse, and I'd been ignoring it. I shouldn't have been.

**BIOLOGICAL HARDWARE**: Cortical Labs by Hon Weng Chong put 800,000 live human neurons on a silicon chip. They rewire themselves in real time when given new tasks, no programming, no reward engineering. The CL1 is a commercial product. It has been trained to play Pong. It has been tested on DOOM. It learns the way neurons learn, because it is a bunch of neurons.

**BIOLOGICAL HARDWARE**: FinalSpark by Dr. Fred Jordan created sixteen brain organoids, grown from human stem cells, connected to electronics. Remote access on subscription. Energy consumption is a fraction of conventional silicon. If you want to run compute on living neural tissue, you can rent it by the hour. This is a real product.

**BIO-INSPIRED SOFTWARE**: Eon Systems by Michael Andregg built a complete digital replica of a fruit fly nervous system (all 140,000 neurons). The system exhibits complex survival behaviours without training data, reward signals, or gradient descent. Intelligence as an emergent property of architectural fidelity, not of learned weights.

**SUBSTRATE-LEVEL BET**: Unconventional, Inc. by Naveen Rao, who previously sold Nervana to Intel for $400M and MosaicML to Databricks for $1.3B, is now building brain-scale efficiency in silicon without biological tissue. Just closed $1B at a $5B pre-product valuation. a16z is leading. Lightspeed and Lux Capital are in.

Although Onto AI, Cortical Labs, FinalSpark, Eon Systems are reportedly worth hundreds of millions or billions from their fundraising numbers, that Unconventional valuation stopped me for a while. Five billion dollars before a single product ships. That's not a number you get by having an interesting idea, it’s a number you get when a room full of people who are very good at evaluating risk decide a thesis is probably right. The thesis, as Naveen Rao has described it, is that the transformer architecture has a ceiling, that biology has already built above it, and that a team with the right background can extract the biological principles and put them in silicon you can manufacture.

Here's the thing that made me laugh slightly when I noticed it: a16z published an essay arguing LLMs need better memory and continuous learning. Then they led the round for a company whose whole argument is that you can't get there from the transformer. I don't think that's hypocrisy, I think it means even the people with the most information are genuinely uncertain which layer the breakthrough comes from, so they're funding both answers simultaneously. Rational. Also slightly unsettling.

The one that impressed me most:

A company called Koniku built a 64-neuron chip for chemical sensing. The chip detects trace molecules, including explosives with sensitivity electronic sensors can't match, using the same biological machinery evolution built for exactly this task. Their chips are going into drones and robotics. It's not a general intelligence play. It's biology deployed in a specific slot where biology is genuinely better than silicon, right now, in the field. I find this more immediately credible than the general-purpose biocomputer pitch, because it's not asking neurons to do everything, just the thing neurons are actually better at.

## The fruit fly is the part I keep coming back to

I want to spend a minute on Eon Systems because I think it's philosophically different from the others, and it's the thing the whitepaper most directly connects to even if it never mentions it.

The fruit fly, Drosophila, has 140,000 neurons. Neuroscientists have spent decades mapping every single connection. It is the most completely charted neural circuit in biology. Eon Systems is betting that this map is a blueprint: that if you faithfully reproduce the architecture in software, the behaviours emerge from the structure rather than from training. Navigational responses, survival instincts, sensorimotor coordination… all of it, without a single gradient descent step.

That's the same bet the whitepaper is making at the software layer. Not "let's train a model to behave like a brain" but "let's build a system that has the structural properties of a brain, and see if the right behaviours emerge from those properties." The knowledge graph with Hebbian weights isn't trained to accumulate institutional knowledge. It accumulates institutional knowledge because that's what a system with those structural properties does. The intelligence is in the wiring, not in the weights.

A fruit fly isn't GPT-4. But the implication is real: if architectural fidelity to biology produces emergent intelligent behaviour at 140,000 neurons, the question of what happens at higher scales of complexity is genuinely open. That's the $5 billion question Unconventional is trying to answer.

## The energy number that should be in every LLM discussion but isn't

The human brain runs on about 20 watts. A single GPU running inference on a frontier model consumes hundreds of watts. A full training run for a frontier model consumes enough electricity to power a small city for weeks. And those training costs grew at roughly 2.5x per year between 2016 and 2024. The inference side is worse: OpenAI's inference spend in 2024 reportedly hit around $2.3 billion, about fifteen times what GPT-4 cost to train.

**20W**: Human brain power consumption. Continuously learning, auditable reasoning, generalises from sparse examples

~ **$2.3B**: OpenAI estimated 2024 inference spend. Roughly 15× what it cost to train GPT-4

**2.5×**: Annual growth in frontier model training costs, 2016–2024. With no obvious plateau

These curves don't bend on their own. And biological systems suggest they don't have to look like this. Not because neurons are magic (they're slow by silicon standards, and they eventually die) but because the architecture doesn't require attending to every token every time you generate one. It doesn't require a KV cache that scales linearly with context length. It compresses experience into structure, which is exactly what the Hippocampus layer is trying to approximate in software. Less energy per inference, not because the chip got cheaper, but because the system doesn't ask the inference layer to do the job of long-term memory.

## Where this leaves the design argument

Reading the whitepaper again after going down this rabbit hole, the architecture makes more sense to me than it did the first time. The four-layer design isn't a branding choice. It's a claim: that the right way to build an intelligence system is to separate the perceptual ingestion layer, the persistent memory and learning layer, the deliberate reasoning layer, and the metacognitive monitoring layer and give each of them the mechanism that's actually appropriate for the job they do. Not ask one transformer to do all four.

The LLM sits on top as a replaceable language interface. Which means when a better LLM releases, you swap an API endpoint. You don't rebuild the memory layer. You don't retrain the reasoning engine. The institutional knowledge in the graph, such as the accumulated evidence, the confirmed relationships, and the provenance trails, persists unchanged.

```
The connection I missed in part one
```

The Onto AI architecture is a deployable, software-layer implementation of the same insight the biocomputing companies are pursuing in hardware: intelligence should have the structural properties of a brain, not the properties of a very large pattern-matcher with memory bolted on. The difference is in the timeline. The knowledge graph approach works today. The biological hardware approach works in some specific domains today and, if the scalability question resolves, could reshape the whole stack in ten years.

## The honest counterargument

I've been building toward a fairly enthusiastic take, so let me say the thing that should also be said.

Why I think the biological direction is right

- Continuous adaptation without retraining. Neurons literally rewire, and the graph learns from every query
- Energy efficiency gap is real and widening. The transformer cost curves have no obvious ceiling
- Architectural fidelity produces emergent behaviour. The fruit fly result is not nothing
- Auditability built in at the software layer. Graph traversal paths are inspectable in a way model weights aren't
- The $5B pre-product bet tells you something about where sophisticated money thinks the ceiling is

What I'm genuinely uncertain about

- 800,000 neurons is not a frontier model. Nobody has demonstrated the scalability yet
- Living tissue is unreliable. Neurons die. Organoids degrade. Manufacturing consistency is unsolved
- Ontology maintenance on knowledge graphs is expensive and ongoing, not a solved problem at scale
- The fruit fly can navigate, but cannot reason abstractly. The jump to general intelligence is not guaranteed
- The whitepaper's own performance figures are internal demos, not independently validated studies
- Timeline to hardware viability could be 3 years or 15, genuinely unclear

The scalability gap is the thing I can't get past entirely. Everything exciting in the biological space has been demonstrated at small scales. Cortical Labs has 800,000 neurons, which is genuinely impressive and also about 0.008% of the human brain. Nobody's shown that the principles that make a small biological system interesting survive being scaled up by several orders of magnitude. That could be a funding problem. It could be a wall. We don't know.

What I'd say is: the uncertainty doesn't mean the bet is wrong. It means it's speculative. A $5 billion pre-product valuation means the people with the most at stake ran the numbers and decided the speculation was worth it. That's not proof. It's the most expensive kind of informed opinion.

Where I'm landing after all this

More confident about the software layer than the hardware layer. The knowledge graph approach is deployed in production environments that don't tolerate failure. The biocomputing approaches are commercially available but not yet at the scale where you could run the tasks we currently ask of frontier models. I'm treating the former as "working now, needs broader validation" and the latter as "directionally correct, timeline genuinely uncertain." These are different things, and I want to be clear about the distinction.

## The thing the whitepaper made me realise I'd been asking wrong

In part one, my question was: how do you add continuous learning to an LLM? It's a reasonable question. It's also maybe the wrong one.

The whitepaper reframed it for me, without quite saying so explicitly. The brain doesn't have the LLM memory problem because it wasn't designed around next-token prediction. It was designed, over billions of years of selection pressure, around the problem of surviving in a world that keeps changing. Memory, continuous learning, generalisation from sparse examples, energy efficiency: these aren't features. They're the thing itself. Every structural property of the biological brain exists because some ancestor that had more of it survived longer than one that had less.

We built a very powerful next-token predictor and then asked it to also be a hippocampus. And we've spent the last three years being surprised that it isn't one.

The companies that are interesting to me right now are the ones that stopped being surprised and started building for the properties they actually want. Onto AI did it at the software layer, by separating memory from reasoning from ingestion and giving each layer the mechanism appropriate to its job. Cortical Labs did it by skipping the silicon entirely for the learning layer and using the biological substrate that actually has those properties. Eon Systems did it by treating the most-mapped neural circuit in biology as a blueprint and asking what happens when you run it.

These are all different approaches, but they're definitely asking a better question than "how do we make the transformer remember things."

How I think the layers fit together now

**Now - software brain:** Knowledge graph as institutional memory, Hebbian self-learning, auditable reasoning layer, LLM as replaceable language interface. Deployable today, addresses the compliance and accountability requirements that regulated industries actually face.

**Near-term - parametric learning:** Sparse memory finetuning and test-time training reach deployment maturity. Models compress inference-time experience into parameters without catastrophic forgetting. The tacit knowledge problem (patterns too high-dimensional for graph edges) starts to close.

**Medium-term - biological edge:** Neuromorphic and biocomputing approaches find their slot: specialist sensing (Koniku-style), continuous learning at the edge, robotic systems that don't need a data centre to adapt. The physical grounding problem gets partial answers.

**Long-term - substrate question resolves:** Either Unconventional or someone after them demonstrates brain-scale efficiency in manufacturable silicon. If that happens, every layer above gets more powerful. If it doesn't, the hybrid stack is what we have and it's probably enough for most of what matters.

I didn't plan to write a second part to the piece I published last week. But the whitepaper sent me somewhere I needed to write about, and this is where I ended up. Part three probably isn't happening. Unless something else lands in my inbox that makes me rethink the stack again.

If you're working on any of the biological hardware approaches and think I've got the scalability picture wrong (which I might ) drop it in the comments. That's the part of this I'm least certain about and most interested in getting better information on.

Part 1: "Your LLM Has Amnesia. And We Built the System That Keeps It That Way.", HackerNoon, May 2026. Companies mentioned: Cortical Labs (CL1 platform, Melbourne), FinalSpark (Neuroplatform, Vevey), Eon Systems (digital fruit fly nervous system), Koniku (64-neuron chemical sensing), Onto AI (Hybrid multi-layered software), Unconventional Inc. (Analog Computing). Source material: Onto AI "The Intelligence Gap" whitepaper (April 2026) via Shanka Jayasinha on Linkedin; a16z "From Memento to Memory" (2026). Benchmark references: Diffbot KG-LM, FalkorDB GraphRAG SDK Q1 2025, Han et al. arXiv:2502.11371. Cost data: Cottier et al. arXiv:2405.21015, OpenAI inference estimates.

---

## Proof of Usefulness

[Proof of Usefulness Hackathon](https://proofofusefulness.com/?ref=hackernoon.com) is a global 6-month developer challenge designed to reward real-world utility projects and initiatives. With 150,000+ in [cash prizes](https://proofofusefulness.com/cash-prizes?ref=hackernoon.com) and [software credits](https://proofofusefulness.com/software-prizes?ref=hackernoon.com) for winners and $1500+ worth of software and inventory for participants, this is undisputedly the biggest contest of the year. Learn more [here](https://proofofusefulness.com/?ref=hackernoon.com).

[![[resources/media/24b1a7f16bfafe41fb0080f98bf8facd_MD5.jpg]]](https://proofofusefulness.com/)
## Deep Read

**Key Insight:** "We built a very powerful next-token predictor and then asked it to also be a hippocampus." The LLM memory problem isn't a missing feature — it's architectural mismatch. The brain is a Hebbian learning system that also does language; we built a language system and are trying to bolt on memory. You can't retrofit a hippocampus.

**What Surprised Me:** The Eon Systems result — a complete software reproduction of the Drosophila fruit fly nervous system (all 140,000 neurons, fully charted) exhibits complex survival behaviors with zero training, zero gradient descent, zero reward engineering. Intelligence as an emergent property of architectural fidelity to biology. This is a different kind of claim than "bio-inspired" — it's structural copying, not metaphorical borrowing.

**Open Questions:**
- The scalability gap is unresolved: Cortical Labs has 800,000 neurons (~0.008% of the human brain). Does the Hebbian/emergent-behavior principle hold as you scale by several orders of magnitude?
- Is the Onto AI 4-layer software stack actually deployed in production regulatory environments, or is "deployed" marketing language? The provenance-trail auditability claim is significant if true.
- a16z funded both the "LLMs need better memory" essay and Unconventional Inc. (thesis: transformers have a ceiling). Is this rational hedging or genuine uncertainty among the best-informed actors?

**Wikilink Candidates:**
- [[Agent Memory Architectures]] — updating with the 4-layer software brain approach
- [[Neuromorphic and Bio-Inspired AI]] — creating this page for the hardware companies
- [[Continuous Learning]] — the property these architectures are trying to achieve; does not yet exist

**Connections:**
- [[Agent Memory Architectures]] — RAG vs. LLM Wiki vs. Fat Skills: this article adds a fourth approach (Hebbian knowledge graph as hippocampus-layer)
- [[LLM Mental Model]] — the transformer architecture and its fundamental limits
- [[Context Rot]] — a symptom of the problem this article diagnoses
- [[Bio-Inspired Computing]] — the hardware companies (Cortical Labs, FinalSpark, Koniku) are doing bio-inspired computing at the substrate level
