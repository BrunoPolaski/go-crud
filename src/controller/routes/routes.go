package routes

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	controller "github.com/BrunoPolaski/go-crud/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	controller controller.UserControllerInterface,
	r *gin.RouterGroup,
) {
	logger.Info("Setting up routes")
	r.GET("/getUserById/:userId", controller.FindUserByIdController)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmailController)

	r.POST("/createUser", controller.CreateUserController)
	r.PUT("/updateUser/:userId", controller.UpdateUserController)
	r.DELETE("/deleteUser/:userId", controller.DeleteUserController)
	r.POST("/login", controller.LoginUserController)
}
