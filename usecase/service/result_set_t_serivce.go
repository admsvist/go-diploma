package service

import (
	"github.com/admsvist/go-diploma/entity"
	"sync"
)

func GetResultSetT() (*entity.ResultSetT, error) {
	var (
		resultSetT = entity.ResultSetT{}
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

	return &resultSetT, err
}
