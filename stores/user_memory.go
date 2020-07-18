package stores

import (
	"apuroda/models"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var UserStore = NewUserMemory()

// UserMemory UserMemory
type UserMemory struct {
	sync.RWMutex
	items []models.User
}

// NewUserMemory NewUserMemory
func NewUserMemory() *UserMemory {
	m := make([]models.User, 0, 100)
	u := &UserMemory{
		items: m,
	}
	return u
}

// All All
func (u *UserMemory) All() []models.User {
	u.Lock()
	defer u.Unlock()
	return u.items
}

// Create Create
func (u *UserMemory) Create(value models.User) {
	u.Lock()
	defer u.Unlock()
	u.items = append(u.items, value)
}

// GetByID GetByID
func (u *UserMemory) GetByID(id uuid.UUID) (*models.User, error) {
	u.Lock()
	defer u.Unlock()
	for _, user := range u.items {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("not found")
}

// GetByName GetByName
func (u *UserMemory) GetByName(name string) (*models.User, error) {
	u.Lock()
	defer u.Unlock()
	for _, user := range u.items {
		if user.Name == name {
			return &user, nil
		}
	}
	return nil, errors.New("not found")
}
