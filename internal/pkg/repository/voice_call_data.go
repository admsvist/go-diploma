package repository

import (
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"log"
	"strconv"
	"strings"
)

type VoiceCallDataRepository struct {
	Data []*entity.VoiceCallData
}

func NewVoiceCallDataRepository() *VoiceCallDataRepository {
	return &VoiceCallDataRepository{}
}

func (s *VoiceCallDataRepository) LoadData(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

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

		s.Data = append(s.Data, &entity.VoiceCallData{
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
}
