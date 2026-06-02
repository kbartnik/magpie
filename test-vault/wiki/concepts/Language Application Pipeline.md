---
title: "Language Application Pipeline"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns.md"
related:
  - "Language Implementation Patterns"
  - "Parsing Patterns"
  - "Bubbletea Elm Architecture"
tags: [compilers, parsers, language-tools, architecture, pipeline]
---

# Language Application Pipeline

The standard architecture for any tool that reads and processes a structured language: characters flow through a multi-stage pipeline, with an Intermediate Representation (IR) as the stable interface between stages.

## The Pipeline

```
Input text
    ↓
  Lexer (tokenizer)         — breaks character stream into tokens
    ↓
  Parser                    — checks syntax, builds IR
    ↓
  IR (usually an AST)       ← the architectural linchpin
    ↓
  Semantic Analyzer         — resolves symbols, type-checks, extracts meaning
    ↓
  Generator                 — emits output (code, report, transformed source)
    ↓
Output
```

The IR decouples the reader from the analyzer and generator. Without it, every stage would need to re-parse the input. With it, multiple passes are cheap and stages can communicate through annotations on the tree nodes.

## Four Application Categories

Every language application is one of:

| Category | Description | Examples |
|----------|-------------|---------|
| **Reader** | Builds a data structure from input | Config file parser, class file loader, cross-reference tool |
| **Generator** | Walks an IR and emits output | Code generator, serializer, HTML renderer |
| **Translator/Rewriter** | Reader + Generator — reads one language, emits another | Formatter, refactorer, language migrator, compiler, assembler |
| **Interpreter** | Reads, decodes, and executes instructions directly | Calculator, scripting language runtime, protocol server |

Most practical developer tools (linters, formatters, DSL processors) are Translators.

## The IR: Parse Tree vs AST

Two common IR shapes:

**Parse Tree** — records every rule applied during parsing. Complete but "noisy." Useful for syntax highlighting editors where you need exact source positions of all constructs.

**AST (Abstract Syntax Tree)** — strips parse-rule noise, keeping only meaningful tokens and operators. Easier to traverse and annotate. The standard choice for analyzers and translators.

Example: the AST for `this.x = y;`

```
    =
   / \
  .   y
 / \
this x
```

The Parse Tree for the same input would include nodes for the statement rule, assignment rule, LHS rule, etc.

## Semantic Analysis

Once the IR is built, the central question is: **for a given symbol reference `x`, what is it?**

Semantic analysis runs in three sequential sub-phases, each annotating the AST with information the next phase consumes:

**1. Definition phase** — walk the AST (descent), build the scope tree, define all symbols into their scopes. After this pass, every declaration node has a symbol object attached.

**2. Resolution phase** — walk the AST again (ascent or second pass), resolve every symbol reference to its definition. Handles forward references because all definitions already exist. After this pass, every reference node has a pointer to its symbol object.

**3. Type-checking phase** — walk the AST a third time, compute the type of every expression bottom-up, check assignment compatibility and argument types.

These three phases are cleanly separable because each only reads what the previous phase wrote onto the AST. See [[Symbol Tables and Scopes]] and [[Static Type Checking]] for implementation detail.

**Two coexisting scope chains in OO languages:**
- *Lexical scope chain* — block → function → global, for local variables and top-level names
- *Class inheritance chain* — class → superclass → ..., for member access via `obj.field`

`obj.field` resolution jumps directly to the class scope; it does not walk the lexical chain.

## Output: Generator and Interpreter

The final stage consumes the annotated AST and produces output or executes the program.

**Generator** — walks the AST and emits text output (source code, HTML, config, reports). Two approaches:
- *Syntax-directed*: output actions embedded in grammar rules; simple, single-pass, no look-ahead
- *Model-driven*: separate walker visits annotated AST nodes and emits; supports multi-pass, context-dependent output
- *StringTemplate*: model-driven with templates separating output structure from traversal logic

**Interpreter** — walks the AST and *executes* it instead of emitting it. Two approaches:
- *Tree-walking*: visit each AST node and execute directly; simple, 10–100× slower than bytecode
- *Bytecode VM*: compile AST to flat instruction stream first, then execute in a tight dispatch loop; portable and fast

See [[Code Generation and Translation]] and [[Interpreters and Bytecode VMs]] for implementation detail.

**The escape hatch:** for simple, regular input languages, the full pipeline is overkill. A streaming lexer with filter rules can translate wiki markup to HTML in one pass — no AST needed. The pipeline adds value when multiple passes, context-dependence, or semantic analysis are required.

## Why This Matters for Tool Builders

Building a linter, formatter, DSL, or code generator means building a language application. The pipeline is the standard playbook:

1. Write a lexer (or use a generator like ANTLR)
2. Write a parser that builds an AST
3. Annotate the AST (symbol resolution, type checking)
4. Walk the AST to emit output or execute

Skipping the IR (trying to do analysis during parsing) couples stages and makes multiple-pass analysis impossible.

## See Also

- [[Language Implementation Patterns]] — Terence Parr's 31-pattern catalog for these stages
- [[Parsing Patterns]] — the lexer/parser layer in detail
- [[Symbol Tables and Scopes]] — scope tree, definition/resolution phases in semantic analysis
- [[Static Type Checking]] — type-checking phase following symbol resolution
- [[Code Generation and Translation]] — model-driven translation, StringTemplate, syntax-directed output
- [[Interpreters and Bytecode VMs]] — tree-walking vs bytecode VM execution
- [[Bubbletea Elm Architecture]] — analogous pipeline: Init→Update (parser/analyzer) with Model as IR
