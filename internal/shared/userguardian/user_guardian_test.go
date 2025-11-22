package userguardian

import (
	"retail_workflow/internal/shared/environment"
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func TestUserGuardian_GenerateAccessToken(t *testing.T) {
	t.Run("when generated", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		userId := ulid.Make().String()
		token, err := GenerateAccessToken(userId)

		require.NoError(t, err)
		require.IsType(t, "", token)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestUserGuardian_GenerateRefreshToken(t *testing.T) {
	t.Run("when generated", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		userId := ulid.Make().String()
		token, err := GenerateRefreshToken(userId)

		require.NoError(t, err)
		require.IsType(t, "", token)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestUserGuardian_Validate(t *testing.T) {
	t.Run("when validated", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		userId := ulid.Make().String()
		accessToken, _ := GenerateAccessToken(userId)
		refreshToken, _ := GenerateRefreshToken(userId)

		_, err1 := Validate(accessToken)
		_, err2 := Validate(refreshToken)

		require.NoError(t, err1)
		require.NoError(t, err2)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}
