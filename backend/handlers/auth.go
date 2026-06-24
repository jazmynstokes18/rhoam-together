package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/mail"
	"rhoam-together/middleware"
	"rhoam-together/utils"
	"strings"
)

// Signup handles user registration
func Signup(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignupRequest

		// Parse JSON request body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		// For now, parse as JSON (better approach)
		if req.Email == "" {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request format")
			return
		}

		// Validate input
		if err := validateSignupInput(&req); err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Hash password
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, "Error processing password")
			return
		}

		// Insert user into database
		var userID string
		err = db.QueryRow(
			"INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3) RETURNING id",
			req.Email, hashedPassword, req.Name,
		).Scan(&userID)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				middleware.RespondWithError(w, http.StatusConflict, "Email already registered")
			} else {
				middleware.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			}
			return
		}

		// Generate JWT token
		token, err := utils.GenerateToken(userID, req.Email, jwtSecret)
		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		// Return response
		middleware.RespondWithJSON(w, http.StatusCreated, AuthResponse{
			Token: token,
			User: UserResponse{
				ID:    userID,
				Email: req.Email,
				Name:  req.Name,
			},
		})
	}
}

// Login handles user authentication
func Login(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		// Parse JSON request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			req.Email = r.FormValue("email")
			req.Password = r.FormValue("password")
		}

		if req.Email == "" || req.Password == "" {
			middleware.RespondWithError(w, http.StatusBadRequest, "Email and password required")
			return
		}

		// Query user from database
		var userID, passwordHash, name string
		err := db.QueryRow(
			"SELECT id, password_hash, name FROM users WHERE email = $1",
			req.Email,
		).Scan(&userID, &passwordHash, &name)

		if err == sql.ErrNoRows {
			middleware.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		} else if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, "Error querying user")
			return
		}

		// Verify password
		if !utils.VerifyPassword(passwordHash, req.Password) {
			middleware.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		// Generate JWT token
		token, err := utils.GenerateToken(userID, req.Email, jwtSecret)
		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		// Return response
		middleware.RespondWithJSON(w, http.StatusOK, AuthResponse{
			Token: token,
			User: UserResponse{
				ID:    userID,
				Email: req.Email,
				Name:  name,
			},
		})
	}
}

// Validation helper
func validateSignupInput(req *SignupRequest) error {
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return error(stringError("Email, password, and name are required"))
	}

	if len(req.Password) < 8 {
		return error(stringError("Password must be at least 8 characters"))
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return error(stringError("Invalid email format"))
	}

	return nil
}

type stringError string

func (e stringError) Error() string {
	return string(e)
}
