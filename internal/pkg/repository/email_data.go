package repository

import (
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/usecase/country_codes"
	"os"
	"strconv"
	"strings"
)

type EmailDataRepository struct {
	Filename string
}

func (s *EmailDataRepository) GetAll() ([]*entity.EmailData, error) {
	// чтение файла
	bytes, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, err
	}

	// создание списка сущностей
	entities := make([]*entity.EmailData, 0)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		params := strings.Split(line, ";")
		if len(params) != 3 {
			continue
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
