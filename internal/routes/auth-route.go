package routes

import (
	"expenze-io.com/internal/config"
	AuthController "expenze-io.com/internal/controllers"
	UserService "expenze-io.com/internal/services"
	"github.com/gin-gonic/gin"
)

func authRoute(router *gin.Engine) {
	userService := UserService.NewUserService(config.DB)
	authController := AuthController.NewAuthController(*userService)

	authRoutes := router.Group(BaseUrl + "/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/verify-otp", authController.VerifyOtp)
		authRoutes.POST("/login", authController.Login)
	}
}
