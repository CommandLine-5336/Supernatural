package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func Mail_sender() error {
	email_host := os.Getenv("EMAIL_HOST")
	email_from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(fmt.Errorf("failed to parse template file: %w", err))
	}
	students := map[int]string{
		101: "denisfolyush@gmail.com",
		102: "soosysoosy69@gmail.com",
		// 103: "enisfolyush@gmail.com",
	}
	d := gomail.NewDialer(email_host, 587, email_from, password)

	for _, value := range students {
		var bodyBuffer bytes.Buffer
		if err := tmpl.Execute(&bodyBuffer, nil); err != nil {
			return err
		}

		m := gomail.NewMessage()

		m.SetHeader("From", "unown@gmail.com")
		m.SetHeader("To", value)
		// m.SetAddressHeader("Cc", "denisfolyush@gmail.com", "Daen")
		m.SetHeader("Subject", "Weather Forecast for this night!")
		m.SetBody("text/html", bodyBuffer.String())
		// m.Attach("/home/Alex/lolcat.jpg")

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			return err
		}
	}
	return nil
}
