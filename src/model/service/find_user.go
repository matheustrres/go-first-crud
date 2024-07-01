package service

import (
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
)

func (u *userDomainService) FindUser(id string) (
	*model.UserDomainInterface,
	*rest_errors.RestError,
) {
	return nil, nil
}
