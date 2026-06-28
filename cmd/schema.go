package cmd

import (
	"fmt"
	"strconv"

	"github.com/kbartnik/magpie/internal/result"
	"github.com/spf13/cobra"
)

// SchemaInfo describes the envelope contract returned by `magpie schema`.
type SchemaInfo struct {
	Version      int               `json:"version"`
	Fields       map[string]string `json:"fields"`
	StatusValues map[string]string `json:"status_values"`
}

var schemaCmd = &cobra.Command{
	Use:   "schema [version]",
	Short: "Print the response envelope schema",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil // vault-exempt
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			v, err := strconv.Atoi(args[0])
			if err != nil || v != SchemaVersion {
				env := WrapStatus(map[string]string{"error": fmt.Sprintf("unknown schema version %q", args[0])},
					result.StatusWarning)
				_ = Print(env)
				return fmt.Errorf("unknown schema version %q", args[0])
			}
		}
		return Print(Wrap(SchemaInfo{
			Version: SchemaVersion,
			Fields: map[string]string{
				"schema_version": "int - envelope contract version; call `magpie schema` when this changes",
				"status":         "string - outcome \"ok\", \"warning\", or \"blocked\"",
				"dry_run":        "bool - when true, no writes were performed; omitted when false",
				"data":           "object - command-specific payload; shape varies per command",
				"effects":        "[]string - files changed, e.g. \"created:path\", \"modified:path\"",
				"delta":          "null - vault health changes; populated once VaultGraph is available",
			},
			StatusValues: map[string]string{
				"ok":      "success - operation completed, caller may proceed",
				"warning": "user error or recoverable condition - caller may proceed with caution",
				"blocked": "hard block - destructive operation refused due to unmet precondition",
			},
		}))
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
