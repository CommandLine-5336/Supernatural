package main

import (
	"database/sql"
	"log"
	"math/rand"
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

	err = rotatePassword(db)
	if err != nil {
		log.Fatal(err)
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

func rotatePassword(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Disable current password
	_, err = tx.Exec(
		"UPDATE website_passwords SET is_active = false WHERE is_active = true",
	)
	if err != nil {
		return err
	}

	// Create new password
	password := PasswordGenerator(64)

	_, err = tx.Exec(
		"INSERT INTO website_passwords (password, is_active) VALUES ($1, $2)",
		password,
		true,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func PasswordGenerator(passwordLength int) string {
	// Character sets for generating passwords
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789Є"
	specialChar := "!@#$%^&*()_-+={}[/?]"

	// Variable for storing password
	password := ""

	// Initialize the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Generate password character by character
	for n := 0; n < passwordLength; n++ {
		// Generate a random number to choose a character set
		randNum := rng.Intn(4)

		switch randNum {
		case 0:
			randCharNum := rng.Intn(len(lowerCase))
			password += string(lowerCase[randCharNum])
		case 1:
			randCharNum := rng.Intn(len(upperCase))
			password += string(upperCase[randCharNum])
		case 2:
			randCharNum := rng.Intn(len(numbers))
			password += string(numbers[randCharNum])
		case 3:
			randCharNum := rng.Intn(len(specialChar))
			password += string(specialChar[randCharNum])
		}
	}

	return password
}
