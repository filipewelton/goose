package routers

import (
	"retail_workflow/internal/drivers/httpdriver/controllers"
	"retail_workflow/internal/drivers/httpdriver/middlewares/adminuserauthentication"

	"github.com/gin-gonic/gin"
)

func SetAdminRouter(r *gin.RouterGroup) {
	r.Use(adminuserauthentication.SetMiddleware())
	r.POST("/users/whitelist", controllers.AddUserToTheWhitelist)
}
