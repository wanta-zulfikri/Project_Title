package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model 
	Name     string   `json:"name" gorm:"type:varchar(100); not null"` 
	Country  string   `json:"country" gorm:"type:varchar(100); not null"` 
	Email    string   `json:"email" gorm:"type:varchar(100); not null"`
	Password string   `json:"password" gorm:"type:varchar(100); not null"`
}