package controller

import (
	"net/http"

	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/validation"
	"github.com/brnocorreia/go-movies-crud/src/controller/model/request"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
