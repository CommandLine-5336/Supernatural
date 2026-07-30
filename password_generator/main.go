package main

import (
	"database/sql"
	"fmt"
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

	// Generate password
	password := PasswordGenerator(64)
	fmt.Printf("Website password: %s\n", password)
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

func PasswordGenerator(passwordLength int) string {
	// Character sets for generating passwords
	lowerCase := "abcdefghijklmnopqrstuvwxyz" // lowercase
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // uppercase
	numbers := "0123456789"                   // numbers
	specialChar := "!@#$%^&*()_-+={}[/?]"     // special characters

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
