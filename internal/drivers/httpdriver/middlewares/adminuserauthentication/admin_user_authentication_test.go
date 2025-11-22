package adminuserauthentication

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"retail_workflow/internal/shared/userguardian"
	"retail_workflow/tests/generators"
	"retail_workflow/tests/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func TestAuthenticateUser(t *testing.T) {
	t.Run("when authenticated", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)

		r := gin.New()
		userRepository := mocks.InMemoryUserRepository{}
		user := generators.CreateUser(userRepository)

		SetUserRepository(userRepository)
		r.Use(SetMiddleware())

		r.GET("/", func(ctx *gin.Context) {
			ctx.Status(http.StatusNoContent)
		})

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		accessToken, _ := userguardian.GenerateAccessToken(user.Id.Get())
		refreshToken, _ := userguardian.GenerateRefreshToken(user.Id.Get())

		req.AddCookie(&http.Cookie{
			Name:  AccessTokenCookieName,
			Value: accessToken,
		})

		req.AddCookie(&http.Cookie{
			Name:  RefreshTokenCookieName,
			Value: refreshToken,
		})

		r.ServeHTTP(rec, req)

		require.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("when the user does not exists", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)

		r := gin.New()
		userRepository := mocks.InMemoryUserRepository{}

		SetUserRepository(userRepository)
		r.Use(SetMiddleware())

		r.GET("/", func(ctx *gin.Context) {
			ctx.Status(http.StatusNoContent)
		})

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		userId := ulid.Make().String()
		accessToken, _ := userguardian.GenerateAccessToken(userId)
		refreshToken, _ := userguardian.GenerateRefreshToken(userId)

		req.AddCookie(&http.Cookie{
			Name:  AccessTokenCookieName,
			Value: accessToken,
		})

		req.AddCookie(&http.Cookie{
			Name:  RefreshTokenCookieName,
			Value: refreshToken,
		})

		r.ServeHTTP(rec, req)

		require.Equal(t, http.StatusNotFound, rec.Code)

		var body map[string]any

		json.NewDecoder(rec.Body).Decode(&body)
		require.Equal(t, "Usuário não encontrado", body["message"])
	})
}
