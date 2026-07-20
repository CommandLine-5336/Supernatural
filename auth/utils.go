package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func generateRandomName() string {
	for {
		adjIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(adjectives))))
		nounIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(nouns))))
		candidate := fmt.Sprintf("%s %s", adjectives[adjIndex.Int64()], nouns[nounIndex.Int64()])

		taken := false
		for _, user := range users {
			if user.DisplayName == candidate {
				taken = true
				break
			}
		}

		if !taken {
			return candidate
		}
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
