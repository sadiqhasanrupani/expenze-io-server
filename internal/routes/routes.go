package routes

import (
	"log"
	"net/http"
	"os"

	"expenze-io.com/internal/handlers"
	"expenze-io.com/internal/services"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api/v1"

func RegisterRoutes(router *gin.Engine) {
	// Initialize WhatsApp service
	waService, err := services.NewWhatsAppService(os.Getenv("PG_CONNSTR"))
	if err != nil {
		log.Fatalf("Failed to initialize WhatsApp service: %v", err)
	}

	// Initialize WhatsApp handler
	waHandler := handlers.NewWhatsAppHandler(waService)

	// Root route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Expenze member!!",
		})
	})

	// Register WhatsApp routes
	router.POST(BaseUrl+"/send", waHandler.SendMessageHandler)

	// Add other routes as needed
	authRoute(router)
}

