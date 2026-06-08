# magpie

The AI harness and executive layer for an LLM-driven wiki — a compiled Go binary that handles every mechanical vault operation so the LLM doesn't have to.

> **Status:** early development (Phase 0 — Foundation). Not yet usable. This is a phase-by-phase Go learning project, built in the open with TDD throughout.

## What It Is

The LLM (Claude, or any future provider) is magpie's primary consumer. Magpie returns information to the LLM in a minimal, machine-friendly way using fixed, versioned schemas (the **oracle**), and handles every purely mechanical wiki operation — vault discovery, file moves, archiving, linting — that the LLM shouldn't have to reason through (the **executor**).

Magpie has zero AI dependency in its core and runs identically with any provider, or none. Designing *for* an LLM caller and depending *on* one are different things — Claude Code integration lives entirely in a separate plugin, [`magpie-claude`](https://github.com/kbartnik/magpie-claude).

## Documentation

| Doc | What's in it |
|---|---|
| [Design Overview](docs/design-overview.md) | A five-minute read: what magpie is, why it exists, and the principles governing every design decision |
| [Design Spec](docs/design-spec.md) | Implementation reference: schemas, config formats, command contracts, phase sequence |

## Building

Requires Go (see [`go.mod`](go.mod) for the version).

```sh
go build .
go test ./...
```

## Project Status

Magpie is built phase by phase — no phase starts until the previous one passes `go test ./...`. The [Design Spec](docs/design-spec.md) lays out the full phase sequence; phase plans are written just-in-time and tracked in the project's planning workspace, not in this repo.
