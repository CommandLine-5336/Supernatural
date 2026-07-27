package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func WriteResponseToJSON(w http.ResponseWriter, status int, payload map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func DBHealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(); err != nil {
		WriteResponseToJSON(w, http.StatusServiceUnavailable, map[string]any{"status": "error"})
		return
	}
	WriteResponseToJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func eraseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteResponseToJSON(w, http.StatusMethodNotAllowed, map[string]any{"status": "error", "message": "use POST"})
		return
	}

	if token := os.Getenv("ERASE_TOKEN"); token != "" && r.Header.Get("X-Erase-Token") != token {
		WriteResponseToJSON(w, http.StatusUnauthorized, map[string]any{"status": "error", "message": "ERASE_TOKEN not provided or provided an incorrect token"})
		return
	}

	rows, err := db.Query(`SELECT tablename FROM pg_tables WHERE schemaname = 'public'`)
	if err != nil {
		WriteResponseToJSON(w, http.StatusInternalServerError, map[string]any{"status": "error", "message": err.Error()})
		return
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			WriteResponseToJSON(w, http.StatusInternalServerError, map[string]any{"status": "error", "message": err.Error()})
			return
		}
		tables = append(tables, `"`+t+`"`)
	}
	if err := rows.Err(); err != nil {
		WriteResponseToJSON(w, http.StatusInternalServerError, map[string]any{"status": "error", "message": err.Error()})
		return
	}
	if len(tables) == 0 {
		WriteResponseToJSON(w, http.StatusOK, map[string]any{"status": "ok", "message": "No Tables Found"})
		return
	}
	query := "TRUNCATE TABLE " + strings.Join(tables, ", ") + " RESTART IDENTITY CASCADE"
	if _, err := db.Exec(query); err != nil {
		WriteResponseToJSON(w, http.StatusInternalServerError, map[string]any{"status": "error", "message": err.Error()})
		return
	}
	WriteResponseToJSON(w, http.StatusOK, map[string]any{
		"status":           "ok",
		"tables_truncated": tables,
	})
}

var db *sql.DB

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set via ENV")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	log.Println("connected to database")

	http.HandleFunc("/erase", eraseHandler)
	http.HandleFunc("/healthdb", DBHealthCheck)

	// Start HTTP server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Println("Error starting the server:", err)
	}
}
