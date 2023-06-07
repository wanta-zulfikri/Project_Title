package repository

import (
	"Project_Title/app/features/books"
	"errors"
	"time"

	"gorm.io/gorm"
)

type BooksRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (er *BooksRepository) CreateBook(book books.Core, userID uint) error {
	var err error 
	tx := er.db.Begin()

	newBook := Books {
	Title         :  book.Title,
	PublishedYear : book.PublishedYear,
	ISBN          : book.ISBN,

	}
	err = tx.Table("books").Create(&newBook).Error 
	if err != nil {
		tx.Rollback()
		return err 
	}
	return err
} 

func (er *BooksRepository) GetBooks() ([]books.Core, error) {
	var cores []books.Core
	if err := er.db.Table("books").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
} 


func (er *BooksRepository) GetBooksByUserID(userid uint) ([]books.Core, error) {
	var cores []books.Core
	if err := er.db.Table("books").Where("user_id = ? AND deleted_at IS NULL", userid).Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
} 

func (er *BooksRepository) GetBook(bookid uint) (books.Core, error) {
	var input Books
	result := er.db.Where("id = ? AND deleted_at IS NULL", bookid).Find(&input)
	if result.Error != nil {
		return books.Core{}, result.Error	
	}
	if result.RowsAffected == 0 { 
		return books.Core{}, result.Error
	}
		return books.Core{
			Title: input.Title,
			PublishedYear: input.PublishedYear,
			ISBN: input.ISBN,
			
		}, nil
}

func (er *BooksRepository) UpdateBook(id uint, updatedBook books.Core) error {
	if err := er.db.Model(&Books{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title"			    : updatedBook.Title, 
		"publishedyear"		: updatedBook.PublishedYear,
		"isbn"			    : updatedBook.ISBN,
		"updated_at"        : time.Now(),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil 
} 

func (er *BooksRepository) DeleteBook(id uint) error {
	input := Books{}
	if err := er.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid:true}

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil 
}