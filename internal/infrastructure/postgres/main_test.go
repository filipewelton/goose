package postgres

import (
	"retail_workflow/internal/shared/environment"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostgres_Connect(t *testing.T) {
	t.Run("when connected", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		require.NoError(t, Connect())
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestPostgres_Disconnect(t *testing.T) {
	t.Run("when disconnected", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		require.NoError(t, Connect())
		require.NotPanics(t, Disconnect)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}
