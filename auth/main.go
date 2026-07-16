package main

import (
	"fmt"
	"net/http"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]Login{}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")

	if _, ok := users[username]; ok {
		status := http.StatusConflict // 409
		http.Error(w, "User already exists", status)
		return
	}

	hashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	fmt.Fprintln(w, "User registered successfully!")
}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func home(w http.ResponseWriter, r *http.Request) {}
