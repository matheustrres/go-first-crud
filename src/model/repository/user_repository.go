package repository

import (
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepositoryInterface {
	return &userRepositoryInterface{
		database,
	}
}

type userRepositoryInterface struct {
	database *mongo.Database
}

type UserRepositoryInterface interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestError)
}
