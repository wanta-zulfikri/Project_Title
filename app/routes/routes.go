package routes

import (
	"Project_Title/app/features/books"
	"Project_Title/app/features/users"
	"Project_Title/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
) 


func Route(e *echo.Echo, uc users.Handler , ec books.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	
	//authentication 
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login()) 
	
	//users 
	e.GET("/users", uc.GetProfile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users", uc.GetProfile(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", uc.DeleteProfile(), middleware.JWT([]byte(config.JWTKey))) 

	//books 
	e.GET("/books", ec.GetBooks(), middleware.JWT([]byte(config.JWTKey))) 
	e.POST("/books", ec.CreateBook(), middleware.JWT([]byte(config.JWTKey))) 
	e.GET("/books/:id", ec.GetBook(), middleware.JWT([]byte(config.JWTKey))) 
	e.PUT("/books/:id", ec.UpdateBook(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/books/:id", ec.DeleteBook(), middleware.JWT([]byte(config.JWTKey)))
}