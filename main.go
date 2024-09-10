package main

import (
	"log"
	"net/http"

	"expenze-io.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

  routes.RegisterRoutes(server)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Expenze member!!",
		})
	})

	// server running on port 8080
	err := server.Run(":8080")

	if err != nil {
		log.Fatal("Failed to run server", err)
	}

}
