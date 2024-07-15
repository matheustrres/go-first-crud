package repository

import (
	"context"
	"os"

	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
	"github.com/matheustrres/go-first-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_USERS_COLLEC = "MONGODB_USERS_COLLEC"
)

func (ur *userRepositoryInterface) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init CreateUserRepository")

	collec_name := os.Getenv(MONGODB_USERS_COLLEC)

	usersCollec := ur.database.Collection(collec_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := usersCollec.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
