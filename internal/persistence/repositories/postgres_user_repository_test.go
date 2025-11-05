package repositories

import (
	"retail_flow/internal/application/dto"
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/shared/lib"
	"retail_flow/tests/generators"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestPostgresUserRepository_Insert(t *testing.T) {
	var repository = PostgresUserRepository{}

	lib.SetupEnvironmentVariables()

	t.Run("when inserted", func(t *testing.T) {
		var user entities.UserEntity

		err := user.New(dto.UserCreationDTO{
			CardNumber: faker.Number().Number(7),
			Name:       faker.Name().Name(),
			Password:   generators.GeneratePassword(),
		})

		require.NoError(t, err)

		err = repository.Insert(user)

		require.NoError(t, err)
	})
}
