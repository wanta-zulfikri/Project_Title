package services

import (
	"Project_Title/app/features/users"
	"Project_Title/helper"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserService struct {
	m users.Repository
}

func New(r users.Repository) users.Service {
	return &UserService{m: r}
}

func (us *UserService) Register(newUser users.Core) error {
	_, err := us.m.Register(newUser)
	if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

func (us *UserService) Login(email string, password string) (users.Core, error) {
	user, err := us.m.Login(email, password)
	if err != nil {
		return users.Core{},err
	}
	return user, nil
} 

func (us *UserService) GetProfile(id uint) (users.Core, error) {
	user, err := us.m.GetProfile(id)
	if err != nil {
		return users.Core{},err
	}
	return user, nil
} 

func (us *UserService) UpdateProfile(id uint, updateUser users.Core) error {
	hashedPassword, err := helper.HashedPassword(updateUser.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	updatedUser := users.Core{
		Name: updateUser.Name,
		Email:    updateUser.Email,
		Country:    updateUser.Country,
		Password: string(hashedPassword),
	}
	if err := us.m.UpdateProfile(id, updatedUser); err != nil {
		return fmt.Errorf("Error while updating %d: %v", id, err)
	}
	return nil
	
}

func (us *UserService) DeleteProfile(id uint) error {
	if id == 0 {
		return fmt.Errorf("Invalid email")
	}
	err := us.m.DeleteProfile(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Email: %v ,not found", id)
		}
		log.Printf("Error while deleting %d: %v", id, err)
		return fmt.Errorf("Terjadi masalah pada server")
	}
	return nil
}