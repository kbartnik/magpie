---
title: "Code Generation and Translation"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns-ch9-13.md"
related:
  - "Language Application Pipeline"
  - "Tree Walking Patterns"
  - "Interpreters and Bytecode VMs"
  - "Language Implementation Patterns"
tags: [compilers, code-generation, translation, templates, language-tools]
---

# Code Generation and Translation

The output stage of a [[Language Application Pipeline]]: given a validated, annotated AST, emit text or bytecode. The Translator and Generator application categories from the pipeline both live here.

## Model-Driven Translation

The dominant pattern for source-to-source translators and code generators:

1. **Parse** input → build AST (the "input model")
2. **Annotate** AST with semantic information (symbol resolution, types)
3. **Walk** the AST with an output-generating visitor
4. **Emit** output text for each node

```java
class JavaToGoTranslator extends ASTVisitor {
    void visitMethodDecl(MethodDeclNode n) {
        emit("func ");
        emit(n.name);
        emit("(");
        visitParams(n.params);
        emit(") ");
        visitReturnType(n.returnType);
        emit(" {\n");
        visit(n.body);
        emit("}\n");
    }
}
```

The AST is the stable interface between recognition and generation. It lets you make multiple passes — type-checking, optimization, then generation — without re-parsing.

## Syntax-Directed Translation

For simpler translators, generation actions are embedded directly in grammar rules. As the parser matches a rule, it fires the corresponding output action immediately.

```antlr
// Grammar rule with embedded translation action
statement
    : 'print' expr  { out.println($expr.text); }
    | assignment
    ;
```

**Advantage:** Simple — no separate AST needed for straightforward 1:1 translations.  
**Disadvantage:** Can't do multi-pass analysis; generation is interleaved with parsing, making it hard to look ahead or back.

**Use when:** The translation is local (output depends only on the current rule, not context from elsewhere in the tree).

## StringTemplate: Separating Structure from Logic

Ad-hoc code generation with print statements mixes output structure with traversal logic, making both harder to read and modify. StringTemplate enforces separation:

```java
// Template file (defines output structure)
// method(name, params, returnType, body) ::= <<
// func <name>(<params>) <returnType> {
//     <body>
// }
// >>

// Generator (traversal logic only)
ST tmpl = group.getInstanceOf("method");
tmpl.add("name", node.name);
tmpl.add("params", visitParams(node.params));
tmpl.add("body", visit(node.body));
emit(tmpl.render());
```

**Why this matters:** Templates enforce that output *structure* is declared in one place (the template file) and output *content* is computed in another (the walker). This is the same principle as CSS/HTML separation — structural changes don't require touching business logic and vice versa.

**Template "holes":** Named attributes filled by the generator. Templates can have conditionals and loops but deliberately exclude arbitrary code execution — this restriction is a feature, not a limitation.

## Choosing Between Approaches

| Situation | Approach |
|-----------|----------|
| Simple 1:1 mappings, no look-ahead needed | Syntax-directed (grammar actions) |
| Multi-pass, context-dependent generation | Model-driven (AST walker) |
| Complex output formats, team separation | StringTemplate |
| Need execution, not output | [[Interpreters and Bytecode VMs]] |

## When You Don't Need a Full Pipeline

The book's final example: a wiki-to-HTML translator implemented as a *lexer only* — no parser, no AST. Wiki syntax is regular enough that a streaming lexer with filter rules can translate in one pass.

The lesson: the full Reader→IR→Analyzer→Generator pipeline is the general solution. For restricted input languages, simpler approaches (pure lexer, regex rewriting, token streaming) may be sufficient. The pipeline adds value when:
- You need multiple passes
- The input has non-trivial nesting or context-dependence
- You need error recovery or semantic analysis

## See Also

- [[Language Application Pipeline]] — code generation is the Generator stage; translation is the Translator category
- [[Tree Walking Patterns]] — the AST walking that drives model-driven translation
- [[Interpreters and Bytecode VMs]] — the alternative to generating output: executing the AST or bytecode directly
- [[Language Implementation Patterns]] — Terence Parr's full 31-pattern catalog; StringTemplate is his own tool
