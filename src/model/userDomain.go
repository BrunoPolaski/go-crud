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

func (ud *UserDomain) EncryptPassword() (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *rest_err.RestErr
	UpdateUser(string, UserDomain) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
