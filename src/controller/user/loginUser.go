package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) LoginUserController(c *gin.Context) {
	var userLogin request.UserLogin
	var ok bool

	userLogin.Email, userLogin.Password, ok = c.Request.BasicAuth()
	if !ok {
		err := rest_err.NewBadRequestError("Invalid basic auth")
		logger.Error("Error parsing basic auth", err)
		c.JSON(err.Code, err)
		return
	}

	userRequest := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domain, token, err := uc.service.LoginUserService(userRequest)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Successfully executed LoginUserController",
		zap.String("ID: ", domain.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": token,
		"user":        domain.GetID(),
	})
}
