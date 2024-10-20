package config

import (
	"log"
	"os"

	"expenze-io.com/internal/services"
)

func ConnectWhatsapp() {
	connStr := os.Getenv("PG_CONNSTR")

	// whatsapp connection
	_, err := services.NewWhatsAppService(connStr)
	if err != nil {
		log.Panic(err)
	}

}
