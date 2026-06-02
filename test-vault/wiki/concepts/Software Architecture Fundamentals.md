---
title: "Software Architecture Fundamentals"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "archive/books/2026-06-01-fundamentals-of-software-architecture.md"
related:
  - "Agentic Workflow Patterns"
  - "Harness Engineering"
  - "Progressive Disclosure Architecture"
tags: [software-architecture, design, trade-offs, modularity]
---

# Software Architecture Fundamentals

Source: [[Fundamentals of Software Architecture]] (Richards & Ford, O'Reilly 2nd ed., 2024).

## The Two Laws

**First Law:** Everything in software architecture is a trade-off.
**Second Law:** Why is more important than how.

The first law frames the discipline: there are no best architectures, only contextually appropriate ones. Every decision optimizes some characteristics at the expense of others. The architect's job is to identify which characteristics matter most, then choose accordingly.

## Architectural Characteristics

Formerly called "non-functional requirements" — a misleading name since they directly shape system function. Architectural characteristics are the "-ilities": scalability, availability, reliability, testability, deployability, security, performance, observability.

Three categories:

| Category | Examples |
|----------|---------|
| **Operational** | Availability, scalability, reliability, performance, elasticity |
| **Structural** | Modularity, deployability, configurability, extensibility |
| **Cross-cutting** | Security, testability, observability, legality, accessibility |

**Key practice:** Never optimize for more than three to five characteristics simultaneously. Every style makes some characteristics easy and others hard — picking the right style starts with knowing which characteristics you're not willing to sacrifice.

### Fitness Functions

Automated, objective metrics that continuously verify the architecture maintains its intended characteristics. The architectural equivalent of unit tests.

Examples:
- "No class in package `com.myapp.service` may import from `com.myapp.data` directly" (enforced via ArchUnit or similar)
- "95th percentile API response time < 200ms" (enforced via monitoring gates in CI)
- "No new circular dependencies introduced" (enforced via dependency analysis)

Fitness functions shift architecture governance from "architect reviews PRs" to "automation catches violations at merge time." See also [[Harness Engineering]] for the analogous pattern in LLM systems.

## Modularity

Modularity is argued to be *more fundamental* than architecture style choice. A well-modularized monolith outperforms a poorly-modularized microservices system.

### Cohesion and Coupling

**Cohesion:** How related are the elements within a module? High cohesion = each module does one thing well.

**Coupling:** How much do modules depend on each other?
- **Afferent coupling** (Ca): how many other modules depend on this one (incoming)
- **Efferent coupling** (Ce): how many modules does this one depend on (outgoing)
- **Instability** = Ce / (Ca + Ce): 0 = maximally stable (nothing depends on anything it imports); 1 = maximally unstable

**Distance from the Main Sequence:** ideal modules are either maximally stable (abstract, widely depended on) or maximally unstable (concrete, depends on many things). Modules in the "zone of pain" (stable but concrete) resist change; modules in the "zone of uselessness" (abstract but nothing depends on them) are dead code.

### Connascence

Generalization of coupling that captures *how* two components are coupled, not just *whether* they are:

| Type | Meaning | Example |
|------|---------|---------|
| **Name** | Must agree on a name | Method name used by caller and callee |
| **Type** | Must agree on a type | Function parameter type |
| **Value** | Must agree on a specific value | Magic constants shared between modules |
| **Algorithm** | Must use the same algorithm | Encryption on both ends of a channel |
| **Position** | Must agree on order | Positional parameters |
| **Execution** | Must execute in specific order | Setup before use |
| **Timing** | Must execute at specific time | Race condition vulnerability |
| **Identity** | Must reference same instance | Shared mutable state |

**Rule:** Prefer static connascence (detectable at compile time) over dynamic (only visible at runtime). Prefer weaker forms (name, type) over stronger (algorithm, timing, identity). Minimize connascence across module boundaries; strong connascence within a module is acceptable.

## Architecture Styles

Styles are named, well-understood patterns with known trade-off profiles. Choosing a style is choosing a trade-off profile.

| Style | Strengths | Weaknesses | When |
|-------|-----------|------------|------|
| **Layered** | Simplicity, cost | Scalability, deployability | Small teams, low complexity, low budget |
| **Modular Monolith** | Modularity + simplicity, easy refactoring | Scalability ceiling | Teams graduating from layered, not yet needing distributed |
| **Pipeline** | Simplicity, composability | Event handling, scalability | ETL, data transformation, linear workflows |
| **Microkernel** | Extensibility, customizability | Scalability, development cost | Plugin-based tools, product lines |
| **Event-Driven** | Scalability, fault tolerance | Complexity, eventual consistency | High-volume async workflows, decoupled producers/consumers |
| **Microservices** | Scalability, team autonomy, deployability | Complexity, operational overhead, distributed system problems | Large teams, high scalability needs |

**The most important question before choosing a style:** How many teams will work on this, and how independently do they need to deploy?

## Architectural Quantum

The smallest independently deployable unit with high functional cohesion. In a monolith, the quantum is the entire application. In microservices, each service is its own quantum. The number of quanta in a system determines deployment coupling and team independence.

## See Also (Systems Thinking Connection)

Architectural trade-off analysis is a linear approximation of systems reasoning — it treats each characteristic as an independent variable. [[Systems Thinking]] provides the nonlinear frame: architectural characteristics interact through feedback loops (e.g., high deployability enables faster feedback, which enables better modularity decisions). Fitness functions are balancing loops: they detect drift and push toward a target state.

## See Also

- [[Agentic Workflow Patterns]] — pipeline and event-driven architecture styles map directly to agentic workflow patterns
- [[Harness Engineering]] — fitness functions parallel harness verification; both enforce architectural invariants automatically
- [[Progressive Disclosure Architecture]] — component-based thinking (Ch 8) grounds the three-tier skill loading model
- [[Fitness Functions Political Viability]], [[Connascence in Distributed Systems]], [[Conceptual Integrity Divergence Detection]] — open questions on software architecture
