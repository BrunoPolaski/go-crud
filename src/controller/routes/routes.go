package routes

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/controller/middlewares"
	controller "github.com/BrunoPolaski/go-crud/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	controller controller.UserControllerInterface,
	r *gin.RouterGroup,
) {
	logger.Info("Setting up routes")
	r.GET("/getUserById/:userId", middlewares.VerifyTokenMiddleware, controller.FindUserByIdController)
	r.GET("/getUserByEmail/:userEmail", middlewares.VerifyTokenMiddleware, controller.FindUserByEmailController)

	r.POST("/createUser", middlewares.VerifyTokenMiddleware, controller.CreateUserController)
	r.PUT("/updateUser/:userId", middlewares.VerifyTokenMiddleware, controller.UpdateUserController)
	r.DELETE("/deleteUser/:userId", middlewares.VerifyTokenMiddleware, controller.DeleteUserController)
	r.POST("/login", controller.LoginUserController)
}
