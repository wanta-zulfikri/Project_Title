package services

import (
	"Project_Title/app/features/books"
	"errors"

	"gorm.io/gorm"
)

type BooksService struct {
	r books.Repository
}

func New(r books.Repository) books.Service {
	return &BooksService{r: r}
}

func (s *BooksService) CreateBook(book books.Core, userID uint) error {
	err := s.r.CreateBook(book, userID)
	if err != nil {
		return err
	}
	return nil 
} 

func (es *BooksService) GetBooks() ([]books.Core, error) {
	lelangs, err := es.r.GetBooks()
	if err != nil {
		return nil, err
	}
	return lelangs, nil
} 

func (es *BooksService) GetBooksByUserID(userid uint) ([]books.Core, error) {
	lelangs, err := es.r.GetBooksByUserID(userid)
	if err != nil {
		return nil, err 
	}
	return lelangs, nil
} 

func (es *BooksService) GetBook(bookid uint) (books.Core, error) {
	book, err := es.r.GetBook(bookid)
	if err != nil {
		return books.Core{}, err
	}
	return book, nil
}

func (es *BooksService) UpdateBook(id uint, updatedBook books.Core) error {
	updatedBook.ID = id 
	if err := es.r.UpdateBook(id, updatedBook); err != nil {
		return nil 
	}
	return nil
}

func (es *BooksService) DeleteBook(id uint) error {
	err := es.r.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil
}