package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

type User struct {
	DisplayName    string
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]User{}

var adjectives = []string{"Swift", "Brave", "Clever", "Quiet", "Bright", "Shadowy", "Wild", "Calm"}
var nouns = []string{"Fox", "Shadow", "Eagle", "Wolf", "River", "Falcon", "Bear", "Lion", "Hawk", "Knight"}

func main() {
	log.Println("Server listening on :8080")
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if _, ok := users[email]; ok {
		status := http.StatusConflict // 409
		http.Error(w, "User already exists", status)
		return
	}

	hashedPassword, _ := hashPassword(password)

	displayName := generateRandomName()

	users[email] = User{
		DisplayName:    displayName,
		HashedPassword: hashedPassword,
	}

	fmt.Fprintf(w, "User registered successfully! Assigned Name: %s\n", displayName)
}

func login(w http.ResponseWriter, r *http.Request) {
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	user, ok := users[email]
	if !ok || !checkPasswordHash(password, user.HashedPassword) {
		status := http.StatusUnauthorized // 401
		http.Error(w, "Invalid email or password", status)
		return
	}

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	// Set CSRF token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false, // Needs to be accessible to the client-side
	})

	// Store tokens in the database
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[email] = user

	fmt.Fprintln(w, "Login successful!")
}

func protected(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		status := http.StatusUnauthorized // 401
		http.Error(w, "Unauthorized", status)
		return
	}

	email := r.FormValue("email")
	fmt.Fprintf(w, "Welcome, %s", email)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		status := http.StatusUnauthorized // 401
		http.Error(w, "Unauthorized", status)
		return
	}

	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	// Delete tokens from the database
	email := r.FormValue("email")
	user, _ := users[email]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[email] = user

	fmt.Fprintln(w, "Logged out successfully!")
}
