package repository

import (
	"context"
	"os"

	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
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

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	result, err := usersCollec.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
