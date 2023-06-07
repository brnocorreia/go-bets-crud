package service

import (
	"fmt"

	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
