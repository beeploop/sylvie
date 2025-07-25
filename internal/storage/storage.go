package storage

type Storage interface {
	Save(path string, data []byte) error
	Open(path string) ([]byte, error)
}
