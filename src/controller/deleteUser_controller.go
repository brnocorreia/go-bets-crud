package controller

import (
	"fmt"
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "DeleteUser"),
	)

	user, errorAuth := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if errorAuth != nil {
		c.JSON(errorAuth.Code, errorAuth)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error trying to call  deleteUser services",
			err,
			zap.String("journey", "FindUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
