package controllers

import (
	"net/http"

	"expenze-io.com/lib"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationBody struct {
	Firstname string `binding:"required" json:"firstname"`
	Lastname  string `binding:"required" json:"lastname"`
	EmailID   string `binding:"required" json:"emailId"`
	Password  string `binding:"required" json:"password"`
}

var validate = validator.New()

func RegisterHandler(r *gin.Context) {
	var userReq RegistrationBody

	// Bind and validate incoming request
	err := r.ShouldBindJSON(&userReq)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "body data is incomplete", "error": err.Error()})
		return
	}

	// password validation
	validatePassErr := lib.ValidatePassword(userReq.Password)

	if validatePassErr != nil {
		r.JSON(http.StatusBadGateway, gin.H{
			"error": validatePassErr,
		})

		return
	}

	// email validation
	if validateEmailErr := lib.ValidateEmail(userReq.EmailID); validateEmailErr != nil {
		r.JSON(http.StatusBadGateway, gin.H{
			"error": validateEmailErr,
		})

		return
	}

	// min firstname
	validateFirstname := lib.MinMaxValidation(lib.MinMaxValidationFields{
		Min:        lib.IntPtr(4),
		FieldName:  "firstname",
		FieldValue: userReq.Firstname,
	})

	if validateFirstname != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": validateFirstname})
	}

	// min firstname
	validateLastName := lib.MinMaxValidation(lib.MinMaxValidationFields{
		Min:        lib.IntPtr(3),
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
