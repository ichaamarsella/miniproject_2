package customer

import (
	"minpro2/entity"
	"minpro2/repositories"
	"time"
)

type UseCaseCustomer interface {
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomer(customer CustomerParam, id uint) (any, error)
	DeleteCustomer(email string) (any, error)
}

type useCaseCustomer struct {
	customerRepo repositories.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) GetCustomerById(id uint) (entities.Customer, error) {
	var customer entities.Customer
	customer, err := uc.customerRepo.GetCustomerById(id)
	return customer, err
}

func (uc useCaseCustomer) UpdateCustomer(customer CustomerParam, id uint) (any, error) {
	var editCustomer *entities.Customer
	editCustomer = &entities.Customer{
		ID:        id,
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.customerRepo.UpdateCustomer(editCustomer)
	if err != nil {
		return *editCustomer, err
	}
	return *editCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomer(email string) (any, error) {
	_, err := uc.customerRepo.DeleteCustomer(email)
	return nil, err
}
