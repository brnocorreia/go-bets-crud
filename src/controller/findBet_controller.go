package controller

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (bc *betControllerInterface) FindBetByID(c *gin.Context) {
	logger.Info("Init findBetByIDController controller",
		zap.String("journey", "FindBetByID"),
	)

	betId := c.Param("betId")

	if _, err := primitive.ObjectIDFromHex(betId); err != nil {
		logger.Error("Error trying to validate betId",
			err,
			zap.String("journey", "FindBetByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"BetID is not an valid id",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	betDomain, err := bc.service.FindBetByIDServices(betId)
	if err != nil {
		logger.Error("Error trying to call  findBetById services",
			err,
			zap.String("journey", "FindBetByID"),
		)
		c.JSON(err.Code, err)
	}

	logger.Info("FindBetByIDController controller executed successfully",
		zap.String("journey", "FindBetByID"),
	)
	c.JSON(http.StatusOK, view.ConvertBetDomainToResponse(betDomain))
}
