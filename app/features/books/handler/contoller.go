package handler

import (
	"Project_Title/app/features/books"
	"Project_Title/helper"
	"math"
	"net/http"
	"strconv"

	"Project_Title/middleware"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	s books.Service
} 

func New(h books.Service) books.Handler {
	return &BookController{s: h}
}


func (ec *BookController) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateBook
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id := claims.ID

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 
		
		newBook := books.Core{
			ID: id,
			Title: input.Title,
			PublishedYear: input.PublishedYear,
			ISBN: input.ISBN,
	
		}

		err = ec.s.CreateBook(newBook, id)
		if err != nil {
			c.Logger().Error("Failed to create Book: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 

		response := BookResponse{
			Code:    http.StatusCreated,
			Message: "Success created an event",
			Data: BookData{
				Title:       newBook.Title,
				PublishedYear: newBook.PublishedYear,
				ISBN: newBook.ISBN,
			},
		}
		return c.JSON(http.StatusCreated, response)
	}
} 



func (ec *BookController) GetBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var books []books.Core
		var err error

		books, err = ec.s.GetBooks()
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}
		

		if len(books) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedBooks := []ResponseGetBooks{}
		for _, book := range books {
			formattedBook := ResponseGetBooks{
				Title:       book.Title,
				PublishedYear: book.PublishedYear,
				ISBN: book.ISBN,
			}
			formattedBooks = append(formattedBooks, formattedBook)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedBooks)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedBooks[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, BooksResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
} 


func (ec *BookController) GetBooksByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		userid := claims.ID
		books, err := ec.s.GetBooksByUserID(userid)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		if len(books) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedBooks := []ResponseGetBooks{}
		for _, book := range books {
			formattedBook := ResponseGetBooks{
				Title:         book.Title,
				PublishedYear: book.PublishedYear,
				ISBN:          book.ISBN,
			}
			formattedBooks = append(formattedBooks, formattedBook)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedBooks)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedBooks[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, BooksResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
}

func (ec *BookController) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		book, err := ec.s.GetBook(uint(bookid))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		response := ResponseGetBook{
				Title: book.Title,
				PublishedYear: book.PublishedYear,
				ISBN: book.ISBN,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    response,
		})
	}
}

func (ec *BookController) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateBook
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		// file, err := c.FormFile("event_picture")
		// var event_picture string
		// if err != nil && err != http.ErrMissingFile {
		// 	c.Logger().Error("Failed to get event_picture from form file: ", err)
		// 	return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		// } else if file != nil {
		// 	event_picture, err = helper.UploadImage(c, file)
		// 	if err != nil {
		// 		c.Logger().Error("Failed to upload event_picture: ", err)
		// 		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		// 	}
		// }

		updatedBook := books.Core{
			ID:            uint(id), 
			Title:         input.Title,
			PublishedYear: input.PublishedYear,
			ISBN:          input.ISBN,
			
		}

		err = ec.s.UpdateBook(updatedBook.ID, updatedBook)
		if err != nil {
			c.Logger().Error("Failed to update book: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := ResponseUpdateBooks{
			Title:         updatedBook.Title,
			PublishedYear: updatedBook.PublishedYear,
			ISBN:          updatedBook.ISBN,
		}
		

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Success updated an event.",
			Data:    response,
		})
	}
}

func (ec *BookController) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if claims.ID != uint(id) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is not valid for this user.", nil))
		}

		err = ec.s.DeleteBook(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an book", nil))
	}
}

