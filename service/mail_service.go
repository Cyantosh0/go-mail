package service

import (
	"fmt"
	"net/smtp"
	"os"
)

type MailService struct {
	SMTPService smtp.Auth
	SenderEmail,
	Host,
	Port string
}

type Mail struct {
	SenderEmail    string
	SenderPassword string
	Message        string
}

func NewMailService() MailService {
	senderEmail := os.Getenv("SENDER_EMAIL")
	host := os.Getenv("SMTP_HOST")

	auth := smtp.PlainAuth("", senderEmail, os.Getenv("SENDER_PASSWORD"), host)

	return MailService{
		SMTPService: auth,
		SenderEmail: senderEmail,
		Host:        host,
		Port:        os.Getenv("SMTP_PORT"),
	}
}

func (m MailService) SendMail(to []string, message []byte) error {
	fmt.Println("Sending Email ---->")

	err := smtp.SendMail(m.Host+":"+m.Port, m.SMTPService, m.SenderEmail, to, message)
	if err != nil {
		return err
	}
	return nil
}
