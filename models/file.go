package models

import (
	"github.com/google/uuid"
)

// File File
type File struct {
	ID        uuid.UUID
	Name      string
	Path      string
	CreatedAt int64
	UpdatedAt int64
}
