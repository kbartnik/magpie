---
title: "The Go Programming Language"
type: books
captured-date: 2026-06-04
source-url: ""
author: "Alan A. A. Donovan, Brian W. Kernighan"
publisher: "Addison-Wesley"
year: 2015
isbn: "978-0-13-419044-0"
---

# The Go Programming Language — Donovan & Kernighan

**Source:** `2026-06-04-go-programming-language.pdf`
**Coverage:** 13 chapters: tutorial, program structure, basic data types, composite types, functions, methods, interfaces, goroutines and channels, concurrency with shared variables, packages and the go tool, testing, reflection

## Deep Read

**Key Insight:** Go is designed for software engineering at scale — for programs written by large teams, maintained over years, with frequent changes. Every language decision (explicit errors, no inheritance, no exceptions, minimal implicit behavior) is an answer to the question: "what makes large codebases maintainable by teams that weren't there when the code was written?" The language is boring by design. Boring is the point.

**What Surprised Me:** Kernighan co-authored this, and the book's precision shows it. The chapter on functions (Ch.5) defines deferred functions in terms of a stack that runs when the surrounding function returns — but then shows that `defer` can be used to measure elapsed time by capturing a time value at function entry and reading it at exit. It's a tiny example that reveals something deep: defer pairs with the function's *return* not with any scope, which means it sees the function's final state. Using named return values, a deferred function can even *modify* what gets returned. Most Go programmers don't know this.

**Open Questions:**
- The book predates generics (Go 1.18). The section on interface-based containers (`sort.Interface`, `http.Handler`) describes patterns that generics now handle more cleanly. Is the interface approach now deprecated, or do both coexist as idiomatic Go?
- Ch.9 on shared-variable concurrency argues for preferring channels over mutexes wherever possible, but acknowledges mutexes are right for some patterns. Has community consensus shifted on where to draw that line since 2015?
- Reflection (Ch.12) is described as a last resort. Given that generics now handle many of the cases where reflection was previously necessary, is reflection use declining in new Go code?

**Wikilink Candidates:**
- [[Go Error Handling]] — Ch.5 functions section; sentinel errors, error wrapping, deferred cleanup; not yet a wiki page
- [[Go Testing Patterns]] — Ch.11 testing; table-driven tests, test helpers, test coverage; not yet a wiki page
- [[Go Reflection]] — Ch.12 reflection; reflect package, Type and Value, unsafe usage; not yet a wiki page

**Connections:**
- [[Go Interfaces]] — Ch.7 is the authoritative treatment of interfaces; implicit satisfaction, the empty interface, type assertions, type switches — deeper than Learning Go 2e on the formal semantics
- [[Go Channel Concurrency Patterns]] — Ch.8 goroutines and channels; the book predates context.Context, so the done-channel pattern is shown in its original form before it became context
- [[Go Modules and Packages]] — Ch.10 covers the pre-modules package system; useful for understanding legacy code and the motivation for the module system that replaced it
