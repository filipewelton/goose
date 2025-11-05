package lib

import (
	"os"
	"retail_flow/internal/domain/constants"
	"retail_flow/internal/shared/exceptions"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestUserGuardian_handleTokeGeneration(t *testing.T) {
	t.Run("when generated", func(t *testing.T) {
		SetupEnvironmentVariables()

		userGuardian := UserGuardian{}
		userID := faker.RandomString(10)
		scopes := []string{constants.CALCULATOR, constants.EXPIRATION_DATE}
		expiration := faker.Time().Forward(time.Hour * 24 * 7).Unix()
		token, err := userGuardian.handleTokeGeneration(userID, scopes, expiration)

		require.NotEmpty(t, token)
		require.NoError(t, err)
	})

	t.Run("when secret is empty", func(t *testing.T) {
		os.Setenv("JWT_SECRET", "")

		userGuardian := UserGuardian{}
		userID := faker.RandomString(10)
		scopes := []string{constants.CALCULATOR, constants.EXPIRATION_DATE}
		expiration := faker.Time().Forward(time.Hour * 24 * 7).Unix()
		token, err := userGuardian.handleTokeGeneration(userID, scopes, expiration)

		require.Empty(t, token)

		require.ErrorIs(
			t,
			exceptions.InternalExceptions.ErrEnvironmentVariableIsMissing,
			err,
		)
	})
}
