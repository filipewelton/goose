package apiauthentication

import (
	"retail_workflow/internal/drivers/httpdriver/utils"
	"retail_workflow/internal/shared/apiguardian"
	"retail_workflow/internal/shared/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		if authorization == "" {
			utils.Render(errors.Error401.ErrUnauthenticatedAPIAccess, ctx)
			return
		}

		token := strings.ReplaceAll(authorization, "Bearer ", "")
		_, err := apiguardian.Validate(token)

		if err != nil {
			utils.Render(errors.Error401.ErrUnauthenticatedAPIAccess, ctx)
			return
		}

		ctx.Next()
	}
}
