package lib

import (
	"os"
	"retail_flow/internal/shared/exceptions"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestAPIGuardian_GenerateToken(t *testing.T) {
	t.Run("when generated", func(t *testing.T) {
		SetupEnvironmentVariables()

		apiGuardian := APIGuardian{}
		hash := faker.RandomString(64)
		err := apiGuardian.GenerateToken(hash)

		require.NoError(t, err)
		require.NotEmpty(t, apiGuardian.token)
	})

	t.Run("when secret is empty", func(t *testing.T) {
		os.Setenv("JWT_SECRET", "")

		apiGuardian := APIGuardian{}
		hash := faker.RandomString(64)
		err := apiGuardian.GenerateToken(hash)

		require.ErrorIs(
			t,
			exceptions.InternalExceptions.ErrEnvironmentVariableIsMissing,
			err,
		)

		t.Cleanup(func() {
			os.Setenv("JWT_SECRET", faker.RandomString(64))
		})
	})
}

func TestAPIGuardian_GetToken(t *testing.T) {
	t.Run("when retrieved", func(t *testing.T) {
		SetupEnvironmentVariables()

		apiGuardian := APIGuardian{}
		hash := faker.RandomString(64)
		err := apiGuardian.GenerateToken(hash)

		require.NoError(t, err)
		require.NotEmpty(t, apiGuardian.token)

		token := apiGuardian.GetToken()

		require.NotEmpty(t, token)
	})

	t.Run("when the token was not generated", func(t *testing.T) {
		SetupEnvironmentVariables()

		apiGuardian := APIGuardian{}
		token := apiGuardian.GetToken()

		require.Empty(t, token)
	})
}

func TestAPIGuardian_Validate(t *testing.T) {
	t.Run("when token is valid", func(t *testing.T) {
		SetupEnvironmentVariables()

		apiGuardian := APIGuardian{}
		hash := faker.RandomString(64)
		err := apiGuardian.GenerateToken(hash)

		require.NoError(t, err)
		require.NotEmpty(t, apiGuardian.token)

		token := apiGuardian.GetToken()
		err = apiGuardian.Validate(token)

		require.NoError(t, err)
	})

	t.Run("when token is invalid", func(t *testing.T) {
		SetupEnvironmentVariables()

		apiGuardian := APIGuardian{}
		token := faker.RandomString(512)
		err := apiGuardian.Validate(token)

		require.ErrorContains(
			t,
			err,
			exceptions.UnauthenticatedExceptions.ErrInvalidAPIToken.Error(),
		)
	})
}
