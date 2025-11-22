package redis

import (
	"retail_workflow/internal/shared/environment"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedis_Connect(t *testing.T) {
	t.Run("when connected", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		require.NoError(t, Connect())
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}
