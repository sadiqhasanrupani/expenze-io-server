package routes

import (
	"net/http"

	"expenze-io.com/internal/controllers"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api/v1"

func RegisterRoutes(server *gin.Engine) {

	// root route
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Expenze member!!",
		})
	})

	// auth route
	server.POST(BaseUrl+"/auth/register", controllers.RegisterHandler)
	server.POST(BaseUrl+"/auth/login", controllers.LoginHandler)
}
