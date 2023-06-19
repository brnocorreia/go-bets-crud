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
	logger.Info("Init CreateUser model",
		zap.String("journey", "CreateUser"))

	email, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if email != nil {
		return nil, rest_err.NewBadRequestError(
			"Email already used to register another account. Try another email")
	}

	nickname, _ := ud.FindUserByNicknameServices(userDomain.GetNickname())
	if nickname != nil {
		return nil, rest_err.NewBadRequestError(
			"Nickname already used. Try another nickname.")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
