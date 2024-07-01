package model

import (
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"go.uber.org/zap"
)

func (u *UserDomain) FindUser(id string) (*UserDomain, *rest_errors.RestError) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	return nil, nil
}
