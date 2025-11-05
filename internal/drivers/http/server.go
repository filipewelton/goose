package http

import (
	"net/http"
	"retail_flow/internal/drivers/http/hooks"
	"retail_flow/internal/drivers/http/routers"

	"github.com/go-chi/chi/v5"
)

func SetupServer() http.Handler {
	r := chi.NewRouter()

	r.Use(hooks.AuthenticateAPIAccess)
	r.Route("/users", routers.SetupUserRouter)

	return r
}
