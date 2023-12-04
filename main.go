package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GuillePantz/Guille_Microservice/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Get PostgreSQL connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize your router using Gorilla Mux
	router := mux.NewRouter()

	// Define your API endpoints and handlers here

	// Example endpoint
	router.HandleFunc("/appointments", func(w http.ResponseWriter, r *http.Request) {
        handlers.CreateAppointment(w, r, db)
    }).Methods("POST")

	// Start the web server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
