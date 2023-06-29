package repository  

import (

	"Project_Title/app/features/users"
	"Project_Title/helper"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
		return &UserRepository{db: db}
}

func (ar *UserRepository) Register(newUser users.Core)  (users.Core, error) {
		var input = User{}
		HashedPassword, err := helper.HashedPassword(newUser.Password)
		if err != nil {
				log.Println("Hashing password error", err.Error())
				return users.Core{}, err
		}

		input.Name = newUser.Name
		input.Email = newUser.Email
		input.Country = newUser.Country
		input.Password = HashedPassword

		if err := ar.db.Table("users").Create(&input).Error; err != nil {
				log.Println("Register error, email "+input.Email+"has been registered", err.Error())
				return users.Core{}, err
		}
		return newUser, nil 
}

func (ar *UserRepository) Login(email, password string) (users.Core, error) {
		var input User 
		if err := ar.db.Where("email = ?", email).Find(&input).Error; err != nil { 
				log.Println(err)
				return users.Core{}, errors.New("Email not found")
		}
		fmt.Println(input.Password,password)
		if err := helper.VerifyPassword(input.Password, password); err != nil {
				log.Println(err)
				return users.Core{}, errors.New("Invalid password")
		}
		
		return users.Core{ID: input.ID, Email: input.Email, Name: input.Name}, nil  
} 

func (ar *UserRepository) GetProfile(id uint) (users.Core, error) {
		var input User 
		result := ar.db.Where("id = ?", id).Find(&input)
		if result.Error != nil {
			return users.Core{}, result.Error
		}
		if result.RowsAffected == 0 {
				return users.Core{}, errors.New("user not found")
		}
		return users.Core{
			Name: input.Name,
			Email:  input.Email,
			Country: input.Country,
			Password: input.Password,

		}, nil
} 

func (ar *UserRepository) UpdateProfile(id uint, updatedUser users.Core) error {
	result := ar.db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{
		"name":       updatedUser.Name,
		"email":      updatedUser.Email,
		"country":    updatedUser.Country,
		"password":   updatedUser.Password,
		"updated_at": time.Now(),
	})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %v not found", id)
		}
		log.Print("Failed to update user", result.Error)
		return result.Error
	}
	return nil
}

func (ar *UserRepository) DeleteProfile(id uint) error {
	input := User{}
	if err := ar.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %v not found", id)
		}
		log.Print("Failed to query user by id", err)
		return err
	}

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := ar.db.Save(&input).Error; err != nil {
		log.Println("Terjadi error saat melakukan user buku", err)
		return err
	}
	return nil
}