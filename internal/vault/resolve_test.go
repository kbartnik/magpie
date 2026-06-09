package vault

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	t.Run("MAGPIE_VAULT set → returns that path (even without .magpie/)", func(t *testing.T) {
		t.Setenv("MAGPIE_VAULT", "/tmp/notmagpie")
		vault, err := Resolve("/Users/kurt/", "")

		require.NoError(t, err)
		assert.Equal(t, "/tmp/notmagpie", vault, "expected /tmp/notmagpie: got %s", vault)
	})

	t.Run("MAGPIE_VAULT set → takes precedence over a vault found via cwd", func(t *testing.T) {
		vaultDir := t.TempDir()
		require.NoError(t, os.Mkdir(filepath.Join(vaultDir, ".magpie"), 0o755))
		t.Setenv("MAGPIE_VAULT", "/tmp/notmagpie")

		vault, resolvErr := Resolve(vaultDir, "")

		require.NoError(t, resolvErr)
		assert.Equal(t, "/tmp/notmagpie", vault, "expected `/tmp/notmagpie`: got %s", vault)
	})

	t.Run("cwd is vault root → returns cwd", func(t *testing.T) {
		vaultDir := t.TempDir()
		require.NoError(t, os.Mkdir(filepath.Join(vaultDir, ".magpie"), 0o755))

		vault, err := Resolve(vaultDir, "")

		require.NoError(t, err)
		assert.Equal(t, vaultDir, vault)
	})

	t.Run("cwd nested inside vault → walks up to innermost vault root", func(t *testing.T) {
		vaultDir := t.TempDir()
		require.NoError(t, os.Mkdir(filepath.Join(vaultDir, ".magpie"), 0o755))
		require.NoError(t, os.MkdirAll(filepath.Join(vaultDir, "notes", "subdirectory", ".magpie"), 0o755))

		vault, err := Resolve(filepath.Join(vaultDir, "notes", "subdirectory", "deeper"), "")

		require.NoError(t, err)
		assert.Equal(t, filepath.Join(vaultDir, "notes", "subdirectory"), vault)
	})

	t.Run("no vault found, no defaultVault → returns ErrNoVault", func(t *testing.T) {
		cwd := t.TempDir()

		_, err := Resolve(cwd, "")

		assert.ErrorIs(t, err, ErrNoVault)
	})

	t.Run("no vault found, defaultVault set → returns defaultVault", func(t *testing.T) {
		cwd := t.TempDir()

		vault, err := Resolve(cwd, "/default/vault")

		require.NoError(t, err)
		assert.Equal(t, "/default/vault", vault)
	})
}
