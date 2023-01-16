package entity

import (
	"github.com/admsvist/go-diploma/pkg/sorts"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

type IncidentDataSlice []*IncidentData

func (s IncidentDataSlice) Len() int {
	return len(s)
}

func (s IncidentDataSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IncidentDataSlice) SortByStatus() {
	sorts.SelectionSort(s, func(i, j int) bool {
		return s[i].Status < s[j].Status
	})
}
