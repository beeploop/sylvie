package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/beeploop/sylvie/internal/metadata"
)

type DiskRepository struct {
	outDir string
}

func NewDiskRepository(outDir string) *DiskRepository {
	if err := os.MkdirAll(outDir, 0777); err != nil {
		log.Fatal(err.Error())
	}

	return &DiskRepository{
		outDir: outDir,
	}
}

func (r *DiskRepository) Save(data metadata.Metadata) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", data.VideoID)
	f, err := os.Create(filepath.Join(r.outDir, filename))
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(b)
	if err != nil {
		return err
	}

	if n != len(b) {
		err := fmt.Errorf("Amount of data written is does not equal to amount of data to write")
		return err
	}

	log.Printf("written %d bytes\n", n)
	return nil
}
