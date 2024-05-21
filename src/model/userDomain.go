package model

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *UserDomain) EncryptPassword() error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		ud.Password = string(hash)
		return nil
	}
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
