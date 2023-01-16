package repository

import (
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/usecase/country_codes"
	"os"
	"strconv"
	"strings"
)

type VoiceCallDataRepository struct {
	Filename string
}

func (s *VoiceCallDataRepository) GetAll() ([]*entity.VoiceCallData, error) {
	// чтение файла
	bytes, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, err
	}

	// создание списка сущностей
	entities := make([]*entity.VoiceCallData, 0)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		params := strings.Split(line, ";")
		if len(params) != 8 {
			continue
		}

		country := params[0]
		if !country_codes.Exists(country) {
			continue
		}

		provider := params[3]
		if !contains([]string{"TransparentCalls", "E-Voice", "JustPhone"}, provider) {
			continue
		}

		bandwidth, _ := strconv.Atoi(params[1])
		responseTime, _ := strconv.Atoi(params[2])
		connectionStability, _ := strconv.ParseFloat(params[4], 32)
		tTFB, _ := strconv.Atoi(params[5])
		voicePurity, _ := strconv.Atoi(params[6])
		medianOfCallsTime, _ := strconv.Atoi(params[7])

		entities = append(entities, &entity.VoiceCallData{
			Country:             country,
			Bandwidth:           bandwidth,
			ResponseTime:        responseTime,
			Provider:            provider,
			ConnectionStability: float32(connectionStability),
			TTFB:                tTFB,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		})
	}

	return entities, err
}
