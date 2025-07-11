package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}

	http.HandleFunc("/drivers", DriversHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
