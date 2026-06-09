---
tags: [concept, go]
cluster: go
aliases: ["go http client", "go http", "http.Client", "go API client"]
related: ["LLM Tool Calling", "Go Context Patterns", "Go Error Handling", "Go Interfaces"]
sources:
  - "[[archive/clippings/2026-06-04-llm-tool-use-in-go]]"
---

# Go HTTP Client Patterns

## The Default Client Trap

`http.DefaultClient` has no timeout. In production, always construct a client with explicit timeouts:

```go
client := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        DialContext: (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
        TLSHandshakeTimeout: 5 * time.Second,
    },
}
```

## Context Propagation

All requests should carry context for cancellation and deadlines:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
```

## JSON Request/Response

```go
// Encode request
buf := &bytes.Buffer{}
json.NewEncoder(buf).Encode(payload)

// Decode response
var result ResponseType
json.NewDecoder(resp.Body).Decode(&result)
defer resp.Body.Close()
```

## Retry with Backoff

For LLM API calls: retry on 429 (rate limit) and 5xx (server error) with exponential backoff. Use `context.WithTimeout` for total operation deadline.

## Connections

- [[LLM Tool Calling]] — Go HTTP clients are the execution layer for LLM tool calls against external APIs
- [[Go Context Patterns]] — `http.NewRequestWithContext` is the standard way to propagate cancellation into HTTP calls
- [[Go Error Handling]] — HTTP errors require checking both the transport error and `resp.StatusCode`
