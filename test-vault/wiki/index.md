# Wiki Index
_Last updated: 2026-06-02_

## Concepts

### Go

- [[Go Knowledge Map]] — MOC: all Go concept pages in dependency order
- [[Go Modules and Packages]] — module path as namespace; MVS; versioning rules; workspaces
- [[Go Error Handling]] — (value, error) returns; %w wrapping; sentinel vs custom error types
- [[Go Testing Patterns]] — TDD discipline; table-driven tests; testify; dependency injection
- [[Go File IO]] — WalkDir; os.Stat stability checks; defer close patterns

### CLI & Tooling

- [[Cobra CLI]] — PersistentPreRunE propagation and override gotcha; RunE vs Run; unknown subcommand handling
- [[XDG Base Directory]] — standard config/data/cache locations; os.UserConfigDir() cross-platform pattern
- [[SSH Config]] — per-host key selection via IdentityFile; wildcard Host * block; testing auth

### Design

- [[Agentic Workflow Patterns]] — catalog of patterns for multi-agent systems
- [[Harness Engineering]] — engineered runtime wrapping an LLM; 6 dimensions; named failure modes

## Entities

- [[magpie]] — vault-tools successor; plugin system, sonar split, vault resolution hierarchy

## Syntheses

- [[magpie-design-signals-from-wiki]] — what the vault's accumulated knowledge implies for magpie's design

## Open Questions

- [[Magpie Claim-Level Provenance]]
- [[Magpie Pipeline vs Hybrid]]
