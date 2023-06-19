package service

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository"
)

func NewBetDomainService(
	betRepository repository.BetRepository,
) BetDomainService {
	return &betDomainService{betRepository}
}

type betDomainService struct {
	betRepository repository.BetRepository
}

type BetDomainService interface {
	CreateBetServices(model.BetDomainInterface) (
		model.BetDomainInterface,
		*rest_err.RestErr)

	FindBetByIDServices(
		id string,
	) (model.BetDomainInterface, *rest_err.RestErr)
}
