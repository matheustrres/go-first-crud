package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheustrres/go-first-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	* Can use gin.New or gin.Default. Both are valid options but has its differences:
	*
	* - gin.New: initialize a router without any middleware
	* - gin.Default: initialize a router with with logger and (panic) recovery middlewares
	 */
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
