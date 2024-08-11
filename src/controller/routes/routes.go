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

	// group routes
	users := r.Group("/user")
	{
		users.GET("/:userId", middlewares.VerifyTokenMiddleware, controller.FindUserByIdController)
		users.GET("/get-by-email/:userEmail", middlewares.VerifyTokenMiddleware, controller.FindUserByEmailController)

		users.POST("/", controller.CreateUserController)
		users.PUT("/:userId", middlewares.VerifyTokenMiddleware, controller.UpdateUserController)
		users.DELETE("/:userId", middlewares.VerifyTokenMiddleware, controller.DeleteUserController)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.LoginUserController)
	}
}
