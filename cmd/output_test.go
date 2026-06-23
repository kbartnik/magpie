package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	t.Run("produces valid JSON on stdout", func(t *testing.T) {
		r, w, err := os.Pipe()
		require.NoError(t, err)

		origStdout := os.Stdout
		os.Stdout = w

		input := map[string]string{"ok": "true"}
		printErr := Print(input)

		w.Close()
		os.Stdout = origStdout

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		require.NoError(t, err)

		assert.NoError(t, printErr)
		assert.True(t, json.Valid(bytes.TrimSpace(buf.Bytes())))

		var got map[string]string
		err = json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &got)
		require.NoError(t, err)
		assert.Equal(t, input, got)
	})

	t.Run("output ends with newline", func(t *testing.T) {
		r, w, err := os.Pipe()
		require.NoError(t, err)

		origStdout := os.Stdout
		os.Stdout = w

		_ = Print("hello")

		w.Close()
		os.Stdout = origStdout

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		require.NoError(t, err)

		assert.True(t, buf.Len() > 0)
		assert.Equal(t, byte('\n'), buf.Bytes()[buf.Len()-1])
	})

	t.Run("returns error for unmarshalable input", func(t *testing.T) {
		err := Print(make(chan int))
		assert.Error(t, err)
	})
}

func TestErr(t *testing.T) {
	t.Run("writes message to stderr", func(t *testing.T) {
		r, w, err := os.Pipe()
		require.NoError(t, err)

		origStderr := os.Stderr
		os.Stderr = w

		Err("something went wrong")

		w.Close()
		os.Stderr = origStderr

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		require.NoError(t, err)

		assert.Equal(t, "something went wrong\n", buf.String())
	})
}
