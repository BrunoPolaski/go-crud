package controller

import (
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/validation"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Creating user",
		zap.String("method", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while binding JSON", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info(fmt.Sprintf("User created: %v", userRequest),
		zap.String("method", "CreateUser"),
	)
}
