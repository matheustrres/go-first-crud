package service

import (
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/model"
	"github.com/matheustrres/go-first-crud/src/model/repository"
)

func NewUserDomainService(repository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{
		userRepository: repository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestError)
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
