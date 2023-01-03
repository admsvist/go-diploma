package repository

import (
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"log"
	"strconv"
	"strings"
)

type EmailDataRepository struct {
	Data []*entity.EmailData
}

func NewEmailDataRepository() *EmailDataRepository {
	return &EmailDataRepository{}
}

func (s *EmailDataRepository) LoadData(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

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

		s.Data = append(s.Data, &entity.EmailData{
			Country:      country,
			Provider:     provider,
			DeliveryTime: deliveryTime,
		})
	}
}
