package controller

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByIDController controller",
		zap.String("journey", "FindUserByID"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "FindUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not an valid id",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call  findUserById services",
			err,
			zap.String("journey", "FindUserByID"),
		)
		c.JSON(err.Code, err)
	}

	logger.Info("FindUserByIDController controller executed successfully",
		zap.String("journey", "FindUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmailController controller",
		zap.String("journey", "FindUserByID"),
	)

	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "FindUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not an valid email",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(email)
	if err != nil {
		logger.Error("Error trying to call  findUserByEmail services",
			err,
			zap.String("journey", "FindUserByEmail"),
		)
		c.JSON(err.Code, err)
	}

	logger.Info("FindUserByEmailController controller executed successfully",
		zap.String("journey", "FindUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
