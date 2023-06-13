package admin

import (
	"errors"
	"minpro2/dto"
	"minpro2/middleware"
	"time"
)

type ControllerAdmin interface {
	CreateAdmin(req AdminParam) (any, error)
	GetAdminById(id uint) (FindAdmin, error)
	UpdateAdmin(req AdminParam, id uint) (any, error)
	DeleteAdmin(email string) (any, error)
	LoginAdmin(admin LoginAdmin) (any, error)
}

type controllerAdmin struct {
	adminUseCase UseCaseAdmin
}

func (uc controllerAdmin) CreateAdmin(req AdminParam) (any, error) {

	admin, err := uc.adminUseCase.CreateAdmin(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create admin",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username: admin.Username,
			Password: admin.Password,
			RoleID:   admin.RoleID,
			Verified: admin.Verified,
			Active:   admin.Active,
		},
	}
	return res, nil
}

func (uc controllerAdmin) GetAdminById(id uint) (FindAdmin, error) {
	var res FindAdmin
	admin, err := uc.adminUseCase.GetAdminById(id)
	if err != nil {
		return FindAdmin{}, err
	}
	res.Data = admin
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get admin",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerAdmin) UpdateAdmin(req AdminParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.adminUseCase.UpdateAdmin(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerAdmin) DeleteAdmin(email string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.adminUseCase.DeleteAdmin(email)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}

func (uc controllerAdmin) LoginAdmin(admin LoginAdmin) (any, error) {
	var res dto.ResponseMeta

	login, err := uc.adminUseCase.LoginByUsername(admin.Username)
	if err != nil {
		return res, errors.New("declined")
	}
	if login.Password != admin.Password {
		return res, errors.New("declined")
	}
	if login.Verified != "true" {
		return res, errors.New("declined")
	}
	if login.Active != "true" {
		return res, errors.New("declined")
	}
	res = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Login!",
		ResponseTime: "",
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = admin.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the token expiration time (e.g., 24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(middleware.JwtSecret))
	if err != nil {
		return res, errors.New("failed to generate JWT token")
	}

	// Return the token in the response
	res.Token = tokenString
	// Other response data

	return res, nil
}
