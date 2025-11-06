package cache

import (
	"os"
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedis_Connect(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when connected", func(t *testing.T) {
		var redis Redis

		require.NoError(t, redis.Connect())

		defer redis.Disconnect()
	})

	t.Run("when username is missing", func(t *testing.T) {
		var redis Redis
		var username = os.Getenv("REDIS_USER")

		os.Setenv("REDIS_USER", "")

		err := redis.Connect()

		require.ErrorIs(
			t,
			exceptions.InternalExceptions.ErrRedisConnectionFailure,
			err,
		)

		t.Cleanup(func() {
			os.Setenv("REDIS_USER", username)
		})
	})
}
