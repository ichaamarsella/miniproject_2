package repositories

import "miniproject_2/internal/models"

type CustomersRepository interface {
	GetCustomers() ([]*models.User, error)
	GetAdmins() ([]*models.User, error)
	FetchUserData(page int) ([]*models.User, error)
}
