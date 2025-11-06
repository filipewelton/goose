package repositories

import (
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestRedisWhitelistRepository_Insert(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when inserted", func(t *testing.T) {
		var repository = RedisWhitelistRepository{}
		var cardNumber = faker.Number().Number(7)
		var err = repository.Insert(cardNumber)

		require.NoError(t, err)
	})
}

func TestRedisWhitelistRepository_Has(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when there is a key", func(t *testing.T) {
		var repository = RedisWhitelistRepository{}
		var cardNumber = faker.Number().Number(7)
		var err = repository.Insert(cardNumber)

		require.NoError(t, err)

		ok, err := repository.Has(cardNumber)

		require.True(t, ok)
		require.NoError(t, err)
	})
}

func TestRedisWhitelistRepository_Delete(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when deleted", func(t *testing.T) {
		var repository = RedisWhitelistRepository{}
		var cardNumber = faker.Number().Number(7)
		var err = repository.Insert(cardNumber)

		require.NoError(t, err)

		err = repository.Delete(cardNumber)

		require.NoError(t, err)

		ok, err := repository.Has(cardNumber)

		require.False(t, ok)
		require.ErrorIs(
			t,
			exceptions.NotFoundExceptions.ErrPreRegisterNotFound,
			err,
		)
	})
}
