package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) DeleteUserController(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	id := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Invalid user ID", err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid user ID"))
		return
	}

	err := uc.service.DeleteUserService(id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted",
		zap.String("method", "DeleteUser"),
	)

	c.Status(http.StatusNoContent)
}
