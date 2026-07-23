package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mail_sending/config"
	"mail_sending/mail"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

type User struct {
	ID     uint `gorm:"primaryKey"`
	Alias  string
	Email  string
	Status string
}

type MailData struct {
	TargetStatus []string `json:"TargetStatus"`
	Subject      string   `json:"Subject"`
	BodyText     string   `json:"BodyText"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error .env not found")
	}
}
func DailyMailScheduler() {
	c := cron.New()

	c.AddFunc("CRON_TZ=Europe/Kyiv 1 0 * * *", func() { // 00:01
		log.Print("Start cron scheduler")
		SendDailyPassword()

	})
	c.Start()
}
func SendDailyPassword() {
	var users []User

	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Fatal("Error fetchin emails", result)

	}
	// send email for each email in db
	for _, user := range users {
		err := mail.SendPasswordMail(user.Email, user.Status, user.Alias)
		if err != nil {
			log.Print("Email not send ", err)
		} else {
			log.Print("Email send ", user.Email)

		}

	}
}

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error db not connected")

	}

	// SendDailyPassword() // to send email immediately

	go DailyMailScheduler() // start by cron

	mux := http.NewServeMux()
	mux.HandleFunc("POST /send_mail", sendMail)

	fmt.Println("server listening to  port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux)) // can be used like ListenAndServeTLS

}

func sendMail(
	w http.ResponseWriter,
	r *http.Request,
) {
	var emailData MailData
	var users []User

	err := json.NewDecoder(r.Body).Decode(&emailData) //decode json into go struct
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	if len(emailData.TargetStatus) == 0 {
		http.Error(
			w,
			"Status required(Cupper,silver,golden)",
			http.StatusBadRequest,
		)
		return
	}
	if emailData.Subject == "" {
		http.Error(
			w,
			"Subject can not be empty",
			http.StatusBadRequest,
		)
		return
	}
	if emailData.BodyText == "" {
		http.Error(
			w,
			"Subject can not be empty",
			http.StatusBadRequest,
		)
		return
	}
	result := config.DB.Where("status IN ?", emailData.TargetStatus).Find(&users)
	if result.Error != nil {
		http.Error(w, "db", http.StatusInternalServerError)
		return
	}
	for _, user := range users {
		err := mail.SendCustomdMail(user.Email, user.Status, emailData.Subject, emailData.BodyText)
		if err != nil {
			log.Print("email not send ", err)
		} else {
			log.Print("email send ", user.Email)

		}

	}

}
