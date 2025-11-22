package utils

import (
	"encoding/json"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"

	"github.com/gin-gonic/gin"
)

func ParseRequestBody[T any](ctx *gin.Context) (T, error) {
	var t T
	var err = json.NewDecoder(ctx.Request.Body).Decode(&t)

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    400,
			Context: "Request body reading",
			Reason:  err.Error(),
		})

		return t, errors.Err400.ErrFailedToReadTheRequestBody
	}

	return t, nil
}
