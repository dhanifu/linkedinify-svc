package main

import (
	"log"
	"net/http"

	"linkedinify-svc/internal/config"
	"linkedinify-svc/internal/gemini"
	"linkedinify-svc/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY is missing!")
	}

	geminiClient := gemini.NewClient(cfg.GeminiAPIKey)

	mux := http.NewServeMux()
	mux.HandleFunc("/translate", handlers.TranslateHandler(geminiClient))

	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
