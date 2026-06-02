---
title: "Multimodal AI"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/clippings/2026-06-01-what-is-multimodal-ai.md"
related:
  - "LLM Mental Model"
  - "Retrieval-Augmented Generation"
tags:
  - llm-fundamentals
  - multimodal
---

# Multimodal AI

A model that can ingest and/or generate multiple data modalities — text, images, audio, video, lidar, thermal imaging, and others. The label describes the model's interface, not a separate category of AI; the underlying mechanism is still token prediction in a high-dimensional vector space.

## Two Architectures

### Feature-Level Fusion (Modular)

The older approach: separate specialist models bolted together.

```
Image → [Vision Encoder (CLIP-based)] → feature vector
                                              ↓
                     Text prompt → [LLM] → response
```

The vision encoder extracts a numerical summary of the image and projects it into a form the LLM can process alongside text. The LLM never sees the raw image — only the encoder's compressed representation.

**The core problem:** compression happens before the model knows the question. If the relevant detail is in a small icon in the corner, the encoder may have already discarded it by the time the question arrives.

Still used for specialized enterprise tasks: cheaper inference, modular components that can be swapped independently, easier to fine-tune one part without retraining the whole system.

### Native Multimodality (Shared Vector Space)

The current standard: all modalities tokenized and embedded into a single shared high-dimensional space.

| Modality | Tokenization |
|----------|-------------|
| Text | Words/subwords → embedding vectors (standard) |
| Images | Chopped into pixel patches (e.g. 16×16) → each patch embedded |
| Audio | Chunked into segments → each segment embedded |
| Video | Spatio-temporal patches (3D cubes: area × time window) → each cube embedded |

**The shared part is load-bearing.** A cat image and the word "cat" end up near each other in the same space because they represent the same concept. No translation between systems; the model reasons about all modalities simultaneously.

With feature-level fusion, the model attends to a summary of the image. With a shared vector space, it attends to image and text *at the same time* — so it can focus on the relevant detail because it knows the question while looking.

## Temporal Reasoning in Video

Early multimodal systems handled video by sampling frames and running them through a vision encoder. This loses temporal information — a single frame can't distinguish "picking up" from "putting down."

Native multimodal models use **spatio-temporal patches**: 3D cubes that encode an area *across a short time window* (e.g. 8 frames). Motion is a property of the token itself, not reconstructed by comparing frames. This makes temporal reasoning structurally sound rather than inferred.

## Any-to-Any Generation

Because all modalities share one vector space, a native multimodal model can output any combination of modalities, not just text. Input can be any mix; output can be any mix. Text in → video out. Image in → audio description out. The coherence comes from the shared representation.

## Image and Video Generation

Text-to-image and text-to-video models use a separate generative architecture: **diffusion models**. The mechanism is not token prediction but iterative denoising — start from pure noise and repeatedly apply a trained network (often a Transformer) to remove noise, steered by CLIP text embeddings and classifier-free guidance. See [[Diffusion Models]] for the full treatment.

## See Also

- [[LLM Mental Model]] — the token/embedding mechanism that shared vector spaces extend to non-text modalities
- [[Retrieval-Augmented Generation]] — multimodal RAG uses the shared vector space for cross-modal similarity search (retrieve images by text query, or vice versa)
- [[Tool Use in AI Systems]] — multimodal input is a tool interface concern: vision, file attachment, and audio inputs enter the model as tokenized context, not as tool calls
- [[Diffusion Models]] — the generative mechanism behind image/video generation; CLIP conditioning; classifier-free guidance
- [[CLIP]] — the OpenAI model that established the shared vision-text embedding space; dedicated page
