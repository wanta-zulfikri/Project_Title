package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string
	PublishedYear string
	ISBN          string
	Users         []User `gorm:"many2many:book_authors"`
}


type Book_Authors struct {
	BookID uint
	UserID uint
}


type User struct {
	gorm.Model 
	Name     string   `json:"name" gorm:"type:varchar(100); not null"` 
	Country  string   `json:"country" gorm:"type:varchar(100); not null"` 
	Email    string   `json:"email" gorm:"type:varchar(100); not null"`
	Password string   `json:"password" gorm:"type:varchar(100); not null"`
}