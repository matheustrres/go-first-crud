package main

import (
	"github.com/matheustrres/go-first-crud/src/controller"
	"github.com/matheustrres/go-first-crud/src/model/repository"
	"github.com/matheustrres/go-first-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(db *mongo.Database) controller.UserControllerInterface {
	repository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repository)

	return controller.NewUserControllerInterface(service)
}
