package controller

import (
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "findUserById"),
	)

	id := c.Param("id")

	if user, err := uc.service.FindUser(id); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info(fmt.Sprintf("User found %v", user),
			zap.String("method", "FindUserById"),
		)
	}
}

func (uc *userController) FindUserByEmail(c *gin.Context) {
	email := c.Param("email")

	if user, err := uc.service.FindUserByEmail(email); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info(fmt.Sprintf("User found %v", user),
			zap.String("method", "FindUserByEmail"),
		)

		c.JSON(200, view.ConvertDomainToResponse(user))
	}
}
