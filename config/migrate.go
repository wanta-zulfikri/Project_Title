package config 

import (
	users "Project_Title/app/features/users/repository"

	"gorm.io/gorm"
)
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&users.User{},
	
	)
	return err
}