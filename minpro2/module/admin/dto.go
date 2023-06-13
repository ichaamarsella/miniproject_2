package admin

import (
	"minpro2/dto"
	"minpro2/entity"
)

type AdminParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data AdminParam `json:"data"`
}

type FindAdmin struct {
	dto.ResponseMeta
	Data entities.Admin `json:"data"`
}

type LoginAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
