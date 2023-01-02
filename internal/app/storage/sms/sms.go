package sms

import (
	"github.com/mishut/go-diploma/entity"
	"github.com/mishut/go-diploma/internal/app/repository/codes"
	"log"
	"os"
	"strings"
)

type CodeRepository interface {
	Contains(code string) bool
}

const filePath = "./../simulator/sms.data"

type SmsDataStorage struct {
	codeRepository CodeRepository
	Data           []*entity.SMSData
}

func New() *SmsDataStorage {
	return &SmsDataStorage{
		codeRepository: codes.New(),
	}
}

func (s *SmsDataStorage) Read() {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

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
