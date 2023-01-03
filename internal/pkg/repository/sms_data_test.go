package repository

import (
	"fmt"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/pkg/filereader"
	"testing"
)

func TestLoadData(t *testing.T) {
	countryCodesBytes := []byte(`{"US": "United States", "BL": "Saint Barthelemy"}`)
	countryCodesFakeReader := filereader.NewFakeFileReader(countryCodesBytes, nil)
	country_codes.Init(countryCodesFakeReader, "codes.json")

	smsData := []byte(fmt.Sprintf("U5;41910;Topol\nUS;36;1576;Rond\nGB28495Topolo\nF2;9;484;Topolo\nBL;68;1594;Kildy"))
	smsDataFakeReader := filereader.NewFakeFileReader(smsData, nil)
	smsDataStorage := NewSMSDataRepository()
	smsDataStorage.LoadData(smsDataFakeReader, "sms.data")

	if len(smsDataStorage.Data) != 2 {
		t.Errorf("The number of structures in the result set is different from 2")
	}
}
