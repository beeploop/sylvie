package storage

import "context"

type Storage interface {
	Write(ctx context.Context, dest string, data []byte) error
	Read(ctx context.Context, pathname string) ([]byte, error)
}
