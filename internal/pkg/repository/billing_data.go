package repository

import (
	"github.com/admsvist/go-diploma/entity"
	"os"
	"strconv"
)

type BillingDataRepository struct {
	Filename string
}

func (s *BillingDataRepository) Get() (*entity.BillingData, error) {
	// чтение файла
	bytes, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, err
	}

	value, _ := strconv.ParseInt(string(bytes), 2, 64)

	billingData := &entity.BillingData{
		CreateCustomer: value&(1<<0) != 0,
		Purchase:       value&(1<<1) != 0,
		Payout:         value&(1<<2) != 0,
		Recurring:      value&(1<<3) != 0,
		FraudControl:   value&(1<<4) != 0,
		CheckoutPage:   value&(1<<5) != 0,
	}

	return billingData, nil
}
