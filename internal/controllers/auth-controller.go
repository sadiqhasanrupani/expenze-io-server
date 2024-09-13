// auth controller

package controllers

import (
	"net/http"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/services"
	"expenze-io.com/pkg"
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

	// Validate password
	if err := pkg.ValidatePassword(userReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Validate email
	if err := pkg.ValidateEmail(userReq.EmailID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Validate firstname
	if err := pkg.MinMaxValidation(pkg.MinMaxValidationFields{
		Min:        pkg.IntPtr(4),
		FieldName:  "firstname",
		FieldValue: userReq.Firstname,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Validate lastname
	if err := pkg.MinMaxValidation(pkg.MinMaxValidationFields{
		Min:        pkg.IntPtr(3),
		FieldName:  "lastname",
		FieldValue: userReq.Lastname,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Call service to register user
	if err := uc.UserService.RegisterUser(&userReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed", "error": err.Error()})
		return
	}

	responseMsg, err := uc.UserService.SendOtpMsg(&userReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get send the otp", "error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": responseMsg})
}

// Login handles user login requests
func (uc *AuthController) Login(c *gin.Context) {
	// Implement login logic
	c.JSON(http.StatusOK, gin.H{"message": "Login done successfully"})
}
