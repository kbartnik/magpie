# Wiki Index
_Last updated: 2026-06-02_

## Concepts

### Go

- [[Go Knowledge Map]] — MOC: all Go concept pages in dependency order
- [[Go Defined Types]] — Go's enum substitute: defined type + iota block; compiler-enforced discriminants
- [[Go Struct Memory Layout]] — field ordering controls padding and size; value vs pointer trade-offs
- [[Go Empty Struct Pattern]] — zero-byte presence marker; sets, channel signals, visited-node tracking
- [[Go Escape Analysis]] — compiler decides stack vs heap; returning pointers to small values is slower
- [[Go Memory Model]] — sequenced-before / synchronized-before / happened-before; what each primitive guarantees
- [[Go Interfaces]] — implicit implementation; accept interfaces, return structs; nil interface gotcha
- [[Go Functional Options Pattern]] — variadic Option func for backward-compatible config
- [[Go Generics Type Constraints]] — ~T underlying-type constraints; type inference; use for repetition, not polymorphism
- [[Go Channel Internals]] — hchan, sudog pooling, send/receive fast paths, scheduler integration
- [[Go Channel Concurrency Patterns]] — select, close, done-channel, worker pool, backpressure
- [[Go sync.Cond Pattern]] — N-dynamic waiters; Wait() atomically releases lock; spurious wakeup loop
- [[Go Context Patterns]] — request-scoped metadata only; cancellation/deadlines; unexported key type
- [[Go Error Handling]] — (value, error) returns; %w wrapping; sentinel vs custom error types
- [[Go File IO]] — WalkDir; os.Stat stability checks; defer close patterns
- [[Go Structured Logging]] — log/slog; Logger/Handler split; LogAttrs for zero-allocation hot paths
- [[Go Testing Patterns]] — TDD discipline; table-driven tests; testify; dependency injection
- [[Go Modules and Packages]] — module path as namespace; MVS; versioning rules; workspaces
- [[Go Workspaces]] — go work for multi-module dev; local only, don't commit go.work
- [[Go Tooling]] — staticcheck, golangci-lint, govulncheck; build tags; go:generate; cross-compilation
- [[Bubbletea Elm Architecture]] — Model/Update/View; tagged union Msg type; Cmd as deferred IO
- [[Bubbletea Async Patterns]] — two-message pattern for flow control; tea.Cmd composition
- [[Fsnotify Filesystem Watching]] — event types; debouncing; two-message pattern for bubbletea integration

### CLI & Tooling

- [[Cobra CLI]] — PersistentPreRunE propagation and override gotcha; RunE vs Run; unknown subcommand handling
- [[XDG Base Directory]] — standard config/data/cache locations; os.UserConfigDir() cross-platform pattern
- [[SSH Config]] — per-host key selection via IdentityFile; wildcard Host * block; testing auth
- [[Tmux]] — session/window/pane model; prefix key; status bar; copy mode
- [[LazyVim]] — Neovim distro; lazy.nvim plugin manager; which-key discoverability
- [[CLAUDE.md Configuration Patterns]] — project vs user config; zone rules; hook integration

### Design & Architecture

- [[Agentic Workflow Patterns]] — catalog of patterns for multi-agent systems
- [[Harness Engineering]] — engineered runtime wrapping an LLM; 6 dimensions; named failure modes
- [[Software Architecture Fundamentals]] — components, connectors, constraints; fitness functions
- [[IDSD]] — Interface Spec Design: separates spec into owned artifacts; ICE framework
- [[Systems Thinking]] — feedback loops, stocks and flows, leverage points
- [[Semantic Diffusion]] — how terms lose precision as they spread; Fowler's definition
- [[Connascence in Distributed Systems]] — coupling metric for distributed services
- [[JSON Schema Discipline]] — schema as contract; validation at boundaries
- [[LLM Wiki Pattern]] — using wikis as structured memory for LLM agents
- [[Progressive Disclosure Architecture]] — surface complexity only when needed
- [[Data Visualization Principles]] — visual encoding; pre-attentive attributes; chart selection
- [[Vibe-Coding Anti-Pattern]] — LLM-driven coding without understanding; failure modes

### Language Implementation

- [[Language Application Pipeline]] — lexer → parser → AST → semantic analysis → codegen pipeline
- [[Parsing Patterns]] — recursive descent, Pratt parsing, precedence climbing
- [[Symbol Tables and Scopes]] — scope chains, shadowing, resolution rules
- [[Static Type Checking]] — type inference, constraint propagation, Hindley-Milner
- [[Tree Walking Patterns]] — visitor pattern, interpreter pattern, AST traversal strategies
- [[Code Generation and Translation]] — IR, SSA, register allocation, instruction selection
- [[Interpreters and Bytecode VMs]] — bytecode design, stack vs register VMs, dispatch loops
- [[Mechanical Sympathy]] — hardware-aware software design; cache lines; branch prediction

### AI & ML

