---
title: "CLIP"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/clippings/2026-06-01-welch-labs-how-ai-images-videos-work.md"
related:
  - "[[Multimodal AI]]"
  - "[[Diffusion Models]]"
  - "[[Transformer Architecture]]"
tags:
  - llm-fundamentals
  - multimodal
  - computer-vision
---

# CLIP

**Contrastive Language-Image Pre-training** (OpenAI, 2021). A model that jointly trains a vision encoder and a text encoder so that matching image-text pairs end up near each other in a shared high-dimensional vector space.

## The Core Idea

Train two encoders — one for images, one for text — on 400 million image-caption pairs scraped from the web. The training objective is **contrastive**: for any batch of (image, text) pairs, maximize similarity between the matched pairs and minimize similarity between mismatched pairs.

```
[Image encoder]  → image vector  ↘
                                    shared embedding space → similarity score
[Text encoder]   → text vector   ↗
```

After training, a photo of a dog and the text "a dog" land near each other; a photo of a dog and the text "a cat" do not.

## Why It Matters

**Zero-shot image classification:** To classify an image, embed it, then embed candidate labels ("a photo of a dog", "a photo of a cat"), and return the highest-similarity label. No per-class training needed — novel categories work immediately if you can describe them.

**Conditioning diffusion models:** CLIP text embeddings are used to steer image generation in Stable Diffusion, DALL-E, and related models. The text prompt is encoded by CLIP; the diffusion model generates an image whose CLIP embedding matches the prompt's. This is what allows text-to-image: the CLIP space is the bridge.

## Architecture

Both encoders are Transformers:
- Image encoder: a **Vision Transformer (ViT)** — the image is split into patches, each patch embedded, then processed by a standard Transformer
- Text encoder: a GPT-style causal Transformer

The shared space dimensionality is typically 512 or 1,024.

## Limitations

CLIP embeddings compress the image into a fixed-size vector *before* the question is asked. If a relevant detail is small or peripheral, the encoder may have already discarded it. This is the "compression before question" problem described in [[Multimodal AI]] § Feature-Level Fusion — native multimodal models (which attend to raw image patches) address this.

## See Also

- [[Multimodal AI]] — CLIP as the prototypical feature-level fusion approach; how native multimodality improves on it
- [[Diffusion Models]] — CLIP conditioning in Stable Diffusion and DALL-E; how text prompts steer image generation
- [[Transformer Architecture]] — the Transformer backbone both CLIP encoders use
