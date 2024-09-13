package main

import (
	"log"

	"expenze-io.com/internal/config"
	"expenze-io.com/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// loading configs
	config.LoadConfigs()

	// define our server engine
	server := gin.Default()

	// routes
	routes.RegisterRoutes(server)

	// server running on port 8080
	err := server.Run(":8080")

	if err != nil {
		log.Panic("failed to run server")
	}
}
