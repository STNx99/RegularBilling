package smtp

import (
	"bytes"
	"github.com/dustin/go-humanize"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"time"
)

type EmailData struct {
	Date    string
	Service []ServiceInfo
	Money   string
}

func SendMail(to []string, money float64, service []ServiceInfo) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_MAIL"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_HOST"),
	)

	now := time.Now()

	var data = EmailData{
		Date:    now.Format("15:04:05 02/01/2006"),
		Service: service,
		Money:   humanize.Comma(int64(money)),
	}
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Không thể đọc template: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Fatalf("Không thể render template: %v", err)
	}

	msg := []byte(
		"Subject: Thông Tin Thanh Toán\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
			body.String(),
	)

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_MAIL"),
		to,
		msg,
	)
	if err != nil {
		log.Fatalf("Lỗi khi gửi email: %v", err)
	}

	log.Println("Email đã được gửi thành công!")
	return nil
}
