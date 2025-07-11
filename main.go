package main

import (
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {
	log.Println("🟢 Starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}
	// mongoURI := os.Getenv("MONGO_URI")

	// if mongoURI == "" {
	// 	log.Fatal("MONGO_URI not set")
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collection = client.Database("eld_data").Collection("drivers")

	http.HandleFunc("/drivers", DriversHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Println("🟢 Server is now listening on port:", port)

}
