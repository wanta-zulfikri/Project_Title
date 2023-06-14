package main 

import (
	"Project_Title/config"
	"fmt"
	"Project_Title/app/routes"
	 userHandler"Project_Title/app/features/users/handler"
	 userRepo"Project_Title/app/features/users/repository"
	 userLogic"Project_Title/app/features/users/services" 
	 bookHandler"Project_Title/app/features/books/handler" 
	 bookRepo"Project_Title/app/features/books/repository"
	 bookLogic"Project_Title/app/features/books/services"


	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg) 
	config.Migrate(db)


	userModel := userRepo.New(db)
	userServices := userLogic.New(userModel)
	userController := userHandler.New(userServices, cfg)  

	bookModel := bookRepo.New(db) 
	bookServices := bookLogic.New(bookModel)
	bookController := bookHandler.New(bookServices, config.GetConfiguration())

	

	routes.Route(e, userController , bookController, cfg) 

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
} 

