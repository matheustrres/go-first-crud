package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheustrres/go-first-crud/src/controller"
	"github.com/matheustrres/go-first-crud/src/controller/routes"
	"github.com/matheustrres/go-first-crud/src/model/service"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	/**
	* Can use gin.New or gin.Default. Both are valid options but has its differences:
	*
	* - gin.New: initialize a router without any middleware
	* - gin.Default: initialize a router with with logger and (panic) recovery middlewares
	 */
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
