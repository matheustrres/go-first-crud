package model

import (
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"go.uber.org/zap"
)

func (u *UserDomain) UpdateUser() *rest_errors.RestError {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
