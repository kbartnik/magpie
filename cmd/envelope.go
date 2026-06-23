package cmd

import "github.com/kbartnik/magpie/internal/result"

// SchemaVersion is the current envelope contract version.
const SchemaVersion = 0

// Envelope wraps every magpie response in the harness contract.
// The LLM checks SchemaVersion on every call; a mismatch triggers `magpie schema`.
type Envelope[T any] struct {
	SchemaVersion int           `json:"schema_version"`
	Status        result.Status `json:"status"`
	DryRun        bool          `json:"dry_run,omitempty"`
	Data          T             `json:"data"`
	Effects       []string      `json:"effects,omitempty"` // what changed: "created:path", "modified:path", "deleted:path"
	Delta         any           `json:"delta,omitempty"`   // vault health changes (null until VaultGraph exists)
}

// Wrap returns data wrapped in an envelope at the current schema version.
func Wrap[T any](data T) Envelope[T] {
	return Envelope[T]{
		SchemaVersion: SchemaVersion,
		Status:        result.StatusOK,
		Data:          data,
	}
}
