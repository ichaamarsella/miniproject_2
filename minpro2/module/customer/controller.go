package customer

import "minpro2/dto"

type ControllerCustomer interface {
	CreateCustomer(req CustomerParam) (any, error)
	GetCustomerById(id uint) (FindCustomer, error)
	UpdateCustomer(req CustomerParam, id uint) (any, error)
	DeleteCustomer(email string) (any, error)
}

type controllerCustomer struct {
	customerUseCase UseCaseCustomer
}

func (uc controllerCustomer) CreateCustomer(req CustomerParam) (any, error) {

	customer, err := uc.customerUseCase.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create customer",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) GetCustomerById(id uint) (FindCustomer, error) {
	var res FindCustomer
	customer, err := uc.customerUseCase.GetCustomerById(id)
	if err != nil {
		return FindCustomer{}, err
	}
	res.Data = customer
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get customer",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerCustomer) UpdateCustomer(req CustomerParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.customerUseCase.UpdateCustomer(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerCustomer) DeleteCustomer(email string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.customerUseCase.DeleteCustomer(email)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}
