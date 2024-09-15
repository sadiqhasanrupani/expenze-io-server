package config

import (
	"log"

	"github.com/joho/godotenv"
)

type WhatsAppService interface {
	SendMessage(phoneNumber string, message string) error
}

func LoadConfigs() {
	// loading .env
	err := godotenv.Load()

	if err != nil {
		log.Fatalf(".env is not able to load, err: %v", err)
	}


  ConnectWhatsapp()
	InitDB()
}
