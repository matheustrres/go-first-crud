package service

import (
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_errors.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
