package service

import (
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_errors.RestError {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}
