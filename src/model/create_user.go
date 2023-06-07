package model

import (
	"fmt"

	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "CreateUser"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
