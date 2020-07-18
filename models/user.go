package models

import (
	"github.com/google/uuid"
)

// User User
type User struct {
	ID   uuid.UUID
	Name string
}
