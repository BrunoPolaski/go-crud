package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/validation"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) LoginUserController(c *gin.Context) {
	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Invalid JSON body", err,
			zap.String("journey", "loginUser"),
		)
		errorMessage := validation.ValidateUserError(err)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userRequest := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domain, err := uc.service.LoginUserService(userRequest)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Successfully executed LoginUserController",
		zap.String("ID: ", domain.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
