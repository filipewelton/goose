package create

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCommand(t *testing.T) {
	t.Run("when created", func(t *testing.T) {
		var filename = "test#1_goose.json"

		os.Setenv("GOOSE_CONFIG_FILE", filename)

		require.NotPanics(t, func() {
			Create.Execute()
		})

		stat, _ := os.Stat(filename)

		require.NotNil(t, stat)

		t.Cleanup(func() {
			os.Remove(filename)
		})
	})

	t.Run("when the configuration file already exists", func(t *testing.T) {
		var filename = "test#2_goose.json"

		os.Setenv("GOOSE_CONFIG_FILE", filename)

		require.NotPanics(t, func() {
			Create.Execute()
		})

		stat, _ := os.Stat(filename)

		require.NotNil(t, stat)

		require.Panics(t, func() {
			Create.Execute()
		})

		t.Cleanup(func() {
			os.Remove(filename)
		})
	})
}
