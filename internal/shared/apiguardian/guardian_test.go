package apiguardian

import (
	"retail_workflow/internal/shared/environment"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGuardian_Generate(t *testing.T) {
	t.Run("when generated", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		token, err := Generate()

		require.NoError(t, err)
		require.IsType(t, "", token)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestAPIGuardian_Validate(t *testing.T) {
	t.Run("when validated", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		token, err := Generate()

		require.NoError(t, err)

		claims, err := Validate(token)

		require.NotEmpty(t, claims)
		require.NoError(t, err)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}
