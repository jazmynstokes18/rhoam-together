package middleware

import (
	"fmt"
	"net/http"
	"rhoam-together/utils"
	"strings"
)

// JWTMiddleware checks for valid JWT token in Authorization header
func JWTMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				RespondWithError(w, http.StatusUnauthorized, "Missing authorization header")
				return
			}

			// Extract token from "Bearer <token>" format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				RespondWithError(w, http.StatusUnauthorized, "Invalid authorization header format")
				return
			}

			tokenString := parts[1]

			// Verify token
			_, err := utils.VerifyToken(tokenString, jwtSecret)
			if err != nil {
				RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
				return
			}

			// Store user info in request context for later use
			// (You can enhance this to pass user ID to handlers)

			next.ServeHTTP(w, r)
		})
	}
}
