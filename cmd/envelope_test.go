package cmd

import (
	"encoding/json"
	"testing"

	"github.com/kbartnik/magpie/internal/result"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testData is a representative T for envelope tests.
type testData struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func TestWrap(t *testing.T) {
	t.Run("sets SchemaVersion and StatusOK on data", func(t *testing.T) {
		envelope := Wrap(testData{})

		assert.Equal(t, SchemaVersion, envelope.SchemaVersion)
		assert.Equal(t, result.StatusOK, envelope.Status)
	})

	t.Run("zero-value DryRun Effects Delta omitted from JSON", func(t *testing.T) {
		envelope := Wrap(testData{Name: "test", Count: 1})

		data, err := json.Marshal(envelope)
		require.NoError(t, err)

		var raw map[string]any
		err = json.Unmarshal(data, &raw)
		require.NoError(t, err)

		assert.Contains(t, raw, "schema_version")
		assert.Contains(t, raw, "status")
		assert.Contains(t, raw, "data")

		assert.NotContains(t, raw, "dry_run")
		assert.NotContains(t, raw, "effects")
		assert.NotContains(t, raw, "delta")
	})

	t.Run("all Status values serialize correctly", func(t *testing.T) {
		for _, status := range []result.Status{result.StatusOK, result.StatusWarning, result.StatusBlocked} {
			envelope := Envelope[testData]{Status: status}

			data, err := json.Marshal(envelope)
			require.NoError(t, err)

			var raw map[string]any
			err = json.Unmarshal(data, &raw)
			require.NoError(t, err)

			assert.Equal(t, string(status), raw["status"])
		}
	})

	t.Run("JSON round-trip preserves all fields", func(t *testing.T) {
		original := Envelope[testData]{
			SchemaVersion: 0,
			Status:        result.StatusOK,
			DryRun:        true,
			Data:          testData{Name: "test", Count: 5},
			Effects:       []string{"created:foo.md"},
		}

		data, err := json.Marshal(original)
		require.NoError(t, err)

		var got Envelope[testData]
		err = json.Unmarshal(data, &got)
		require.NoError(t, err)

		assert.Equal(t, original, got)
	})
}
