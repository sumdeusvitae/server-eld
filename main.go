package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env in non-production
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, continuing without it")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local dev
	}

	http.HandleFunc("/drivers", DriversHandler)

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
