package stores

import (
	"apuroda/db"
	"apuroda/models"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserDB UserDB
type UserDB struct {
	db *sqlx.DB
}

// NewUserDB NewUserDB
func NewUserDB() *UserDB {
	u := &UserDB{
		db: db.GetDB(),
	}
	return u
}

// All All
func (u *UserDB) All() (users []models.User) {
	users = []models.User{}
	err := u.db.Select(&users, "SELECT id, name, created_at, updated_at FROM users")
	if err != nil {
		panic(err)
	}
	return
}

// Create Create
func (u *UserDB) Create(user models.User) {
	u.db.MustExec("INSERT INTO users (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)", &user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
}

// GetByID GetByID
func (u *UserDB) GetByID(id uuid.UUID) (user *models.User, err error) {
	user = new(models.User)
	err = u.db.Get(user, "SELECT id, name, created_at, updated_at FROM users WHERE id = $1 LIMIT 1", id.String())
	if err != nil {
		return nil, errors.New("not found")
	}
	return
}

// GetByName GetByName
func (u *UserDB) GetByName(name string) (user *models.User, err error) {
	user = new(models.User)
	err = u.db.Get(user, "SELECT id, name, created_at, updated_at FROM users WHERE name = $1 LIMIT 1", name)
	if err != nil {
		return nil, errors.New("not found")
	}
	return
}
