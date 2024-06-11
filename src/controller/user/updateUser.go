package controller

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/configuration/validation"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUserController(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userUpdateRequest request.UserUpdateRequest
	id := c.Param("userId")
	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("Error converting id to primitive.ObjectID", err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid ID"))
	}

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)

	if err := uc.service.UpdateUserService(userDomain, id); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info(fmt.Sprintf("ID updated: %s", id),
			zap.String("method", "UpdateUser"),
		)

		c.Status(http.StatusOK)
	}
}
