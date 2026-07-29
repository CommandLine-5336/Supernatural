package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mail_sending/config"
	"mail_sending/mail"
	"net/http"
	netmail "net/mail"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
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

type InviteData struct {
    Email string `json:"email"`
}

type InquisitorMailRequest struct {
	Email string `json:"email"`
	Alias string `json:"alias"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
	}
}
func DailyMailScheduler() {
	c := cron.New()

	c.AddFunc("CRON_TZ=Europe/Kyiv 1 0 * * *", func() { // 00:01
		log.Println("Start scheduled password mailing")
		SendDailyPassword()

	})
	c.Start()
}
func SendDailyPassword() {
	var users []User

	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Fatal("Error fetching emails", result)

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

var jwtSecret = []byte(os.Getenv("JWT_KEY"))
func CreateInviteJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func CreateInvite(
	w http.ResponseWriter,
	r *http.Request,
) {
	var emailData InviteData
	err := json.NewDecoder(r.Body).Decode(&emailData) //decode json into go struct
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

    if emailData.Email == "" {
		http.Error(
			w,
			"Email is required",
			http.StatusBadRequest,
		)
		return
	}

	if _, err := netmail.ParseAddress(emailData.Email); err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest) // 400
		return
	}

    token, err := CreateInviteJWT(emailData.Email)

	if err != nil {
		http.Error(w, "could not create invite token", http.StatusInternalServerError) // 500
		return
	}

    link := fmt.Sprintf("http://127.0.0.1:8100/invite/%s", token)

	err = mail.SendInvite(emailData.Email, link)
	if err != nil {
		log.Print("email not send ", err)
	} else {
		log.Print("email send ", emailData.Email)
	}
}



func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error db not connected")

	}

	// SendDailyPassword() // to send email immediately

	DailyMailScheduler() // start by cron
	log.Println("Cron job started")

	mux := http.NewServeMux()
	mux.HandleFunc("POST /mail", sendMail)
	mux.HandleFunc("POST /inquisitor_mail", sendInquisitorMail)

	log.Println("server listening to  port 8074")
	log.Fatal(http.ListenAndServe(":8074", CORS(mux))) // can be used like ListenAndServeTLS

}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4040")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
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


func sendInquisitorMail(
	w http.ResponseWriter,
	r *http.Request,
) {
	var inquisitorMail InquisitorMailRequest

	err := json.NewDecoder(r.Body).Decode(&inquisitorMail) //decode json into go struct
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	if inquisitorMail.Email == "" {
		http.Error(
			w,
			"Subject can not be empty",
			http.StatusBadRequest,
		)
		return
	}
	if inquisitorMail.Alias == "" {
		http.Error(
			w,
			"Subject can not be empty",
			http.StatusBadRequest,
		)
		return
	}

	err = mail.SendInquisitorMail(inquisitorMail.Email, inquisitorMail.Alias)
	if err != nil {
		log.Print("email not send ", err)
	} else {
		log.Print("email send ", inquisitorMail.Email)
	}
}
