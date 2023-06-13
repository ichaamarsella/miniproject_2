package repositories

import (
	"gorm.io/gorm"
	"minpro2/entity"
)

type Admin struct {
	db *gorm.DB
}

func NewAdmin(dbCrud *gorm.DB) Admin {
	return Admin{
		db: dbCrud,
	}
}

type AdminInterfaceRepo interface {
	CreateAdmin(admin *entities.Admin) (*entities.Admin, error)
	GetAdminsById(id uint) (entities.Admin, error)
	UpdateAdmin(admin *entities.Admin) (any, error)
	DeleteAdmin(email string) (any, error)
	LoginByUsername(username string) (entities.Admin, error)
}

func (repo Admin) CreateAdmin(admin *entities.Admin) (*entities.Admin, error) {
	err := repo.db.Model(&entities.Admin{}).Create(admin).Error
	return admin, err
}

func (repo Admin) GetAdminsById(id uint) (entities.Admin, error) {
	var admin entities.Admin
	repo.db.First(&admin, "id = ? ", id)
	return admin, nil
}

// UpdateUser multiple fields
func (repo Admin) UpdateAdmin(admin *entities.Admin) (any, error) {
	err := repo.db.Save(admin).Error
	return nil, err
}

// DeleteUser by Id and email
func (repo Admin) DeleteAdmin(email string) (any, error) {
	err := repo.db.Model(&entities.Admin{}).
		Where("email = ?", email).
		Delete(&entities.Admin{}).
		Error
	return nil, err
}

func (repo Admin) LoginByUsername(username string) (entities.Admin, error) {
	var admin entities.Admin
	repo.db.First(&admin, "username = ? ", username)
	return admin, nil
}
