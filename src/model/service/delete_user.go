package service

import (
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUser(userId string) *rest_errors.RestError {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
