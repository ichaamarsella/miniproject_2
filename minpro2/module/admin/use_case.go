package admin

import (
	"minpro2/entity"
	"minpro2/middleware"
	"minpro2/repositories"
	"time"
)

type UseCaseAdmin interface {
	CreateAdmin(admin AdminParam) (entities.Admin, error)
	GetAdminById(id uint) (entities.Admin, error)
	UpdateAdmin(admin AdminParam, id uint) (any, error)
	DeleteAdmin(email string) (any, error)
	LoginByUsername(username string) (entities.Admin, error)
}

type useCaseAdmin struct {
	adminRepo repositories.AdminInterfaceRepo
}

func (uc useCaseAdmin) CreateAdmin(admin AdminParam) (entities.Admin, error) {
	var newAdmin *entities.Admin

	// Hash the password
	hashedPassword, err := middleware.HashPassword(admin.Password)
	if err != nil {
		return entities.Admin{}, err
	}

	newAdmin = &entities.Admin{
		Username: admin.Username,
		Password: hashedPassword,
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err = uc.adminRepo.CreateAdmin(newAdmin)
	if err != nil {
		return *newAdmin, err
	}
	return *newAdmin, nil
}

func (uc useCaseAdmin) GetAdminById(id uint) (entities.Admin, error) {
	var admin entities.Admin
	admin, err := uc.adminRepo.GetAdminsById(id)
	return admin, err
}

func (uc useCaseAdmin) UpdateAdmin(admin AdminParam, id uint) (any, error) {
	var editAdmin *entities.Admin
	editAdmin = &entities.Admin{
		ID:       id,
		Username: admin.Username,
		Password: admin.Password,
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err := uc.adminRepo.UpdateAdmin(editAdmin)
	if err != nil {
		return *editAdmin, err
	}
	return *editAdmin, nil
}

func (uc useCaseAdmin) DeleteAdmin(email string) (any, error) {
	_, err := uc.adminRepo.DeleteAdmin(email)
	return nil, err
}

func (uc useCaseAdmin) LoginByUsername(username string) (entities.Admin, error) {
	var admin entities.Admin
	admin, err := uc.adminRepo.LoginByUsername(username)
	return admin, err
}
