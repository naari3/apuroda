package stores

import (
	"apuroda/models"
	"errors"
	"sync"

	"github.com/google/uuid"
)

// FileMemoryStore FileMemoryStore
var FileMemoryStore = NewFileMemory()

// FileMemory FileMemory
type FileMemory struct {
	sync.RWMutex
	items []models.File
}

// NewFileMemory NewFileMemory
func NewFileMemory() *FileMemory {
	m := make([]models.File, 0, 100)
	u := &FileMemory{
		items: m,
	}
	return u
}

// All All
func (u *FileMemory) All() []models.File {
	u.Lock()
	defer u.Unlock()
	return u.items
}

// Create Create
func (u *FileMemory) Create(value models.File) {
	u.Lock()
	defer u.Unlock()
	u.items = append(u.items, value)
}

// GetByID GetByID
func (u *FileMemory) GetByID(id uuid.UUID) (*models.File, error) {
	u.Lock()
	defer u.Unlock()
	for _, file := range u.items {
		if file.ID == id {
			return &file, nil
		}
	}
	return nil, errors.New("not found")
}
