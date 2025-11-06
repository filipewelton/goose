package handlers

import (
	"net/http"
	"retail_flow/internal/application/dto"
	"retail_flow/internal/application/usecases/user"
	"retail_flow/internal/drivers/http/lib"
	"retail_flow/internal/persistence/repositories"

	"github.com/go-chi/render"
)

var userRepository = repositories.PostgresUserRepository{}
var whitelistRepository = repositories.RedisWhitelistRepository{}

func HandlerUserCreation(w http.ResponseWriter, r *http.Request) {
	payload, err := lib.ParseRequestBody[dto.UserCreationDTO](r)

	if err != nil {
		lib.RenderErrorResponse(err, w, r)
		return
	}

	result, err := user.Create(user.UserCreationParams{
		UserRepository:      userRepository,
		WhitelistRepository: whitelistRepository,
		Payload:             payload,
	})

	if err != nil {
		lib.RenderErrorResponse(err, w, r)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, result)
}
