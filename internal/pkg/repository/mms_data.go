package repository

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"io"
	"log"
	"net/http"
)

type MMSDataRepository struct {
	Data []*entity.MMSData
}

func NewMMSDataRepository() *MMSDataRepository {
	return &MMSDataRepository{}
}

func (s *MMSDataRepository) LoadData(url string) {
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
	data := []*entity.MMSData{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalln(err)
	}

	// Провалидировать каждую структуру MMSData
	for i, mmsData := range data {
		if !country_codes.Exists(mmsData.Country) {
			continue
		}

		if !contains([]string{"Topolo", "Rond", "Kildy"}, mmsData.Provider) {
			continue
		}

		s.Data = append(s.Data, data[i])
	}
}
