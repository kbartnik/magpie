// Package result defines the value objects returned by magpie domain operations.
package result

// Status represents the outcome of a magpie operation.
type Status string

const (
	// StatusOK indicates success — caller proceeds.
	StatusOK Status = "ok"
	// StatusWarning indicates a user error or warning — caller may proceed.
	StatusWarning Status = "warning"
	// StatusBlocked indicates a hard block — destructive op with unmet precondition.
	StatusBlocked Status = "blocked"
)
