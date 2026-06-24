package handlers

import (
	"database/sql"
	"net/http"
	"rhoam-together/middleware"
	"rhoam-together/utils"
	"strings"
)

// GetCurrentUser returns the authenticated user's information
func GetCurrentUser(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			middleware.RespondWithError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.VerifyToken(tokenString, jwtSecret)
		if err != nil {
			middleware.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Query user from database
		var name string
		err = db.QueryRow(
			"SELECT name FROM users WHERE id = $1",
			claims.UserID,
		).Scan(&name)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		middleware.RespondWithJSON(w, http.StatusOK, UserResponse{
			ID:    claims.UserID,
			Email: claims.Email,
			Name:  name,
		})
	}
}
