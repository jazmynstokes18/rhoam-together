package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rhoam-together/config"
	"rhoam-together/handlers"
	"rhoam-together/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	// Connect to database
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("✓ Database connection successful")

	router := mux.NewRouter()
	corsHandler := middleware.CORSMiddleware(cfg.FrontendURL, cfg.Environment)(router)
	// Health check endpoint
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Auth endpoints (Phase 4)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", handlers.Signup(db, cfg.JWTSecret)).Methods("POST")
	auth.HandleFunc("/login", handlers.Login(db, cfg.JWTSecret)).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	protected.HandleFunc("/me", handlers.GetCurrentUser(db, cfg.JWTSecret)).Methods("GET")

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
	err = http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Printf("🚨 Error at the listen and Serve: %s", err.Error())
	}

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	middleware.RespondWithError(w, http.StatusNotImplemented, "Endpoint not yet implemented")
}
