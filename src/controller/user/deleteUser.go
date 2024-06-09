package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) DeleteUser(c *gin.Context) {
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

	if err := uc.service.DeleteUser(id); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		c.Status(http.StatusOK)
	}

	logger.Info("User deleted",
		zap.String("method", "DeleteUser"),
	)
}
