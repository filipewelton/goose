package controllers

import (
	"net/http"
	application "retail_workflow/internal/application/user"
	domain "retail_workflow/internal/domain/user"
	"retail_workflow/internal/drivers/httpdriver/utils"
	"retail_workflow/internal/persistence/repositories"

	"github.com/gin-gonic/gin"
)

var whitelistRepository = repositories.RedisWhitelistRepository{}

func AddUserToTheWhitelist(ctx *gin.Context) {
	payload, err := utils.ParseRequestBody[domain.WhitelistInclusionDTO](ctx)

	if err != nil {
		utils.Render(err, ctx)
		return
	}

	err = application.AddUserToTheWhitelist(application.WhitelistInclusion{
		WhitelistRepository: whitelistRepository,
		Payload:             payload,
	})

	if err != nil {
		utils.Render(err, ctx)
		return
	}

	ctx.Status(http.StatusNoContent)
}
