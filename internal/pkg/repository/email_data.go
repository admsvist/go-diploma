package repository

import (
	"encoding/csv"
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/usecase/country_codes"
	"os"
	"strconv"
)

type EmailDataRepository struct {
	Filename string
}

func (s *EmailDataRepository) GetAll() ([]*entity.EmailData, error) {
	// чтение файла
	file, err := os.Open(s.Filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = 3

	// создание списка сущностей
	entities := make([]*entity.EmailData, 0)
	for {
		params, err := reader.Read()
		if err != nil {
			break
		}

		country := params[0]
		if !country_codes.Exists(country) {
			continue
		}

		validProviders := []string{
			"Gmail",
			"Yahoo",
			"Hotmail",
			"MSN",
			"Orange",
			"Comcast",
			"AOL",
			"Live",
			"RediffMail",
			"GMX",
			"Proton Mail",
			"Yandex",
			"Mail.ru",
		}
		provider := params[1]
		if !contains(validProviders, provider) {
			continue
		}

		deliveryTime, _ := strconv.Atoi(params[2])

		entities = append(entities, &entity.EmailData{
			Country:      country,
			Provider:     provider,
			DeliveryTime: deliveryTime,
		})
	}

	return entities, nil
}
