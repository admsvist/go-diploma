package repository

import (
	"encoding/json"
	"errors"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"net/http"
)

type SupportDataRepository struct {
	Url string
}

func (s *SupportDataRepository) GetAll() ([]*entity.SupportData, error) {
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

	// Декодировать JSON-массив в слайс структуры SupportData
	entities := make([]*entity.SupportData, 0)
	_ = json.Unmarshal(bytes, &entities)

	return entities, nil
}
