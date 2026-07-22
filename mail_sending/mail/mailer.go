package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func SendPasswordMail(emailAddress string, userStatus string, userAlias string) error {
	email_host := os.Getenv("EMAIL_HOST")
	email_from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	var AWSpassword = "ab&uOJ^ILDgg,dslOD" //this part need to be done via AWS secretmager
	templateData := struct {
		Status     string
		SecretPswd string
		Alias      string
	}{
		Status:     userStatus,
		SecretPswd: AWSpassword,
		Alias:      userAlias,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(fmt.Errorf("failed to parse template file: %w", err))
	}

	d := gomail.NewDialer(email_host, 587, email_from, password)

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, templateData); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "unown@gmail.com")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", "Weather Forecast for this night!")
	m.SetBody("text/html", bodyBuffer.String())

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}
