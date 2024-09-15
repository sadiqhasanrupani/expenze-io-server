package services

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsAppService struct {
	Client *whatsmeow.Client
}

func NewWhatsAppService(dsn string) (*WhatsAppService, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("postgres", dsn, dbLog)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %w", err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			return nil, fmt.Errorf("failed to connect: %w", err)
		}

		go func() {
			for evt := range qrChan {
				if evt.Event == "code" {
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					fmt.Println("QR code:", evt.Code)
				} else {
					fmt.Println("Login event:", evt.Event)
				}
			}
		}()
	} else {
		err = client.Connect()
		if err != nil {
			return nil, fmt.Errorf("failed to connect: %w", err)
		}
	}

	return &WhatsAppService{Client: client}, nil
}

func (s *WhatsAppService) SendMessage(phoneNumber, message string) error {
	// 917400262640:3@s.whatsapp.net

	// Parse recipient JID
  recipientJID, err := types.ParseJID(phoneNumber + "@s.whatsapp.net")
	if err != nil {
		return fmt.Errorf("invalid phone number: %w", err)
	}

	// Prepare the message using waProto.Message
	msg := &waProto.Message{
		Conversation: &message,
  }

	// Send the message
	_, err = s.Client.SendMessage(context.Background(), recipientJID, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Println("Message sent to", phoneNumber)
	return nil
}
