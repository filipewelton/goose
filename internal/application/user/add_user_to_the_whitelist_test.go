package user

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/shared/environment"
	"retail_workflow/tests/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestAddUserToTheWhitelist(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		var repository = mocks.InMemoryWhitelistRepository{}
		var err = AddUserToTheWhitelist(WhitelistInclusion{
			WhitelistRepository: repository,
			Payload: user.WhitelistInclusionDTO{
				EmployeeId: faker.Number().Number(7),
			},
		})

		require.NoError(t, err)
	})
}
