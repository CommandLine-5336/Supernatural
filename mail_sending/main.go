package main

import (
	"log"
	"mail_sending/config"
	"mail_sending/mail"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env not finded")
	}
}

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error db not connected")

	}

	err = mail.Mail_sender()
	if err != nil {
		log.Fatal("Error sending emails")

	}

}
