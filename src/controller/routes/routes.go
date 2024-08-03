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

	users := r.Group("/users").Use(middlewares.VerifyTokenMiddleware)
	users.GET("/:userId", controller.FindUserByIdController)
	users.GET("/", controller.FindAllUsersController)
	users.PUT("/:userId", controller.UpdateUserController)
	users.DELETE("/:userId", controller.DeleteUserController)

	r.POST("/", controller.CreateUserController)
	auth := r.Group("/auth")
	auth.POST("/login", controller.LoginUserController)
}
