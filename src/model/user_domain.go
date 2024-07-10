package model

import "encoding/json"

type userDomain struct {
	ID       string
	Name     string
	Email    string
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
