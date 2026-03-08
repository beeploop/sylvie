package storage

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskStorage(t *testing.T) {
	t.Run("test read and write", func(t *testing.T) {
		tests := []struct {
			Name     string
			Input    []byte
			Filepath string
			Expected []byte
		}{
			{
				Name:     "test hello world content",
				Input:    []byte("hello world"),
				Filepath: "hello_world.txt",
				Expected: []byte("hello world"),
			},
			{
				Name:     "test empty content",
				Input:    []byte{},
				Filepath: "empty.txt",
				Expected: []byte{},
			},
			{
				Name:     "test nested path",
				Input:    []byte("nested path content"),
				Filepath: "a/b/c/file.txt",
				Expected: []byte("nested path content"),
			},
		}

		s := NewDiskStorage(DiskStorageConfig{
			BaseDir:    "./tmp",
			Permission: os.FileMode(0777),
		})

		for _, tc := range tests {
			t.Run(tc.Name, func(t *testing.T) {
				err := s.Write(context.Background(), tc.Filepath, tc.Input)
				assert.NoError(t, err)

				b, err := s.Read(context.Background(), tc.Filepath)
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, b)
			})
		}
	})
}
