package services

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type WhatsAppService struct {
	Client *whatsmeow.Client
}

func NewWhatsAppService(connStr string) (*WhatsAppService, error) {

	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New("postgres", connStr, dbLog)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %w", err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
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

func (s *WhatsAppService) SendOtpButtonMessage(phoneNumber, title, content, footer, buttonText, otpCode string) error {

	// Parse recipient JID
	recipientJID, err := types.ParseJID(phoneNumber + "@s.whatsapp.net")
	if err != nil {
		return fmt.Errorf("invalid phone number: %w", err)
	}

	// Build the URL for copying the OTP
	copyToClipboardUrl := fmt.Sprintf("https://www.whatsapp.com/otp/code/?otp_type=COPY_CODE&code_expiration_minutes=10&code=%v", otpCode)
	timestamp, _ := strconv.ParseUint("1724818629", 10, 64)

	index := uint32(0)

	// Create Hydrated Template Button
	hyderatedButtons := []*waProto.HydratedTemplateButton{
		{
			Index: &index,
			HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
				UrlButton: &waProto.HydratedTemplateButton_HydratedURLButton{
					DisplayText: &buttonText,
					URL:         &copyToClipboardUrl,
				},
			},
		},
	}

	// Create Hydrated Template Message
	hydratedTemplate := &waProto.TemplateMessage_HydratedFourRowTemplate{
		HydratedContentText: &content,
		HydratedFooterText:  &footer,
		HydratedButtons:     hyderatedButtons,
		Title: &waProto.TemplateMessage_HydratedFourRowTemplate_HydratedTitleText{
			HydratedTitleText: title,
		},
	}

	templateMessage := &waE2E.TemplateMessage{
		HydratedTemplate: hydratedTemplate,
		TemplateID:       proto.String("4194019344155670"),
	}

	deviceListMetaData := &waE2E.DeviceListMetadata{
		RecipientKeyHash:    []byte{},
		RecipientTimestamp:  &timestamp,
		RecipientKeyIndexes: []uint32{},
	}

	messageContextinfo := &waE2E.MessageContextInfo{
		DeviceListMetadataVersion: proto.Int32(2),
		DeviceListMetadata:        deviceListMetaData,
	}

	msg := &waE2E.Message{
		TemplateMessage:    templateMessage,
		MessageContextInfo: messageContextinfo,
	}

	msgViewOnce := &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: msg,
		},
	}

	// Send the message
	_, err = s.Client.SendMessage(context.Background(), recipientJID, msgViewOnce)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Println("Message sent to", phoneNumber)
	return nil
}

func (s *WhatsAppService) SendMessage(phoneNumber, content string) error {
	// Parse recipient JID
	recipientJID, err := types.ParseJID(phoneNumber + "@s.whatsapp.net")
	if err != nil {
		return fmt.Errorf("invalid phone number: %w", err)
	}

	message := &waProto.Message{
		Conversation: proto.String(content),
	}

	// Send the message
	_, err = s.Client.SendMessage(context.Background(), recipientJID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Println("Message sent to", phoneNumber)
	return nil
}
