---
title: "Tokenization"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/videos/2026-06-01-3b1b-transformers-tech-behind-llms.md"
related:
  - "[[Transformer Architecture]]"
  - "[[LLM Mental Model]]"
tags:
  - llm-fundamentals
  - transformers
---

# Tokenization

The process of converting raw text into discrete tokens before feeding it to a language model. A token is not necessarily a word — it can be a word, a word-piece, a punctuation mark, or any common character combination that the tokenizer has learned.

## Why Not Characters or Words?

- **Characters only**: sequences become very long (10,000-char document → 10,000 steps); model must learn to compose characters into words at every step
- **Words only**: vocabulary explodes with morphological variants; rare words and typos require an "unknown" token
- **Subword tokenization**: the practical middle ground — common words are single tokens; rare words are split into pieces

## Byte Pair Encoding (BPE)

The dominant method (used in GPT models). Algorithm:
1. Start with a vocabulary of individual characters
2. Count all adjacent character pairs in the training corpus
3. Merge the most frequent pair into a new single token
4. Repeat until the vocabulary reaches the target size (GPT-3: 50,257 tokens)

Result: common English words become single tokens; rare or foreign words split into subword pieces.

## GPT-3 Numbers

| Property | Value |
|----------|-------|
| Vocabulary size | 50,257 tokens |
| Context window | 2,048 tokens |
| Embedding dimension | 12,288 per token |

## Implications

**Non-intuitive splits:** "tokenization" might be one token; "Tokenization" (capital T) might be two. Arithmetic fails partly because numbers like "573" are split unpredictably into subwords.

**Cross-lingual asymmetry:** English dominates training data, so English words tokenize efficiently (1 token ≈ 1 word). Non-Latin scripts may use 2–5 tokens per word, effectively shrinking the usable context window for those languages.

**Images and audio:** Tokenization generalizes to other modalities — images split into pixel patches, audio into time-domain chunks. Each patch/chunk becomes an embedding vector in the same space as text tokens. See [[Transformer Architecture]] § Ch. 5 and [[Multimodal AI]].

## See Also

- [[Transformer Architecture]] § Ch. 5 — where tokenization fits in the full data flow
- [[LLM Mental Model]] — why "the model processes tokens, not words" matters for understanding LLM behavior
- [[Multimodal AI]] — tokenization extended to images, audio, and video
- [[Softmax]] — the operation applied to the final token logits to produce a probability distribution
