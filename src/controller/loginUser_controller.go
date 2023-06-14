package controller

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/validation"
	"github.com/brnocorreia/go-movies-crud/src/controller/model/request"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller",
		zap.String("journey", "loginUser"),
	)
	var UserRequest request.UserLogin

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		UserRequest.Email,
		UserRequest.Password,
	)

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call  loginUser services",
			err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User logged successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
