package controller

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Not implemented yet")
	c.JSON(err.Code, err)
}
