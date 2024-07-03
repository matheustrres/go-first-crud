package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetJSONValue() (string, error)

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}
}

type userDomain struct {
	ID       string
	Name     string
	Email    string `json:"email"`
	Password string
	Age      int8
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) SetID(id string) {
	u.ID = id
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))
}
