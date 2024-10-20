package main

import (
	"fmt"
	"log"
	"os"

	"expenze-io.com/internal/config"
	cronjobs "expenze-io.com/internal/cron-jobs"
	"expenze-io.com/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// loading configs
	config.LoadConfigs()

	PORT := os.Getenv("SERVER_PORT")
	port := fmt.Sprintf(":%v", PORT)

	// define our server engine
	server := gin.Default()

	// cron jobs
	cronjob := cronjobs.New(config.DB)
	cronjob.Start()

	// routes
	routes.RegisterRoutes(server)

	err := server.Run(port)

	if err != nil {
		log.Panic("failed to run server")
	}
}
