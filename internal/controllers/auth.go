package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationBody struct {
	Firstname string `binding:"required" json:"firstname" validate:"required"`
	Lastname  string `binding:"required" json:"lastname" validate:"required"`
	EmailID   string `binding:"required" json:"emailId" validate:"required"`
	Password  string `binding:"required" json:"password" validate:"required"`
}

func RegisterHandler(ctx *gin.Context) {
	var registrationDetail RegistrationBody

	// Bind JSON to struct
	err := ctx.ShouldBindJSON(&registrationDetail)
	if err != nil {

		// creating a validate from validator
		validate := validator.New()
		err := validate.Struct(registrationDetail)

		if err != nil {
			// checking exceptional check
			err, ok := err.(validator.ValidationErrors)

			if ok {
				// Handle the validation errors
				ctx.JSON(http.StatusBadRequest, gin.H{
					"errors":  err,
					"message": "Validation failed",
				})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registration done successfully",
		"data":    registrationDetail,
	})
}

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login done successfully"})
}
