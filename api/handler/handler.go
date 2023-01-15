package handler

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/internal/pkg/repository"
	"net/http"
	"sync"
)

const smsDataPath = "./../simulator/sms.data"
const voiceCallDataPath = "./../simulator/voice.data"
const emailDataPath = "./../simulator/email.data"
const billingDataPath = "./../simulator/billing.data"
const mmsUrl = "http://127.0.0.1:8383/mms"
const supportUrl = "http://127.0.0.1:8383/support"
const incidentUrl = "http://127.0.0.1:8383/accendent"

var (
	resultSetT entity.ResultSetT
	die        error
	mutex      sync.Mutex
)

func handleError(e error) {
	die = e
}

func getResponse() entity.ResultT {
	var response entity.ResultT

	if die != nil {
		response.Status = false
		response.Data = nil
		response.Error = die.Error()
	} else {
		response.Status = true
		response.Data = &resultSetT
		response.Error = ""
	}

	return response
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(7)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sms, e := prepareSMSData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.SMS = sms
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		mms, e := prepareMMSData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.MMS = mms
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		voiceCall, e := prepareVoiceCallData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.VoiceCall = voiceCall
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		email, e := prepareEmailData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.Email = email
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		billing, e := prepareBillingData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.Billing = billing
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		support, e := prepareSupportData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.Support = support
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		incident, e := prepareIncidentData()
		mutex.Lock()
		defer mutex.Unlock()
		if e != nil {
			handleError(e)
		}
		resultSetT.Incidents = incident
	}(&wg)

	wg.Wait()

	// сериализация сущностей в JSON
	response := getResponse()
	data, e := json.Marshal(response)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// возврат ответа сервера
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
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
		entities[i].Country = country_codes.GetFullCountryName(v.Country)
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
