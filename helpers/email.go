package helpers

import (
	// "bytes"
	"bytes"
	"fmt"
	"goshaka/configs"
	"strconv"

	"html/template"

	"gopkg.in/gomail.v2"
)

// To send email to recipient
//
//	param recipient string
//	param subject string
//	param emailType string
//	param data interface{}
//	return	error
func SendMail(recipient string, subject string, emailType string, data interface{}) error {
	emailPort, _ := strconv.Atoi(configs.GetEnv("MAIL_PORT"))
	dialer := gomail.NewDialer(configs.GetEnv("MAIL_HOST"), emailPort, configs.GetEnv("MAIL_USERNAME"), configs.GetEnv("MAIL_PASSWORD"))
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", configs.GetEnv("MAIL_FROM"))
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", subject)

	html := template.Must(template.ParseFiles("templates/" + emailType + ".html"))

	var b bytes.Buffer
	if err := html.Execute(&b, data); err != nil {
		return fmt.Errorf("error parse html")
	}

	mailer.SetBody("text/html", b.String())

	go func() {
		dialer.DialAndSend(mailer)
	}()

	return nil
}
