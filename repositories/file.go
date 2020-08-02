package repositories

import (
	"apuroda/models"
	"apuroda/storages"
	"apuroda/stores"
	"io"
	"time"

	"github.com/google/uuid"
)

// FileRepository FileRepository
type FileRepository struct{}

// All All
func (f FileRepository) All() []models.File {
	files := stores.FileStore.All()

	return files
}

// GetByID GetByID
func (f FileRepository) GetByID(id uuid.UUID) (*models.File, error) {
	file, err := stores.FileStore.GetByID(id)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Create Create
func (f FileRepository) Create(name string, binary io.Reader) (*models.File, error) {
	id, _ := uuid.NewRandom()
	file := models.File{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	stores.FileStore.Create(file)
	err := storages.LegacyLocalStorage.Set(id.String(), binary)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
