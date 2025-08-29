package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/beeploop/sylvie/internal/metadata"
)

type DiskRepository struct {
	outDir   string
	filename string
}

func NewDiskRepository(outDir string) *DiskRepository {
	return &DiskRepository{
		outDir:   outDir,
		filename: "result.json",
	}
}

func (r *DiskRepository) Save(data metadata.Metadata) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(r.filename)
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
