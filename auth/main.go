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
		http.Error(w, "Invalid email format", http.StatusBadRequest) // 400
		return
	}

	if _, ok := users[email]; ok {
		http.Error(w, "User already exists", http.StatusConflict) // 409
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
		http.Error(w, "Invalid email or password", http.StatusUnauthorized) // 401
		return
	}

	token, err := createJWT(email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError) // 500
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized) // 401
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
