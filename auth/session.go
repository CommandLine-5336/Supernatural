package main

import (
	"errors"
	"net/http"
)

var AuthError = errors.New("unauthorized")

func Authorize(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil || cookie.Value == "" {
		return "", AuthError
	}

	email, err := parseJWT(cookie.Value)
	if err != nil {
		return "", AuthError
	}

	return email, nil
}
