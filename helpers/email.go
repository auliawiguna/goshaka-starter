package helpers

import (
	// "bytes"
	"goshaka/configs"

	"gopkg.in/gomail.v2"
)

func SendMail(recipient string, subject string, data struct{}) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", configs.GetEnv("MAIL_FROM"))
	mailer.SetHeader("From", recipient)
	mailer.SetHeader("Subject", subject)

	// var b bytes.Buffer
}
