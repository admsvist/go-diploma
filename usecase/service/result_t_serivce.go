package service

import (
	"github.com/admsvist/go-diploma/entity"
)

func GetResultT() *entity.ResultT {
	resultSetT, err := GetResultSetT()
	if err != nil {
		return &entity.ResultT{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		}
	}

	return &entity.ResultT{
		Status: true,
		Data:   resultSetT,
		Error:  "",
	}
}
