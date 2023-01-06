package repository

import (
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"os"
	"strings"
)

type SMSDataRepository struct {
	Filename string
}

func (s *SMSDataRepository) GetAll() ([]*entity.SMSData, error) {
	// чтение файла
	bytes, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, err
	}

	// создание списка сущностей
	entities := make([]*entity.SMSData, 0)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		params := strings.Split(line, ";")
		if len(params) != 4 {
			continue
		}

		country := params[0]
		if !country_codes.Exists(country) {
			continue
		}

		provider := params[3]
		if !contains([]string{"Topolo", "Rond", "Kildy"}, provider) {
			continue
		}

		entities = append(entities, &entity.SMSData{
			Сountry:      country,
			Bandwidth:    params[1],
			ResponseTime: params[2],
			Provider:     provider,
		})
	}

	return entities, nil
}
