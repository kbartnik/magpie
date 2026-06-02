---
title: "Static Type Checking"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns-ch6-8.md"
related:
  - "Symbol Tables and Scopes"
  - "Language Application Pipeline"
  - "Tree Walking Patterns"
tags: [compilers, type-system, type-checking, semantic-analysis, language-tools]
---

# Static Type Checking

Static type checking is the semantic analysis phase that verifies type rules at compile time. It runs after [[Symbol Tables and Scopes]] have been fully built — types are themselves symbols, and checking requires knowing what type each expression resolves to.

## Types as Symbols

Built-in types (`int`, `float`, `bool`) and user-defined types (`class T`) all live in the symbol table as type symbols:

```java
class BuiltInTypeSymbol extends Symbol implements Type { … }
class ClassSymbol extends ScopedSymbol implements Type { … }
```

This uniformity means type lookup uses the same scope-chain mechanism as variable lookup: `ref("int")` walks up the scope tree and finds the built-in type symbol.

## Computing Expression Types

Type checking happens during the resolution phase tree walk. For each expression node, compute its type bottom-up:

| Expression | Type rule |
|------------|-----------|
| Integer literal | `int` |
| Float literal | `float` |
| Variable ref `x` | `x.symbol.type` |
| Binary op `a + b` | `promote(type(a), type(b))` |
| Method call `f()` | `f.symbol.returnType` |
| Field access `obj.x` | `x.symbol.type` |

The resolved type is stored back on the AST node for use by later passes.

## Type Promotion

When two types appear in an arithmetic expression, the "smaller" type is promoted to the "larger":

```
int + float  →  float   (int promoted to float)
int + int    →  int
```

Promotion rules are encoded as a precedence table or as methods on type objects. The `promote(t1, t2)` function returns the dominant type.

## Assignment Compatibility

Type checking for assignment (`x = expr`) requires asking: can the expression's type be assigned to the variable's declared type?

For procedural languages, this is usually exact match or numeric promotion:
```java
boolean canAssign(Type valueType, Type destType) {
    return valueType == destType || isPromotion(valueType, destType);
}
```

For OO languages, subtype compatibility is added: a `Dog` can be assigned to an `Animal` variable if `Dog` extends `Animal`. This is delegated to the type objects themselves:

```java
// On Type interface
boolean canAssignTo(Type destType);

// On ClassSymbol
boolean canAssignTo(Type destType) {
    if (this == destType) return true;
    if (superClass != null) return superClass.canAssignTo(destType);
    return false;
}
```

Walking the superclass chain in `canAssignTo()` implements subtype polymorphism as a recursive parent-chain traversal — the same structure as lexical scope lookup but over the class hierarchy instead.

## A Complete Type-Checking Pass

```
// After definition + resolution phases have run:

Walk AST with type-checking visitor:

  visit BinaryOp(+, left, right):
    tLeft  = visit(left)   // recurse, get type
    tRight = visit(right)  // recurse, get type
    result = promote(tLeft, tRight)
    node.evalType = result
    return result

  visit Assign(target, value):
    tTarget = target.symbol.type
    tValue  = visit(value)
    if !tValue.canAssignTo(tTarget):
        error("type mismatch: cannot assign %s to %s", tValue, tTarget)

  visit MethodCall(f, args):
    method = f.symbol  // already resolved in resolution phase
    checkArgTypes(method.parameters, args)
    return method.returnType
```

## Relationship to Scope Trees

Type checking is a consumer of the scope tree built in [[Symbol Tables and Scopes]]:
- Every symbol reference was resolved to a symbol object in the resolution phase
- Type checking reads `symbol.type` from those resolved symbols
- No additional scope lookups are needed — the type info is already on the AST nodes

This is why the three phases (definition → resolution → type-checking) form a clean pipeline: each phase annotates the AST with information the next phase consumes.

## See Also

- [[Symbol Tables and Scopes]] — prerequisite; types live in the symbol table; resolution phase must run first
- [[Language Application Pipeline]] — type checking is part of the Semantic Analyzer stage
- [[Tree Walking Patterns]] — the type-checking pass is an external visitor with bottom-up (post-order) traversal
