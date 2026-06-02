---
title: "Data Visualization Principles"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/books/2026-06-01-designing-data-visualizations.md"
  - "archive/books/2026-06-01-making-data-visual.md"
related: []
tags: [data-visualization, design, communication]
---

# Data Visualization Principles

Frameworks for deciding what to visualize, how to encode data, and what to avoid. Source: [[Designing Data Visualizations]] (Iliinsky & Steele, O'Reilly 2011).

## The Designer-Reader-Data Trinity

Every visualization exists at the intersection of three forces:

- **Designer** — why are you making this? What do you want to communicate?
- **Reader** — who will see it? What context do they bring? What do they need to do with it?
- **Data** — what does the data actually support? What relationships are real vs. artifacts?

Visualization fails when any vertex is ignored:
- Ignoring the reader → designed for self, incomprehensible to others
- Ignoring the data → persuasive but misleading
- Ignoring the designer's goal → exploratory scatter with no point

## Exploration vs. Explanation

Two fundamentally different modes:

| Mode | Goal | Reader state | Design priority |
|------|------|-------------|-----------------|
| **Exploration** | Find unknown patterns | Analyst, patient, data-literate | Density, interaction, filtering |
| **Explanation** | Communicate a known finding | General audience, time-constrained | Clarity, focus, narrative |

Hybrids exist but require explicit design choices for both audiences — the defaults for one mode actively harm the other.

## The Encoding Power Hierarchy

Visual channels ranked by accuracy of human perception (position is most accurate):

1. **Position** (spatial placement on a shared axis) — most powerful; always use for the most important dimension
2. **Length** (bar height, line length)
3. **Angle/Slope**
4. **Area** (bubble size)
5. **Color hue** — categorical only; cannot encode quantitative magnitude reliably
6. **Color saturation/lightness** — limited quantitative range
7. **Shape** — categorical only; ~6 distinct shapes maximum
8. **Texture, motion** — weakest

**Rule:** Map the most important data dimension to position. Every step down the hierarchy is an accuracy penalty.

## Redundant Encoding

Mapping the same data dimension to multiple channels simultaneously (e.g., color + position, shape + size) improves comprehension and accessibility without adding noise — as long as both channels agree. Readers who can't perceive one channel (colorblindness, printing in grayscale) can still read the other.

## Structural Pitfalls

These chart structures are problematic regardless of aesthetics:

| Structure | Problem |
|-----------|---------|
| **3D charts** | Adds a dimension that encodes nothing; distorts perception of the encoded dimensions |
| **Pie charts** | Humans are poor at comparing angles; bar charts of the same data are always more accurate |
| **Gradients** | Suggest continuous variation where none exists; readers read meaning into the gradient |
| **Drop shadows** | Create false depth cues; make readers perceive spatial relationships that aren't there |
| **Excel defaults** | Optimized for output speed, not comprehension |

"Some structures are just inherently bad" — not a matter of taste but of perceptual accuracy.

## Choosing Structure

Before choosing a chart type:
1. Know the data type: categorical, ordinal, quantitative, temporal, spatial
2. Know the relationship to show: comparison, distribution, composition, relationship, trend
3. Verify the structure encodes the most important dimension with the most powerful channel available

**Comparisons need to compare:** data meant to be compared must share a baseline or axis. Floating bars, separated pie slices, and dual-axis charts make comparison physically harder.

## The Process: Questions to Tasks

Source: [[Making Data Visual]] (Fisher & Meyer, O'Reilly 2018).

The most common failure isn't choosing the wrong chart — it's not decomposing the question into tasks first.

**Step 1 — Refine the question into tasks.** A vague question ("how is our system doing?") cannot be visualized. Break it into concrete, answerable tasks:
- "Compare 95th-percentile latency across 5 services for the last 30 days"
- "Identify which services have the highest error rate on weekends vs. weekdays"

Each task specifies: *what* to show, *how* to compare, *what* time range.

**Step 2 — Data counseling.** Before choosing a chart, understand what data the requester *actually has* vs. what they *think they have*. Field names, granularity, missing values, and freshness are routinely misunderstood. Skipping data counseling produces beautiful charts of the wrong thing.

**Step 3 — Choose marks and channels.**

| Term | Definition |
|------|-----------|
| **Mark** | Geometric primitive used to represent a data item — point, line, bar, area |
| **Channel** | Visual property applied to a mark to encode a data field — position, color, size, shape, opacity |

Every visualization is a mapping: data fields → channels on marks. The encoding power hierarchy (position > length > color > shape) tells you which channels to use for which fields.

**Step 4 — Single view or coordinated views.** For a single question, use a single view. For questions that require comparing across subsets or drilling down, use coordinated views — multiple views linked so that selecting in one highlights in others. Coordinated views are powerful but add cognitive load; use only when the analytical question requires it.

## See Also

- No other vault pages connect directly; this page is the vault's data visualization hub.
