package model

import (
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"go.uber.org/zap"
)

func (u *UserDomain) CreateUser() *rest_errors.RestError {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	u.EncryptPassword()

	return nil
}
