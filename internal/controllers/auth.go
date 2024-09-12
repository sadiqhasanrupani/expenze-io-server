package controllers

import (
	"net/http"

	"expenze-io.com/internal/validators"
	"expenze-io.com/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RegisterHandler(r *gin.Context) {
	var userReq validators.RegistrationBody

	// Bind and validate incoming request
	err := r.ShouldBindJSON(&userReq)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "body data is incomplete", "error": err.Error()})
		return
	}

	// password validation
	validatePassErr := pkg.ValidatePassword(userReq.Password)

	if validatePassErr != nil {
		r.JSON(http.StatusBadGateway, gin.H{
			"error": validatePassErr,
		})

		return
	}

	// email validation
	if validateEmailErr := pkg.ValidateEmail(userReq.EmailID); validateEmailErr != nil {
		r.JSON(http.StatusBadGateway, gin.H{
			"error": validateEmailErr,
		})

		return
	}

	// min firstname
	validateFirstname := pkg.MinMaxValidation(pkg.MinMaxValidationFields{
		Min:        pkg.IntPtr(4),
		FieldName:  "firstname",
		FieldValue: userReq.Firstname,
	})

	if validateFirstname != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": validateFirstname})
	}

	// min firstname
	validateLastName := pkg.MinMaxValidation(pkg.MinMaxValidationFields{
		Min:        pkg.IntPtr(3),
		FieldName:  "lastname",
		FieldValue: userReq.Lastname,
	})

	if validateLastName != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": validateLastName})
	}

	// Call service to register user
	// if err := userService.RegisterUser(&userReq); err != nil {
	// 	r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	r.JSON(http.StatusOK, gin.H{
		"message": "Registration done successfully",
	})
}

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login done successfully"})
}
