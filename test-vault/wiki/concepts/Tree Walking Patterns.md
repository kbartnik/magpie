---
title: "Tree Walking Patterns"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns-ch3-5.md"
  - "archive/books/2026-06-01-crafting-interpreters.md"
related:
  - "Language Application Pipeline"
  - "Parsing Patterns"
  - "Language Implementation Patterns"
tags: [compilers, AST, tree-walking, visitor-pattern, language-tools]
---

# Tree Walking Patterns

Once a parser builds an AST, downstream stages need to traverse it — to extract information, check semantics, or produce output. The choice of tree-walking pattern affects how cleanly you can add passes, handle different node types, and rewrite subtrees.

## Two Core Approaches

### Pattern 12: Embedded Heterogeneous Tree Walker

Traversal methods live directly on each node class.

```java
class AssignNode extends ASTNode {
    void walk() {
        target.walk();  // recurse into children
        value.walk();
        // action: do something with this node
    }
}
```

**When to use:** Simple tools with one pass, where tight coupling between node structure and traversal logic is acceptable.

**Drawback:** Adding a new kind of traversal (e.g., type-checking pass after a symbol-resolution pass) requires modifying every node class.

### Pattern 13: External Tree Visitor

Traversal logic lives in a separate visitor object. Node classes only need an `accept(visitor)` hook.

```java
class TypeChecker implements TreeVisitor {
    void visitAssign(AssignNode n) {
        visit(n.target);
        visit(n.value);
        // check type compatibility
    }
    void visitReturn(ReturnNode n) { … }
}
```

**When to use:** Any tool with multiple independent passes, or when you can't (or don't want to) modify node classes.

**Advantage:** Each pass is a self-contained class. Adding a new pass doesn't touch existing nodes. This is the standard approach for real language tools.

## Automating Visitors: Tree Grammars and Pattern Matchers

Rather than building visitors by hand, ANTLR lets you specify them declaratively.

### Pattern 14: Tree Grammar

Describes the **entire valid structure** of all trees. The generator uses it to build a complete walker.

```antlr
expr : ^('+' expr expr)
     | ^('*' expr expr)
     | INT
     | ID
     ;
```

The `^(op child child)` notation means "node `op` with children." The grammar is exhaustive — it must handle every possible subtree.

**Use when:** You need to process all nodes and the tree structure is fully specified.

### Pattern 15: Tree Pattern Matcher

Describes only the **subtrees you care about**. The matcher fires actions when patterns match anywhere in the tree; non-matching subtrees are traversed silently.

```antlr
// Only fires for multiply-by-zero; ignores everything else
^('*' . INT[0]) -> INT[0]    // rewrite 4*0 → 0
^('*' INT[0] .) -> INT[0]    // rewrite 0*4 → 0
```

**Use when:** You're doing targeted rewrites or extracting specific constructs. Less specification than a full tree grammar — you only describe what you care about.

**Decision threshold:** If you need to handle every node (type-checking, code generation), use a tree grammar. If you're doing localized rewrites or analysis (constant folding, dead code elimination), use a pattern matcher.

## Rewriting Direction

Some transformations must happen in a specific tree traversal order.

### Top-down (pre-order)

Rewrite a node before descending into its children. Use for:
- Dead code elimination: if a conditional is statically false, prune the subtree before traversing it
- Macro expansion: replace a node before processing its contents

### Bottom-up (post-order)

Rewrite a node after all its children have been processed. Use for:
- Constant folding: `4 * 0 * 2` — reduce the `4*0` subtree to `0` before checking the parent `*`
- Type inference: a node's type depends on its children's resolved types

### Mixed

Complex optimizers often combine both: top-down passes for structural simplifications, bottom-up passes for value propagation. Tree grammars support separate rule sets for descent and ascent.

## The "Launch Missiles" Problem

Backtracking parsers try alternatives speculatively and back up on failure. Any side effects executed during a failed alternative can't be undone. This is why real language tools defer actions to separate tree-walking passes rather than embedding them in the parser:

1. **Parse** — build the AST; no actions, no side effects
2. **Walk** — traverse the AST in one or more passes; perform actions here

This pattern separates recognition (does the input conform to the grammar?) from analysis (what does it mean?), making both stages simpler and independently testable.

## AST Construction Notes

### Imaginary Tokens

When no input token serves as a natural subtree root (variable declarations, function signatures), invent one:

```
// Input: int i;
// Natural subtree root: none (int and i are both operands, not operators)

VARDECL
├── int
└── i
```

Common imaginary tokens: `VARDECL`, `FUNCDECL`, `CLASSDECL`, `BLOCK`, `ARGS`.

### Node Children: List vs Named Fields

| Homogeneous | Heterogeneous |
|-------------|--------------|
| `children []*ASTNode` (generic list) | Named fields (`Target Expr`, `Value Expr`) |
| Simple; any number of children | Type-safe; compiler checks field access |
| Requires type-checking at traversal time | More node types to define |

Homogeneous is faster to prototype. Heterogeneous pays off for tools with many passes over a stable grammar.

## Reference Implementation: jlox

[[Crafting Interpreters]] Part II (jlox) is the most complete worked example of the External Visitor pattern applied to a full language. jlox implements:

- A generated `Expr` hierarchy with `accept(Visitor)` on each node
- Four visitor passes: pretty-printer, interpreter, resolver, and class checker
- The resolver pass as a pre-execution walk that annotates variables with their static depth (see [[Symbol Tables and Scopes]])

jlox's structure makes the separation of passes concrete: each pass is a separate Java class implementing `Expr.Visitor<R>`, traversing the same heterogeneous AST. Comparing jlox to clox (the bytecode VM in Part III) shows directly what you gain and give up by moving from tree-walking to a compiled representation.

## See Also

- [[Parsing Patterns]] — the layer that builds the AST this page walks
- [[Language Application Pipeline]] — tree walking is the Semantic Analyzer stage of the pipeline
- [[Language Implementation Patterns]] — source; Terence Parr's Pattern 12–15
