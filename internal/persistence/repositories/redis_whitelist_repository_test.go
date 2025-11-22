package repositories

import (
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestRedisWhitelistRepository_Insert(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var repository = RedisWhitelistRepository{}
		var err = repository.Insert(faker.Number().Number(7))

		require.NoError(t, err)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestRedisWhitelistRepository_Has(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var repository = RedisWhitelistRepository{}
		var employeeId = faker.Number().Number(7)
		var err = repository.Insert(employeeId)

		require.NoError(t, err)

		has, err := repository.Has(employeeId)

		require.True(t, has)
		require.NoError(t, err)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}

func TestRedisWhitelistRepository_Delete(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var repository = RedisWhitelistRepository{}
		var employeeId = faker.Number().Number(7)
		var err = repository.Insert(employeeId)

		require.NoError(t, err)

		err = repository.Delete(employeeId)

		require.NoError(t, err)

		has, err := repository.Has(employeeId)

		require.False(t, has)
		require.ErrorIs(
			t,
			errors.Err404.ErrEmployeeIdNotFound,
			err,
		)
		t.Cleanup(environment.UnloadEnvironmentVariables)
	})
}
