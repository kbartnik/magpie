---
title: "ADHD in Software Engineering"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/papers/2026-05-31-liebel-2023-software-engineers-adhd.pdf"
  - "archive/papers/2026-05-31-shah-2025-tether-adhd-se-assistant.pdf"
  - "archive/papers/2026-05-31-das-2021-neurodivergent-remote-work.pdf"
related:
  - "ADHD"
  - "Digital Biomarkers for ADHD"
  - "Digital Therapeutics for ADHD"
  - "Harness Engineering"
tags:
  - adhd
  - software-engineering
  - neurodiversity
  - workplace
---

# ADHD in Software Engineering

An estimated 5.0–7.1% of software engineers have ADHD. Software engineering is a domain that both amplifies certain ADHD challenges (ambiguous task scopes, long planning horizons, constant context-switching, always-on communication) and rewards certain ADHD strengths (creativity, nonlinear thinking, hyperfocus on engaging problems). The research is sparse but growing.

## Challenges (Liebel et al. 2023 — 19 SE practitioners, qualitative)

### Task, Deadline & Estimation
- Difficulty estimating how long tasks will take — a known ADHD executive function deficit (time blindness), compounded by software engineering's inherent estimation uncertainty
- Prioritizing across multiple tasks, especially when all feel simultaneously urgent or none feel urgent enough
- Starting tasks — particularly tasks that require sustained effort before any visible progress

### Attention Management
- Sustaining attention during meetings, code reviews, and documentation work — activities without immediate feedback loops
- Hyperfocus as double-edged: deep engagement on interesting problems can yield exceptional output, but at the cost of neglecting adjacent tasks and deadlines
- Context-switching costs are magnified: each interruption triggers a longer recovery time than for neurotypical colleagues

### Interpersonal and Communication
- Meeting dynamics: difficulty tracking multiple speakers, following agenda, filtering irrelevant information
- Communication style mismatches: direct/blunt communication creates friction; social nuance requires cognitive overhead
- Non-disclosure: most participants had not disclosed ADHD to employers, fearing stigma and professional consequences

### Physical and Mental Health
- Burnout risk elevated by chronic compensatory effort
- Co-occurring conditions (anxiety, depression) common; SE stress interacts with these

## Strengths

The same participants consistently reported distinctive strengths:
- **Creativity and lateral thinking** — nonlinear association generates novel approaches to design problems
- **Puzzle-mode engagement** — when a problem is sufficiently interesting and bounded, ADHD can produce exceptional focused output
- **Systems thinking** — ability to rapidly see architectural implications and long-range dependencies
- **Pattern recognition** — noticing anomalies and connections that more focused attention might miss

These strengths are recognized but rarely structurally leveraged by organizations.

## Strategies

Self-engineered compensatory strategies (no organizational support was described):
- Time-boxing and external deadlines (Pomodoro-style)
- Todo lists and external memory systems
- Working in quieter or controlled sensory environments
- Asynchronous communication preferred over synchronous meetings
- Pair programming as a focus anchor (social accountability)
- Choosing tasks by interest to leverage hyperfocus

## Remote Work (Das et al. 2021 — 36 neurodivergent professionals, qualitative)

Remote work during COVID surfaced the accessibility mismatch of office environments:

**What WFH improved:**
- Eliminated open-plan office noise and sensory overload
- Restored control over workspace customization
- Removed mandatory in-person social performance burden
- Enabled asynchronous communication norms by default

**What WFH worsened or introduced:**
- Social isolation — informal connection with colleagues, which can anchor attention, disappeared
- Work/rest boundary collapse — without physical separation, the "done" signal is absent
- Video call fatigue — on-camera expectations add social performance load
- Productivity guilt — without observable office presence, self-monitoring pressure increases

Most neurodivergent participants preferred WFH, but not because WFH is designed for neurodivergent workers — because offices are designed against them.

## Tether: An LLM-Powered Support Tool (Shah et al. 2025)

Tether addresses the tooling gap Liebel identified. Architecture:
- **OS-level activity monitoring** — tracks active window, idle state, app switching in real time
- **Context assembly** — behavioral signals assembled into context without requiring self-report from the user (addresses ADHD metacognitive deficit)
- **RAG pipeline** — LangChain indexes ADHD-specific resources + conversation history
- **Gentle notifications** — when idle too long, personalized prompts based on recent activity
- **Chatbot** — supports emotional regulation and task management via ADHD-aware conversational guidance
- **Gamification** — points and badges for focus goals; engagement mechanism

