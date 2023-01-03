package repository

import (
	"github.com/admsvist/go-diploma/entity"
	"log"
	"strconv"
	"strings"
)

type BillingDataRepository struct {
	Data []*entity.BillingData
}

func NewBillingDataRepository() *BillingDataRepository {
	return &BillingDataRepository{}
}

func (s *BillingDataRepository) LoadData(reader FileReader, path string) {
	bytes, err := reader.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		value, _ := strconv.ParseInt(line, 2, 64)

		s.Data = append(s.Data, &entity.BillingData{
			CreateCustomer: value&(1<<0) != 0,
			Purchase:       value&(1<<1) != 0,
			Payout:         value&(1<<2) != 0,
			Recurring:      value&(1<<3) != 0,
			FraudControl:   value&(1<<4) != 0,
			CheckoutPage:   value&(1<<5) != 0,
		})
	}
}
