---
title: "Connascence in Distributed Systems"
question: "Does Page-Jones's connascence taxonomy extend cleanly to distributed systems, or do distributed systems require a different coupling framework?"
type: question
status: open
domain: architecture
created: 2026-06-02
updated: 2026-06-02
sources:
  - "archive/books/2026-06-01-fundamentals-of-software-architecture.md"
related:
  - "[[Software Architecture Fundamentals]]"
tags:
  - software-architecture
---

# Connascence in Distributed Systems

*Does Page-Jones's connascence taxonomy extend cleanly to distributed systems, or do distributed systems require a different coupling framework?*

Connascence measures coupling: how much a change in one component requires changes in another. The taxonomy (name, type, meaning, position, algorithm, execution order, timing, identity, value) was developed for monolithic object-oriented systems where components communicate through function calls.

In distributed systems, connascence of timing (one component's behavior depends on when another acts) and connascence of execution order become structural properties of the network protocol, not local implementation details. They're much harder to instrument, test, or enforce through fitness functions. A distributed system might have strong connascence of timing due to a race condition in an event stream — and this may only manifest under specific load patterns.

Whether connascence as a framework is useful for distributed systems, needs extension, or should be replaced with a different coupling vocabulary (event coupling, temporal coupling as distinct categories) is unresolved in the architecture literature.

## See Also

- [[Software Architecture Fundamentals]] — connascence taxonomy; modularity primacy principle
