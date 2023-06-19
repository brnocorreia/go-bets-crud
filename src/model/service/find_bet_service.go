package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (bd *betDomainService) FindBetByIDServices(
	id string,
) (model.BetDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByIDServices model",
		zap.String("journey", "findUserByID"),
	)
	return bd.betRepository.FindBetByID(id)
}
