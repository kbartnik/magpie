---
tags: [concept, agents, llm, transformers]
cluster: agents
aliases: ["prompting", "few-shot prompting", "chain-of-thought", "system prompts", "prompt design"]
related: ["Agentic Workflow Patterns", "JSON Schema Discipline", "LLM Tool Calling", "Context Engineering", "Transformer Architecture"]
sources:
  - "[[archive/books/2026-06-04-prompt-engineering-for-llms]]"
---

# Prompt Engineering

Applied transformer psychology — every effective prompting technique works because of something specific about how next-token prediction was trained.

## Core Techniques

**Zero-shot:** Ask directly, no examples. Works when the task is well-represented in training data.

**Few-shot:** Provide 2-5 examples of (input, desired output) before the actual query. Activates the in-context learning pattern from training data.

**Chain-of-thought:** Instruct or demonstrate step-by-step reasoning before the answer. Works because training data contains reasoning traces — the model has learned that intermediate steps precede correct conclusions.

**System prompt:** Placed at the top of context; trained on specifically for instruction-following. Has greater "weight" than mid-conversation instructions — system prompts are load-bearing, not cosmetic.

## Why Techniques Work

Each technique is a lever on the training distribution. Few-shot works because the model has seen millions of Q&A examples. CoT works because training data contains reasoning traces. The prompt activates a distribution, not a hardcoded behavior.

## Structured Output

Prefilling the assistant's response with `{` biases the model toward JSON completion. Combined with a JSON Schema constraint, this is reliable structured output. See [[JSON Schema Discipline]].

## Connections

- [[Transformer Architecture]] — understanding why techniques work requires understanding next-token prediction and the training distribution
- [[Context Engineering]] — prompt engineering optimizes within a fixed context; context engineering shapes the context itself
- [[LLM Tool Calling]] — tool descriptions are prompts; quality of description determines selection accuracy
- [[Retrieval-Augmented Generation]] — RAG prepends retrieved context; prompt engineering determines how to instruct the model to use it
