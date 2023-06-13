package repositories

import (
	"gorm.io/gorm"
	"minpro2/entity"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}
}

type CustomerInterfaceRepo interface {
	CreateCustomer(user *entities.Customer) (*entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomer(user *entities.Customer) (any, error)
	DeleteCustomer(email string) (any, error)
}

// CreateCustomer new customer
func (repo Customer) CreateCustomer(user *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(user).Error
	return user, err
}

// GetCustomerById get single customer by id
func (repo Customer) GetCustomerById(id uint) (entities.Customer, error) {
	var user entities.Customer
	repo.db.First(&user, "id = ? ", id)
	return user, nil
}

// UpdateCustomer multiple fields
// func (repo Customer) UpdateCustomer(customer *entities.Customer) (any, error) {
func (repo Customer) UpdateCustomer(user *entities.Customer) (any, error) {
	err := repo.db.Save(user).Error
	return nil, err
}

// DeleteCustomer by Id and email
func (repo Customer) DeleteCustomer(email string) (any, error) {
	err := repo.db.Model(&entities.Customer{}).
		Where("email = ?", email).
		Delete(&entities.Customer{}).
		Error
	return nil, err
}
