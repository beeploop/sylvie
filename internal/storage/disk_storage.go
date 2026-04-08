package storage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type DiskStorageConfig struct {
	BaseDir    string
	Permission os.FileMode
}

type diskStorage struct {
	BaseDir    string
	Permission os.FileMode
}

func NewDiskStorage(config DiskStorageConfig) *diskStorage {
	return &diskStorage{
		BaseDir:    config.BaseDir,
		Permission: config.Permission,
	}
}

func (s *diskStorage) Write(ctx context.Context, dest string, data []byte) (string, error) {
	// Create basedir and sub-directories of the dest
	subdir := filepath.Dir(dest)
	fullpath := filepath.Join(s.BaseDir, subdir)
	if err := os.MkdirAll(fullpath, s.Permission); err != nil {
		return "", err
	}

	pathname := filepath.Join(s.BaseDir, dest)
	if err := os.WriteFile(pathname, data, s.Permission); err != nil {
		return "", err
	}

	return pathname, nil
}

func (s *diskStorage) Read(ctx context.Context, pathname string) ([]byte, error) {
	fullpath := filepath.Join(s.BaseDir, pathname)

	_, err := os.Stat(fullpath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("file does not exist")
	}

	return os.ReadFile(fullpath)
}
