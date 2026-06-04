---
type: concept
tags: [ml, llm, nlp]
---

# Tokenization

The process of breaking input text into **tokens** — the discrete units an LLM operates on. Tokens are not characters or words; they're the vocabulary items the model was trained with.

**Source:** [[archive/videos/2026-06-04-3b1b-transformers-tech-behind-llms|3Blue1Brown — Transformers Ch. 5]]

---

## What a Token Is

Tokens are words, word-pieces, punctuation, or other common character sequences — whatever unit minimizes vocabulary size while keeping sequences manageable. Common words are single tokens; rare words are split into sub-word pieces.

- GPT-3 vocabulary: **50,257 tokens**
- Common words: one token (`cat`, `the`, `running`)
- Rare words: multiple tokens (`antidisestablishmentarianism` → several pieces)
- Whitespace and punctuation: their own tokens

For images and audio, tokens are patches or chunks of the raw signal.

## Why It Matters

The context window limit is measured in **tokens**, not characters or words. GPT-3: 2,048 tokens. GPT-4: up to 128k. The tokenization choice affects:

- How much text fits in context
- How the model "sees" code vs. prose (code often tokenizes less efficiently)
- Surprising edge cases (single characters sometimes split across tokens, affecting counting tasks)

## Byte Pair Encoding (BPE)

The most common tokenization algorithm. Starts with individual bytes/characters, then iteratively merges the most frequent adjacent pairs into new vocabulary items. Produces a vocabulary that covers the training distribution efficiently.

---

## Related

- [[Transformer Architecture]] — tokenization is the first step in the data flow
- [[Softmax]] — the final layer converts to a probability distribution over the token vocabulary
