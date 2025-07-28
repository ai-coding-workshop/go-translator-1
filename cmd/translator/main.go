package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"translator-service/internal/handlers"
)

func main() {
	fmt.Println("Translation Service Starting...")

	// Create a new serve mux for routing
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/translate", handlers.TranslateHandler)
	mux.HandleFunc("/api/translate", handlers.APIHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
