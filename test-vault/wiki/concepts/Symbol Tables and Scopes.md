---
title: "Symbol Tables and Scopes"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns-ch6-8.md"
  - "archive/books/2026-06-01-crafting-interpreters.md"
related:
  - "Language Application Pipeline"
  - "Tree Walking Patterns"
  - "Parsing Patterns"
  - "Static Type Checking"
tags: [compilers, symbol-table, scopes, semantic-analysis, language-tools]
---

# Symbol Tables and Scopes

The infrastructure underlying semantic analysis: tracking what symbols (variables, functions, types) exist, where they're defined, and how to look them up during analysis.

## What is a Symbol?

A symbol is a name for a program entity. Every named thing in a program — variable, function, parameter, type, class — is a symbol. A symbol object stores:

- Its name
- Its type (for typed languages)
- Its kind (variable, function, class, etc.)
- A link back to the AST node that defined it

```java
class VariableSymbol extends Symbol {
    Type type;
    // name inherited from Symbol
}
class MethodSymbol extends Symbol implements Scope {
    Type returnType;
    Map<String, Symbol> parameters;
}
```

## Scopes and Scope Trees

A **scope** is a region of the program where a set of symbols is visible. Scopes nest: a block scope is inside a function scope, which is inside the global scope.

The **scope tree** mirrors the program's nesting structure. Each scope node holds:
- A map of `name → Symbol` for symbols defined in this scope
- A parent pointer to the enclosing scope

```
GlobalScope
├── int, float, bool  (built-in types)
├── f: MethodSymbol
│   └── LocalScope (function body)
│       ├── x: VariableSymbol
│       └── y: VariableSymbol
└── g: MethodSymbol
    └── LocalScope
        └── BlockScope (inner block)
            └── z: VariableSymbol
```

**Symbol lookup** walks up the parent chain until a match is found or the root is reached. This implements lexical scoping automatically.

## Flat vs Nested Scopes

**Flat scopes** (e.g., simple scripting languages): all symbols defined at one global level. Lookup is a single map lookup. Simple but no block-local variables, no shadowing.

**Nested scopes** (most real languages): each block, function, and class creates a new scope. Symbols in inner scopes can shadow outer ones. Lookup walks up the parent chain.

**When a scope is entered** (e.g., `{`): push a new scope onto the current scope chain.  
**When a scope is exited** (e.g., `}`): pop back to the parent scope.

## The Definition/Resolution Two-Pass Pattern

Separating "collect all definitions" from "resolve all references" solves the forward-reference problem:

```
Pass 1 — Definition phase (tree descent):
  upon VAR_DECL node   → create VariableSymbol, add to current scope
  upon METHOD_DECL node → create MethodSymbol, push new scope
  upon block {         → push new LocalScope
  upon block }         → pop scope

Pass 2 — Resolution phase (tree ascent or second walk):
  upon variable reference x → look up x in current scope chain
  upon type reference T     → look up T in current scope chain
  upon call f()             → look up f, check arg count/types
```

Because pass 1 runs first across the entire tree, pass 2 can resolve any forward reference — a function called before its definition, a type used before it's declared.

**Implementation:** Use `downup()` from [[Tree Walking Patterns]] — fire definition actions on descent, resolution actions on ascent. Or run two explicit tree walks in sequence.

## Data Aggregates: Class Scopes

Struct and class members live in a **class scope** attached to the type's symbol — not on the lexical scope chain. This is a completely separate lookup mechanism.

```
obj.field   →  1. look up obj in lexical scope chain → get VariableSymbol(type=ClassT)
               2. look up field in ClassT's scope → get FieldSymbol
               NOT: walk block → function → global
```

Class inheritance adds a superclass chain: if `field` isn't found in `ClassT`'s scope, walk up to `ClassT`'s superclass scope, then its superclass, etc.

```
ClassScope(Dog)
  ├── bark: MethodSymbol
  └── superclass → ClassScope(Animal)
                    ├── eat: MethodSymbol
                    └── name: FieldSymbol
```

This means two scope chains coexist in an OO language:
1. **Lexical scope chain** — for local variables, global names
2. **Class inheritance chain** — for member access

## Building the Scope Tree from the AST

The scope tree is built during the definition phase tree walk. Grammar rules produce imaginary token roots (`METHOD_DECL`, `VARDECL`) that trigger scope-tree actions:

| AST node | Definition phase action |
|----------|------------------------|
| `METHOD_DECL` | Create `MethodSymbol`, push new `LocalScope` |
| `VARDECL` | Create `VariableSymbol`, add to current scope |
| `CLASS_DECL` | Create `ClassSymbol` (itself a `Scope`), push class scope |
| Block `{` | Push new `LocalScope` |
| Block `}` | Pop scope |

After the definition phase, every AST node that defines a symbol has a `symbol` field pointing to the corresponding symbol object.

## The Resolver Pass (Pre-Run Semantic Analysis)

In a tree-walk interpreter, variable lookup happens at runtime by walking up an environment chain. This causes a subtle bug: a closure can capture a variable from a scope that has since exited — the resolution depends on the dynamic call stack, not the lexical structure.

The solution is a **resolver pass** — a fast tree walk that runs *after* parsing but *before* interpretation:

```java
void resolveLocal(Expr expr, Token name) {
    for (int i = scopes.size() - 1; i >= 0; i--) {
        if (scopes.get(i).containsKey(name.lexeme)) {
            interpreter.resolve(expr, scopes.size() - 1 - i);
            return;
        }
    }
    // assume global
}
```

Each resolved variable is annotated with its **static distance** — how many environment hops to the defining scope. The interpreter then uses this depth number directly instead of walking the chain by name at runtime.

This is the canonical example of why semantic analysis is a separate pipeline stage: it solves a *correctness* problem (not just a performance problem) that would be very messy to fix inside the parser or the interpreter.

Source: [[Crafting Interpreters]] Ch 11 (jlox resolver).

## See Also

- [[Language Application Pipeline]] — semantic analysis is the stage where the scope tree is built and queried
- [[Tree Walking Patterns]] — `downup()` implements definition (descent) + resolution (ascent) in one pass
- [[Static Type Checking]] — uses the scope tree and symbol types to enforce type rules
- [[Parsing Patterns]] — imaginary tokens (`VARDECL`, `METHOD_DECL`) are defined in Ch.4 and consumed here
