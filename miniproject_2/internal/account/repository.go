package account

import (
	"miniproject_2/internal/models"
)

// AccountRepository defines the interface for the account repository.
type AccountRepository interface {
	RegisterCustomer(customer *models.User) error
	RegisterAdmin(admin *models.User) error
	ApproveAdmin(admin *models.User) error
	GetApprovalRequests() ([]*models.User, error)
	FindAdminByUsername(username string) (*models.User, error)
	RemoveCustomer(userID int) error
	RemoveAdmin(userID int) error
	ActivateDeactivateAdmin(userID int, active bool) error
}

// MySQLAccountRepository is an implementation of the AccountRepository interface for MySQL.
type MySQLAccountRepository struct {
	// Include your MySQL database connection or any other dependencies here
}

// NewMySQLAccountRepository creates a new instance of MySQLAccountRepository.
func NewMySQLAccountRepository() *MySQLAccountRepository {
	// Initialize and return a new instance of MySQLAccountRepository
}

// Implement the methods of the AccountRepository interface for MySQLAccountRepository.
// ...

// Alternatively, you can have other repository implementations for different databases or storage mechanisms.
