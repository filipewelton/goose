package httpdriver

import (
	"retail_workflow/internal/drivers/httpdriver/middlewares/apiauthentication"
	"retail_workflow/internal/drivers/httpdriver/routers"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery(), apiauthentication.SetMiddleware())

	// Routers
	routers.SetAdminRouter(r.Group("/"))
	routers.SetupUsersRouter(r.Group("/users"))

	return r
}
