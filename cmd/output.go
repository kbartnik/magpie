package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

// Print marshals v to JSON and writes it to stdout.
func Print(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(os.Stdout, string(data))
	return err
}

// Err writes a human-readable diagnostic to stderr.
func Err(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}
