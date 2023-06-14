package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailServices model",
		zap.String("journey", "findUserByEmail"),
	)
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByIDServices model",
		zap.String("journey", "findUserByID"),
	)
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPasswordServices model",
		zap.String("journey", "findUserByEmailAndPassword"),
	)
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
