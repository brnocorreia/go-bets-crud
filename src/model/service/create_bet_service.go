package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (bd *betDomainService) CreateBetServices(
	betDomain model.BetDomainInterface,
) (model.BetDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateBet model",
		zap.String("journey", "CreateBet"))

	betDomainRepository, err := bd.betRepository.CreateBet(betDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "betUser"))
		return nil, err
	}

	logger.Info("CreateBet service executed successfully",
		zap.String("betId", betDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return betDomainRepository, nil
}
