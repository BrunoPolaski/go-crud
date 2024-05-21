package controller

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/validation"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	"github.com/BrunoPolaski/go-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := userDomain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User created: %v", userDomain),
		zap.String("method", "CreateUser"),
	)

	c.String(http.StatusOK, "User created")
}
