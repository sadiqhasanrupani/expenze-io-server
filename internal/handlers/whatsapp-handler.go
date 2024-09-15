package handlers

import (
	"net/http"

	"expenze-io.com/internal/services"
	"github.com/gin-gonic/gin"
)

type WhatsAppHandler struct {
	Service *services.WhatsAppService
}

func NewWhatsAppHandler(service *services.WhatsAppService) *WhatsAppHandler {
	return &WhatsAppHandler{Service: service}
}

func (h *WhatsAppHandler) SendMessageHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Message from Whatsapp"})
}

func RegisterRoutes(router *gin.Engine) {
	// router.post("/send", handler.SendMessageHandler).Methods("POST")
	router.POST("/send")
}