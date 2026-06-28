package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/kbartnik/magpie/internal/result"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func captureStdout(t *testing.T) func() *bytes.Buffer {
	t.Helper()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	orig := os.Stdout
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = orig })
	return func() *bytes.Buffer {
		w.Close()
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r)
		require.NoError(t, err)
		return &buf
	}
}

func TestSchemaCommand(t *testing.T) {
	t.Run("no args returns current schema envelope", func(t *testing.T) {
		collect := captureStdout(t)

		rootCmd.SetArgs([]string{"schema"})
		require.NoError(t, rootCmd.Execute())

		var env Envelope[SchemaInfo]
		require.NoError(t, json.Unmarshal(collect().Bytes(), &env))

		assert.Equal(t, SchemaVersion, env.SchemaVersion)
		assert.Equal(t, result.StatusOK, env.Status)
		assert.NotEmpty(t, env.Data.Fields)
	})

	t.Run("unknown schema version returns warning", func(t *testing.T) {
		collect := captureStdout(t)

		rootCmd.SetArgs([]string{"schema", "99"})
		assert.Error(t, rootCmd.Execute())

		var env Envelope[map[string]string]
		require.NoError(t, json.Unmarshal(collect().Bytes(), &env))

		assert.Equal(t, SchemaVersion, env.SchemaVersion)
		assert.Equal(t, result.StatusWarning, env.Status)
	})
}
