package middlewares

import (
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET")
	tokenValue := model.RemoveBearerPrefix(c.Request.Header.Get("Authorization"))
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		errRest := rest_err.NewBadRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code, errRest)
		return
	}

}
