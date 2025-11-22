package repositories

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/shared/environment"
	"retail_workflow/tests/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestPostgresUserRepository_Insert(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var entity user.UserEntity
		var err = entity.New(user.UserCreationDTO{
			Name:       faker.Name().Name(),
			EmployeeId: faker.Number().Number(7),
			Password:   utils.GeneratePassword(),
		})

		require.NoError(t, err)

		var repository = PostgresUserRepository{}
		err = repository.Insert(entity)

		require.NoError(t, err)

		t.Cleanup(func() {
			repository.DeleteByEmployeeId(entity.EmployeeId)
			environment.UnloadEnvironmentVariables()
		})
	})
}

func TestPostgresUserRepository_FindByEmployeeId(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var entity user.UserEntity
		var err = entity.New(user.UserCreationDTO{
			Name:       faker.Name().Name(),
			EmployeeId: faker.Number().Number(7),
			Password:   utils.GeneratePassword(),
		})

		require.NoError(t, err)

		var repository = PostgresUserRepository{}
		err = repository.Insert(entity)

		require.NoError(t, err)

		user, err := repository.FindByEmployeeId(entity.EmployeeId)

		require.NotEmpty(t, user)
		require.NoError(t, err)

		t.Cleanup(func() {
			repository.DeleteByEmployeeId(entity.EmployeeId)
			environment.UnloadEnvironmentVariables()
		})
	})
}
