---
title: "Parsing Patterns"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns.md"
  - "archive/books/2026-05-31-language-implementation-patterns-ch3-5.md"
  - "archive/books/2026-06-01-crafting-interpreters.md"
related:
  - "Language Application Pipeline"
  - "Language Implementation Patterns"
tags: [compilers, parsers, lexer, AST, language-tools]
---

# Parsing Patterns

The lexer/parser layer of a [[Language Application Pipeline]]. Covers tokenization, the parser power hierarchy, and the tree representations parsers build.

## Lexing (Tokenizing)

Before parsing, a **lexer** (tokenizer) converts the raw character stream into a stream of typed **tokens**. Whitespace and comments are typically discarded here.

```
Input:  "return x + 1;"
Tokens: RETURN, NAME("x"), PLUS, INT(1), SEMI
```

The lexer is implemented as a set of methods, one per token type, driven by a `nextToken()` function that uses the lookahead character to route to the right method.

## Parser Power Hierarchy

Parsers are ranked by how much lookahead they need to make decisions. Each level handles more languages at the cost of complexity or speed.

### Pattern 3: LL(1) Recursive-Descent Parser
- Looks at **1 token** of lookahead to decide which rule to apply
- Implemented as one function per grammar rule — the call tree implicitly traces the parse tree
- Fast and simple; handles most practical languages

```go
// Grammar: stat → returnstat | ifstat
func stat() {
    switch lookahead {
    case RETURN: returnstat()
    case IF:     ifstat()
    }
}
func returnstat() { match(RETURN); expr(); match(SEMI) }
```

### Pattern 4: LL(k) Recursive-Descent Parser
- Looks at **k tokens** of lookahead
- Needed when alternatives share a common prefix longer than 1 token
- Implemented with a circular lookahead buffer of size k

Example: distinguishing `name` from `name = value` requires 2-token lookahead (see the `=` after the name).

### Pattern 5: Backtracking Parser
- Tries alternatives speculatively in order, backs up on failure
- Can look at **arbitrary amounts of input** — handles ambiguous grammars LL(k) can't resolve
- Uses a **mark/release stack**: `mark()` saves the current token buffer position; `release()` rewinds to it
- Worst-case O(2^n) — exponential without memoization
- **The "launch missiles" problem**: parser actions (side effects) during a failed speculative path can't be undone. Solutions: disallow actions during speculation; or defer all actions to a separate tree-walking pass (preferred — also cleaner architecture)

### Pattern 6: Memoizing Parser (Packrat)
- Backtracking parser + a result cache keyed on `(rule, input-position)`
- If a rule has already been tried at a position, return the cached result instead of re-parsing
- Converts backtracking from exponential to **linear time O(n)**
- The classic Packrat parsing algorithm — guaranteed linear but uses more memory

### Pattern 7: Predicated Parser
- Adds arbitrary boolean **semantic predicates** after lookahead tests: `if (lookahead == NAME && isType(LT(1)))`
- The predicate can call the symbol table, check version flags, or inspect any runtime state
- Most powerful — handles context-sensitive constructs like `T(i)` (function call vs. cast depending on T's declaration)
- Requires the symbol table to be partially populated before parsing completes — the standard solution is a pre-pass to collect declarations

## Choosing a Parser

| If your language needs... | Use... |
|--------------------------|--------|
| Simple rules, no shared prefixes | LL(1) |
| Alternatives with shared k-token prefixes | LL(k) |
| Arbitrary lookahead, rare edge cases | Backtracking |
| Fast arbitrary lookahead | Memoizing (Packrat) |
| Context-dependent syntax | Predicated |

In practice, most languages (including subsets of C, Java, Python) can be handled with LL(1) or LL(2). Backtracking is needed for genuinely ambiguous grammars.

**ANTLR** automates building LL(k)/predicated parsers from grammars — the patterns are what ANTLR generates.

### Pratt Parsing (Top-Down Operator Precedence)

Recursive descent handles most constructs cleanly but becomes awkward for **expression precedence** — you need one function per precedence level, leading to deeply nested call chains. Pratt parsing solves this elegantly.

Each token type is assigned:
- A **prefix parse function** — what to do when the token appears at the start of an expression (e.g., `-` for unary negation)
- An **infix parse function** — what to do when the token appears between two expressions (e.g., `-` for subtraction)
- A **binding power** (precedence) — how tightly the token binds to adjacent expressions

The core loop:

```c
static ParseRule* getRule(TokenType type); // table lookup

static void parsePrecedence(Precedence minimum) {
    advance();
    ParseFn prefixFn = getRule(parser.previous.type)->prefix;
    prefixFn();  // parse the left operand

    while (minimum <= getRule(parser.current.type)->precedence) {
        advance();
        ParseFn infixFn = getRule(parser.previous.type)->infix;
        infixFn();  // parse operator + right operand
    }
}
```

The binding power comparison (`minimum <= current precedence`) is what handles precedence and associativity: pass `PREC_UNARY + 1` for left-associative operators, `PREC_UNARY` for right-associative.

**Advantage over recursive descent for expressions:** a single flat loop replaces a pyramid of `equality()` → `comparison()` → `term()` → `factor()` → `unary()` calls. Adding a new operator requires only a new row in the parse rule table, not a new function and a new call site.

Source: [[Crafting Interpreters]] Ch 17 (clox compiler).

## Tree Representations

### Parse Tree
Records every grammar rule applied during parsing. Includes "noise" — intermediate rule nodes that don't correspond to meaningful program structure.

Use when: you need exact source positions for every syntactic construct (e.g., syntax highlighting, precise error recovery).

### Homogeneous AST
All nodes are the same type; children stored in a generic list.

```go
type ASTNode struct {
    Token    Token
    Children []*ASTNode
}
```

Simple to build; requires type-checking children at traversal time.

### Heterogeneous AST (Irregular)
Different node types for different constructs; children stored as named fields.

```go
type AssignNode struct { Target Expr; Value Expr }
type ReturnNode struct { Value Expr }
```

More verbose to define but type-safe at compile time; natural fit for visitor patterns.

**Rule of thumb:** If you have many node types and traverse them extensively, heterogeneous ASTs pay off. For small tools or prototypes, homogeneous is faster to build.

## Tree Walking

See [[Tree Walking Patterns]] for the full treatment — embedded walkers vs external visitors, tree grammars vs pattern matchers, top-down vs bottom-up rewriting.

## See Also

- [[Language Application Pipeline]] — where lexing and parsing fit in the full pipeline
- [[Tree Walking Patterns]] — the next stage: traversing and rewriting the AST
- [[Language Implementation Patterns]] — Terence Parr's source; 31 patterns covering the full pipeline
