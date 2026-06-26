package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kbartnik/magpie/internal/vault"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// makeVault creates a temp vault with .magpie/config.yaml containing the given YAML.
// Returns the vault root path.
func makeVault(t *testing.T, localYAML string) string {
	t.Helper()
	root := t.TempDir()
	magpieDir := filepath.Join(root, ".magpie")
	require.NoError(t, os.MkdirAll(magpieDir, 0o755))
	if localYAML != "" {
		require.NoError(t, os.WriteFile(
			filepath.Join(magpieDir, "config.yaml"),
			[]byte(localYAML),
			0o644,
		))
	}
	return root
}

// makeGlobalConfig creates a temp global config file containing the given YAML.
// Sets XDG_CONFIG_HOME so ConfigPath() points to it.
func makeGlobalConfig(t *testing.T, globalYAML string) {
	t.Helper()
	xdgDir := t.TempDir()
	magpieDir := filepath.Join(xdgDir, "magpie")
	require.NoError(t, os.MkdirAll(magpieDir, 0o755))
	if globalYAML != "" {
		require.NoError(t, os.WriteFile(
			filepath.Join(magpieDir, "config.yaml"),
			[]byte(globalYAML),
			0o644,
		))
	}
	t.Setenv("XDG_CONFIG_HOME", xdgDir)
}

// execRoot runs PersistentPreRunE via a temporary subcommand that captures
// the resolved context. Returns the cobra.Command (with context set) and any error.
func execRoot(t *testing.T, args ...string) (*cobra.Command, error) {
	t.Helper()
	var capturedCmd *cobra.Command

	probe := &cobra.Command{
		Use: "probe",
		RunE: func(cmd *cobra.Command, args []string) error {
			capturedCmd = cmd
			return nil
		},
	}
	vaultFlag = ""
	rootCmd.AddCommand(probe)
	t.Cleanup(func() {
		rootCmd.RemoveCommand(probe)
		vaultFlag = ""
	})

	rootCmd.SetArgs(append([]string{"probe"}, args...))
	err := rootCmd.Execute()
	return capturedCmd, err
}

func TestRootIntegration(t *testing.T) {
	t.Run("vault CWD with merged config", func(t *testing.T) {
		vaultRoot := makeVault(t, `
inbox_path: vault-inbox
log_path: vault-log
`)
		makeGlobalConfig(t, `
inbox_path: global-inbox
archive_path: global-archive
`)
		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(vaultRoot))
		t.Cleanup(func() { require.NoError(t, os.Chdir(origDir)) })
		t.Setenv("MAGPIE_VAULT", "")

		cmd, err := execRoot(t)
		require.NoError(t, err)

		merged := MergedConfig(cmd)
		root := VaultRoot(cmd)

		wantRoot, err := filepath.EvalSymlinks(vaultRoot)
		require.NoError(t, err)
		assert.Equal(t, wantRoot, root)
		assert.Equal(t, "vault-inbox", merged.InboxPath)      // local overrides global
		assert.Equal(t, "vault-log", merged.LogPath)          // local only
		assert.Equal(t, "global-archive", merged.ArchivePath) // global fills gap
	})

	t.Run("--vault flag overrides CWD resolution", func(t *testing.T) {
		vaultRoot := makeVault(t, "")
		makeGlobalConfig(t, "")
		t.Setenv("MAGPIE_VAULT", "")

		// CWD is NOT inside a vault — resolution relies on the flag
		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(t.TempDir()))
		t.Cleanup(func() { require.NoError(t, os.Chdir(origDir)) })

		cmd, err := execRoot(t, "--vault", vaultRoot)
		require.NoError(t, err)

		root := VaultRoot(cmd)

		assert.Equal(t, vaultRoot, root)
	})

	t.Run("no vault and no config returns error", func(t *testing.T) {
		makeGlobalConfig(t, "")
		t.Setenv("MAGPIE_VAULT", "")

		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(t.TempDir()))
		t.Cleanup(func() { require.NoError(t, os.Chdir(origDir)) })

		_, err := execRoot(t)

		assert.NotNil(t, err)
		require.ErrorIs(t, err, vault.ErrNoVault)
	})
}
