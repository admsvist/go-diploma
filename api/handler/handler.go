package handler

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/internal/pkg/repository"
	"net/http"
)

const smsDataPath = "./../simulator/sms.data"
const voiceCallDataPath = "./../simulator/voice.data"
const emailDataPath = "./../simulator/email.data"
const billingDataPath = "./../simulator/billing.data"
const mmsUrl = "http://127.0.0.1:8383/mms"
const supportUrl = "http://127.0.0.1:8383/support"
const incidentUrl = "http://127.0.0.1:8383/accendent"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	result := entity.ResultT{}

	entities, err := getResultData()
	if err != nil {
		result.Status = false
		result.Error = err.Error()
	} else {
		result.Status = true
		result.Data = entities
	}

	// сериализация сущностей в JSON
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// возврат ответа сервера
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getResultData() (*entity.ResultSetT, error) {
	sms, err := prepareSMSData()
	if err != nil {
		return nil, err
	}

	mms, err := prepareMMSData()
	if err != nil {
		return nil, err
	}

	voiceCall, err := prepareVoiceCallData()
	if err != nil {
		return nil, err
	}

	email, err := prepareEmailData()
	if err != nil {
		return nil, err
	}

	billing, err := prepareBillingData()
	if err != nil {
		return nil, err
	}

	support, err := prepareSupportData()
	if err != nil {
		return nil, err
	}

	incident, err := prepareIncidentData()
	if err != nil {
		return nil, err
	}

	return &entity.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: voiceCall,
		Email:     email,
		Billing:   billing,
		Support:   support,
		Incidents: incident,
	}, nil
}

func prepareSMSData() ([][]*entity.SMSData, error) {
	repo := repository.SMSDataRepository{
		Filename: smsDataPath,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	byProviderData := make([]*entity.SMSData, len(entities))
	copy(byProviderData, entities)
	entity.SMSDataSlice(byProviderData).SortByProvider()

	byCountryData := make([]*entity.SMSData, len(entities))
	copy(byCountryData, entities)
	entity.SMSDataSlice(byCountryData).SortByCountry()

	for i, v := range entities {
		entities[i].Сountry = country_codes.GetFullCountryName(v.Сountry)
	}

	result := make([][]*entity.SMSData, 0)
	result = append(result, byProviderData, byCountryData)

	return result, nil
}

func prepareMMSData() ([][]*entity.MMSData, error) {
	repo := repository.MMSDataRepository{
		Url: mmsUrl,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	byProviderData := make([]*entity.MMSData, len(entities))
	copy(byProviderData, entities)
	entity.MMSDataSlice(byProviderData).SortByProvider()

	byCountryData := make([]*entity.MMSData, len(entities))
	copy(byCountryData, entities)
	entity.MMSDataSlice(byCountryData).SortByCountry()

	for i, v := range entities {
		entities[i].Country = country_codes.GetFullCountryName(v.Country)
	}

	result := make([][]*entity.MMSData, 0)
	result = append(result, byProviderData, byCountryData)

	return result, nil
}

func prepareVoiceCallData() ([]*entity.VoiceCallData, error) {
	repo := repository.VoiceCallDataRepository{
		Filename: voiceCallDataPath,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func prepareEmailData() (map[string][][]*entity.EmailData, error) {
	repo := repository.EmailDataRepository{
		Filename: emailDataPath,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	entity.EmailDataSlice(entities).SortByDeliveryTime()

	sortedByCountry := make(map[string][]*entity.EmailData, 0)
	for i, v := range entities {
		countryCode := v.Country
		sortedByCountry[countryCode] = append(sortedByCountry[countryCode], entities[i])
	}

	result := make(map[string][][]*entity.EmailData, 0)
	for i, v := range sortedByCountry {
		result[i] = append(result[i], v[:3], v[len(v)-3:])
	}

	return result, nil
}

func prepareBillingData() (*entity.BillingData, error) {
	repo := repository.BillingDataRepository{
		Filename: billingDataPath,
	}

	entities, err := repo.Get()
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func prepareSupportData() ([]int, error) {
	repo := repository.SupportDataRepository{
		Url: supportUrl,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

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

	result = append(result, ticketsCount*(60/18))

	return result, nil
}

func prepareIncidentData() ([]*entity.IncidentData, error) {
	repo := repository.IncidentDataRepository{
		Url: incidentUrl,
	}

	entities, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	entity.IncidentDataSlice(entities).SortByStatus()

	return entities, nil
}
