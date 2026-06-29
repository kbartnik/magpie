package cmd

import "github.com/kbartnik/magpie/internal/result"

// ExitCoder is implemented by errors that carry a specific process exit code.
type ExitCoder interface {
	ExitCode() int
}

type statusError struct {
	status result.Status
	err    error
}

func (e *statusError) Error() string { return e.err.Error() }
func (e *statusError) ExitCode() int { return statusExitCode(e.status) }
func (e *statusError) Unwrap() error { return e.err }

func statusExitCode(s result.Status) int {
	switch s {
	case result.StatusOK:
		return 0
	case result.StatusBlocked:
		return 2
	default:
		return 1
	}
}

// StatusErr wraps err with a result.Status so the process exits with the corresponding code.
func StatusErr(status result.Status, err error) error {
	return &statusError{status: status, err: err}
}
