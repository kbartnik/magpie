---
title: "Interpreters and Bytecode VMs"
type: concept
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - "archive/books/2026-05-31-language-implementation-patterns-ch9-13.md"
  - "archive/books/2026-06-01-crafting-interpreters.md"
related:
  - "Language Application Pipeline"
  - "Tree Walking Patterns"
  - "Symbol Tables and Scopes"
  - "Language Implementation Patterns"
tags: [compilers, interpreters, bytecode, vm, language-tools]
---

# Interpreters and Bytecode VMs

The execution layer of a [[Language Application Pipeline]]: given a validated, annotated AST, how do you run it? Two main approaches, same semantics, very different performance characteristics.

## Pattern 1: High-Level (Tree-Walking) Interpreter

Walk the AST with an external visitor; execute each node directly as you visit it.

```java
class Interpreter extends ASTVisitor {
    Object visitAssign(AssignNode n) {
        Object val = visit(n.value);   // evaluate RHS
        memory.set(n.target.name, val); // store to memory
        return null;
    }
    Object visitAdd(BinaryOpNode n) {
        int left  = (int) visit(n.left);
        int right = (int) visit(n.right);
        return left + right;
    }
    Object visitCall(CallNode n) {
        // push new scope, execute body, pop scope
        ...
    }
}
```

**Advantages:** Simple — the AST structure mirrors the language semantics; no compilation step needed. Easy to debug and extend.

**Disadvantages:** Recursive tree traversal is slow. Each node dispatch involves pointer chasing, virtual method calls, and deep call stacks. 10–100× slower than bytecode.

**Use when:** Prototyping, scripting languages where startup time matters more than throughput, or when the language is simple enough that performance isn't a concern.

## Pattern 2: Bytecode Compiler + Stack-Based VM

Compile the AST to a flat linear sequence of bytecode instructions, then run a tight interpreter loop over those instructions.

### The Compilation Step

Walk the AST and emit bytecode instructions into a code buffer:

```java
void compileExpr(BinaryOpNode n) {
    compile(n.left);   // emit: push left value
    compile(n.right);  // emit: push right value
    emit(ADD);         // emit: pop two, push sum
}
void compileAssign(AssignNode n) {
    compile(n.value);          // emit: evaluate RHS onto stack
    emit(STORE, n.target.idx); // emit: pop and store to local
}
```

### Instruction Set

A minimal stack-based bytecode VM needs:

| Instruction | Effect |
|-------------|--------|
| `iconst N` | Push integer constant N |
| `load i` | Push local variable at index i |
| `store i` | Pop top of stack into local i |
| `gload a` | Push global variable at address a |
| `gstore a` | Pop into global at address a |
| `add`, `sub`, `mul` | Pop two, push result |
| `br a` | Unconditional branch to address a |
| `brt a` | Branch to a if top-of-stack is true |
| `brf a` | Branch to a if top-of-stack is false |
| `call f` | Push call frame, jump to function f |
| `ret` | Pop call frame, return to caller |

### The Interpreter Loop

```java
while (true) {
    int op = code[ip++];
    switch (op) {
        case ADD:   sp--; operands[sp] = operands[sp] + operands[sp+1]; break;
        case LOAD:  operands[++sp] = callStack[fp].locals[code[ip++]]; break;
        case STORE: callStack[fp].locals[code[ip++]] = operands[sp--]; break;
        case CALL:  pushFrame(code[ip++]); break;
        case RET:   popFrame(); break;
        // ...
    }
}
```

The tight switch/dispatch loop over a flat array is dramatically faster than recursive AST traversal — no pointer chasing between tree nodes, no recursive call overhead.

### Call Frames

Each function call pushes a **call frame** onto the call stack:
- Return address (where to resume after `ret`)
- Local variable slots (indexed, not named — names resolved at compile time)
- Operand stack pointer

Function calls become: push frame → jump to function start → execute → `ret` → pop frame → resume.

### Why Bytecode Over Machine Code?

- **Portability:** Bytecodes run on any machine with a compatible VM; machine code is CPU-specific
- **Safety:** The VM can bounds-check, type-check, and garbage collect in ways native code cannot
- **Simplicity:** Generating correct machine code (register allocation, instruction scheduling, ABI compliance) is far harder than generating bytecodes

This is why JVM, CPython, Lua, and Ruby all use bytecode VMs.

### Stack-Based vs Register-Based

Stack-based VMs (JVM, CPython): all operands flow through a single operand stack. Simple to generate code for; compact bytecode.

Register-based VMs (Lua 5, Dalvik): operands addressed by register number in each instruction. ~32% faster in practice (fewer push/pop operations), but larger instruction encoding and harder code generation.

## Runtime Symbol Tracking

At runtime, the scope tree concept from [[Symbol Tables and Scopes]] maps to the call stack:
- The call stack is the runtime version of the scope chain
- Each call frame holds the local variables for one scope
- Global variables live in a flat global memory array (resolved at compile time to numeric addresses)

Names are resolved at compile time; at runtime, all accesses are by index — no hash table lookups.

## See Also

- [[Language Application Pipeline]] — interpreters and VMs are the Interpreter application category
- [[Tree Walking Patterns]] — tree-walking interpreter is the direct application of the external visitor pattern to execution
- [[Symbol Tables and Scopes]] — compile-time scope tree maps to the runtime call stack structure
- [[Code Generation and Translation]] — the alternative to interpretation: emit source or target language output instead of executing
