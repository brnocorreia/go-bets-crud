package controller

import (
	"fmt"
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/configuration/validation"
	"github.com/brnocorreia/go-movies-crud/src/controller/model/request"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "UpdateUser"),
	)

	user, errorAuth := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if errorAuth != nil {
		c.JSON(errorAuth.Code, errorAuth)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	var UserRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "UpdateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		UserRequest.Name,
		UserRequest.Age,
		UserRequest.Nationality,
	)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error trying to call  updateUser services",
			err,
			zap.String("journey", "FindUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
