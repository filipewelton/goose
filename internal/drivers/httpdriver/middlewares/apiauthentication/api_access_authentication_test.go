package apiauthentication

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"retail_workflow/internal/shared/apiguardian"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthenticateAPIAccess(t *testing.T) {
	t.Run("when authenticated", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)

		r := gin.New()

		r.Use(SetMiddleware())

		r.GET("/", func(ctx *gin.Context) {
			ctx.Status(http.StatusNoContent)
		})

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		token, _ := apiguardian.Generate()

		req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
		r.ServeHTTP(rec, req)

		require.Equal(t, http.StatusNoContent, rec.Code)
	})
}
