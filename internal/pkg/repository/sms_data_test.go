package repository

import (
	"fmt"
	"github.com/admsvist/go-diploma/country_codes"
	"os"
	"testing"
)

func TestGetAll(t *testing.T) {
	// создание временного файла
	codesFile, err := os.CreateTemp("", "codes.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(codesFile.Name())

	countryCodesBytes := []byte(`{"US": "United States", "BL": "Saint Barthelemy"}`)
	if _, err := codesFile.Write(countryCodesBytes); err != nil {
		t.Fatal(err)
	}
	if err := codesFile.Close(); err != nil {
		t.Fatal(err)
	}

	country_codes.Init(codesFile.Name())

	// создание временного файла
	smsDataFile, err := os.CreateTemp("", "sms.data")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(smsDataFile.Name())

	smsDataBytes := []byte(fmt.Sprintf("U5;41910;Topol\nUS;36;1576;Rond\nGB28495Topolo\nF2;9;484;Topolo\nBL;68;1594;Kildy"))
	if _, err := smsDataFile.Write(smsDataBytes); err != nil {
		t.Fatal(err)
	}
	if err := smsDataFile.Close(); err != nil {
		t.Fatal(err)
	}

	smsDataStorage := SMSDataRepository{
		Filename: smsDataFile.Name(),
	}

	entities, err := smsDataStorage.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(entities) != 2 {
		t.Errorf("The number of structures in the result set is different from 2")
	}
}
