package service

import (
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/internal/pkg/repository"
)

const smsDataPath = "./../simulator/sms.data"
const voiceCallDataPath = "./../simulator/voice.data"
const emailDataPath = "./../simulator/email.data"
const billingDataPath = "./../simulator/billing.data"
const mmsUrl = "http://127.0.0.1:8383/mms"
const supportUrl = "http://127.0.0.1:8383/support"
const incidentUrl = "http://127.0.0.1:8383/accendent"

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
