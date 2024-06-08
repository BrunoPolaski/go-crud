package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/validation"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userUpdateRequest request.UserUpdateRequest

	id := c.Param("userId")

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil || strings.TrimSpace(id) == "" {
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

	err := uc.service.UpdateUser(userDomain, id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("ID updated: %s", id),
		zap.String("method", "UpdateUser"),
	)

	c.Status(http.StatusOK)
}
