package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration done successfully"})
}

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login done successfully"})
}
