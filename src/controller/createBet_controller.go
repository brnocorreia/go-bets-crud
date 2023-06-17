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

var (
	_ model.BetDomainInterface
)

func (bc *betControllerInterface) CreateBet(c *gin.Context) {
	logger.Info("Init CreateBet controller",
		zap.String("journey", "createBet"),
	)
	var BetRequest request.BetRequest

	if err := c.ShouldBindJSON(&BetRequest); err != nil {
		logger.Error("Error trying to validate bet info", err,
			zap.String("journey", "createBet"))
		restErr := validation.ValidateUserError(err)

		logger.Info("Chegou aqui.")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewBetDomain(
		BetRequest.Nickname,
		BetRequest.Date,
		BetRequest.Bookmaker,
		BetRequest.Sport,
		BetRequest.Description,
		BetRequest.Odd,
		BetRequest.Investment,
		BetRequest.Tipster,
		BetRequest.Winner,
	)

	domainResult, err := bc.service.CreateBetServices(domain)
	if err != nil {
		logger.Error("Error trying to call  createBet services",
			err,
			zap.String("journey", "createBet"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Bet created successfully",
		zap.String("journey", "createBet"))

	c.JSON(http.StatusOK, view.ConvertBetDomainToResponse(domainResult))

}
