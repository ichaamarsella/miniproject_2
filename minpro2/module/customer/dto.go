package customer

import (
	"minpro2/dto"
	"minpro2/entity"
)

type CustomerParam struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entities.Customer `json:"data"`
}
