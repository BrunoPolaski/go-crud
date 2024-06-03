package controller

import (
	"github.com/BrunoPolaski/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
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
