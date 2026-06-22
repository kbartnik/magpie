package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Run("local value overrides global when set", func(t *testing.T) {
		global := &Config{DefaultVault: "/global/vault", InboxPath: "/global/inbox"}
		local := &Config{DefaultVault: "/local/vault"}

		got := Merge(global, local)

		assert.Equal(t, "/local/vault", got.DefaultVault)
		assert.Equal(t, "/global/inbox", got.InboxPath)
	})

	t.Run("zero local value does not override global", func(t *testing.T) {
		global := &Config{DefaultVault: "/global/vault", InboxPath: "/global/inbox"}
		local := &Config{DefaultVault: "/local/vault", InboxPath: ""}

		got := Merge(global, local)

		assert.Equal(t, "/local/vault", got.DefaultVault)
		assert.Equal(t, "/global/inbox", got.InboxPath)
	})

	t.Run("local Plugins entries win on collision, global entries preserved", func(t *testing.T) {
		global := &Config{
			Plugins: map[string]string{
				"archive": "builtin",
				"search":  "ripgrep",
			},
		}
		local := &Config{
			Plugins: map[string]string{
				"archive": "custom-archive",
				"export":  "pdf-tools",
			},
		}

		got := Merge(global, local)

		assert.Equal(t, "custom-archive", got.Plugins["archive"])
		assert.Equal(t, "ripgrep", got.Plugins["search"])
		assert.Equal(t, "pdf-tools", got.Plugins["export"])
	})

	t.Run("Merge(nil, nil) returns empty Config without panicking", func(t *testing.T) {
		got := Merge(nil, nil)

		assert.NotNil(t, got)
		assert.Equal(t, &Config{}, got)
	})
}
