package controller

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/validation"
	"github.com/brnocorreia/go-movies-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
}
