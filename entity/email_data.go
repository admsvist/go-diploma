package entity

import "github.com/admsvist/go-diploma/internal/pkg/sorts"

type EmailData struct {
	Country      string `json:"country"`       // alpha-2 — код страны;
	Provider     string `json:"provider"`      // провайдер;
	DeliveryTime int    `json:"delivery_time"` // среднее время доставки писем в миллисекундах.
}

type EmailDataSlice []*EmailData

func (s EmailDataSlice) Len() int {
	return len(s)
}

func (s EmailDataSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s EmailDataSlice) SortByDeliveryTime() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].DeliveryTime < s[j].DeliveryTime
	})
}
