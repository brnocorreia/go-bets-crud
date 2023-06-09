package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface,
		*rest_err.RestErr)

	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
