package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("✅ Starting server on port:", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ Hello from Cloud Run"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
