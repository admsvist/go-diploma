package repository

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"log"
	"net/http"
)

type IncidentDataRepository struct {
	Data []*entity.IncidentData
}

func NewIncidentDataRepository() *IncidentDataRepository {
	return &IncidentDataRepository{}
}

func (s *IncidentDataRepository) LoadData(url string) {
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

	// Декодировать JSON-массив в слайс структуры IncidentData
	_ = json.Unmarshal(bytes, &s.Data)
}
