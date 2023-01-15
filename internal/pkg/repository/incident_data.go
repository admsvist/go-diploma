package repository

import (
	"encoding/json"
	"errors"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"net/http"
)

type IncidentDataRepository struct {
	Url string
}

func (s *IncidentDataRepository) GetAll() ([]*entity.IncidentData, error) {
	// Отправить GET-запрос по указанному URL
	response, err := http.Get(s.Url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code")
	}

	// Прочитать содержимое ответа в байтовый срез
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Декодировать JSON-массив в слайс структуры IncidentData
	entities := make([]*entity.IncidentData, 0)
	_ = json.Unmarshal(bytes, &entities)

	return entities, nil
}
