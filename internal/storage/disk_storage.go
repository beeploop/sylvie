package storage

import (
	"io/fs"
	"os"
)

type DiskStorage struct {
	permission fs.FileMode
}

func NewDiskStorage() *DiskStorage {
	return &DiskStorage{
		permission: 0777,
	}
}

func (s *DiskStorage) Save(path string, data []byte) error {
	return os.WriteFile(path, data, s.permission)
}

func (s *DiskStorage) Open(path string) ([]byte, error) {
	return os.ReadFile(path)
}
