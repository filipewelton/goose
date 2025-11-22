package logger

import (
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/typings"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestLogger_Info(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		Info(faker.RandomString(100))

		cmd := client.Keys(ctx, "*")

		require.NotEmpty(t, cmd.Val())
	})
}

func TestLogger_Error(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
		errorResult := typings.ErrorResult{
			Code:    faker.Number().NumberInt(3),
			Context: faker.RandomString(10),
			Reason:  faker.RandomString(10),
		}

		Error(errorResult)

		cmd := client.Keys(ctx, "*")

		require.NotEmpty(t, cmd.Val())
	})
}
