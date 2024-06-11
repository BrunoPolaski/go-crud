package controller

import (
	"github.com/BrunoPolaski/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	CreateUserController(c *gin.Context)
	FindUserByIdController(c *gin.Context)
	FindUserByEmailController(c *gin.Context)
	DeleteUserController(c *gin.Context)
	UpdateUserController(c *gin.Context)
	LoginUserController(c *gin.Context)
}

type userController struct {
	service service.UserDomainService
}

func NewUserController(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userController{
		service: serviceInterface,
	}
}
