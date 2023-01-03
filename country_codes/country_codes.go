package country_codes

import (
	"encoding/json"
	"log"
)

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

var countryCodes map[string]string

func Init(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(bytes, &countryCodes); err != nil {
		log.Fatalln(err)
	}
}

func Exists(code string) bool {
	if len(countryCodes) == 0 {
		log.Fatalln("country codes not loaded")
	}

	_, ok := countryCodes[code]

	return ok
}
