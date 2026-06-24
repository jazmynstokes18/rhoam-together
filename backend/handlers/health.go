package handlers

import (
	"net/http"
	"rhoam-together/middleware"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	middleware.RespondWithJSON(w, http.StatusOK, HealthResponse{
		Status:  "ok",
		Message: "Rhoam Together API is running",
	})
}
