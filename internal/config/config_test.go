package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testdata returns the absolute path to a file under testdata/.
// Using runtime.Caller keeps the path correct regardless of where
// `go test` is invoked from.
func testdata(name string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "testdata", name)
}

func TestLoad(t *testing.T) {
	t.Run("valid YAML file returns correct fields", func(t *testing.T) {
		cfg, err := Load(testdata("valid.yaml"))
		require.NoError(t, err)
		require.NotNil(t, cfg)

		assert.Equal(t, "/home/user/notes", cfg.DefaultVault)
		assert.Equal(t, map[string]string{"archive": "builtin"}, cfg.Plugins)
		assert.Equal(t, "/home/user/notes/inbox", cfg.InboxPath)
		assert.Equal(t, "/home/user/notes/archive", cfg.ArchivePath)
		assert.Equal(t, "/home/user/notes/log.md", cfg.LogPath)
	})

	t.Run("missing file returns empty Config, no error", func(t *testing.T) {
		cfg, err := Load(testdata("missing.yaml"))
		require.NoError(t, err)
		require.NotNil(t, cfg)

		assert.Equal(t, "", cfg.InboxPath)
		assert.Equal(t, "", cfg.ArchivePath)
		assert.Equal(t, "", cfg.LogPath)
	})

	t.Run("malformed YAML returns an error", func(t *testing.T) {
		cfg, err := Load(testdata("invalid.yaml"))

		assert.Error(t, err)
		assert.Nil(t, cfg)
	})
}

func TestConfigPath(t *testing.T) {
	t.Run("XDG_CONFIG_HOME unset → ~/.config/magpie/config.yaml", func(t *testing.T) {
		t.Setenv("XDG_CONFIG_HOME", "")

		expected, err := os.UserConfigDir()
		require.NoError(t, err)
		require.NotEmpty(t, expected)

		assert.Equal(t, filepath.Join(expected, "magpie", "config.yaml"), ConfigPath())
	})

	t.Run("XDG_CONFIG_HOME set → $XDG_CONFIG_HOME/magpie/config.yaml", func(t *testing.T) {
		t.Setenv("XDG_CONFIG_HOME", "/tmp/xdg")

		assert.Equal(t, "/tmp/xdg/magpie/config.yaml", ConfigPath())
	})
}
