package frontmatter_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kbartnik/magpie/internal/frontmatter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	t.Run("file with frontmatter returns fields and body", func(t *testing.T) {
		path := writeFile(t, "---\ntitle: Hello\ncount: 3\n---\n\n# Body\n")

		fields, body, err := frontmatter.Read(path)
		require.NoError(t, err)
		assert.Equal(t, "Hello", fields["title"])
		assert.Equal(t, 3, fields["count"])
		assert.Equal(t, "\n# Body \n", string(body))
	})

	t.Run("file without frontmatter returns empty fields and full content as body", func(t *testing.T) {
		path := writeFile(t, "")

		fields, body, err := frontmatter.Read(path)
		require.NoError(t, err)
		assert.Empty(t, fields)
		assert.Empty(t, body)
	})

	t.Run("malformed YAML returns error", func(t *testing.T) {
		path := writeFile(t, "--\nkey: [unclosed\n---\n")

		_, _, err := frontmatter.Read(path)
		assert.Error(t, err)
	})
}

func TestWrite(t *testing.T) {
	t.Run("creates new file with frontmatter and body", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "note.md")

		err := frontmatter.Write(path, map[string]any{"title": "Hello"}, []byte("\n# Body\n"))
		require.NoError(t, err)

		fields, body, err := frontmatter.Read(path)
		require.NoError(t, err)

		assert.Equal(t, "Hello", fields["title"])
		assert.Equal(t, "\n# Body\n", string(body))
	})

	t.Run("overwrites existing file", func(t *testing.T) {
		path := writeFile(t, "---\ntitle: Old\n---\n\nold body\n")

		err := frontmatter.Write(path, map[string]any{"title": "New"}, []byte("\nnew body\n"))
		require.NoError(t, err)

		fields, body, err := frontmatter.Read(path)
		require.NoError(t, err)

		assert.Equal(t, "New", fields["title"])
		assert.Equal(t, "\nnew body\n", string(body))
	})

	t.Run("round-trip preserves fields and body", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "note.md")
		fields := map[string]any{"title": "Test", "count": 42}
		body := []byte("\n# Content\n")

		require.NoError(t, frontmatter.Write(path, fields, body))

		got, gotBody, err := frontmatter.Read(path)
		require.NoError(t, err)

		assert.Equal(t, "Test", got["title"])
		assert.Equal(t, 42, got["count"])
		assert.Equal(t, body, gotBody)
	})
}

func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "note.md")
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	return path
}
