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
		done       = make(chan bool, 1)
		die        = make(chan error, 1)
	)

	go func() {
		wg.Add(7)

		go func() {
			defer wg.Done()
			entities, e := GetSMSDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareSMSData(entities); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			entities, e := GetMMSDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareMMSData(entities); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			entities, e := GetVoiceCallDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareVoiceCallData(entities); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			entities, e := GetEmailDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareEmailData(entities); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			data, e := GetBillingData()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareBillingData(data); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			entities, e := GetSupportDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareSupportData(entities); e != nil {
				die <- e
			}
		}()

		go func() {
			defer wg.Done()
			entities, e := GetIncidentDataEntities()
			if e != nil {
				die <- e
			}
			mutex.Lock()
			defer mutex.Unlock()
			if e = resultSetT.PrepareIncidentData(entities); e != nil {
				die <- e
			}
		}()

		wg.Wait()

		close(done)
	}()

	select {
	case err := <-die:
		return nil, err
	case <-done:
		return &resultSetT, nil
	}
}
