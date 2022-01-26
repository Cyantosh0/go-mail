package messages

import (
	"bytes"
	"fmt"
	"text/template"
)

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

func MailUsingTemplate(templatePath string, data interface{}) []byte {
	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: my subject\n%s\n\n", headers)))

	t.Execute(&body, data)
	return body.Bytes()
}
