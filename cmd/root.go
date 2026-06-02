package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var vaultFlag string

var rootCmd = &cobra.Command{
	Use:          "magpie",
	Short:        "Obsidian vault CLI",
	Args:         cobra.ArbitraryArgs,
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// vault resolution wired in Task 7
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			fmt.Fprintf(os.Stderr, "unknown subcommand %q (plugin dispatch not yet implemented)\n", args[0])
			return nil
		}
		return cmd.Help()
	},
}

// Execute is the entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&vaultFlag, "vault", "", "vault path (overrides MAGPIE_VAULT)")
}
