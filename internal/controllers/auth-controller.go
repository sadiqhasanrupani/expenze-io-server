// auth controller

package controllers

import (
	"net/http"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/services"
	"expenze-io.com/internal/validations"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

// Register handles user registration requests
func (uc *AuthController) Register(c *gin.Context) {
	var userReq body.RegistrationBody

	// Bind and validate incoming request
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Body data is incomplete", "error": err.Error()})
		return
	}

	// validation
	registerValdiation := validations.New(&userReq)
	if err := registerValdiation.ValidateRegistration(); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error,
			"error":   err,
		})
		return
	}

	// Call service to register user
	userId, err := uc.UserService.RegisterUser(&userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed", "error": err.Error()})
		return
	}

	responseMsg, err := uc.UserService.SendOtpMsg(&userReq, *userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get send the otp", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": responseMsg})
}

// Login handles user login requests
func (uc *AuthController) Login(c *gin.Context) {
	// Implement login logic
	c.JSON(http.StatusOK, gin.H{"message": "Login done successfully"})
}
