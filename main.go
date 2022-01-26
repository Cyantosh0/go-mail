package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/Cyantosh0/go-mail/messages"
	"github.com/Cyantosh0/go-mail/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup Mail Service
	mailer := service.NewMailService()

	// Receiver email address
	to := []string{
		"receiver1@gmail.com",
		"receiver2@gmail.com",
	}

	// Message
	// message := messages.ProperMail("receiver1@gmail.com", "Testing", "This is the email test body.")
	body := struct {
		Name,
		Message string
	}{
		Name:    "Reciever",
		Message: "This is the test message.",
	}
	message := messages.MailUsingHTMLTemplate("templates/email_template.html", body)

	// Send Email
	err = mailer.SendMail(to, message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Email Sent Successfully!")
	}
}
