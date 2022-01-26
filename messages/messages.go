package messages

import "fmt"

func MailWithBodyOnly() []byte {
	body := "Hello Receiver! I hope you are fine."
	return []byte(body)
}

func ProperMail(to, subject, body string) []byte {
	mail := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	return []byte(mail)
}
