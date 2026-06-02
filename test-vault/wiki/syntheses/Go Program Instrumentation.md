---
title: "Go Program Instrumentation"
type: synthesis
status: active
created: 2026-05-31
updated: 2026-05-31
sources:
  - archive/books/2026-05-31-learning-go-2e.md
related:
  - Go Testing Patterns
  - Go Structured Logging
  - Go Escape Analysis
  - Go Memory Model
  - Go Tooling
  - preflight-sync-go
tags:
  - go
  - testing
  - observability
  - performance
  - synthesis
---

# Go Program Instrumentation

Go's instrumentation tools form a mandatory sequence, not a menu. The discipline is: establish correctness, then measure, then diagnose, then observe at runtime. Skipping the sequence produces optimization that either doesn't hold or can't be explained.

## The Sequence

```
-race → testing.B → -benchmem → pprof → slog
  ↑           ↑            ↑          ↑        ↑
correct    measure      allocs    explain   observe
```

Each stage has a specific question it answers. Later stages depend on earlier ones being resolved first.

## Stage 1 — Correctness Before Measurement (`-race`)

Run `-race` before interpreting any performance number. A data race means the benchmark is measuring undefined behavior — values read under a race can be arbitrary, making timing results meaningless. Concurrency bugs that pass unit tests reliably appear under the race detector because it observes actual memory access patterns at runtime, not just logical outcomes.

Fix races before benchmarking. There are no exceptions.

## Stage 2 — Benchmarks for Before/After Comparisons (`testing.B`)

The benchmark number in isolation is meaningless. The *delta between two implementations* is the signal. This reframes what a benchmark is for: not "how fast is this?" but "did this change make it faster, and by how much?"

```go
func BenchmarkProcess(b *testing.B) {
    b.ReportAllocs()  // always include allocation count
    for b.Loop() {    // Go 1.24+; use b.N loop for earlier versions
        result = Process(input)
    }
}
```

Stabilize results with `-count=5 -benchtime=3s` to average across GC pauses and scheduler noise. A single benchmark run is not a data point — it's an anecdote.

The `-benchmem` flag (or `b.ReportAllocs()` inline) is non-optional. In Go, allocation count often tells you more than nanoseconds: an optimization that saves 40ns but adds two allocations per call is usually a regression at scale due to GC pressure accumulating across many calls.

## Stage 3 — The Allocation Insight

Returning a pointer to a small struct is *slower* than returning the value (~30ns pointer vs ~10ns for a 100-byte struct) because pointer returns force heap allocation — the compiler's escape analysis determines the data can't stay on the stack, so it lands on the heap and adds GC work.

This means: a benchmark showing a "faster" implementation that doubles allocations is trading latency now for GC pauses later. `-benchmem` makes the allocation cost visible at the macro level; [[Go Escape Analysis]]'s `gcflags=-m` flag confirms the cause at the source level:

```bash
go build -gcflags="-m" ./...  # shows escape decisions per variable
```

The two tools compose: benchmark flags tell you *that* allocations changed, escape analysis tells you *why*.

## Stage 4 — pprof for Diagnosis (Gap)

`pprof` is the missing middle layer between "I know allocations changed" and "I know why and where." It produces CPU, memory, goroutine, and mutex contention profiles that pinpoint which functions are consuming resources under real load — not in isolated benchmarks, but in the full call graph.

```go
import _ "net/http/pprof"  // registers /debug/pprof endpoints

// Or for CLI tools:
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

This stage is not yet documented in the vault. It's the primary gap in the instrumentation sequence — worth ingesting before doing serious optimization work on any Go program, including preflight-sync-go.

## Stage 5 — Runtime Observability (`log/slog`)

Benchmarks measure isolated functions in controlled conditions. `log/slog` with structured latency fields measures the running program under real load — different users, different file sizes, different concurrency patterns.

The key instrumentation pattern: inject a `time.Now` function parameter so latency measurement is testable without real clocks (connects to [[Go Testing Patterns]] dependency injection via function parameters):

```go
func ProcessFiles(paths []string, now func() time.Time) error {
    start := now()
    // ...
    slog.LogAttrs(ctx, slog.LevelInfo, "processed",
        slog.Int("count", len(paths)),
        slog.Duration("elapsed", now().Sub(start)),
    )
}

// Production
ProcessFiles(paths, time.Now)

// Test — deterministic timing
ProcessFiles(paths, func() time.Time { return fixedTime })
```

`Logger.With` for request-scoped correlation means every log line carries the relevant fields without repetition. `slog.Duration` is a typed attribute — no string formatting, no parsing required downstream. See [[Go Structured Logging]] for the full API.

## How the Stages Compose for preflight-sync-go

For a file-sync TUI with concurrency:

1. `-race` on the full test suite — channels and goroutines are present; race detector is mandatory
2. Benchmark the core transfer loop with `-benchmem` — allocation count per file copied is the primary signal
3. `gcflags=-m` on the hot path to confirm escape behavior matches expectations
4. `pprof` CPU profile under a realistic workload (large directory, many small files) to find the actual bottleneck
5. `slog` structured latency fields on the transfer completion event for production observability

## Open Questions

- At what file count / concurrency level does pprof profiling overhead become non-negligible? Profile overhead is ~5-10% CPU for production systems.
- Does the Go scheduler's integration with the network/filesystem poller affect benchmark stability for I/O-bound code? `-benchtime` may need to be longer for I/O benchmarks than CPU benchmarks.
- `b.Loop()` (Go 1.24) vs the traditional `for i := 0; i < b.N; i++` loop — the new form handles warmup automatically; worth verifying which Go version preflight-sync-go targets.

## See Also

- [[Go Testing Patterns]] — unit testing foundation; DI via function parameters for testable instrumentation
- [[Go Structured Logging]] — slog API, LogAttrs, Logger.With for runtime observability
- [[Go Escape Analysis]] — gcflags=-m, stack vs heap decisions, GC pressure
- [[Go Memory Model]] — what the concurrent benchmarks actually guarantee
- [[Go Tooling]] — staticcheck, linters; build tags for benchmark-only files
- [[preflight-sync-go]] — active project this sequence applies to directly
