package config 

import (
	users "Project_Title/app/features/users/repository"
	books "Project_Title/app/features/books/repository"

	"gorm.io/gorm"
)
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&users.User{},
		&books.Book{},

	
	)
	return err
}