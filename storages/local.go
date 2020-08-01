package storages

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// LocalStorage LocalStorage
type LocalStorage struct {
	basePath string
}

// LegacyLocalStorage LegacyLocalStorage
var LegacyLocalStorage = NewLocalStorage("./tmp")

// NewLocalStorage NewLocalStorage
func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: filepath.Clean(basePath),
	}
}

// Set Set
func (u *LocalStorage) Set(key string, binary io.Reader) error {
	file, err := os.OpenFile(filepath.Join(u.basePath, key), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, binary)

	return nil
}

// Get Get
func (u *LocalStorage) Get(key string) (io.Reader, error) {
	file, err := os.OpenFile(filepath.Join(u.basePath, key), os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	stdout := new(bytes.Buffer)
	io.Copy(stdout, file)

	return stdout, nil
}
