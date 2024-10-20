package routes

import (
	"expenze-io.com/internal/config"
	"expenze-io.com/internal/controllers"
	"expenze-io.com/internal/services"
	"github.com/gin-gonic/gin"
)

func authRoute(router *gin.Engine) {
	userService := services.New(config.DB)
	authController := controllers.New(*userService)

	authRoutes := router.Group(BaseUrl + "/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/verify-otp", authController.VerifyOtp)
		authRoutes.POST("/login", authController.Login)
	}
}
