package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
)

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &UserDomain{
		name, email, password, age,
	}
}

type UserDomain struct {
	Name     string
	Email    string
	Password string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestError
	UpdateUser() *rest_errors.RestError
	FindUser(string) (*UserDomain, *rest_errors.RestError)
	DeleteUser() *rest_errors.RestError
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))
}
