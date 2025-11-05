package database

import (
	"os"
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostgres_Connect(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when connected", func(t *testing.T) {
		var postgres Postgres

		require.NoError(t, postgres.Connect())
		require.NotNil(t, postgres.Client)
	})

	t.Run("when database DSN is missing", func(t *testing.T) {
		var postgres Postgres
		var dsn = os.Getenv("DATABASE_DSN")

		os.Setenv("DATABASE_DSN", "")

		err := postgres.Connect()

		require.ErrorIs(
			t,
			exceptions.InternalExceptions.ErrPostgresConnectionFailure,
			err,
		)

		t.Cleanup(func() {
			os.Setenv("DATABASE_DSN", dsn)
		})
	})
}

func TestPostgres_Disconnect(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when disconnected", func(t *testing.T) {
		var postgres Postgres

		require.NoError(t, postgres.Connect())
		require.NoError(t, postgres.Disconnect())
		require.Nil(t, postgres.Client)
	})
}
