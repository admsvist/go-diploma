package entity

import "github.com/admsvist/go-diploma/internal/pkg/sorts"

type SMSData struct {
	Сountry      string `json:"сountry"`      // alpha-2 — код страны;
	Bandwidth    string `json:"bandwidth"`    // пропускная способность канала от 0 до 100%;
	ResponseTime string `json:"responseTime"` // среднее время ответа в миллисекундах;
	Provider     string `json:"provider"`     // название компании-провайдера.
}

type SMSDataSlice []*SMSData

func (s SMSDataSlice) Len() int {
	return len(s)
}

func (s SMSDataSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SMSDataSlice) SortByCountry() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].Сountry < s[j].Сountry
	})
}

func (s SMSDataSlice) SortByProvider() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].Provider < s[j].Provider
	})
}
