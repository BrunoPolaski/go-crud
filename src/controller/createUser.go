package controller

import (
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintln("There are incorrect fields"))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
