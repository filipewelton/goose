package user

import (
	"retail_flow/internal/application/dto"
	"retail_flow/internal/shared/lib"
	"retail_flow/tests/generators"
	"retail_flow/tests/repositories"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestUserCreation(t *testing.T) {
	t.Run("when created", func(t *testing.T) {
		lib.SetupEnvironmentVariables()

		userRepository := repositories.InMemoryUserRepository{}
		whitelistRepository := repositories.InMemoryWhitelistRepository{}
		cardNumber := faker.Number().Number(7)

		whitelistRepository.Insert(cardNumber)

		payload := dto.UserCreationDTO{
			CardNumber: cardNumber,
			Name:       faker.Name().Name(),
			Password:   generators.GeneratePassword(),
		}

		result, err := Create(UserCreationParams{
			UserRepository:      userRepository,
			WhitelistRepository: whitelistRepository,
			Payload:             payload,
		})

		require.NoError(t, err)
		require.NotEmpty(t, result.AccessToken)
		require.NotEmpty(t, result.RefreshToken)
		require.NotEmpty(t, result.User)
	})
}