**Key design insight:** Context comes from *observed behavior*, not self-report. Users with ADHD often struggle to accurately report their own current state; passive monitoring bypasses this.

**Connection to this vault:** Tether's architecture is functionally identical to the vault's SessionStart hook + context injection pattern — passive monitoring → context assembly → LLM with behavioral context. Different domain, same harness engineering pattern.

**Limitation:** As of 2025, Tether has not been evaluated with actual ADHD users. Validation against the target population is pending.

## Get Me In The Groove (Newman et al. — ICSE 2025)

The most comprehensive quantitative study of ADHD professional programmers to date. Two-phase mixed methods: (1) qualitative analysis of 99 posts + 1,659 comments from r/ADHD_Programmers (first academic study of this 61,000-member forum), (2) survey of 493 professional programmers (239 ADHD / 254 non-ADHD).

**Key findings:**
- **10.6% of programmers have ADHD** (Stack Overflow Developer Survey 2022, N≈70,000) — higher than the 5–7% general population figure
- ADHD programmers are **1.8–4.4× more likely** to struggle with all measured challenges
- Hardest challenges by severity: **time management** and **design** (not just implementation)
- Creative strengths validated objectively: ADHD programmers score significantly higher on divergent thinking tests (medium-to-large effect size)
- **Only 1 of 9 neurodivergent programmers** who wanted accommodations actually requested them — disclosure risk is the primary barrier

**The strategic argument for organizational change:** challenges affect *non-ADHD programmers too*, just less severely. Organizational accommodations can be framed as productivity improvements rather than disability accommodations — lowering the disclosure requirement.

**Method note:** Replication package is public (GitHub). First SE paper to analyze r/ADHD_Programmers. Survey includes both ADHD and non-ADHD groups for direct comparison.

Source: [[archive/papers/2026-06-01-get-me-in-the-groove-adhd-programmers-notes|Newman et al. 2025]]

## Neurodiverse Tech Employees (Morris, Begel & Wiedermann — ASSETS 2015)

The first systematic study of neurodiverse *software developers specifically* (Microsoft Research). Methods: 10 qualitative interviews + 846-engineer survey (59 neurodiverse / 781 neurotypical).

Key findings that extend or qualify the Liebel 2023 qualitative data:
- Neurodiverse workers often have not disclosed to employers — the study sample is self-selected toward those comfortable enough to join an internal autism email list, likely underestimating true prevalence and severity
- Open-plan offices, ambiguous communication norms, and social performance expectations were the most consistently reported structural barriers
- Organizational changes are rarely offered; accommodation is individual and often self-engineered
- ASD and ADHD are conflated under "neurodiversity" — the study doesn't cleanly separate which findings belong to which diagnosis

**Methodological note:** The 2015 sample predates remote-first norms and AI coding tools. The landscape for neurodiverse SE workers in 2025 is meaningfully different.

Source: [[archive/papers/2026-06-01-neurodiverse-tech-employees-2015-notes|Morris et al. 2015]]

## Open Research Questions

- What organizational (not just individual) changes would reduce the ADHD accommodation burden in SE workplaces?
- How does the ADHD profile interact with SE role (architect vs. engineer vs. QA)?
- Does remote work maintain its accessibility benefits in stable hybrid arrangements, or do the benefits erode as "office norms" infiltrate remote settings?
- Can passive behavioral sensing (keyboard, mouse, idle detection) identify when an ADHD developer needs support — without requiring disclosure?

## See Also

- [[ADHD]] — clinical overview and domain map
- [[Digital Biomarkers for ADHD]] — the passive sensing layer; keyboard/mouse data relevant to workplace monitoring
- [[Digital Therapeutics for ADHD]] — intervention designs that inform workplace tool design (gamification, adaptive difficulty)
- [[Harness Engineering]] — Tether is a harness around an LLM for cognitive support; the pattern transfers
- [[Claude Code Hooks]] — SessionStart and activity-monitoring hooks share structural DNA with Tether's monitoring architecture
