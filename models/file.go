package models

import (
	"apuroda/storages"
	"io"

	"github.com/google/uuid"
)

// File File
type File struct {
	ID        uuid.UUID
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

// Get Get
func (f *File) Get() (io.Reader, error) {
	return storages.LegacyLocalStorage.Get(f.ID.String())
}
