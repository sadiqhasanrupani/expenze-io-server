package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api/v1"

func RegisterRoutes(router *gin.Engine) {

	// Root route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Expenze member!!",
		})
	})

	// Add other routes as needed
	authRoute(router)
}
