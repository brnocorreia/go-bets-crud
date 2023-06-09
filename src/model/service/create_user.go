package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser model", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
