package controller

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	id := c.Param("id")

	if err := uc.service.DeleteUserService(id); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted",
		zap.String("method", "DeleteUser"),
	)
}
