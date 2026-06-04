---
title: "Prompt Engineering for LLMs"
type: books
captured-date: 2026-06-04
source-url: ""
author: "John Berryman, Albert Ziegler"
publisher: "O'Reilly"
year: 2024
isbn: "978-1-098-15303-8"
---

# Prompt Engineering for LLMs — Berryman & Ziegler

**Source:** `2026-06-04-prompt-engineering-for-llms.pdf`
**Coverage:** 12 chapters: text generation basics, zero-shot and few-shot prompting, chain-of-thought, system prompt design, structured output, tool use, context management, evaluation, fine-tuning vs. prompting, retrieval-augmented generation, multi-turn conversations, production patterns

## Deep Read

**Key Insight:** Prompt engineering is applied transformer psychology. Every effective technique (few-shot examples, chain-of-thought, structured output prefilling) works because of something specific about how next-token prediction was trained. Few-shot works because the model has seen millions of Q&A examples and a few examples in the prompt activate that pattern. Chain-of-thought works because training data contains reasoning traces — the model has learned that intermediate steps precede correct conclusions. The techniques are not magic; they are levers on the training distribution.

**What Surprised Me:** The book argues that system prompts are structurally different from user prompts in a way that goes beyond role assignment. Because the system prompt always appears at the top of the context and is trained on specifically for instruction-following, it has a different "weight" in the model's attention — instructions placed there are more reliably followed than the same instructions placed mid-conversation. This is empirically testable and has real engineering implications: system prompts are load-bearing, not cosmetic.

**Open Questions:**
- Chain-of-thought prompting improves performance on reasoning tasks — but the chain of thought itself is not verified. The model can generate a plausible-looking but incorrect reasoning trace that leads to the right answer, or a correct trace that leads to the wrong answer. Is there a reliable way to detect when the chain of thought is confabulated vs. genuinely explanatory?
- The RAG chapter advocates for semantic chunking over fixed-size chunking. But semantic chunk boundaries (by paragraph, section, or topic) depend on document structure. How does RAG performance degrade for poorly structured documents (raw PDFs, transcripts, code)?
- The book was written before extended context models (100K+ tokens). Does context management advice (compress, retrieve, summarize) still apply when the context window is large enough to hold entire codebases?

**Wikilink Candidates:**
- [[Prompt Engineering]] — zero-shot, few-shot, chain-of-thought, system prompt design; primary source; not yet a wiki page
- [[Retrieval-Augmented Generation]] — RAG architecture, chunking strategies, semantic search; not yet a wiki page

**Connections:**
- [[Transformer Architecture]] — every prompt engineering technique is a lever on the training distribution; understanding why techniques work requires understanding next-token prediction
- [[Agentic Workflow Patterns]] — prompt chaining, tool use, and multi-turn conversation patterns in this book map directly onto the agentic patterns vocabulary
- [[JSON Schema Discipline]] — structured output section covers the same JSON Schema contract approach; reinforces the diagnostic question from the agentic patterns source
- [[LLM Tool Calling]] — Ch.6 tool use is the prompt-engineering perspective on tool calling; pairs with the Go implementation perspective
