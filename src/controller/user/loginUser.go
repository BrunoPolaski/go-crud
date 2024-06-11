package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) LoginUserController(c *gin.Context) {
	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Invalid JSON body", err,
			zap.String("journey", "loginUser"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"Invalid JSON body")
		c.JSON(errorMessage.Code, errorMessage)
	}

	if err := uc.service.LoginUserService(userLogin); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
