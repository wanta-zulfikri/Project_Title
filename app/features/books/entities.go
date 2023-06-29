package books 

import (
	"github.com/labstack/echo/v4"
)

type Core struct { 
	ID            uint 
	Title         string 
	PublishedYear string 
	ISBN          string 
	Image         string 
	BookAuthors   []Book_Authorscore `gorm:"many2many:book_authors"`
} 

type Book_Authorscore struct {
	BookID  uint 
	UserID  uint  
} 


type Repository interface {
	CreateBook(books Core, userID uint) error
	GetBooks() ([]Core, error)
	GetBooksByUserID(userid uint) ([]Core, error)
	GetBook(bookid uint) (Core, error)
	UpdateBook(id uint, updatedBook Core) error
	DeleteBook(id uint) error
}

type Service interface {
	CreateBook(books Core, userID uint) error
	GetBooks() ([]Core, error)
	GetBooksByUserID(userid uint) ([]Core, error)
	GetBook(bookid uint) (Core, error)
	UpdateBook(id uint, updatedBook Core) error
	DeleteBook(id uint) error
}

type Handler interface {
	CreateBook() echo.HandlerFunc
	GetBooks() echo.HandlerFunc
	GetBooksByUserID() echo.HandlerFunc
	GetBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}