package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api/v1"

func RegisterRoutes(router *gin.Engine) {
	// root route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Expenze member!!",
		})
	})

	authRoute(router)
}
