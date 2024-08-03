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

	users := r.Group("/users")
	users.POST("/", middlewares.ApiKeyAuth, controller.CreateUserController)
	users.GET("/", middlewares.BearerAuth, controller.FindAllUsersController)
	users.GET("/:userId", middlewares.BearerAuth, controller.FindUserByIdController)
	users.PUT("/:userId", middlewares.BearerAuth, controller.UpdateUserController)
	users.DELETE("/:userId", middlewares.BearerAuth, controller.DeleteUserController)

	auth := r.Group("/auth")
	auth.POST("/login", middlewares.ApiKeyAuth, controller.LoginUserController)
}
