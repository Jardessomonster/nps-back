package main

import (
	"log"
	"net/http"

	functions "github.com/Jardessomonster/nps-back"
	"github.com/Jardessomonster/nps-back/internal/infra/middlewares"
)

func main() {
	http.HandleFunc("/create-survey", middlewares.CorsMiddleware(functions.CreateSurvey))

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
