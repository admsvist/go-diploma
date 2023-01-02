package codes

import (
	"encoding/json"
	"log"
	"os"
)

type CodeRepository struct {
	codes map[string]string
}

func New() *CodeRepository {
	content, err := os.ReadFile("./codes.json")
	if err != nil {
		log.Fatalln(err)
	}

	codes := make(map[string]string)

	if err := json.Unmarshal(content, &codes); err != nil {
		log.Fatalln(err)
	}

	return &CodeRepository{codes: codes}
}

func (r *CodeRepository) Contains(code string) bool {
	_, ok := r.codes[code]

	return ok
}
