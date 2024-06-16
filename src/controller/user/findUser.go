package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

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

func (uc *userController) FindUserByEmailController(c *gin.Context) {
	email := c.Param("userEmail")

	_, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		logger.Error("Bearer not valid", err)
		c.JSON(err.Code, err)
		return
	}

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Invalid user email", err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"Invalid user email",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if user, err := uc.service.FindUserByEmailService(email); err != nil {
		c.JSON(err.Code, err)
		return
	} else {
		logger.Info(fmt.Sprintf("User found %v", user),
			zap.String("method", "FindUserByEmail"),
		)

		c.JSON(http.StatusFound, view.ConvertDomainToResponse(user))
	}
}
