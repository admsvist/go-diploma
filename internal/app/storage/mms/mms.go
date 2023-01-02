package mms

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"log"
	"net/http"
)

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type CodeRepository interface {
	Contains(code string) bool
}

type MmsDataStorage struct {
	codeRepository CodeRepository
	Data           []*entity.MMSData
}

func New(codeRepository CodeRepository) *MmsDataStorage {
	return &MmsDataStorage{
		codeRepository: codeRepository,
	}
}

func (s *MmsDataStorage) Read(url string) {
	// Отправить GET-запрос по указанному URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return
	}

	// Прочитать содержимое ответа в байтовый срез
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Декодировать JSON-массив в слайс структуры MMSData
	err = json.Unmarshal(bytes, &s.Data)
	if err != nil {
		log.Fatalln(err)
	}

	//// Провалидировать каждую структуру User
	//for _, user := range users {
	//	if err := user.validate(); err != nil {
	//		return nil, err
	//	}
	//}
}
