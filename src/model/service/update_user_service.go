package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser model",
		zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "UpdateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	return nil

}