- [[Agent Memory Architectures]] — episodic, semantic, procedural, working memory in agents
- [[Agentic Identity and Zero Trust]] — trust boundaries in multi-agent systems
- [[AI Productivity Research]] — empirical studies on AI-assisted developer productivity
- [[AI Scraping Resistance]] — techniques to detect and deter AI training scraping
- [[Autonomous Learning Architecture]] — self-directed learning systems; curriculum design
- [[Claude Code Hooks]] — lifecycle hooks; pre/post-tool events; stop hooks
- [[Claude Code Memory Architecture]] — memory types; persistence; recall strategies
- [[Claude Code Skills]] — skill authoring; trigger conditions; rigid vs flexible skills
- [[Adversarial Examples]] — inputs designed to fool ML models; attack taxonomy
- [[Behavioral Poisoning Adoption Threshold]] — how much data poisoning is needed to shift behavior
- [[Bio-Inspired Computing]] — neural, evolutionary, swarm approaches; strengths and limits
- [[CLIP]] — contrastive image-language pretraining; zero-shot classification
- [[Community Self-Defense]] — collective security practices; mutual aid networks
- [[Diffusion Models]] — score matching, DDPM, DDIM; latent diffusion
- [[Emergence and Complexity]] — emergent behavior in complex systems; phase transitions
- [[MCP Protocol]] — Model Context Protocol; tool calling standard; server/client architecture
- [[Multimodal AI]] — vision-language models; modality fusion; benchmark gaps
- [[Neural Network Mechanics]] — backprop, activation functions, weight initialization
- [[Neuromorphic and Bio-Inspired AI]] — spiking neural nets; energy efficiency; event-driven compute
- [[Prompting as Specification]] — prompts as executable specs; precision requirements
- [[Retrieval-Augmented Generation]] — retrieval pipeline; chunk strategies; reranking
- [[Softmax]] — temperature scaling; logit interpretation; numerical stability
- [[Sparse Autoencoder]] — feature extraction; superposition hypothesis connection
- [[Stochastic Gradient Descent]] — minibatch dynamics; learning rate schedules; momentum
- [[Superposition Hypothesis]] — features as directions in activation space; polysemanticity
- [[Tokenization]] — BPE, WordPiece, SentencePiece; token boundary artifacts
- [[Tool Use in AI Systems]] — function calling; structured output; tool selection heuristics
- [[Transformer Architecture]] — attention mechanism; positional encoding; KV cache
- [[Vanishing Gradient Problem]] — gradient flow in deep nets; ResNet/LSTM solutions
- [[Threat Modeling]] — STRIDE; attack surface; threat actor modeling
- [[Operational Security]] — compartmentalization; need-to-know; tradecraft basics
- [[Security Culture]] — organizational security posture; behavior change; trust networks
- [[Digital Surveillance Resistance]] — counter-surveillance practices; metadata hygiene
- [[LLM Mental Model]] — how to reason about LLM capabilities and failure modes
- [[Context Rot]] — self-reinforcing LLM context degradation; remedy is deep modules with simple interfaces
- [[AI Without Illusions (Series)]] — synthesis series on realistic AI capabilities and limits
- [[Probability and Statistics Foundations]] — foundations for ML: distributions, inference, Bayes

### Society & Politics

- [[ADHD]] — executive function; dopamine regulation; diagnosis and management
- [[ADHD in Software Engineering]] — hyperfocus, context-switching costs, tooling adaptations
- [[Digital Biomarkers for ADHD]] — passive sensing; smartphone data; clinical validation
- [[Digital Phenotyping]] — continuous behavioral measurement via device sensors
- [[Digital Therapeutics for ADHD]] — software-based interventions; FDA regulation
- [[Mutual Aid]] — reciprocal community support; distinction from charity
- [[Qualitative Populism]] — research approach centering lived experience; methodological notes
- [[Race Class Narrative]] — Haney López's framework; race and class as linked, not competing
- [[Strategic Racism]] — political weaponization of race to divide class interests
- [[Stoicism and AI Disruption]] — Stoic frameworks applied to technological displacement
- [[Transformative Justice]] — community-based accountability; alternatives to punitive systems
- [[Ur-Fascism]] — Eco's 14 features of eternal fascism

## Entities

