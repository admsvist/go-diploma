package repository

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"log"
	"net/http"
)

type SupportDataRepository struct {
	Data []*entity.SupportData
}

func NewSupportDataRepository() *SupportDataRepository {
	return &SupportDataRepository{}
}

func (s *SupportDataRepository) LoadData(url string) {
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

	// Декодировать JSON-массив в слайс структуры SupportData
	_ = json.Unmarshal(bytes, &s.Data)
}
