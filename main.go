package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Starting HTTP server on port", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
