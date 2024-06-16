package model

import "github.com/BrunoPolaski/go-crud/src/configuration/rest_err"

type UserDomainInterface interface {
	ComparePassword(password string) error
	EncryptPassword() error
	GenerateToken() (string, *rest_err.RestErr)
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string
	SetID(string)
	SetEmail(string)
	SetPassword(string)
	SetName(string)
	SetAge(int8)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
