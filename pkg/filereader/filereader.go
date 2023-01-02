package filereader

import (
	"os"
)

type FileReader struct{}

func New() *FileReader {
	return &FileReader{}
}

func (r *FileReader) ReadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
