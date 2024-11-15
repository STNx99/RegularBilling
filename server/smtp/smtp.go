package smtp

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(to []string, money float32) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_MAIL"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_HOST"),
	)

msg := []byte(fmt.Sprintf("To: %s\r\n"+ "Subject: Monthly Services Bill\r\n"+ "\r\n"+ "Your monthly services bill is: %.2f\r\n", to[0], money))
err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_MAIL"),
		to,
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
