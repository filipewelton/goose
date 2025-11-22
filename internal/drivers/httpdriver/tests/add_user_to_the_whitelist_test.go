package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"retail_workflow/internal/drivers/httpdriver"
	"retail_workflow/internal/drivers/httpdriver/middlewares/adminuserauthentication"
	"retail_workflow/internal/persistence/repositories"
	"retail_workflow/internal/shared/apiguardian"
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/userguardian"
	"retail_workflow/tests/generators"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestAddUserToTheWhitelist(t *testing.T) {
	t.Run("when the status code is 204", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		router := httpdriver.SetupServer()
		userRepository := repositories.PostgresUserRepository{}

		adminuserauthentication.SetUserRepository(userRepository)

		user := generators.CreateUser(userRepository)
		userId := user.Id.Get()
		employeeId := faker.Number().Number(7)

		body := bytes.NewReader(fmt.Appendf(nil, `{
			"employeeId": "%s"
		}`, employeeId))

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users/whitelist", body)
		apiToken, _ := apiguardian.Generate()
		accessToken, _ := userguardian.GenerateAccessToken(userId)
		refreshToken, _ := userguardian.GenerateRefreshToken(userId)

		req.Header.Add("Authorization", fmt.Sprint("Bearer ", apiToken))

		req.AddCookie(&http.Cookie{
			Name:  adminuserauthentication.AccessTokenCookieName,
			Value: accessToken,
		})

		req.AddCookie(&http.Cookie{
			Name:  adminuserauthentication.RefreshTokenCookieName,
			Value: refreshToken,
		})

		router.ServeHTTP(rec, req)
		require.Equal(t, http.StatusNoContent, rec.Code)
	})
}
