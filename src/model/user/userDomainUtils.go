package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomain) EncryptPassword() error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		ud.password = string(hash)
		return nil
	}
}

func (ud *userDomain) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(ud.password), []byte(password)); err != nil {
		return err
	} else {
		return nil
	}
}

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenWithSecret, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error trying to create jwt,err: %s", err.Error()))
	}

	return tokenWithSecret, nil
}

func VerifyToken(token string) (UserDomainInterface, *rest_err.RestErr)

func (ud *userDomain) RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
