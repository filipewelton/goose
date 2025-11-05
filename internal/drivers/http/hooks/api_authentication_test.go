package hooks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"retail_flow/internal/shared/lib"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestAuthenticateAPIAccess(t *testing.T) {
	lib.SetupEnvironmentVariables()

	t.Run("when is authenticated", func(t *testing.T) {
		var guardian lib.APIGuardian

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		handler := AuthenticateAPIAccess(next)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		guardian.GenerateToken(faker.RandomString(64))

		req.Header.Add(
			"Authorization",
			fmt.Sprint("Bearer ", guardian.GetToken()),
		)

		handler.ServeHTTP(rec, req)

		res := rec.Result()

		defer res.Body.Close()

		require.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("when API token is invalid", func(t *testing.T) {
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		handler := AuthenticateAPIAccess(next)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		req.Header.Add(
			"Authorization",
			fmt.Sprint("Bearer ", faker.RandomString(64)),
		)

		handler.ServeHTTP(rec, req)

		res := rec.Result()

		defer res.Body.Close()

		require.Equal(t, http.StatusUnauthorized, res.StatusCode)

		expectedMsg := "{\"message\":\"Token de acesso à API é inválido\"}\n"

		require.Equal(t, expectedMsg, rec.Body.String())
	})
}
