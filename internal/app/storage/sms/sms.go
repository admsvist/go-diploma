package sms

import (
	"github.com/admsvist/go-diploma/entity"
	"log"
	"strings"
)

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type CodeRepository interface {
	Contains(code string) bool
}

type SmsDataStorage struct {
	codeRepository CodeRepository
	Data           []*entity.SMSData
}

func New(codeRepository CodeRepository) *SmsDataStorage {
	return &SmsDataStorage{
		codeRepository: codeRepository,
	}
}

func (s *SmsDataStorage) Read(reader FileReader, path string) {
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
		if !s.codeRepository.Contains(country) {
			continue
		}

		provider := params[3]
		if !contains(getSupportedProviders(), provider) {
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

func getSupportedProviders() []string {
	return []string{"Topolo", "Rond", "Kildy"}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
