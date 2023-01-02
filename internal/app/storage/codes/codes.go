package codes

import (
	"encoding/json"
	"log"
)

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type CodeStorage struct {
	codes map[string]string
}

func New() *CodeStorage {
	return &CodeStorage{}
}

func (r *CodeStorage) Read(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(bytes, &r.codes); err != nil {
		log.Fatalln(err)
	}
}

func (r *CodeStorage) Contains(code string) bool {
	_, ok := r.codes[code]

	return ok
}
