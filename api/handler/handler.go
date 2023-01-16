package handler

import (
	"encoding/json"
	"github.com/admsvist/go-diploma/entity"
	"github.com/admsvist/go-diploma/internal/pkg/service"
	"net/http"
	"time"
)

var resultT *entity.ResultT

func TestHandler(w http.ResponseWriter, r *http.Request) {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			resultT = nil
		}
	}()

	// заполняем объект
	if resultT == nil {
		resultT = &entity.ResultT{
			Data: &entity.ResultSetT{},
		}
		service.Fill(resultT)
	}

	// сериализация сущностей в JSON
	data, e := json.Marshal(resultT)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// возврат ответа сервера
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}
