package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	// Init database connection
	var err error
	db, err = connectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Close database connection when main exits
	defer db.Close()

	// Register HTTP routes
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/session", session)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/invite/", setTrespassingCookie)

	// Start HTTP server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://frontend:8080", "http://127.0.0.1:8080", "http://mail_service:8074"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Println("Error starting the server:", err)
	}
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func setTrespassingCookie(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.URL.Path, "/invite/")
	if token == "" {
		http.Error(w, "no invite token", http.StatusBadRequest) // 400
		return
	}
	_, err := parseInviteJWT(token)
	if err != nil {
		http.Error(w, "couldn't get email", http.StatusBadRequest) // 400
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "trespass",
		Value:    "true",
		Expires:  time.Now().Add(1 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	})

	w.Header().Set("Content-Type", "application/json")

	register_url := fmt.Sprintf("http://frontend:8080/register/%s", token)
	http.Redirect(w, r, register_url, http.StatusSeeOther)
}

func register(w http.ResponseWriter, r *http.Request) {
	invite_token := r.FormValue("invite_token")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if invite_token != "" && invite_token != "null" && invite_token != "undefined" {
		token_email, err := parseInviteJWT(invite_token)
		if err != nil {
			http.Error(w, "couldn't get email", http.StatusBadRequest) // 400
			return
		}
		email = token_email
	}

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest) // 400
		return
	}

	hashedPassword, _ := hashPassword(password)
	displayName := generateRandomName()

	_, err := db.Exec(
		"INSERT INTO users (alias, email, password, status, inquisitor, banned) VALUES ($1, $2, $3, $4, $5, $6)",
		displayName,
		email,
		hashedPassword,
		"copper",
		false,
		false,
	)

	if err != nil {
		http.Error(w, "user already exists", http.StatusConflict) // 409
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var (
		userID       int
		displayName  string
		passwordHash string
		status       string
		inquisitor   bool
		banned       bool
	)

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := db.QueryRow(
		`SELECT id, alias, password, status, inquisitor, banned FROM users WHERE email=$1`,
		email,
	).Scan(
		&userID,
		&displayName,
		&passwordHash,
		&status,
		&inquisitor,
		&banned,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "no user found with the given email address", http.StatusUnauthorized) // 401
		return
	}

	if !checkPasswordHash(password, passwordHash) {
		http.Error(w, "incorrect password", http.StatusUnauthorized) // 401
		return
	}

	token, err := createJWT(userID)

	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError) // 500
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{
		"id": %d,
		"alias": "%s",
		"status": "%s",
		"inquisitor": %t,
		"banned": %t
	}`, userID, displayName, status, inquisitor, banned)
}

func session(w http.ResponseWriter, r *http.Request) {
	var (
		userID      int
		displayName string
		status      string
		inquisitor  bool
		banned      bool
	)

	userID, err := Authorize(r)

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized) // 401
		return
	}

	err = db.QueryRow(
		`SELECT alias, status, inquisitor, banned FROM users WHERE id=$1`,
		userID,
	).Scan(
		&displayName,
		&status,
		&inquisitor,
		&banned,
	)

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound) // 404
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{
		"id": %d,
		"alias": "%s",
		"status": "%s",
		"inquisitor": %t,
		"banned": %t
	}`, userID, displayName, status, inquisitor, banned)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	fmt.Fprintf(w, "Logout successful!")
}
