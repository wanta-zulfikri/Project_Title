package routes

import (
	"Project_Title/app/features/books"
	"Project_Title/app/features/users"
	"Project_Title/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
) 


func Route(e *echo.Echo, uc users.Handler , ec books.Handler, config *config.Configuration) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	
	//authentication 
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login()) 
	
	//users 
	e.GET("/users", uc.GetProfile())
	e.PUT("/users", uc.GetProfile())
	e.DELETE("/users", uc.DeleteProfile()) 

	//books 
	e.GET("/books", ec.GetBooks()) 
	e.POST("/books", ec.CreateBook()) 
	e.GET("/books/:id", ec.GetBook()) 
	e.PUT("/books/:id", ec.UpdateBook())
	e.DELETE("/books/:id", ec.DeleteBook())
}