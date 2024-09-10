package routes

import (
	"expenze-io.com/controllers"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api/v1"

func RegisterRoutes(server *gin.Engine) {
	// auth route

	server.POST(BaseUrl+"/auth/register", controllers.RegisterHandler)
	server.POST(BaseUrl+"/auth/login", controllers.LoginHandler)
}
