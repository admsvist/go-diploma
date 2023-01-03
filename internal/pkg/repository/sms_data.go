package repository

import (
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"log"
	"strings"
)

type SMSDataRepository struct {
	Data []*entity.SMSData
}

func NewSMSDataRepository() *SMSDataRepository {
	return &SMSDataRepository{}
}

func (s *SMSDataRepository) LoadData(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

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

		s.Data = append(s.Data, &entity.SMSData{
			Сountry:      country,
			Bandwidth:    params[1],
			ResponseTime: params[2],
			Provider:     provider,
		})
	}
}
