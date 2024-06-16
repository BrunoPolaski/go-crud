package routes

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	controller "github.com/BrunoPolaski/go-crud/src/controller/user"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	controller controller.UserControllerInterface,
	r *gin.RouterGroup,
) {
	logger.Info("Setting up routes")
	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, controller.FindUserByIdController)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, controller.FindUserByEmailController)

	r.POST("/createUser", model.VerifyTokenMiddleware, controller.CreateUserController)
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, controller.UpdateUserController)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, controller.DeleteUserController)
	r.POST("/login", controller.LoginUserController)
}
