package models

import (
	"apuroda/storages"
	"io"
	"time"

	"github.com/google/uuid"
)

// File File
type File struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Get Get
func (f *File) Get() (io.Reader, error) {
	return storages.LegacyLocalStorage.Get(f.ID.String())
}
