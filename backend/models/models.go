package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Trip represents a travel trip
type Trip struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

// TripMember represents a user's access to a trip
type TripMember struct {
	ID          uuid.UUID `json:"id"`
	TripID      uuid.UUID `json:"trip_id"`
	UserID      uuid.UUID `json:"user_id"`
	AccessLevel string    `json:"access_level"` // read_only, editor, admin
	JoinedAt    time.Time `json:"joined_at"`
}

// Suggestion represents a suggestion/plan for a trip
type Suggestion struct {
	ID            uuid.UUID  `json:"id"`
	TripID        uuid.UUID  `json:"trip_id"`
	CreatedBy     uuid.UUID  `json:"created_by"`
	Title         string     `json:"title"`
	Description   string     `json:"description,omitempty"`
	SuggestedDate *time.Time `json:"suggested_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}
