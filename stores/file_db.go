package stores

import (
	"apuroda/db"
	"apuroda/models"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// FileDB FileDB
type FileDB struct {
	db *sqlx.DB
}

// NewFileDB NewFileDB
func NewFileDB() *FileDB {
	f := &FileDB{
		db: db.GetDB(),
	}
	return f
}

// All All
func (f *FileDB) All() (files []models.File) {
	files = []models.File{}
	err := f.db.Select(&files, "SELECT id, name, created_at, updated_at FROM files;")
	if err != nil {
		panic(err)
	}
	return
}

// Create Create
func (f *FileDB) Create(file models.File) {
	f.db.MustExec("INSERT INTO files (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)", &file.ID, &file.Name, &file.CreatedAt, &file.UpdatedAt)
}

// GetByID GetByID
func (f *FileDB) GetByID(id uuid.UUID) (file *models.File, err error) {
	file = new(models.File)
	err = f.db.Get(file, "SELECT id, name, created_at, updated_at FROM files WHERE id = $1 LIMIT 1", id.String())
	if err != nil {
		return nil, errors.New("not found")
	}
	return
}

// GetByName GetByName
func (f *FileDB) GetByName(name string) (file *models.File, err error) {
	file = new(models.File)
	err = f.db.Get(file, "SELECT id, name, created_at, updated_at FROM files WHERE name = $1 LIMIT 1", name)
	if err != nil {
		return nil, errors.New("not found")
	}
	return
}
