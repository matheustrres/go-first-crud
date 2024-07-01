package service

import (
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
)

func (u *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_errors.RestError {
	return nil
}
