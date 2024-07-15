package model

import "encoding/json"

type userDomain struct {
	id       string
	name     string
	email    string
	password string
	age      int8
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) GetID() string {
	return u.id
}

func (u *userDomain) SetID(id string) {
	u.id = id
}
