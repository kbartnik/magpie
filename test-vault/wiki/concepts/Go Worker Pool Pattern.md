---
tags: [concept, go, concurrency]
cluster: go
aliases: ["goroutine pool", "worker pool", "bounded concurrency", "work queue pattern"]
related: ["Go Channel Concurrency Patterns", "Go Select Statement", "Go Context Patterns", "Go Memory Model"]
sources:
  - "[[archive/videos/2026-06-04-go-concurrency-patterns]]"
---

# Go Worker Pool Pattern

N goroutines consuming from a shared buffered channel. The buffer capacity is the concurrency bound — it controls how many work items can be queued without blocking the producer.

## Implementation

```go
func workerPool(ctx context.Context, numWorkers int, jobs <-chan Job) <-chan Result {
    results := make(chan Result, numWorkers)
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                select {
                case job, ok := <-jobs:
                    if !ok { return }
                    results <- process(job)
                case <-ctx.Done():
                    return
                }
            }
        }()
    }
    go func() { wg.Wait(); close(results) }()
    return results
}
```

## Buffer Capacity = Concurrency Bound

Formal basis from the Go Memory Model: the kth receive from a channel of capacity C happens-before the (k+C)th send. Setting buffer size = N workers ensures N items in flight at once.

## Connections

- [[Go Channel Concurrency Patterns]] — worker pool is the canonical application of buffered channels and fan-in
- [[Go Select Statement]] — workers use select to drain the job channel while respecting context cancellation
- [[Go Context Patterns]] — context cancellation is the standard shutdown mechanism for worker pools
- [[Go Memory Model]] — the buffered channel guarantee is the formal basis for buffer-capacity-as-concurrency-bound
