// Package cmd implements the magpie command-line interface.
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/kbartnik/magpie/internal/config"
	"github.com/kbartnik/magpie/internal/result"
	"github.com/kbartnik/magpie/internal/vault"
	"github.com/spf13/cobra"
)

type contextKey string

const (
	vaultRootKey contextKey = "vaultRoot"
	mergedCfgKey contextKey = "mergedCfg"
)

var vaultFlag string

var rootCmd = &cobra.Command{
	Use:               "magpie",
	Short:             "Obsidian vault CLI",
	Args:              cobra.ArbitraryArgs,
	SilenceUsage:      true,
	SilenceErrors:     true,
	PersistentPreRunE: initVaultContext,
	RunE:              runRoot,
}

func VaultRoot(cmd *cobra.Command) string {
	return cmd.Context().Value(vaultRootKey).(string)
}

func MergedConfig(cmd *cobra.Command) *config.Config {
	return cmd.Context().Value(mergedCfgKey).(*config.Config)
}

// Execute is the entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		code := 1
		if ec, ok := err.(ExitCoder); ok {
			code = ec.ExitCode()
		}
		os.Exit(code)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&vaultFlag, "vault", "", "vault path (overrides MAGPIE_VAULT)")
}

func initVaultContext(cmd *cobra.Command, _ []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}

	vaultRoot, err := vault.Resolve(cwd, vaultFlag)
	if err != nil {
		return fmt.Errorf("resolving vault directory: %w", err)
	}

	globalCfg, err := config.Load(config.ConfigPath())
	if err != nil {
		return fmt.Errorf("loading global configuration: %w", err)
	}

	localCfg, err := config.Load(filepath.Join(vaultRoot, ".magpie", "config.yaml"))
	if err != nil {
		return fmt.Errorf("loading local configuration: %w", err)
	}

	merged := config.Merge(globalCfg, localCfg)

	ctx := context.WithValue(cmd.Context(), vaultRootKey, vaultRoot)
	ctx = context.WithValue(ctx, mergedCfgKey, merged)
	cmd.SetContext(ctx)

	return nil
}

func resolvePlugin(subcommand string) (string, error) {
	return exec.LookPath("magpie-" + subcommand)
}

func runRoot(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		path, err := resolvePlugin(args[0])
		if err != nil {
			env := WrapStatus(map[string]string{"error": "plugin not found: magpie-" + args[0]}, result.StatusWarning)
			_ = Print(env)
			return StatusErr(result.StatusWarning, fmt.Errorf("plugin not found: magpie-%s", args[0]))
		}

		envv := append(os.Environ(), "MAGPIE_VAULT="+VaultRoot(cmd))
		if err := syscall.Exec(path, append([]string{"magpie-" + args[0]}, args[1:]...), envv); err != nil {
			return fmt.Errorf("exec %s: %w", path, err)
		}
		return nil // unreachable on success
	}
	return cmd.Help()
}