- [[magpie]] — vault-tools successor; plugin system, sonar split, vault resolution hierarchy
- [[preflight-sync-go]] — Go TUI + file-sync learning project; uses bubbletea and fsnotify
- [[Learning Go]] — Jon Bodner, O'Reilly 2024; canonical source for the Go knowledge cluster
- [[Jon Bodner]] — author of Learning Go
- [[Dave Cheney]] — originator of the Functional Options pattern
- [[Terence Parr]] — creator of ANTLR; author of Language Implementation Patterns
- [[Language Implementation Patterns]] — Terence Parr; definitive guide to building language tools
- [[Robert Nystrom]] — author of Crafting Interpreters
- [[Crafting Interpreters]] — Robert Nystrom; free online book; tree-walk then bytecode VM
- [[Brian Hogan]] — author of tmux 2 and other developer tools books
- [[tmux 2]] — Brian Hogan; practical tmux; O'Reilly
- [[Drew Neil]] — author of Practical Vim and Modern Vim
- [[Modern Vim]] — Drew Neil; Neovim-focused; plugin ecosystem
- [[Dusty Phillips]] — author of Python books; OOP and design patterns
- [[Martin Fowler]] — refactoring patterns; semantic diffusion; enterprise architecture
- [[Diana Montalion]] — author of Fundamentals of Software Architecture (O'Reilly)
- [[Fundamentals of Software Architecture]] — Mark Richards & Neal Ford; architecture styles and trade-offs
- [[Andrej Karpathy]] — researcher; neural nets; LLM mechanistic interpretability
- [[Daniel Kunin]] — creator of Seeing Theory; visual probability/statistics
- [[Seeing Theory]] — visual intro to probability and statistics; Brown University
- [[Making Data Visual]] — Danyel Fisher; data visualization for data science
- [[Designing Data Visualizations]] — Noah Iliinsky; design principles for charts and graphs
- [[rtk]] — Ruby toolkit; (context TBD from page)
- [[Complete Guide to Building Skills for Claude]] — official Anthropic skills guide
- [[Nexus Vault]] — this vault; the knowledge system being built
- [[Learning Systems Thinking]] — Donella Meadows; Thinking in Systems
- [[LazyVim for Ambitious Developers]] — book on LazyVim configuration and workflow
- [[Thomas Byern]] — author; context TBD from page
- [[Umberto Eco]] — semiotician; author of Ur-Fascism essay and fiction

## Syntheses

- [[magpie-design-signals-from-wiki]] — what the vault's accumulated knowledge implies for magpie's design
- [[Go Program Instrumentation]] — full sequence: -race → testing.B → -benchmem → pprof → slog
- [[Nexus Vault Template]] — template design for replicable vault structure

## Open Questions

- [[Magpie Claim-Level Provenance]]
- [[Magpie Pipeline vs Hybrid]]
- [[Go Generics for context.WithValue]]
- [[errgroup All-Errors Collection]]
- [[CGo Goroutine Scheduler Interaction]]
- [[GODEBUG Governance Model]]
- [[Entropy-Degraded Codebase Recovery]]
- [[Interface Design Responsibility at Scale]]
- [[Ubiquitous Language Ownership]]
- [[AI-Checking-AI Stop Hooks]]
- [[Behavioral Poisoning Adoption Threshold]]
- [[CLAUDE.md Degradation Mechanism]]
- [[Data Poisoning Legal Status]]
- [[Diffused Responsibility at Scale]]
- [[Dumb Zone Threshold Universality]]
- [[Ghosts Framing Engineering Prescriptions]]
- [[Hard Deny Hooks Interaction Order]]
- [[ICE Framework Vault Ingest Mapping]]
- [[Karpathy-METR Productivity Gap]]
- [[Poisoning Counter-Adaptation Timeline]]
- [[SLM Performance with Fine-Tuning Cost]]
- [[Stoic Acceptance as Rationalization]]
- [[System M Achievability]]
- [[Training Poisoning Degradation at Scale]]
- [[Vault Security Gate Exit Code]]
- [[Agentic RAG vs Advanced RAG Threshold]]
- [[Bio-Inspired vs Gradient Descent]]
- [[Episodic Memory Forgetting Problem]]
- [[Leverage Points for Software Systems]]
- [[Multimodal Quality Gap for Enterprise]]
- [[Onto AI Scalability and Auditability]]
- [[Typed Wikilinks Tractability]]
- [[Vault Frozen-Agentic Position]]
- [[Vault LLM Wiki Failure Modes]]
- [[Vault Retrieval Layer Threshold]]
- [[Vault Threat Model if Compromised]]
- [[Wolfram Irreducibility and AI]]
- [[ADHD Disclosure Risk and Policy]]
- [[AI Firewall Attack Surface]]
- [[Alignment as Iterative Design]]
- [[Compartmentalization vs Movement Building]]
- [[Computational Irreducibility Falsifiability]]
- [[Conceptual Integrity Divergence Detection]]
- [[Config Drift vs Adaptive Learning]]
- [[Connascence in Distributed Systems]]
- [[Credential Renewal for Long-Running Tasks]]
- [[Disappearing Messages vs Documentation]]
- [[Dopamine Transfer Deficit Subtypes]]
- [[Employer Monitoring as ADHD Sensing]]
- [[Encrypted Comms as Surveillance Signal]]
- [[Fitness Functions Political Viability]]
- [[Neurodiverse SE Research Post-LLM Validity]]
- [[Security Culture in Smartphone Era]]
- [[Verification vs OpSec Paper Trail]]
- [[RCN and Organic Racism]]
- [[RCN Messaging vs Security Culture]]
- [[Visibility Without Pressure Failure Mode]]
