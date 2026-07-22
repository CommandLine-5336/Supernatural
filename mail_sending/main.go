package main

import (
	"fmt"
	"log"
	"mail_sending/config"
	"mail_sending/mail"

	"github.com/joho/godotenv"
)

type User struct {
	ID     uint `gorm:"primaryKey"`
	Alias  string
	Email  string
	Status string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env not found")
	}
}

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error db not connected")

	}
	var users []User

	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Fatal("Error fetchin emails")

	}
	// send email for each email in db
	for _, user := range users {
		fmt.Printf("Here is our users %s his status is %s\n", user.Email, user.Status)
		err = mail.SendPasswordMail(user.Email, user.Status, user.Alias)
		if err != nil {
			log.Fatalf("email not send", err)

		}
	}
	// err = mail.Mail_sender()
	if err != nil {
		log.Fatal("Error sending emails")

	}

}
