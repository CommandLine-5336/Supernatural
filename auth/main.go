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
	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest) // 400
		return
	}

	hashedPassword, _ := hashPassword(password)
	displayName := generateRandomName()

	_, err := db.Exec(
		"INSERT INTO users (display_name, email, password, status) VALUES ($1, $2, $3, $4)",
		displayName,
		email,
		hashedPassword,
		"Copper",
	)

	if err != nil {
		http.Error(w, "user already exists", http.StatusConflict) // 409
		return
	}

	fmt.Fprintf(w, "Registration successful! Assigned Name: %s", displayName)
}

func login(w http.ResponseWriter, r *http.Request) {
	var (
		userID       int
		displayName  string
		passwordHash string
		status       string
	)

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := db.QueryRow(
		`SELECT id, display_name, password, status FROM users WHERE email=$1`,
		email,
	).Scan(
		&userID,
		&displayName,
		&passwordHash,
		&status,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "no user found with the given email address", http.StatusUnauthorized) // 401
		return
	}

	if !checkPasswordHash(password, passwordHash) {
		http.Error(w, "incorrect password", http.StatusUnauthorized) // 401
		return
	}

	token, err := createJWT(
		userID,
		status,
	)

	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError) // 500
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	fmt.Fprintf(w, "Login successful, %s!", displayName)
}

func protected(w http.ResponseWriter, r *http.Request) {
	userID, err := Authorize(r)

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var displayName string

	err = db.QueryRow(
		`SELECT display_name FROM users WHERE id=$1`,
		userID,
	).Scan(&displayName)

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound) // 404
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", displayName)
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
