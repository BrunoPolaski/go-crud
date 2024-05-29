package controller

import (
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
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

}
