package repository

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title         string
	PublishedYear string
	ISBN          string
	BookAuthors   []Book_Authors `gorm:"many2many:book_authors"`
}

type Book_Authors struct {
	BookID uint
	UserID uint
}
