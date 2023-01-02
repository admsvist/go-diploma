package sms

import (
	"fmt"
	"github.com/admsvist/go-diploma/internal/app/storage/codes"
	"github.com/admsvist/go-diploma/pkg/filereader"
	"testing"
)

func TestRead(t *testing.T) {
	codesBytes := []byte(`{"US": "United States", "BL": "Saint Barthelemy"}`)
	codesFakeReader := filereader.NewFakeFileReader(codesBytes, nil)
	codeStorage := codes.New()
	codeStorage.Read(codesFakeReader, "codes.json")

	smsData := []byte(fmt.Sprintf("U5;41910;Topol\nUS;36;1576;Rond\nGB28495Topolo\nF2;9;484;Topolo\nBL;68;1594;Kildy"))
	smsDataFakeReader := filereader.NewFakeFileReader(smsData, nil)
	smsDataStorage := New(codeStorage)
	smsDataStorage.Read(smsDataFakeReader, "sms.data")

	if len(smsDataStorage.Data) != 2 {
		t.Errorf("The number of structures in the result set is different from 2")
	}
}
