package routers

import (
	"retail_flow/internal/drivers/http/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupUserRouter(r chi.Router) {
	r.Post("/", handlers.HandlerUserCreation)
}
