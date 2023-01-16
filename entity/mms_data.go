package entity

import (
	"github.com/admsvist/go-diploma/pkg/sorts"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type MMSDataSlice []*MMSData

func (s MMSDataSlice) Len() int {
	return len(s)
}

func (s MMSDataSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MMSDataSlice) SortByCountry() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].Country < s[j].Country
	})
}

func (s MMSDataSlice) SortByProvider() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].Provider < s[j].Provider
	})
}
