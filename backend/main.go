package main

import (
	"fmt"
	"log"
	"net/http"
	"rhoam-together/config"
	"rhoam-together/handlers"
	"rhoam-together/middleware"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	router := mux.NewRouter()

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware(cfg.FrontendURL))

	// Health check endpoint
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// API routes (will be added in later phases)
	api := router.PathPrefix("/api").Subrouter()

	// Auth endpoints (Phase 4)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", notImplementedHandler).Methods("POST")
	auth.HandleFunc("/login", notImplementedHandler).Methods("POST")

	// Trips endpoints (Phase 6)
	trips := api.PathPrefix("/trips").Subrouter()
	trips.HandleFunc("", notImplementedHandler).Methods("GET", "POST")
	trips.HandleFunc("/{id}", notImplementedHandler).Methods("GET", "PUT", "DELETE")

	// WebSocket endpoint (Phase 6)
	router.HandleFunc("/ws", notImplementedHandler).Methods("GET")

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Rhoam Together API starting on http://localhost%s", addr)
	log.Printf("📍 Environment: %s", cfg.Environment)
	log.Printf("🌐 CORS allowed origin: %s", cfg.FrontendURL)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	middleware.RespondWithError(w, http.StatusNotImplemented, "Endpoint not yet implemented")
}
