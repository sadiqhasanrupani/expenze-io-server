package mails

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type SendMail struct {
	Subject      string
	TemplatePath string
	To           []string
	TemplateData any
}

func SendMailTemplate(sm *SendMail, doneChan chan bool, errorChan chan error) {
	companyEmail := os.Getenv("COMPANY_EMAIL")
	appPass := os.Getenv("APP_PASS")
	smtpEmail := os.Getenv("SMTP_EMAIL")

	var body bytes.Buffer

	t, err := template.ParseFiles(sm.TemplatePath)
	if err != nil {
		errorChan <- err
		return
	}

	if err := t.Execute(&body, sm.TemplateData); err != nil {
		errorChan <- err
		return
	}

	auth := smtp.PlainAuth(
		"",
		companyEmail,
		appPass,
		smtpEmail,
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := fmt.Sprintf("Subject: %v \n%v \n\n %v", sm.Subject, headers, body.String())

	// Send the email
	err = smtp.SendMail(smtpEmail+":587", auth, companyEmail, sm.To, []byte(msg))
	if err != nil {
		errorChan <- err
		return
	}

	doneChan <- true
}
