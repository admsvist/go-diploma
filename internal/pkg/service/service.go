package service

import (
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/internal/pkg/repository"
	"sync"
)

const smsDataPath = "./../simulator/sms.data"
const voiceCallDataPath = "./../simulator/voice.data"
const emailDataPath = "./../simulator/email.data"
const billingDataPath = "./../simulator/billing.data"
const mmsUrl = "http://127.0.0.1:8383/mms"
const supportUrl = "http://127.0.0.1:8383/support"
const incidentUrl = "http://127.0.0.1:8383/accendent"

func Fill(t *entity.ResultT) {
	var (
		resultSetT = &entity.ResultSetT{}
		wg         sync.WaitGroup
		mutex      sync.Mutex
		err        error
	)

	wg.Add(7)

	go func() {
		defer wg.Done()
		entities, e := GetSMSDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareSMSData(entities); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		entities, e := GetMMSDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareMMSData(entities); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		entities, e := GetVoiceCallDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareVoiceCallData(entities); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		entities, e := GetEmailDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareEmailData(entities); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		data, e := GetBillingData()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareBillingData(data); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		entities, e := GetSupportDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareSupportData(entities); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		entities, e := GetIncidentDataEntities()
		if e != nil {
			err = e
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		if e = resultSetT.PrepareIncidentData(entities); e != nil {
			err = e
		}
	}()

	wg.Wait()

	if err != nil {
		t.Status = false
		t.Data = nil
		t.Error = err.Error()
	} else {
		t.Status = true
		t.Data = resultSetT
		t.Error = ""
	}
}

func GetSMSDataEntities() ([]*entity.SMSData, error) {
	repo := repository.SMSDataRepository{
		Filename: smsDataPath,
	}

	return repo.GetAll()
}

func GetMMSDataEntities() ([]*entity.MMSData, error) {
	repo := repository.MMSDataRepository{
		Url: mmsUrl,
	}

	return repo.GetAll()
}

func GetVoiceCallDataEntities() ([]*entity.VoiceCallData, error) {
	repo := repository.VoiceCallDataRepository{
		Filename: voiceCallDataPath,
	}

	return repo.GetAll()
}

func GetEmailDataEntities() ([]*entity.EmailData, error) {
	repo := repository.EmailDataRepository{
		Filename: emailDataPath,
	}

	return repo.GetAll()
}

func GetBillingData() (*entity.BillingData, error) {
	repo := repository.BillingDataRepository{
		Filename: billingDataPath,
	}

	return repo.Get()
}

func GetSupportDataEntities() ([]*entity.SupportData, error) {
	repo := repository.SupportDataRepository{
		Url: supportUrl,
	}

	return repo.GetAll()
}

func GetIncidentDataEntities() ([]*entity.IncidentData, error) {
	repo := repository.IncidentDataRepository{
		Url: incidentUrl,
	}

	return repo.GetAll()
}
