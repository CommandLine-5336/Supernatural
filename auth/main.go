package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type User struct {
	DisplayName    string
	HashedPassword string
}

var users = map[string]User{}

var adjectives = []string{"Swift", "Brave", "Clever", "Quiet", "Bright", "Shadowy", "Wild", "Calm"}
var nouns = []string{"Fox", "Shadow", "Eagle", "Wolf", "River", "Falcon", "Bear", "Lion", "Hawk", "Knight"}

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
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)

	// Start HTTP server
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func connectDB() (*sql.DB, error) {
	connStr := os.Getenv("DSN")
	return sql.Open("postgres", connStr)
}

func register(w http.ResponseWriter, r *http.Request) {
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest) // 400
		return
	}

	if _, ok := users[email]; ok {
		http.Error(w, "user already exists", http.StatusConflict) // 409
		return
	}

	hashedPassword, _ := hashPassword(password)

	displayName := generateRandomName()

	users[email] = User{
		DisplayName:    displayName,
		HashedPassword: hashedPassword,
	}

	fmt.Fprintf(w, "Registration successful! Assigned Name: %s", displayName)
}

func login(w http.ResponseWriter, r *http.Request) {
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	user, ok := users[email]
	if !ok || !checkPasswordHash(password, user.HashedPassword) {
		http.Error(w, "invalid email or password", http.StatusUnauthorized) // 401
		return
	}

	token, err := createJWT(email)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError) // 500
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	fmt.Fprintf(w, "Login successful, %s!", users[email].DisplayName)
}

func protected(w http.ResponseWriter, r *http.Request) {
	email, err := Authorize(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized) // 401
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", users[email].DisplayName)
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
