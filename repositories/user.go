package repositories

import (
	"apuroda/models"
	"apuroda/stores"
	"time"

	"github.com/google/uuid"
)

// UserRepository UserRepository
type UserRepository struct{}

// All All
func (u UserRepository) All() []models.User {
	users := stores.UserStore.All()

	return users
}

// GetByID GetByID
func (u UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user, err := stores.UserStore.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByName GetByName
func (u UserRepository) GetByName(name string) (*models.User, error) {
	user, err := stores.UserStore.GetByName(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Save Save
func (u UserRepository) Save(user *models.User) error {
	return nil
}

// Create Create
func (u UserRepository) Create(name string) (*models.User, error) {
	id, _ := uuid.NewRandom()
	user := models.User{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	stores.UserStore.Create(user)
	return &user, nil
}
