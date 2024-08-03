package controller

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/controller/model/response"
	"github.com/BrunoPolaski/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) FindAllUsersController(c *gin.Context) {
	logger.Info("Init FindAllUsers controller",
		zap.String("journey", "findAllUsers"),
	)

	email := c.Query("userEmail")

	if users, err := uc.service.FindAllUsersService(email); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info("Users found",
			zap.String("method", "FindAllUsers"),
		)

		usersResponse := make([]response.UserResponse, 0)

		for _, user := range users {
			usersResponse = append(usersResponse, view.ConvertDomainToResponse(user))
		}

		c.JSON(http.StatusOK, usersResponse)
	}
}

func (uc *userController) FindUserByIdController(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "findUserById"),
	)

	id := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Invalid user ID", err,
			zap.String("journey", "findUserById"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"Invalid user ID",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if user, err := uc.service.FindUserByIDService(id); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info(fmt.Sprintf("User found %v", user),
			zap.String("method", "FindUserById"),
		)

		c.JSON(200, view.ConvertDomainToResponse(user))
	}
}
