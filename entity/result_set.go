package entity

import (
	"github.com/admsvist/go-diploma/country_codes"
)

type ResultSetT struct {
	SMS       [][]*SMSData              `json:"sms"`
	MMS       [][]*MMSData              `json:"mms"`
	VoiceCall []*VoiceCallData          `json:"voice_call"`
	Email     map[string][][]*EmailData `json:"email"`
	Billing   *BillingData              `json:"billing"`
	Support   []int                     `json:"support"`
	Incidents []*IncidentData           `json:"incident"`
}

func (t *ResultSetT) PrepareSMSData(entities []*SMSData) error {
	byProviderData := make([]*SMSData, len(entities))
	copy(byProviderData, entities)
	SMSDataSlice(byProviderData).SortByProvider()

	byCountryData := make([]*SMSData, len(entities))
	copy(byCountryData, entities)
	SMSDataSlice(byCountryData).SortByCountry()

	for i, v := range entities {
		entities[i].Country = country_codes.GetFullCountryName(v.Country)
	}

	t.SMS = append(make([][]*SMSData, 0), byProviderData, byCountryData)

	return nil
}

func (t *ResultSetT) PrepareMMSData(entities []*MMSData) error {
	byProviderData := make([]*MMSData, len(entities))
	copy(byProviderData, entities)
	MMSDataSlice(byProviderData).SortByProvider()

	byCountryData := make([]*MMSData, len(entities))
	copy(byCountryData, entities)
	MMSDataSlice(byCountryData).SortByCountry()

	for i, v := range entities {
		entities[i].Country = country_codes.GetFullCountryName(v.Country)
	}

	t.MMS = append(make([][]*MMSData, 0), byProviderData, byCountryData)

	return nil
}

func (t *ResultSetT) PrepareVoiceCallData(entities []*VoiceCallData) error {
	t.VoiceCall = entities

	return nil
}

func (t *ResultSetT) PrepareEmailData(entities []*EmailData) error {
	EmailDataSlice(entities).SortByDeliveryTime()

	sortedByCountry := make(map[string][]*EmailData, 0)
	for i, v := range entities {
		countryCode := v.Country
		sortedByCountry[countryCode] = append(sortedByCountry[countryCode], entities[i])
	}

	result := make(map[string][][]*EmailData, 0)
	for i, v := range sortedByCountry {
		result[i] = append(result[i], v[:3], v[len(v)-3:])
	}

	t.Email = result

	return nil
}

func (t *ResultSetT) PrepareBillingData(data *BillingData) error {
	t.Billing = data

	return nil
}

func (t *ResultSetT) PrepareSupportData(entities []*SupportData) error {
	ticketsCount := 0
	for _, data := range entities {
		ticketsCount += data.ActiveTickets
	}

	result := make([]int, 0)
	if ticketsCount < 9 {
		result = append(result, 1)
	} else if ticketsCount < 17 {
		result = append(result, 2)
	} else {
		result = append(result, 3)
	}

	t.Support = append(result, ticketsCount*(60/18))

	return nil
}

func (t *ResultSetT) PrepareIncidentData(entities []*IncidentData) error {
	IncidentDataSlice(entities).SortByStatus()

	t.Incidents = entities

	return nil
}
