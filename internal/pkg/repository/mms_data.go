package repository

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"net/http"
)

type MMSDataRepository struct {
	Url string
}

func (s *MMSDataRepository) GetAll() ([]*entity.MMSData, error) {
	// Отправить GET-запрос по указанному URL
	response, err := http.Get(s.Url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	// Прочитать содержимое ответа в байтовый срез
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Декодировать JSON-массив в слайс структуры MMSData
	data := make([]*entity.MMSData, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	// Провалидировать каждую структуру MMSData
	entities := make([]*entity.MMSData, 0)
	for i, mmsData := range data {
		if !country_codes.Exists(mmsData.Country) {
			continue
		}

		if !contains([]string{"Topolo", "Rond", "Kildy"}, mmsData.Provider) {
			continue
		}

		entities = append(entities, data[i])
	}

	return entities, nil
}
