package external

import (
	"backend-test/helpers"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (e *Email) SendEmail() error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("To", e.To)
	mailer.SetHeader("From", helpers.GetEnv("SMTP_AUTH_EMAIL", ""))
	mailer.SetHeader("Subject", e.Subject)
	mailer.SetBody("text/html", e.Body)

	smtpPort := helpers.GetEnv("SMTP_PORT", "")
	intSmtpPort, _ := strconv.Atoi(smtpPort)

	dialer := gomail.NewDialer(
		helpers.GetEnv("SMTP_HOST", ""),
		intSmtpPort,
		helpers.GetEnv("SMTP_AUTH_EMAIL", ""),
		helpers.GetEnv("SMTP_AUTH_PASSWORD", ""),
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
