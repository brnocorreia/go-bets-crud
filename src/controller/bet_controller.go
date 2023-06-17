package controller

import (
	"github.com/brnocorreia/go-movies-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewBetControllerInterface(
	serviceInterface service.BetDomainService,
) BetControllerInterface {
	return &betControllerInterface{
		service: serviceInterface,
	}
}

type BetControllerInterface interface {
	CreateBet(c *gin.Context)
}

type betControllerInterface struct {
	service service.BetDomainService
}
